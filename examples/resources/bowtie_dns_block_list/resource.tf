# Use a third-party DNS blocklist for content blocking:

resource "bowtie_dns_block_list" "example" {
  name     = "Infosec Block List"
  upstream = "https://raw.githubusercontent.com/hagezi/dns-blocklists/main/domains/tif.txt"
  override_to_allow = [
    "permitted.example.com"
  ]
}
