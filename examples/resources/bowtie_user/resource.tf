resource "bowtie_user" "jane" {
  name  = "Jane Doe"
  email = "jane.doe@example.com"
}

resource "bowtie_user" "owner" {
  name                = "J. Jonah Jameson"
  email               = "jjj@example.com"
  role                = "Owner"
  authz_devices       = true
  authz_policies      = true
  authz_control_plane = true
  authz_users         = true
}

resource "bowtie_user" "disabled" {
  name    = "John Doe"
  email   = "john.doe@example.com"
  role    = "User"
  enabled = false
}

