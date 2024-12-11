package fortipam

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
	"strconv"
)

func ResourceSecretFolder() *schema.Resource {
	folder_importer := ResourceFolderImport()

	return &schema.Resource{
		Read:     ResourceFolderRead,
		Create:   ResourceFolderCreate,
		Update:   ResourceFolderUpdate,
		Delete:   ResourceFolderDelete,
		Schema:   ResourceFolderSchema(),
		Importer: &folder_importer,
	}
}

func ResourceFolderImport() schema.ResourceImporter {
	return schema.ResourceImporter{
		StateContext: ResourceFolderStateContext,
	}
}

func ResourceFolderSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"parent": {
			Description:      "full folder path to the parent folder	e.g. public_folder/local",
			Required:         true,
			Type:             schema.TypeString,
			DiffSuppressFunc: ResourcePathdSuppressDiff,
		},
		"name": {
			Description:  "folder name under parent folder",
			Required:     true,
			Type:         schema.TypeString,
			ValidateFunc: ResourceNameValid,
		},
		"path": {
			Description: "full path to the folder",
			Computed:    true,
			Type:        schema.TypeString,
		},
	}
}

func ResourceFolderRead(schema *schema.ResourceData, meta interface{}) error {
	folder_id, id_err := strconv.Atoi(schema.Id())
	if id_err != nil {
		return fmt.Errorf("Id missing from schema")
	}
	ResourceFolderReadHelper(schema, meta, int32(folder_id))
	return nil
}

func ResourceFolderReadHelper(schema *schema.ResourceData, meta interface{}, id int32) error {
	terra_conf := meta.(TerraformConfiguration)
	client := *terra_conf.FpamClient
	auth := terra_conf.AuthObj

	fld_intf, _, _ := client.SecretfolderApi.CmdbSecretFolderIdGet(auth, id, nil)
	if fld_intf == nil {
		schema.SetId("")
		return fmt.Errorf("Folder ID retrieval failed. Access to folder [%d] denied", int(id))
	}
	json_fld_resp := fld_intf.(map[string]interface{})
	sec_http_status := int(json_fld_resp["http_status"].(float64))
	if sec_http_status != 200 {
		schema.SetId("")
		return fmt.Errorf("Folder ID retrieval failed. Access to folder [%d] denied", int(id))
	}
	result := json_fld_resp["results"].([]interface{})[0].(map[string]interface{})
	folder_name := result["name"].(string)
	fld_id := result["parent-folder"].(float64)

	schema.Set("name", folder_name)
	schema.Set("parent", getFolderPath(client, auth, int32(fld_id)))
	return nil
}

func ResourceFolderCreate(schema *schema.ResourceData, meta interface{}) error {
	parent_folder := schema.Get("parent").(string)
	folder := schema.Get("name").(string)
	terra_conf := meta.(TerraformConfiguration)
	client := *terra_conf.FpamClient
	auth := terra_conf.AuthObj

	parent_id := getId(client, auth, parent_folder, "folder")
	if parent_id == 0 {
		return fmt.Errorf("Parent folder ID retrieval failed. Access to folder [%s] denied", parent_folder)
	}
	folder_id := getId(client, auth, parent_folder+"/"+folder, "folder")
	if folder_id > 0 {
		schema.SetId("")
		return fmt.Errorf("Folder:%s already exist under %s", folder, parent_folder)
	}
	err := ResourceFolderCreateHelper(schema, meta, folder, parent_id, parent_folder)
	if err != nil {
		return err
	}

	return nil
}

// Creation assumes pre-check is always being done
func ResourceFolderCreateHelper(schema *schema.ResourceData, meta interface{}, folder_name string, root_id int32, root_name string) error {
	terra_conf := meta.(TerraformConfiguration)
	client := *terra_conf.FpamClient
	auth := terra_conf.AuthObj

	path_str := root_name + "/" + folder_name
	parent_id := root_id

	folder_body := buildFolderCreateBody(parent_id, folder_name)
	folder_intf, _, intf_err := client.SecretfolderApi.CmdbSecretFolderPost(auth, *folder_body, nil)
	if folder_intf == nil {
		return fmt.Errorf("Create API connection error	%s", intf_err)
	}
	json_sec_resp := folder_intf.(map[string]interface{})
	sec_http_status := int(json_sec_resp["http_status"].(float64))
	if sec_http_status != 200 {
		return fmt.Errorf("Folder Create failed:	%s", folder_intf)
	}

	folder_id := int32(json_sec_resp["mkey"].(float64))
	log.Println("New folder:%d is created", folder_id)
	schema.SetId(strconv.Itoa(int(folder_id)))
	schema.Set("path", path_str)
	return nil
}

func ResourceFolderDelete(schema *schema.ResourceData, meta interface{}) error {
	id, id_err := strconv.Atoi(schema.Id())

	if id_err != nil {
		return fmt.Errorf("Id missing from schema")
	}
	terra_conf := meta.(TerraformConfiguration)
	client := *terra_conf.FpamClient
	auth := terra_conf.AuthObj

	folder_intf, _, intf_err := client.SecretfolderApi.CmdbSecretFolderIdDelete(auth, int32(id), nil)
	if folder_intf == nil {
		return fmt.Errorf("Delete API connection error	%s. Please make sure no secret and sub-folder is under current folder", intf_err)
	}
	log.Println("folder:%d is deleted", id)
	json_sec_resp := folder_intf.(map[string]interface{})
	sec_http_status := int(json_sec_resp["http_status"].(float64))
	if sec_http_status != 200 {
		return fmt.Errorf("Folder delete failed:	%s", folder_intf)
	}

	schema.SetId("")
	return nil
}

func ResourceFolderUpdate(schema *schema.ResourceData, meta interface{}) error {
	parent_folder := schema.Get("parent").(string)
	folder := schema.Get("name").(string)
	terra_conf := meta.(TerraformConfiguration)
	client := *terra_conf.FpamClient
	auth := terra_conf.AuthObj
	id, id_err := strconv.Atoi(schema.Id())

	if id_err != nil {
		return fmt.Errorf("Id missing from schema")
	}
	parent_id := getId(client, auth, parent_folder, "folder")
	if parent_id == 0 {
		schema.SetId("")
		return fmt.Errorf("Parent folder ID retrieval failed. Access to folder [%s] denied", parent_folder)
	}

	folder_body := buildFolderPutBody(int32(id), parent_id, folder)
	folder_intf, _, intf_err := client.SecretfolderApi.CmdbSecretFolderIdPut(auth, *folder_body, int32(id), nil)
	if folder_intf == nil {
		return fmt.Errorf("Update API connection error	%s", intf_err)
	}
	json_sec_resp := folder_intf.(map[string]interface{})
	sec_http_status := int(json_sec_resp["http_status"].(float64))
	if sec_http_status != 200 {
		schema.SetId("")
		return fmt.Errorf("Folder Create failed:	%s", folder_intf)
	}
	path_str := parent_folder + "/" + folder
	schema.Set("path", path_str)
	return nil
}

func ResourceFolderStateContext(ctx context.Context, sch *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	folder_id, id_err := strconv.Atoi(sch.Id())
	if id_err != nil {
		return nil, fmt.Errorf("Id missing from schema")
	}
	read_err := ResourceFolderReadHelper(sch, meta, int32(folder_id))
	return []*schema.ResourceData{sch}, read_err
}
