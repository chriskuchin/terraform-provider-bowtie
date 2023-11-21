# Attach only a single resource to this resource group:
resource "bowtie_resource_group" "single" {
  name      = "Internal Address"
  resources = ["a2f99c6e-e9a2-401e-ab2d-3b62e02a2f5d"]
}

# Define HTTP and HTTPS access to all addresses under the CIDR range
# 10.0.0.0/16:
resource "bowtie_resource" "cidr" {
  name     = "Private Range"
  protocol = "http"
  location = {
    cidr = "10.0.0.0/16"
  }
  ports = {
    collection = [80, 443]
  }
}

# Control HTTP-based access to an internal DNS name:
resource "bowtie_resource" "dns" {
  name     = "Access to test.example.com"
  protocol = "https"
  location = {
    dns = "test.example.com"
  }
  ports = {
    collection = [443, 80, 8080]
  }
}

# Reference the `bowtie_resource.corp` resource to place it into a
# group directly.
#
# First, create the resource:
resource "bowtie_resource" "corp" {
  name     = "Internal Corporate Range"
  protocol = "http"
  location = {
    cidr = "10.0.0.0/16"
  }
  ports = {
    range = [
      0, 65535
    ]
  }
}
# Then, reference the resource ID:
resource "bowtie_resource_group" "corp" {
  name      = "Corporate Resources"
  resources = [bowtie_resource.corp.id]
}

# Combine resources and inherit from other resource groups:
resource "bowtie_resource_group" "combined" {
  name      = "Internal resources and corporate network."
  resources = [bowtie_resource.dns.id]
  inherited = [bowtie_resource_group.corp.id]
}
