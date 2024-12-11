/*
 * FortiPAM
 *
 * REST API for configuring FortiPAM objects and settings
 *
 * API version: v1.1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package fpam_go

type CmdbsecretfolderidUserpermission struct {
	// Permission ID.
	Id int32 `json:"id,omitempty"`
	// User name.
	UserName []CmdbsecretdatabaseidUsername `json:"user-name,omitempty"`
	// Folder permission option.    none:No permission.    view:View permission.    add-secret:Add secret permission.    edit:Edit permission.    owner:Owner permission.
	FolderPermission string `json:"folder-permission,omitempty"`
	// Secret permission option.    none:No Permissions.    list:List permission.    view:View permission.    edit:Edit permission.    owner:Owner permission.
	SecretPermission string `json:"secret-permission,omitempty"`
}
