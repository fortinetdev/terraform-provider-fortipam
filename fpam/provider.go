package fortipam

import (
	"github.com/terraform-providers/terraform-provider-fortipam/fpamGo"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"context"
	"net/http"
	"crypto/tls"
)

type TerraformConfiguration struct {
	BaseConfiguration	*fpam_go.Configuration
	AuthObj				context.Context
	FpamClient 			*fpam_go.APIClient
}

func FpamConfig(ctx context.Context, schema *schema.ResourceData) (interface{}, diag.Diagnostics) {
	tConf := TerraformConfiguration{}
	conf := fpam_go.NewConfiguration()

	conf.BasePath =  schema.Get("base_url").(string) + "/api/v2"
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: !schema.Get("verify_ssl").(bool)},
	}
	conf.HTTPClient = &http.Client{Transport: tr}
	client := fpam_go.NewAPIClient(conf)
	auth := context.WithValue(context.Background(), fpam_go.ContextAPIKey, fpam_go.APIKey{
		Key: schema.Get("access_token").(string),
		Prefix: "Bearer",
	})

	tConf.BaseConfiguration = conf
	tConf.FpamClient = client
	tConf.AuthObj = auth
	return tConf, nil
}

func Provider() *schema.Provider {
	return &schema.Provider{
		DataSourcesMap: map[string]*schema.Resource{
			"fortipam_secret":  dataSourceSecretCred(),
		},
		ResourcesMap: map[string]*schema.Resource{
			"fortipam_resource_secret": ResourceSecretCred(),
			"fortipam_resource_folder": ResourceSecretFolder(),
		},
		Schema: map[string]*schema.Schema{
			"base_url": {
				Type:			schema.TypeString,
				Required:		true,
				DefaultFunc:	schema.EnvDefaultFunc("FPAM_URL", nil),
				Description:	"FortiPAM server URL https://192.168.1.99",
			},
			"access_token": {
				Type:			schema.TypeString,
				Required:		true,
				DefaultFunc:	schema.EnvDefaultFunc("FPAM_TOKEN", nil),
				Description:	"FortiPAM API user access token",
			},
			"verify_ssl": {
				Type:			schema.TypeBool,
				Optional:		true,
				Default: 		false,
				DefaultFunc:	schema.EnvDefaultFunc("FPAM_VERIFY_SSL", nil),
				Description:	"FortiPAM API require verify SSL",
			},
		},
		ConfigureContextFunc: FpamConfig,
	}
}