package bowtie

import (
	"context"
	"os"

	"github.com/bowtieworks/terraform-provider-bowtie/internal/bowtie/client"
	"github.com/bowtieworks/terraform-provider-bowtie/internal/bowtie/data_sources"
	"github.com/bowtieworks/terraform-provider-bowtie/internal/bowtie/resources"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type BowtieProvider struct{}

type bowtieProviderModel struct {
	Host     types.String `tfsdk:"host"`
	Username types.String `tfsdk:"username"`
	Password types.String `tfsdk:"password"`
}

func New() provider.Provider {
	return &BowtieProvider{}
}

func (b *BowtieProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Bowtie Provider",
		Attributes: map[string]schema.Attribute{
			"host": schema.StringAttribute{
				Description: "The url to communicate with the bowtie control plane at example: https://bowtie.example.com",
				Optional:    true,
			},
			"username": schema.StringAttribute{
				Description: "The username/email to login to the bowtie control plane",
				Optional:    true,
			},
			"password": schema.StringAttribute{
				Description: "The password to login to the bowtie control plane",
				Sensitive:   true,
				Optional:    true,
			},
		},
	}
}

func (b *BowtieProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "bowtie"
}

func (b *BowtieProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	var config bowtieProviderModel

	diags := req.Config.Get(ctx, &config)

	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	if config.Host.IsUnknown() {
		resp.Diagnostics.AddAttributeError(
			path.Root("host"),
			"Unknown Bowtie API Host",
			"The provider cannot create the Bowtie API Client as the host value is unknown",
		)
	}

	if config.Username.IsUnknown() {
		resp.Diagnostics.AddAttributeError(
			path.Root("username"),
			"Unknown Bowtie API Username",
			"The provider cannot create the Bowtie API Client as the username value is unknown",
		)
	}

	if config.Password.IsUnknown() {
		resp.Diagnostics.AddAttributeError(
			path.Root("password"),
			"Unknown Bowtie API Password",
			"The provider cannot create the Bowtie API Client as the password value is unknown",
		)
	}

	if resp.Diagnostics.HasError() {
		return
	}

	host := os.Getenv("BOWTIE_HOST")
	username := os.Getenv("BOWTIE_USERNAME")
	password := os.Getenv("BOWTIE_PASSWORD")

	if !config.Host.IsNull() {
		host = config.Host.ValueString()
	}

	if !config.Username.IsNull() {
		username = config.Username.ValueString()
	}

	if !config.Password.IsNull() {
		password = config.Password.ValueString()
	}

	if host == "" {
		resp.Diagnostics.AddAttributeError(
			path.Root("host"),
			"Missing Bowtie API Host",
			"The provider cannot create the Bowtie API client without a host being set",
		)
	}

	if username == "" {
		resp.Diagnostics.AddAttributeError(
			path.Root("username"),
			"Missing Bowtie API Username",
			"The provider cannot create the Bowtie API Client without a username",
		)
	}

	if password == "" {
		resp.Diagnostics.AddAttributeError(
			path.Root("password"),
			"Missing Bowtie API Password",
			"The provider cannot create the Bowtie API Client without a password",
		)
	}

	if resp.Diagnostics.HasError() {
		return
	}

	client, err := client.NewClient(ctx, host, username, password)
	if err != nil {
		resp.Diagnostics.AddError(
			"Failed to create Bowtie API Client",
			"An unexpected error creating the Bowtie API Client:  "+err.Error(),
		)
		return
	}

	resp.DataSourceData = client
	resp.ResourceData = client
}

func (b *BowtieProvider) Resources(_ context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		resources.NewDNSResource,
		resources.NewGroupResource,
		resources.NewSiteRangeResource,
		resources.NewSiteResource,
		resources.NewResourceResource,
		resources.NewResourceGroupResource,
		resources.NewGroupMembershipResource,
	}
}

func (b *BowtieProvider) DataSources(_ context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		data_sources.NewUserDataSource,
	}
}
