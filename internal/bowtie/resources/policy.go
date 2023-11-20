package resources

import (
	"context"
	"fmt"

	"github.com/bowtieworks/terraform-provider-bowtie/internal/bowtie/client"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &policyResource{}
var _ resource.ResourceWithImportState = &policyResource{}

type policyResource struct {
	client *client.Client
}

type policyResourceModel struct {
	ID     types.String      `tfsdk:"id"`
	Source policySourceModel `tfsdk:"source"`
	Dest   types.String      `tfsdk:"dest"`
	Action types.String      `tfsdk:"action"`
}

type policySourceModel struct {
	ID        types.String   `tfsdk:"id"`
	Predicate predicateModel `tfsdk:"predicate"`
	Always    types.Bool     `tfsdk:"always"`
}

type predicateModel struct {
	And           types.List   `tfsdk:"and"`
	Or            types.List   `tfsdk:"or"`
	Nor           types.List   `tfsdk:"nor"`
	User          types.String `tfsdk:"user"`
	Device        types.String `tfsdk:"device"`
	InUserGroup   types.String `tfsdk:"user_group"`
	InDeviceGroup types.String `tfsdk:"device_group"`
}

type AndPredicateModel struct {
	And types.List `tfsdk:"And"`
}

func NewPolicyResource() resource.Resource {
	return &policyResource{}
}

func (p *policyResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_policy"
}

func (p *policyResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "",
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"source": schema.SingleNestedAttribute{
				MarkdownDescription: "",
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						MarkdownDescription: "",
						Computed:            true,
					},
					"predicate": schema.SingleNestedAttribute{
						Attributes: map[string]schema.Attribute{
							"and": schema.ListNestedAttribute{
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"user":         schema.StringAttribute{},
										"device":       schema.StringAttribute{},
										"user_group":   schema.StringAttribute{},
										"device_group": schema.StringAttribute{},
									},
								},
							},
							"or": schema.ListNestedAttribute{
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"user":         schema.StringAttribute{},
										"device":       schema.StringAttribute{},
										"user_group":   schema.StringAttribute{},
										"device_group": schema.StringAttribute{},
									},
								},
							},
							"nor": schema.ListNestedAttribute{
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"user":         schema.StringAttribute{},
										"device":       schema.StringAttribute{},
										"user_group":   schema.StringAttribute{},
										"device_group": schema.StringAttribute{},
									},
								},
							},
							"user":         schema.StringAttribute{},
							"device":       schema.StringAttribute{},
							"user_group":   schema.StringAttribute{},
							"device_group": schema.StringAttribute{},
							"predicate": schema.SingleNestedAttribute{
								Attributes: map[string]schema.Attribute{
									"user":         schema.StringAttribute{},
									"device":       schema.StringAttribute{},
									"user_group":   schema.StringAttribute{},
									"device_group": schema.StringAttribute{},
								},
							},
						},
					},
				},
			},
			"dest": schema.StringAttribute{
				MarkdownDescription: "",
			},
			"action": schema.StringAttribute{
				MarkdownDescription: "",
			},
		},
	}
}

func (p *policyResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

	p.client = client
}

func (p *policyResource) Create(ctx context.Context, req resource.CreateRequest, respo *resource.CreateResponse) {

}

func (p *policyResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {

}

func (p *policyResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {

}

func (p *policyResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {

}

func (p *policyResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {

}
