---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "bowtie_group Resource - terraform-provider-bowtie"
subcategory: ""
description: |-
  Manage user groups which assign access policies to groups of users.
---

# bowtie_group (Resource)

Manage user groups which assign access policies to groups of users.

## Example Usage

```terraform
resource "bowtie_group" "admins" {
  name = "Administrators"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) The human-readable name of the group.

### Read-Only

- `id` (String) Internal resource ID.
- `last_updated` (String) Metadata about the last time a write API was called by this provider for this resource.

## Import

Import is supported using the following syntax:

```shell
terraform import bowtie_group.admins 47480e17-e7a2-4f7d-a0c0-3db8fd86c4ff
```
