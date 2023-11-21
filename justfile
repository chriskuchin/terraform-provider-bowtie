envvars := "config.env"

set dotenv-filename := "config.env"

container_cmd := env_var_or_default("COMPOSE_CMD", "docker-compose")

# Print usage
help:
	@just --list

# Generate user documentation
generate:
	go generate ./...

# Ensure documentation is up-to-date
stale-docs: generate
	#!/usr/bin/env bash

	if git diff --no-ext-diff --quiet --exit-code docs
	then
		echo "Documentation is up-to-date"
	else
		echo -e "\n[ ! ] Documentation is out-of-date with source.\n"
		echo "Regenerate and commit updated docs with 'just generate'."
		exit 1
	fi

# Perform documentation checks
stylecheck: generate
	vale docs

# Run the tests
test:
	go test ./... -count=1

# Run all tests, including acceptance tests
acceptance-test: container
	#!/usr/bin/env bash
	# Ensure the container has had time to come up
	sleep 5
	# Run the tests
	TF_ACC=1 just test
	# Shut down the container
	just stop-container || true

# Generate a SITE_ID for the test container in config.env
site-id:
	#!/usr/bin/env bash

	conf={{envvars}}
	if grep SITE_ID $conf &>/dev/null
	then
		echo "SITE_ID present in $conf"
	else
		set -x
		echo "export SITE_ID=$(uuidgen)" >> $conf
	fi

# Generate an init-users file for bootstrapping
init-users:
	#!/usr/bin/env bash

	users_file=container/init-users

	if [[ -e $users_file ]]
	then
		echo "$users_file  exists; use 'just clean' to purge container state"
		exit
	fi

	sed -i '/BOWTIE_USERNAME/d;/BOWTIE_PASSWORD/d' {{envvars}}
	username=admin@example.com
	password=$(openssl rand -hex 16)
	hash=$(echo -n $password | argon2 $(uuidgen) -i -t 3 -p 1 -m 12 -e)
	echo $username:$hash > $users_file
	echo "export BOWTIE_PASSWORD=$password" >> {{envvars}}
	echo "export BOWTIE_USERNAME=$username" >> {{envvars}}
	echo "Generated user $username"

# Start a background container for bowtie-server
container cmd=container_cmd: site-id init-users
	#!/usr/bin/env bash
	source {{envvars}}
	{{cmd}} up --detach

# Stop the background container
stop-container cmd=container_cmd:
	{{cmd}} down

# Remove build and container artifacts
clean:
	git clean -f -d -x container/
