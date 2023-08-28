resource "bowtie_resource_group" "example" {
  name      = "dev"
  resources = ["a2f99c6e-e9a2-401e-ab2d-3b62e02a2f5d"]
}

# Full Example
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

resource "bowtie_resource_group" "corp" {
  name      = "corp"
  resources = [bowtie_resource.corp.id]
}

resource "bowtie_resource_group" "dns" {
  name      = "dns record"
  resources = [bowtie_resource.dns.id]
  inherited = [bowtie_resource_group.corp.id]
}