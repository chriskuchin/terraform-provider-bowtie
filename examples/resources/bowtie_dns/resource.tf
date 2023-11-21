# Resolve `example.com` names using the DNS host at 192.0.2.1 but do
# not pass `wrong.example.com` upstream.

resource "bowtie_dns" "example" {
  name = "example.com"
  servers = [{
    addr = "192.0.2.1"
  }]
  excludes = [{
    name = "wrong.example.com"
  }]
}
