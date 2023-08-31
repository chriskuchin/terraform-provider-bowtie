package resources

import (
	"context"

	"github.com/chriskuchin/terraform-provider-bowtie/internal/bowtie/client"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &resourceGroupResource{}
var _ resource.ResourceWithImportState = &resourceGroupResource{}

type resourceGroupResource struct {
	client *client.Client
}

type resourceGroupResourceModel struct {
	ID        types.String `tfsdk:"id"`
	Name      types.String `tfsdk:"name"`
	Inherited types.List   `tfsdk:"inherited"`
	Resources types.List   `tfsdk:"resources"`
}

func NewResourceGroupResource() resource.Resource {
	return &resourceGroupResource{}
}

func (rg *resourceGroupResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_resource_group"
}

func (rg *resourceGroupResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id for the resource group in the api",
				Computed:            true,
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "The human readable name/description of the resource group",
				Required:            true,
			},
			"inherited": schema.ListAttribute{
				MarkdownDescription: "The list of resource groups to include in this resource group",
				ElementType:         types.StringType,
				Required:            true,
			},
			"resources": schema.ListAttribute{
				MarkdownDescription: "The resources that should directly be included in this resource group",
				ElementType:         types.StringType,
				Required:            true,
			},
		},
	}
}

func (rg *resourceGroupResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*client.Client)
	if !ok {
		resp.Diagnostics.AddError(
			"Incorrect provider data",
			"The provider data was not appropiate and failed to resolve as *client.Client",
		)
	}

	rg.client = client
}

func (rg *resourceGroupResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan resourceGroupResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resources := []string{}
	resp.Diagnostics.Append(plan.Resources.ElementsAs(ctx, &resources, false)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resource_groups := []string{}
	resp.Diagnostics.Append(plan.Inherited.ElementsAs(ctx, &resource_groups, false)...)
	if resp.Diagnostics.HasError() {
		return
	}

	id, err := rg.client.CreateResourceGroup(ctx, plan.Name.ValueString(), resources, resource_groups)
	if err != nil {
		resp.Diagnostics.AddError(
			"Failed to create the resource group",
			"Unexpected error creating resource group: "+err.Error(),
		)
		return
	}

	plan.ID = types.StringValue(id)

	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
}

func (rg *resourceGroupResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state resourceGroupResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resourceGroup, err := rg.client.GetResourceGroup(state.ID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Failed to read the resource group",
			"Unexpected error reading the resource group: "+state.ID.ValueString()+" err: "+err.Error(),
		)
		return
	}

	state.Name = types.StringValue(resourceGroup.Name)

	inherited, diags := types.ListValueFrom(ctx, types.StringType, resourceGroup.Inherited)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	state.Inherited = inherited

	resources, diags := types.ListValueFrom(ctx, types.StringType, resourceGroup.Resources)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	state.Resources = resources

	resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
}

func (rg *resourceGroupResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan resourceGroupResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resources := []string{}
	resp.Diagnostics.Append(plan.Resources.ElementsAs(ctx, &resources, false)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resource_groups := []string{}
	resp.Diagnostics.Append(plan.Inherited.ElementsAs(ctx, &resource_groups, false)...)
	if resp.Diagnostics.HasError() {
		return
	}

	err := rg.client.UpsertResourceGroup(ctx, plan.ID.ValueString(), plan.Name.ValueString(), resources, resource_groups)
	if err != nil {
		resp.Diagnostics.AddError(
			"Failed updating the resource group",
			"Unexpected error updating the resource group: "+plan.ID.ValueString()+" err: "+err.Error(),
		)
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
}

func (rg *resourceGroupResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var plan resourceGroupResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	err := rg.client.DeleteResourceGroup(plan.ID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Failed deleting the resource group",
			"Unexpected error calling bowtie api to delete resource group: "+plan.ID.ValueString()+" error: "+err.Error(),
		)
	}
}

func (rg *resourceGroupResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
