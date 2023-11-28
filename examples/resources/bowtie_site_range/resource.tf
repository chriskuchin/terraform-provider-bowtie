# Create a new site named "Corporate"
resource "bowtie_site" "corp" {
  name = "Corporate"
}

# Associate the CIDR range 10.0.0.0/16 with the `corp` site:
resource "bowtie_site_range" "office" {
  site_id = bowtie_site.corp.id

  name        = "Office"
  description = "The office internal network range"
  ipv4_range  = "10.0.0.0/16"
}

# Associate a datacenter IPv6 range with the site as well:
resource "bowtie_site_range" "dc_v6" {
  site_id = bowtie_site.corp.id

  name        = "Datacenter"
  description = "The datacenter internal network range"
  ipv6_range  = "64:ff9b:1::/48"
}
