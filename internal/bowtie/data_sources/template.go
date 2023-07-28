package data_sources

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

var (
	_ datasource.DataSource = &templateDataSource{}
)

func NewTemplateDataSource() datasource.DataSource {
	return &templateDataSource{}
}

type templateDataSource struct{}

func (d *templateDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_template"
}

func (d *templateDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{}
}

func (d *templateDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
}
