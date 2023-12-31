---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "bowtie_resource Resource - terraform-provider-bowtie"
subcategory: ""
description: |-
  Bowtie resources represent network properties like address ranges that may be targeted by policies.
  Note that defining these resources does not implicitly grant or deny access to them - resources must be collected into resource groups and then referenced by policies.
---

# bowtie_resource (Resource)

Bowtie *resources* represent network properties like address ranges that may be targeted by *policies*.

Note that defining these resources does not implicitly grant or deny access to them - resources must be collected into resource groups and then referenced by policies.

## Example Usage

```terraform
resource "bowtie_resource" "ip" {
  name     = "example"
  protocol = "all"
  location = {
    ip = "127.0.0.1"
  }
  ports = {
    range = [
      0, 65535
    ]
  }
}

resource "bowtie_resource" "cidr" {
  name     = "example"
  protocol = "http"
  location = {
    cidr = "10.0.0.0/16"
  }
  ports = {
    collection = [80, 443]
  }

}

resource "bowtie_resource" "dns" {
  name     = "example"
  protocol = "https"
  location = {
    dns = "test.example.com"
  }
  ports = {
    collection = [443, 80, 8080]
  }
}

# Default Resources
resource "bowtie_resource" "all_ipv6" {
  name     = "All IPv6"
  protocol = "all"
  location = {
    cidr = "::/0"
  }
  ports = {
    range = [0, 65535]
  }
}

resource "bowtie_resource" "all_ipv4" {
  name     = "All IPv4"
  protocol = "all"
  location = {
    cidr = "0.0.0.0/0"
  }
  ports = {
    range = [0, 65535]
  }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `location` (Attributes) The address of the resource. May be a CIDR address, single IP, or DNS name. (see [below for nested schema](#nestedatt--location))
- `name` (String) Human readable name of the resource.
- `ports` (Attributes) Which ports to include in this resource. (see [below for nested schema](#nestedatt--ports))
- `protocol` (String) Matching connection protocol.

### Read-Only

- `id` (String) Internal resource ID.

<a id="nestedatt--location"></a>
### Nested Schema for `location`

Optional:

- `cidr` (String) A CIDR address reachable from behind your Bowtie Controller.
- `dns` (String) A DNS name pointing to a resource reachable from behind your Bowtie Controller.
- `ip` (String) The IP address of a resource reachable from behind your Bowtie Controller.


<a id="nestedatt--ports"></a>
### Nested Schema for `ports`

Optional:

- `collection` (List of Number) List of allowed ports.
- `range` (List of Number) First element is the low port and second is the high port (range is inclusive).

## Import

Import is supported using the following syntax:

```shell
terraform import bowtie_resource.example 47480e17-e7a2-4f7d-a0c0-3db8fd86c4ff
```
