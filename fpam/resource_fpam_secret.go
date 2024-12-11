package fortipam

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
	"strconv"
)

func ResourceSecretCred() *schema.Resource {
	sec_importer := ResourceSecretImport()

	return &schema.Resource{
		Read:     ResourceSecretRead,
		Create:   ResourceSecretCreate,
		Update:   ResourceSecretUpdate,
		Delete:   ResourceSecretDelete,
		Schema:   ResourceSecretSchema(),
		Importer: &sec_importer,
	}
}

func ResourceSecretImport() schema.ResourceImporter {
	return schema.ResourceImporter{
		StateContext: ResourceSecretStateContext,
	}
}

func ResourceSecretSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"path": {
			Description:      "full folder path to the secret	e.g. public_folder/local",
			Required:         true,
			Type:             schema.TypeString,
			DiffSuppressFunc: ResourcePathdSuppressDiff,
		},
		"name": {
			Description:  "secret name under path",
			Required:     true,
			Type:         schema.TypeString,
			ValidateFunc: ResourceNameValid,
		},
		"template": {
			Description: "secret template name",
			Required:    true,
			Type:        schema.TypeString,
		},
		"fields": {
			Description: "Secret field info",
			Required:    true,
			Type:        schema.TypeList,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"name": {
						Type:     schema.TypeString,
						Required: true,
					},
					"value": {
						Type:     schema.TypeString,
						Required: true,
					},
				},
			},
		},
	}
}

func ResourceSecretRead(schema *schema.ResourceData, meta interface{}) error {
	sec_id, id_err := strconv.Atoi(schema.Id())
	if id_err != nil {
		return fmt.Errorf("Id missing from schema")
	}
	ResourceSecretReadHelper(schema, meta, int32(sec_id))
	return nil
}

func ResourceSecretReadHelper(schema *schema.ResourceData, meta interface{}, id int32) error {
	terra_conf := meta.(TerraformConfiguration)
	client := *terra_conf.FpamClient
	auth := terra_conf.AuthObj

	sec_intf, _, _ := client.SecretdatabaseApi.CmdbSecretDatabaseIdGet(auth, id, nil)
	if sec_intf == nil {
		schema.SetId("")
		return fmt.Errorf("Secret ID retrieval failed. Access to secret [%d] denied", int(id))
	}
	json_sec_resp := sec_intf.(map[string]interface{})
	sec_http_status := int(json_sec_resp["http_status"].(float64))
	if sec_http_status != 200 {
		schema.SetId("")
		return fmt.Errorf("Secret ID retrieval failed. Access to secret [%d] denied", int(id))
	}

	result := json_sec_resp["results"].([]interface{})[0].(map[string]interface{})
	sec_name := result["name"].(string)
	template_name := result["template"].(string)
	sec_field := result["field"].([]interface{})
	fld_id := result["folder"].(float64)

	schema.Set("name", sec_name)
	schema.Set("template", template_name)
	var parsed_field []interface{}
	for _, v := range sec_field {
		field := v.(map[string]interface{})
		fld_name := field["name"]
		fld_val := field["value"]
		parsed_map := make(map[string]interface{})
		parsed_map["name"] = fld_name
		parsed_map["value"] = fld_val
		parsed_field = append(parsed_field, parsed_map)
	}
	schema.Set("fields", parsed_field)
	schema.Set("path", getFolderPath(client, auth, int32(fld_id)))
	return nil
}

