package data_sources

import (
	"context"
	"fmt"

	"github.com/bowtieworks/terraform-provider-bowtie/internal/bowtie/client"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ datasource.DataSource              = &userDataSource{}
	_ datasource.DataSourceWithConfigure = &userDataSource{}
)

func NewUserDataSource() datasource.DataSource {
	return &userDataSource{}
}

type userDataSource struct {
	client *client.Client
}

type userModel struct {
	ID                types.String `tfsdk:"id"`
	Name              types.String `tfsdk:"name"`
	Email             types.String `tfsdk:"email"`
	AuthzDevices      types.Bool   `tfsdk:"authz_devices"`
	AuthzPolicies     types.Bool   `tfsdk:"authz_policies"`
	AuthzControlPanel types.Bool   `tfsdk:"authz_control_plane"`
	AuthzUsers        types.Bool   `tfsdk:"authz_users"`
	Status            types.String `tfsdk:"status"`
}

func (u *userDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_user"
}

func (u *userDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"name": schema.StringAttribute{
				Computed: true,
			},
			"email": schema.StringAttribute{
				Required: true,
			},
			"authz_devices": schema.BoolAttribute{
				Computed: true,
			},
			"authz_policies": schema.BoolAttribute{
				Computed: true,
			},
			"authz_control_plane": schema.BoolAttribute{
				Computed: true,
			},
			"authz_users": schema.BoolAttribute{
				Computed: true,
			},
			"status": schema.StringAttribute{
				Computed: true,
			},
		},
	}
}

func (u *userDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*client.Client)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configuration Type",
			fmt.Sprintf("Expected *client.Client, got: %T, please report this to the provider.", req.ProviderData),
		)
	}

	u.client = client
}

func (u *userDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state userModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &state)...)

	user, err := u.client.GetUserByEmail(ctx, state.Email.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"failed to retrieve user",
			"Unexptected error retrieving user: "+state.Email.ValueString()+" err: "+err.Error(),
		)
		return
	}

	state.ID = types.StringValue(user.ID)
	state.Name = types.StringValue(user.Name)
	state.Status = types.StringValue(user.Status)

	state.AuthzControlPanel = types.BoolValue(user.AuthzControlPanel)
	state.AuthzDevices = types.BoolValue(user.AuthzDevices)
	state.AuthzPolicies = types.BoolValue(user.AuthzPolicies)
	state.AuthzUsers = types.BoolValue(user.AuthzUsers)

	resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
}
