package resources

import (
	"context"
	"time"

	"github.com/chriskuchin/terraform-provider-bowtie/internal/bowtie/client"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &siteRangeResource{}
var _ resource.ResourceWithImportState = &siteRangeResource{}

type siteRangeResource struct {
	client *client.Client
}

type siteRangeResourceModel struct {
	ID          types.String `tfsdk:"id"`
	SiteID      types.String `tfsdk:"site_id"`
	Name        types.String `tfsdk:"name"`
	Description types.String `tfsdk:"description"`
	IPV4Range   types.String `tfsdk:"ipv4_range"`
	IPV6Range   types.String `tfsdk:"ipv6_range"`
	Weight      types.Int64  `tfsdk:"weight"`
	Metric      types.Int64  `tfsdk:"metric"`
	LastUpdated types.String `tfsdk:"last_updated"`
}

func (srm siteRangeResourceModel) getRange() types.String {
	if !srm.IPV4Range.IsNull() {
		return srm.IPV4Range
	}

	return srm.IPV6Range
}

func NewSiteRangeResource() resource.Resource {
	return &siteRangeResource{}
}

func (sr *siteRangeResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_site_range"
}

func (sr *siteRangeResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
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

func (sr *siteRangeResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

	sr.client = client
}

func (sr *siteRangeResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan siteRangeResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	id, err := sr.client.CreateSiteRange(plan.SiteID.ValueString(), plan.Name.ValueString(), plan.Description.ValueString(), plan.getRange().ValueString(), !plan.IPV4Range.IsNull(), !plan.IPV4Range.IsNull(), int(plan.Weight.ValueInt64()), int(plan.Metric.ValueInt64()))
	if err != nil {
		resp.Diagnostics.AddError(
			"",
			"",
		)
	}

	plan.ID = types.StringValue(id)
	plan.LastUpdated = types.StringValue(time.Now().Format(time.RFC850))

	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
}

func (sr *siteRangeResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state siteRangeResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// sr.client.Get
	info, err := sr.client.GetSiteRange(state.SiteID.ValueString(), state.ID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"",
			"",
		)
	}

	state.Name = types.StringValue(info.Name)
	state.Description = types.StringValue(info.Description)
	state.Weight = types.Int64Value(int64(info.Weight))
	state.Metric = types.Int64Value(int64(info.Metric))

	if info.ISV6 {
		state.IPV6Range = types.StringValue(info.Range)
	} else if info.ISV4 {
		state.IPV4Range = types.StringValue(info.Range)
	}
}

func (sr *siteRangeResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {

}

func (sr *siteRangeResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {

}

func (sr *siteRangeResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {

}
