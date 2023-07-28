resource "bowtie_site" "corp" {
  name = "corporate"
}

resource "bowtie_site_range" "office" {
  site_id = bowtie_site.corp.id

  name        = "office"
  description = "the office internal network range"
  ipv4_range  = "10.0.0.0/16"
}

resource "bowtie_site_range" "dc_v6" {
  site_id = bowtie_site.corp.id

  name        = "datacenter"
  description = "the office internal network range"
  ipv6_range  = "64:ff9b:1::/48"
}
