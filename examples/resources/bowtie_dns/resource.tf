resource "bowtie_dns" "example" {
  name = "example.com"
  servers = ["192.0.2.1"]
  exclude = ["wrong.example.com"]
}
