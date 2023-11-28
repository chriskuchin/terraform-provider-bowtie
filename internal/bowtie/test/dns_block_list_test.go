package test

import (
	"fmt"
	"testing"

	"github.com/bowtieworks/terraform-provider-bowtie/internal/bowtie/provider"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const (
	resourceType = "bowtie_dns_block_list"
	resourceName = "test"

	blName       = "Test DNS Block List"
	blNameChange = "Different DNS Block List name"
	blUrl        = "https://gist.githubusercontent.com/tylerjl/a98e82a7c62207dcd91aad47110e135d/raw/409e482fa067f0ca21c62f916a2bfb8f8b83bcc4/block.txt"
	blUrlChange  = "https://gist.githubusercontent.com/tylerjl/a98e82a7c62207dcd91aad47110e135d/raw/c128c2d15ddaf787eed4efa61960f984ca61995a/block.txt"
	blOverride   = `ipchicken.com
downloadmoreram.com
`
	blOverrideChange = `ipchicken.com
downloadmoreram.com
neopets.com
`
)

func TestDNSBlockListResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: provider.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Basic tests for upstream URLs
			{
				Config: provider.ProviderConfig + fmt.Sprintf(`
resource "%s" "%s" {
  name = "%s"
  upstream = "%s"
  override_to_allow = <<EOF
%sEOF
}`, resourceType, resourceName, blName, blUrl, blOverride),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceType+"."+resourceName, "name", blName),
					resource.TestCheckResourceAttr(resourceType+"."+resourceName, "upstream", blUrl),
					resource.TestCheckResourceAttr(resourceType+"."+resourceName, "override_to_allow", blOverride),
					resource.TestCheckResourceAttrSet(resourceType+"."+resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceType+"."+resourceName, "last_updated"),
				),
			},
			// ImportState testing
			{
				ResourceName:      resourceType + "." + resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				// The last_updated attribute does not exist in the HashiCups
				// API, therefore there is no value for it during import.
				ImportStateVerifyIgnore: []string{"last_updated"},
			},
			// Update and Read testing
			{
				Config: provider.ProviderConfig + fmt.Sprintf(`
resource "%s" "%s" {
  name = "%s"
  upstream = "%s"
  override_to_allow = <<EOF
%sEOF
}`, resourceType, resourceName, blNameChange, blUrlChange, blOverrideChange),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceType+"."+resourceName, "name", blNameChange),
					resource.TestCheckResourceAttr(resourceType+"."+resourceName, "upstream", blUrlChange),
					resource.TestCheckResourceAttr(resourceType+"."+resourceName, "override_to_allow", blOverrideChange),
					resource.TestCheckResourceAttrSet(resourceType+"."+resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceType+"."+resourceName, "last_updated"),
				),
			},
		},
	})
}
