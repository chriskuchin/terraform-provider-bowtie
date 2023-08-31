resource "bowtie_policy" "example" {
  source = {
    user_id         = ""
    user_group_id   = ""
    device_id       = ""
    device_group_id = ""
    always          = true
  }
  dest   = ""
  action = "Accept"
}