func ResourceSecretCreate(schema *schema.ResourceData, meta interface{}) error {
	path := schema.Get("path").(string)
	secret := schema.Get("name").(string)
	terra_conf := meta.(TerraformConfiguration)
	client := *terra_conf.FpamClient
	auth := terra_conf.AuthObj

	folder_id := getId(client, auth, path, "folder")
	if folder_id == 0 {
		schema.SetId("")
		return fmt.Errorf("Folder ID retrieval failed. Access to folder [%s] denied", path)
	}
	//pre-check
	sec_id := getId(client, auth, path+"/"+secret, "secret")
	if sec_id > 0 {
		schema.SetId("")
		return fmt.Errorf("Secret:%s already exist under %s", secret, path)
	}
	body := buildSecretCreateBody(schema, folder_id)
	sec_intf, _, sec_err := client.SecretdatabaseApi.CmdbSecretDatabasePost(auth, *body, nil)
	if sec_intf == nil {
		schema.SetId("")
		return fmt.Errorf("Create API connection error	%s", sec_err)
	}
	json_sec_resp := sec_intf.(map[string]interface{})
	sec_http_status := int(json_sec_resp["http_status"].(float64))
	if sec_http_status != 200 {
		schema.SetId("")
		return fmt.Errorf("Secret Create failed:	%s", sec_intf)
	}

	secret_id := int(json_sec_resp["mkey"].(float64))
	log.Println("New secret:%d is created", secret_id)
	schema.SetId(strconv.Itoa(secret_id))
	return ResourceSecretRead(schema, meta)
}

func ResourceSecretDelete(schema *schema.ResourceData, meta interface{}) error {
	id, id_err := strconv.Atoi(schema.Id())

	if id_err != nil {
		return fmt.Errorf("Id missing from schema")
	}
	terra_conf := meta.(TerraformConfiguration)
	client := *terra_conf.FpamClient
	auth := terra_conf.AuthObj

	sec_intf, _, sec_err := client.SecretdatabaseApi.CmdbSecretDatabaseIdDelete(auth, int32(id), nil)
	if sec_intf == nil {
		return fmt.Errorf("Delete API connection error	%s", sec_err)
	}
	log.Println("Secret:%d is deleted", id)
	json_sec_resp := sec_intf.(map[string]interface{})
	sec_http_status := int(json_sec_resp["http_status"].(float64))
	if sec_http_status != 200 {
		return fmt.Errorf("Secret delete failed:	%s", sec_intf)
	}
	schema.SetId("")
	return nil
}

func ResourceSecretUpdate(schema *schema.ResourceData, meta interface{}) error {
	terra_conf := meta.(TerraformConfiguration)
	client := *terra_conf.FpamClient
	auth := terra_conf.AuthObj

	secret := schema.Get("name").(string)
	path := schema.Get("path").(string)
	id, id_err := strconv.Atoi(schema.Id())
	if id_err != nil {
		return fmt.Errorf("Id missing from schema")
	}
	folder_id := getId(client, auth, path, "folder")
	if folder_id == 0 {
		schema.SetId("")
		return fmt.Errorf("Folder ID retrieval failed. Access to folder [%s] denied", path)
	}
	//pre-check
	sec_id := getId(client, auth, path+"/"+secret, "secret")
	if sec_id != 0 && sec_id != int32(id) {
		schema.SetId("")
		return fmt.Errorf("Secret:%s already exist under %s", secret, path)
	}
	body := buildSecretPutBody(schema, folder_id, int32(id))
	sec_intf, _, sec_err := client.SecretdatabaseApi.CmdbSecretDatabaseIdPut(auth, *body, int32(id), nil)
	if sec_intf == nil {
		schema.SetId("")
		return fmt.Errorf("Update API connection error	%s", sec_err)
	}
	json_sec_resp := sec_intf.(map[string]interface{})
	sec_http_status := int(json_sec_resp["http_status"].(float64))
	if sec_http_status != 200 {
		schema.SetId("")
		return fmt.Errorf("Secret Create failed:	%s", sec_intf)
	}

	return ResourceSecretRead(schema, meta)
}

func ResourceSecretStateContext(ctx context.Context, sch *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	sec_id, id_err := strconv.Atoi(sch.Id())
	if id_err != nil {
		return nil, fmt.Errorf("Id missing from schema")
	}
	read_err := ResourceSecretReadHelper(sch, meta, int32(sec_id))
	return []*schema.ResourceData{sch}, read_err
}
