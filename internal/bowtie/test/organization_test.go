package test

import (
	"context"
	"os"
	"strings"
	"testing"
	"text/template"

	"github.com/bowtieworks/terraform-provider-bowtie/internal/bowtie/client"
	"github.com/bowtieworks/terraform-provider-bowtie/internal/bowtie/provider"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

const (
	resourceName = "bowtie_organization.org"
	orgName      = "Test Organization"
	orgDomain    = "Different DNS Block List name"
)

func TestAccOrganizationResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: provider.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Import testing. Note that organizations cannot be
			// created or destroyed, so we rely purely on import to create
			// the initial resource.
			{
				Config:       getOrganizationConfig(resourceName, orgName, orgDomain),
				ResourceName: resourceName,
				ImportState:  true,
				// We can’t control what Id the API comes up with for the
				// organization ID, so derive it dynamically:
				ImportStateIdFunc: getOrgId(),
				// The last_updated attribute does not exist in the HashiCups
				// API, therefore there is no value for it during import.
				ImportStateVerifyIgnore: []string{"last_updated"},
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "name"),
					resource.TestCheckResourceAttrSet(resourceName, "domain"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "last_updated"),
				),
			},
		},
	})
}

func getOrganizationConfig(resource string, name string, domain string) string {
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
	err = tmpl.ExecuteTemplate(output, "organization.tmpl", map[string]interface{}{
		"provider":      provider.ProviderConfig,
		"resource_type": strings.Split(resource, ".")[0],
		"resource":      strings.Split(resource, ".")[1],
		"name":          name,
		"domain":        domain,
	})

	if err != nil {
		panic("Failed to render template")
	}

	return output.String()
}

func getOrgId() resource.ImportStateIdFunc {
	return func(state *terraform.State) (string, error) {
		ctx := context.Background()

		username := os.Getenv("BOWTIE_USERNAME")
		password := os.Getenv("BOWTIE_PASSWORD")

		client, err := client.NewClient(ctx, "http://localhost:3000", username, password, false)

		if err != nil {
			return "", err
		}

		org, err := client.GetOrganization()
		if err != nil {
			return "", err
		}

		return org.ID, nil
	}
}
