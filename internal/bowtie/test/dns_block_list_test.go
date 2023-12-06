package test

import (
	"strings"
	"testing"
	"text/template"

	"github.com/bowtieworks/terraform-provider-bowtie/internal/bowtie/provider"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const (
	resourceName = "bowtie_dns_block_list.test"

	blName       = "Test DNS Block List"
	blNameChange = "Different DNS Block List name"
	blUrl        = "https://gist.githubusercontent.com/tylerjl/a98e82a7c62207dcd91aad47110e135d/raw/409e482fa067f0ca21c62f916a2bfb8f8b83bcc4/block.txt"
	blUrlChange  = "https://gist.githubusercontent.com/tylerjl/a98e82a7c62207dcd91aad47110e135d/raw/c128c2d15ddaf787eed4efa61960f984ca61995a/block.txt"
)

var blOverride = []string{"ipchicken.com", "downloadmoreram.com"}
var blOverrideChange = []string{"ipchicken.com", "downloadmoreram.com", "neopets.com"}

func TestDNSBlockListResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: provider.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Basic tests for upstream URLs
			{
				Config: getDNSBlockListConfig(resourceName, blName, blUrl, blOverride),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", blName),
					resource.TestCheckResourceAttr(resourceName, "upstream", blUrl),
					resource.TestCheckResourceAttr(resourceName, "override_to_allow.0", blOverride[0]),
					resource.TestCheckResourceAttr(resourceName, "override_to_allow.1", blOverride[1]),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "last_updated"),
				),
			},
			// ImportState testing
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				// The last_updated attribute does not exist in the HashiCups
				// API, therefore there is no value for it during import.
				ImportStateVerifyIgnore: []string{"last_updated"},
			},
			// Update and Read testing
			{
				Config: getDNSBlockListConfig(resourceName, blNameChange, blUrlChange, blOverrideChange),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", blNameChange),
					resource.TestCheckResourceAttr(resourceName, "upstream", blUrlChange),
					resource.TestCheckResourceAttr(resourceName, "override_to_allow.0", blOverrideChange[0]),
					resource.TestCheckResourceAttr(resourceName, "override_to_allow.1", blOverrideChange[1]),
					resource.TestCheckResourceAttr(resourceName, "override_to_allow.2", blOverrideChange[2]),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "last_updated"),
				),
			},
		},
	})
}

func getDNSBlockListConfig(resource string, name string, url string, overrides []string) string {
	funcMap := template.FuncMap{
		"notNil": func(val any) bool {
			return val != nil
		},
	}

	tmpl, err := template.New("").Funcs(funcMap).ParseGlob("testdata/*.tmpl")
	if err != nil {
		return ""
	}

	var output *strings.Builder = &strings.Builder{}
	err = tmpl.ExecuteTemplate(output, "dns_block_list.tmpl", map[string]interface{}{
		"provider":  provider.ProviderConfig,
		"resource":  strings.Split(resource, ".")[1],
		"name":      name,
		"upstream":  url,
		"overrides": overrides,
	})

	if err != nil {
		panic("Failed to render template")
	}

	return output.String()
}
