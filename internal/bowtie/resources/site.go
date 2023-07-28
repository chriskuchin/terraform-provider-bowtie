package resources

import (
	"context"
	"time"

	"github.com/chriskuchin/terraform-provider-bowtie/internal/bowtie/client"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &siteResource{}
var _ resource.ResourceWithImportState = &siteResource{}

type siteResource struct {
	client *client.Client
}

type siteResourceModel struct {
	ID          types.String `tfsdk:"id"`
	Name        types.String `tfsdk:"name"`
	LastUpdated types.String `tfsdk:"last_updated"`
}

func NewSiteResource() resource.Resource {
	return &siteResource{}
}

func (s *siteResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_site"
}

func (s *siteResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"last_updated": schema.StringAttribute{
				Computed: true,
			},
			"name": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "",
			},
		},
	}
}

func (s *siteResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*client.Client)
	if !ok {
		resp.Diagnostics.AddError(
			"",
			"",
		)
	}

	s.client = client
}

func (s *siteResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan siteResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	id, err := s.client.CreateSite(plan.Name.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"",
			"",
		)
	}

	plan.ID = types.StringValue(id)
	plan.LastUpdated = types.StringValue(time.Now().Format(time.RFC850))

	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

}

func (s *siteResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state siteResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	site, err := s.client.GetSite(state.ID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"",
			"",
		)
	}

	state.Name = types.StringValue(site.Name)

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (s *siteResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan siteResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	err := s.client.UpsertSite(plan.ID.ValueString(), plan.Name.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"",
			"",
		)
		return
	}

	plan.LastUpdated = types.StringValue(time.Now().Format(time.RFC850))

	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (s *siteResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state groupResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// TODO Implement Delete on API support
}

func (s *siteResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
