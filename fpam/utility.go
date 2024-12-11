package fortipam

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/terraform-providers/terraform-provider-fortipam/fpamGo"
	"log"
	"strings"
)

func ResourceNameValid(val any, key string) (warns []string, errs []error) {
	v := val.(string)
	if strings.Contains(v, "/") {
		errs = append(errs, fmt.Errorf("Name cannot contain \"/\""))
	}
	return
}

func ResoucePathSplitFunc(c rune) bool {
	return c == '/'
}

func ResourcePathdSuppressDiff(k, old, new string, d *schema.ResourceData) bool {
	old_parts := strings.FieldsFunc(old, ResoucePathSplitFunc)
	new_parts := strings.FieldsFunc(new, ResoucePathSplitFunc)
	old_len := len(old_parts)
	new_len := len(new_parts)

	if old_len != new_len {
		return false
	}
	for i := 0; i < old_len; i++ {
		if old_parts[i] != new_parts[i] {
			return false
		}
	}

	return true
}

func getId(client fpam_go.APIClient, auth context.Context, path string, _type string) int32 {
	id_intf, _, id_err := client.PathToSecretfolderIDApi.UtilityIdPathGet(auth, _type, path)
	if id_intf == nil {
		fmt.Errorf("API connection error	%s", id_err)
		return 0
	}
	json_id_resp := id_intf.(map[string]interface{})
	id_http_status := int(json_id_resp["http_status"].(float64))
	if id_http_status != 200 {
		fmt.Errorf("ID retrieval failed	http resp returned `%s`", id_http_status)
		return 0
	}

	id := int32(json_id_resp["results"].(map[string]interface{})["id"].(float64))
	return id
}

func getFolderPath(client fpam_go.APIClient, auth context.Context, id int32) string {
	curr_id := id
	path := ""

	for curr_id != 0 {
		fld_intf, _, _ := client.SecretfolderApi.CmdbSecretFolderIdGet(auth, curr_id, nil)
		if fld_intf == nil {
			curr_id = 0
			path = "<unknown folder>/" + path
		} else {
			json_fld_resp := fld_intf.(map[string]interface{})
			sec_http_status := int(json_fld_resp["http_status"].(float64))
			if sec_http_status != 200 {
				curr_id = 0
				path = "<unknown folder>/" + path
			} else {
				result := json_fld_resp["results"].([]interface{})[0].(map[string]interface{})
				curr_path := result["name"].(string)
				curr_id = int32(result["parent-folder"].(float64))
				if path == "" {
					path = curr_path
				} else {
					path = curr_path + "/" + path
				}
			}
		}
	}
	return path
}

func buildSecretPutBody(schema *schema.ResourceData, folder_id int32, sec_id int32) *fpam_go.Body {
	schema_fields := schema.Get("fields").([]interface{})
	var put_fields []fpam_go.CmdbsecretdatabaseidField

	for _, field := range schema_fields {
		put_field := fpam_go.CmdbsecretdatabaseidField{}
		field_name := field.(map[string]interface{})["name"].(string)
		field_value := field.(map[string]interface{})["value"].(string)

		put_field.Id = 0
		if field_value == "" || field_name == "" {
			log.Println("Skipping invalid field pair: %s:%s", field_name, field_value)
			continue
		}
		put_field.Name = field_name
		put_field.Value = field_value

		put_fields = append(put_fields, put_field)
	}
	put_body := &fpam_go.Body{
		Id:                    sec_id,
		Name:                  schema.Get("name").(string),
		Folder:                folder_id,
		Template:              schema.Get("template").(string),
		AvScan:                "disable",
		DlpStatus:             "disable",
		SshFilter:             "disable",
		RdpEventFilter:        "disable",
		AvProfile:             "",
		SshFilterProfile:      "",
		RdpEventFilterProfile: "",
		DlpSensor:             "",
		Field:                 put_fields,
	}

	return put_body
}

func buildSecretCreateBody(schema *schema.ResourceData, folder_id int32) *fpam_go.Body1 {
	schema_fields := schema.Get("fields").([]interface{})
	var post_fields []fpam_go.CmdbsecretdatabaseField

	for _, field := range schema_fields {
		post_field := fpam_go.CmdbsecretdatabaseField{}
		field_name := field.(map[string]interface{})["name"].(string)
		field_value := field.(map[string]interface{})["value"].(string)

		post_field.Id = 0
		if field_value == "" || field_name == "" {
			log.Println("Skipping invalid field pair: %s:%s", field_name, field_value)
			continue
		}
		post_field.Name = field_name
		post_field.Value = field_value

		post_fields = append(post_fields, post_field)
	}
	post_body := &fpam_go.Body1{
		Id:                0,
		Name:              schema.Get("name").(string),
		InheritPermission: "enable",
		Folder:            folder_id,
		Template:          schema.Get("template").(string),
		Field:             post_fields,
	}

	return post_body
}

func buildFolderPutBody(id int32, parent_id int32, folder_name string) *fpam_go.Body4 {

	post_body := &fpam_go.Body4{
		Id:                id,
		Name:              folder_name,
		ParentFolder:      parent_id,
		InheritPolicy:     "enable",
		InheritPermission: "enable",
		InheritFolderZtna: "enable",
		UserPermission:    []fpam_go.CmdbsecretfolderidUserpermission{},
		GroupPermission:   []fpam_go.CmdbsecretfolderidGrouppermission{},
	}
	return post_body
}

func buildFolderCreateBody(parent_id int32, folder_name string) *fpam_go.Body5 {

	post_body := &fpam_go.Body5{
		Id:                0,
		Name:              folder_name,
		ParentFolder:      parent_id,
		InheritPolicy:     "enable",
		InheritPermission: "enable",
		InheritFolderZtna: "enable",
		UserPermission:    []fpam_go.CmdbsecretfolderidUserpermission{},
		GroupPermission:   []fpam_go.CmdbsecretfolderidGrouppermission{},
	}
	return post_body
}
