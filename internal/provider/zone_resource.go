// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"MyTest/internal/sdk"
	"MyTest/internal/sdk/pkg/models/operations"
	"MyTest/internal/sdk/pkg/models/shared"
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &ZoneResource{}
var _ resource.ResourceWithImportState = &ZoneResource{}

func NewZoneResource() resource.Resource {
	return &ZoneResource{}
}

// ZoneResource defines the resource implementation.
type ZoneResource struct {
	client *sdk.MyTest
}

// ZoneResourceModel describes the resource data model.
type ZoneResourceModel struct {
	AccountID     types.Int64        `tfsdk:"account_id"`
	APIURL        types.String       `tfsdk:"api_url"`
	ApplianceURL  types.String       `tfsdk:"appliance_url"`
	Code          types.String       `tfsdk:"code"`
	Config        *ZoneVcenterConfig `tfsdk:"config"`
	Credential    *ZoneCredential    `tfsdk:"credential"`
	Datacenter    types.String       `tfsdk:"datacenter"`
	Description   types.String       `tfsdk:"description"`
	Enabled       types.Bool         `tfsdk:"enabled"`
	GroupID       types.Int64        `tfsdk:"group_id"`
	Groups        []ZoneGroups       `tfsdk:"groups"`
	ID            types.Int64        `tfsdk:"id"`
	Name          types.String       `tfsdk:"name"`
	ScalePriority types.Int64        `tfsdk:"scale_priority"`
	Type          types.String       `tfsdk:"type"`
	Username      types.String       `tfsdk:"username"`
	Visibility    types.String       `tfsdk:"visibility"`
	ZoneType      *ZoneZoneType      `tfsdk:"zone_type"`
}

func (r *ZoneResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_zone"
}

func (r *ZoneResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Zone Resource",

		Attributes: map[string]schema.Attribute{
			"account_id": schema.Int64Attribute{
				Computed: true,
				Optional: true,
			},
			"api_url": schema.StringAttribute{
				Optional: true,
			},
			"appliance_url": schema.StringAttribute{
				Optional: true,
			},
			"code": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
			"config": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"api_url": schema.StringAttribute{
						Computed: true,
					},
					"appliance_url": schema.StringAttribute{
						Computed: true,
					},
					"datacenter": schema.StringAttribute{
						Computed: true,
					},
					"username": schema.StringAttribute{
						Computed: true,
					},
				},
			},
			"credential": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"type": schema.StringAttribute{
						Computed: true,
					},
				},
			},
			"datacenter": schema.StringAttribute{
				Optional: true,
			},
			"description": schema.StringAttribute{
				Optional:    true,
				Description: `Optional description field if you want to put more info there`,
			},
			"enabled": schema.BoolAttribute{
				Computed: true,
				Optional: true,
			},
			"group_id": schema.Int64Attribute{
				Required:    true,
				Description: `Specifies which Server group this cloud should be assigned to`,
			},
			"groups": schema.ListNestedAttribute{
				Computed: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"account_id": schema.Int64Attribute{
							Computed: true,
						},
						"id": schema.Int64Attribute{
							Computed: true,
						},
						"name": schema.StringAttribute{
							Computed: true,
						},
					},
				},
			},
			"id": schema.Int64Attribute{
				Computed: true,
			},
			"name": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
			"scale_priority": schema.Int64Attribute{
				Computed: true,
				Optional: true,
			},
			"type": schema.StringAttribute{
				Optional: true,
			},
			"username": schema.StringAttribute{
				Optional: true,
			},
			"visibility": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
			"zone_type": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"code": schema.StringAttribute{
						Computed: true,
					},
				},
				Description: `Map containing the Cloud (zone) code name. See the zone-types API to fetch a list of all available Cloud (zone) types and their codes.`,
			},
		},
	}
}

func (r *ZoneResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*sdk.MyTest)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *sdk.MyTest, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = client
}

