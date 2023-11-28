package test

import (
	"context"
	"os"
	"testing"

	"github.com/bowtieworks/terraform-provider-bowtie/internal/bowtie/client"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestMain(m *testing.M) {
	resource.TestMain(m)
}

func getBowtieClient(ctx context.Context, host string) (*client.Client, error) {
	username := os.Getenv("BOWTIE_USERNAME")
	password := os.Getenv("BOWTIE_PASSWORD")

	c, err := client.NewClient(ctx, host, username, password)
	return c, err

}
