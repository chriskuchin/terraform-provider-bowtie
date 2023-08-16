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
