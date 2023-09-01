resource "bowtie_group" "admins" {
  name = "admins"
}

resource "bowtie_group_membership" "admin_memberships" {
  group_id = bowtie_group.admins.id
  users = [
    "814db1a1-777e-4552-b0c9-bbb69de32cb5",
    "example@example.com"
  ]
}