func (r *ZoneResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *ZoneResourceModel
	var item types.Object

	resp.Diagnostics.Append(req.Plan.Get(ctx, &item)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(item.As(ctx, &data, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})...)

	if resp.Diagnostics.HasError() {
		return
	}

	accountID := new(int64)
	if !data.AccountID.IsUnknown() && !data.AccountID.IsNull() {
		*accountID = data.AccountID.ValueInt64()
	} else {
		accountID = nil
	}
	code := new(string)
	if !data.Code.IsUnknown() && !data.Code.IsNull() {
		*code = data.Code.ValueString()
	} else {
		code = nil
	}
	var config *shared.ZoneVcenterConfig
	if data != nil {
		apiURL := new(string)
		if !data.APIURL.IsUnknown() && !data.APIURL.IsNull() {
			*apiURL = data.APIURL.ValueString()
		} else {
			apiURL = nil
		}
		applianceURL := new(string)
		if !data.ApplianceURL.IsUnknown() && !data.ApplianceURL.IsNull() {
			*applianceURL = data.ApplianceURL.ValueString()
		} else {
			applianceURL = nil
		}
		datacenter := new(string)
		if !data.Datacenter.IsUnknown() && !data.Datacenter.IsNull() {
			*datacenter = data.Datacenter.ValueString()
		} else {
			datacenter = nil
		}
		username := new(string)
		if !data.Username.IsUnknown() && !data.Username.IsNull() {
			*username = data.Username.ValueString()
		} else {
			username = nil
		}
		config = &shared.ZoneVcenterConfig{
			APIURL:       apiURL,
			ApplianceURL: applianceURL,
			Datacenter:   datacenter,
			Username:     username,
		}
	}
	var credential *shared.ZoneCreateCredential
	if data != nil {
		typeVar := new(string)
		if !data.Type.IsUnknown() && !data.Type.IsNull() {
			*typeVar = data.Type.ValueString()
		} else {
			typeVar = nil
		}
		credential = &shared.ZoneCreateCredential{
			Type: typeVar,
		}
	}
	description := new(string)
	if !data.Description.IsUnknown() && !data.Description.IsNull() {
		*description = data.Description.ValueString()
	} else {
		description = nil
	}
	enabled := new(bool)
	if !data.Enabled.IsUnknown() && !data.Enabled.IsNull() {
		*enabled = data.Enabled.ValueBool()
	} else {
		enabled = nil
	}
	groupID := data.GroupID.ValueInt64()
	name := data.Name.ValueString()
	scalePriority := new(int64)
	if !data.ScalePriority.IsUnknown() && !data.ScalePriority.IsNull() {
		*scalePriority = data.ScalePriority.ValueInt64()
	} else {
		scalePriority = nil
	}
	visibility := new(shared.ZoneCreateVisibility)
	if !data.Visibility.IsUnknown() && !data.Visibility.IsNull() {
		*visibility = shared.ZoneCreateVisibility(data.Visibility.ValueString())
	} else {
		visibility = nil
	}
	code1 := new(string)
	if !data.Code.IsUnknown() && !data.Code.IsNull() {
		*code1 = data.Code.ValueString()
	} else {
		code1 = nil
	}
	zoneType := shared.ZoneCreateZoneType{
		Code: code1,
	}
	zone := shared.ZoneCreate{
		AccountID:     accountID,
		Code:          code,
		Config:        config,
		Credential:    credential,
		Description:   description,
		Enabled:       enabled,
		GroupID:       groupID,
		Name:          name,
		ScalePriority: scalePriority,
		Visibility:    visibility,
		ZoneType:      zoneType,
	}
	request := operations.AddCloudsRequestBody{
		Zone: zone,
	}
	res, err := r.client.Clouds.AddClouds(ctx, request)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
		if res != nil && res.RawResponse != nil {
			resp.Diagnostics.AddError("unexpected http request/response", debugResponse(res.RawResponse))
		}
		return
	}
	if res == nil {
		resp.Diagnostics.AddError("unexpected response from API", fmt.Sprintf("%v", res))
		return
	}
	if res.StatusCode != 200 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res.StatusCode), debugResponse(res.RawResponse))
		return
	}
	if res.AddClouds200ApplicationJSONObject == nil || res.AddClouds200ApplicationJSONObject.Zone == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromCreateResponse(res.AddClouds200ApplicationJSONObject.Zone)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ZoneResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *ZoneResourceModel
	var item types.Object

	resp.Diagnostics.Append(req.State.Get(ctx, &item)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(item.As(ctx, &data, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})...)

	if resp.Diagnostics.HasError() {
		return
	}

	id := data.ID.ValueInt64()
	request := operations.GetCloudsRequest{
		ID: id,
	}
	res, err := r.client.Clouds.GetClouds(ctx, request)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
		if res != nil && res.RawResponse != nil {
			resp.Diagnostics.AddError("unexpected http request/response", debugResponse(res.RawResponse))
		}
		return
	}
	if res == nil {
		resp.Diagnostics.AddError("unexpected response from API", fmt.Sprintf("%v", res))
		return
	}
	if res.StatusCode != 200 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res.StatusCode), debugResponse(res.RawResponse))
		return
	}
	if res.GetClouds200ApplicationJSONObject == nil || res.GetClouds200ApplicationJSONObject.Zone == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromGetResponse(res.GetClouds200ApplicationJSONObject.Zone)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ZoneResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *ZoneResourceModel
	merge(ctx, req, resp, &data)
	if resp.Diagnostics.HasError() {
		return
	}

	var requestBody *operations.UpdateCloudsRequestBody
	zone := *data.ToUpdateSDKType()
	requestBody = &operations.UpdateCloudsRequestBody{
		Zone: zone,
	}
	id := data.ID.ValueInt64()
	request := operations.UpdateCloudsRequest{
		RequestBody: requestBody,
		ID:          id,
	}
	res, err := r.client.Clouds.UpdateClouds(ctx, request)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
		if res != nil && res.RawResponse != nil {
			resp.Diagnostics.AddError("unexpected http request/response", debugResponse(res.RawResponse))
		}
		return
	}
	if res == nil {
		resp.Diagnostics.AddError("unexpected response from API", fmt.Sprintf("%v", res))
		return
	}
	if res.StatusCode != 200 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res.StatusCode), debugResponse(res.RawResponse))
		return
	}
	if res.UpdateClouds200ApplicationJSONObject == nil || res.UpdateClouds200ApplicationJSONObject.Zone == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromUpdateResponse(res.UpdateClouds200ApplicationJSONObject.Zone)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ZoneResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *ZoneResourceModel
	var item types.Object

	resp.Diagnostics.Append(req.State.Get(ctx, &item)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(item.As(ctx, &data, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})...)

	if resp.Diagnostics.HasError() {
		return
	}

	id := data.ID.ValueInt64()
	request := operations.RemoveCloudsRequest{
		ID: id,
	}
	res, err := r.client.Clouds.RemoveClouds(ctx, request)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
		if res != nil && res.RawResponse != nil {
			resp.Diagnostics.AddError("unexpected http request/response", debugResponse(res.RawResponse))
		}
		return
	}
	if res == nil {
		resp.Diagnostics.AddError("unexpected response from API", fmt.Sprintf("%v", res))
		return
	}
	if res.StatusCode != 200 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res.StatusCode), debugResponse(res.RawResponse))
		return
	}

}

func (r *ZoneResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
