package test

import (
	"strings"
	"testing"
	"text/template"

	"github.com/bowtieworks/terraform-provider-bowtie/internal/bowtie/provider"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
)

func TestAccSiteRangeResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: provider.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read testing
			{
				Config: getSiteRangeConfig("Test Site", "Office", "Office network CIDR", "10.0.0.0/16", 1, 255),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectNonEmptyPlan(),
					},
					PostApplyPostRefresh: []plancheck.PlanCheck{
						plancheck.ExpectEmptyPlan(),
					},
				},
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("bowtie_site_range.test", "name", "Office"),
					resource.TestCheckResourceAttr("bowtie_site_range.test", "description", "Office network CIDR"),
					resource.TestCheckResourceAttr("bowtie_site_range.test", "ipv4_range", "10.0.0.0/16"),
					resource.TestCheckResourceAttr("bowtie_site_range.test", "metric", "255"),
					resource.TestCheckResourceAttr("bowtie_site_range.test", "weight", "1"),
					resource.TestCheckResourceAttrSet("bowtie_site_range.test", "id"),
					resource.TestCheckResourceAttrSet("bowtie_site_range.test", "last_updated"),
				),
			},
			{
				Config: getSiteRangeConfig("Test Site", "LA Office", "LA Office network CIDR", "10.0.0.0/16", 1, 255),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectNonEmptyPlan(),
					},
					PostApplyPostRefresh: []plancheck.PlanCheck{
						plancheck.ExpectEmptyPlan(),
					},
				},
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("bowtie_site_range.test", "name", "LA Office"),
					resource.TestCheckResourceAttr("bowtie_site_range.test", "description", "LA Office network CIDR"),
					resource.TestCheckResourceAttr("bowtie_site_range.test", "ipv4_range", "10.0.0.0/16"),
					resource.TestCheckResourceAttr("bowtie_site_range.test", "metric", "255"),
					resource.TestCheckResourceAttr("bowtie_site_range.test", "weight", "1"),
					resource.TestCheckResourceAttrSet("bowtie_site_range.test", "id"),
					resource.TestCheckResourceAttrSet("bowtie_site_range.test", "last_updated"),
				),
			},
		},
	})
}

func getSiteRangeConfig(siteName, rangeName, rangeDescription, rangeCIDR string, weight, metric int) string {
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
	err = tmpl.ExecuteTemplate(output, "site_range.tmpl", map[string]interface{}{
		"provider":          provider.ProviderConfig,
		"site_name":         siteName,
		"range_name":        rangeName,
		"range_description": rangeDescription,
		"range_ipv4_cidr":   rangeCIDR,
		"range_weight":      weight,
		"range_metric":      metric,
	})
	if err != nil {
		panic("Failed to render template")
	}

	return output.String()

}
