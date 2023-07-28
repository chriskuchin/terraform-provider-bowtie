package resources

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &TemplateResource{}
var _ resource.ResourceWithImportState = &TemplateResource{}

type TemplateResource struct {
}

type TemplateResourceModel struct {
}

func NewTemplateResource() resource.Resource {
	return &TemplateResource{}
}

func (g *TemplateResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_group"
}

func (g *TemplateResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{}
}

func (g *TemplateResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {

}

func (g *TemplateResource) Create(ctx context.Context, req resource.CreateRequest, respo *resource.CreateResponse) {

}

func (g *TemplateResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {

}

func (g *TemplateResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {

}

func (g *TemplateResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {

}

func (g *TemplateResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {

}
