package fortipam

import (
	"fmt"
	"log"
	"strconv"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/terraform-providers/terraform-provider-fortipam/fpamGo"
	"github.com/antihax/optional"
)

func dataSourceCredRead(schema *schema.ResourceData, meta interface{}) error {
	path := schema.Get("path").(string) + "/" + schema.Get("name").(string)
	field := schema.Get("field").(string)
	terra_conf := meta.(TerraformConfiguration)
	client := *terra_conf.FpamClient
	auth := terra_conf.AuthObj

	sec_id := getId(client, auth, path, "secret")
	log.Printf("Sec id: %d", sec_id)
	if sec_id <= 0 {
		return fmt.Errorf("ID retrieval failed");
	}
	sec_intf, _, sec_err := client.SecretdatabaseApi.CmdbSecretDatabaseIdGet(auth, sec_id,
										&fpam_go.SecretdatabaseApiCmdbSecretDatabaseIdGetOpts{
											Fieldname: optional.NewString(field),
										})
	if sec_intf == nil {
		return fmt.Errorf("API connection error	%s", sec_err)
	}
	json_sec_resp := sec_intf.(map[string]interface{})
	sec_http_status := int(json_sec_resp["http_status"].(float64))
	if (sec_http_status != 200) {
		return fmt.Errorf("Secret retrieval failed:	%s", sec_intf)
	}
	result := json_sec_resp["results"].([]interface{})[0]
	value := result.(map[string]interface{})[field]
	if value == nil {
		return fmt.Errorf("the secret does not contain a '%s' field", field)
	}
	schema.SetId(strconv.Itoa(int(sec_id))) //resource is destroyed if id not set
	schema.Set("value", value)
	return nil
}

func dataSourceSecretCred() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceCredRead,

		Schema: map[string]*schema.Schema{
			"value": {
				Computed:		true,
				Description:	"the value of the field of the secret",
				Sensitive:		true,
				Type:			schema.TypeString,
			},
			"field": {
				Description:	"the field to extract from the secret based on name",
				Required:		true,
				Type:			schema.TypeString,
			},
			"path": {
				Description:	"full folder path to the secret	e.g. public_folder/local",
				Required:		true,
				Type:			schema.TypeString,
			},
			"name": {
				Description:	"secret name under path",
				Required:		true,
				Type:			schema.TypeString,
			},
		},
	}
}