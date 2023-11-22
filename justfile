envvars := "config.env"

# Print usage
help:
	@just --list

# Generate a SITE_ID for the test container
site-id:
	#!/usr/bin/env bash

	conf={{envvars}}
	if grep SITE_ID $conf &>/dev/null
	then
		echo "SITE_ID present in $conf"
	else
		set -x
		echo "SITE_ID=$(uuidgen)" >> $conf
	fi

# Generate an init-users file for bootstrapping
init-users:
	#!/usr/bin/env bash

	sed -i '/BOWTIE_USERNAME/d;/BOWTIE_PASSWORD/d' {{envvars}}
	username=admin@example.com
	password=$(openssl rand -hex 16)
	hash=$(echo -n $password | argon2 $(uuidgen) -i -t 3 -p 1 -m 12 -e)
	echo $username:$hash > container/init-users
	echo "BOWTIE_PASSWORD=$password" >> {{envvars}}
	echo "BOWTIE_USERNAME=$username" >> {{envvars}}

# Start a background container for bowtie-server
container cmd="docker-compose": site-id
	{{cmd}} up --detach

podman: (container "podman-compose")

# Remove build and container artifacts
clean:
	git clean -f -d -x container/
