package data_sources

import (
	"context"

	"github.com/chriskuchin/terraform-provider-bowtie/internal/bowtie/client"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ datasource.DataSource = &groupDataSource{}
)

func NewGroupDataSource() datasource.DataSource {
	return &groupDataSource{}
}

type groupDataSource struct {
	client *client.Client
}

type groupDataSourceModel struct {
	ID types.String `tfsdk:"id"`
}

func (d *groupDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_group"
}

func (d *groupDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "TODO",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "TODO",
				Required:            true,
			},
		},
	}
}

func (d *groupDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
}
