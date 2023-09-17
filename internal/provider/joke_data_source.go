package provider

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	//	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
  _ datasource.DataSource = &JokeDataSource{}
  _ datasource.DataSourceWithConfigure = &JokeDataSource{}
)

func NewJokeDataSource() datasource.DataSource {
	return &JokeDataSource{}
}

// ExampleDataSource defines the data source implementation.
type JokeDataSource struct {
	client *http.Client
}

// ExampleDataSourceModel describes the data source data model.
type JokeDataSourceModel struct {
  Categories []string `tfsdk:"categories"`
	CreatedAt  types.String   `tfsdk:"created_at"`
	IconURL    types.String   `tfsdk:"icon_url"`
	ID         types.String   `tfsdk:"id"`
	UpdatedAt  types.String   `tfsdk:"updated_at"`
	URL        types.String   `tfsdk:"url"`
	Value      types.String   `tfsdk:"value"`
	JokeID     string   `tfsdk:"joke_id"`
}

type basicJokeDataModel struct {
  Categories []string `json:"categories"`
	CreatedAt  string   `json:"created_at"`
	IconURL    string   `json:"icon_url"`
	ID         string   `json:"id"`
	UpdatedAt  string   `json:"updated_at"`
	URL        string   `json:"url"`
	Value      string   `json:"value"`
}

func (d *JokeDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_joke"
}

func (d *JokeDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "Joke data source",

    Attributes: map[string]schema.Attribute{
		"categories": schema.ListAttribute {
      ElementType: types.StringType,
      Computed: true,
      Optional: true,
		},
		"created_at": schema.StringAttribute {
      Computed: true,
		},
		"icon_url": schema.StringAttribute {
      Computed: true,
		},
		"id": schema.StringAttribute {
      Computed: true,
		},
		"updated_at": schema.StringAttribute {
      Computed: true,
		},
		"url": schema.StringAttribute {
      Computed: true,
		},
		"value": schema.StringAttribute {
      Computed: true,
		},
    "joke_id": schema.StringAttribute {
      Required: true,
    },
	},


	}
}

func (d *JokeDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*http.Client)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected *http.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	d.client = client
}

func (d *JokeDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data JokeDataSourceModel
	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

  url := "https://api.chucknorris.io/jokes/"+data.JokeID

  fmt.Printf("URL : %v\n", url)
  // Actual call to the API
  response, err := d.client.Get(url)

  if err != nil {
      resp.Diagnostics.AddError("Can't take a joke", err.Error())
      return
  }
  defer response.Body.Close()

  // Check the HTTP response status code
  if response.StatusCode != http.StatusOK {
      resp.Diagnostics.AddError("Failed to fetch joke", fmt.Sprintf("HTTP status: %s", response.Status))
      return
  }

  // Read the http response and get the Body
//  body, err := ioutil.ReadAll(response.Body)
  if err != nil {
      resp.Diagnostics.AddError("Failed to read response body", err.Error())
      return
  }
  
 
  //marshall to regular struct (not TF Crap)
  var basic basicJokeDataModel
  bodytemp, err := ioutil.ReadAll(response.Body)
  
  err = json.Unmarshal(bodytemp, &basic)



//  fmt.Println(string(body))
//  fmt.Printf("Before unmarshal: %+v\n", data)
//  // Unmarshal the JSON response into the data struct
//  if err := json.Unmarshal(body, &data); err != nil {
//      resp.Diagnostics.AddError("Failed to unmarshal JSON response", err.Error())
//      return
//  }  
//  fmt.Printf("After unmarshal: %+v\n", data.ID)



	// If applicable, this is a great opportunity to initialize any necessary
	// provider client data and make a call using it.
	// httpResp, err := d.client.Do(httpReq)
	// if err != nil {
	//     resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read example, got error: %s", err))
	//     return
	// }
  
	// For the purposes of this example code, hardcoding a response value to
	// save into the Terraform state.

  fmt.Printf("%+v\n", basic)
  
  data.ID = types.StringValue(basic.ID)
  data.Value = types.StringValue(basic.Value)
  data.UpdatedAt = types.StringValue(basic.UpdatedAt)
  data.Categories = basic.Categories
  data.URL = types.StringValue(basic.URL)
  data.IconURL = types.StringValue(basic.IconURL)
  data.CreatedAt = types.StringValue(basic.CreatedAt)


  fmt.Printf("%v\n", &data)
//  fmt.Printf("STATE: %v\n",resp.State.Raw)
	// Write logs using the tflog package
	// Documentation: https://terraform.io/plugin/log
	tflog.Trace(ctx, "read a data source")

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
