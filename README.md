# Terraform Provider [Bowtie](https://docs.bowtie.works)

## Using this provider:

In our example we'll set up these variables

    export TF_VAR_bowtie_admin_user=issac@bowtie.works
    export TF_VAR_bowtie_admin_password=hunter2
    export TF_VAR_bowtie_host=https://canary-8.net.rock.associates.example

And then configuring a few dns records would look something like this:

    terraform {
    required_providers {
        bowtie = {
            source = "bowtie.works/bowtie/bowtie"
            version = "0.1.3"
            }
        }
    }

    variable bowtie_host {
        type = string
        nullable = false
    }

    variable bowtie_admin_user {
        type = string
        nullable = false
    }

    variable bowtie_admin_password {
        type = string
        nullable = false
    }


    provider "bowtie" {
        host     = var.bowtie_host
        username = var.bowtie_admin_user
        password = var.bowtie_admin_password
    }

    resource "bowtie_dns" "freshbooks" {
        name = "rock.associates.com"
        servers = [{
            addr = "172.128.40.78",
        },
        {
            addr = "172.16.40.199",
        }]
        is_dns64 = true
        is_drop_all = false
        is_drop_a = true
        is_log = true
        excludes = []
    }

    resource "bowtie_dns" "contoso" {
        name = "contoso.com"
        servers = [{
            addr = "9.9.9.9",
        },
        {
            addr = "1.1.1.1",
        }]
        is_drop_all = true
        is_drop_a = true
        is_log = true
        excludes = []
    }

Then you can run `terraform plan` and `terraform apply` as usual

## Building

Setup your dev environment:

    terraform-provider-bowtie on  main [!?] via 🐹 v1.21.2 via 🐍 v3.10.12 (env) 
    ❯ cat ~/.terraformrc 
    provider_installation {
      dev_overrides {
        "bowtie.works/bowtie/bowtie" = "/home/issac/Projects/bowtie/terraform-provider-bowtie"
      }
    }


go build -o terraform-provider-bowtie
