/*
 * FortiPAM
 *
 * REST API for configuring FortiPAM objects and settings
 *
 * API version: v1.1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package fpam_go

type CmdbsecretdatabaseidGrouppermission struct {
	// Permision ID.
	Id int32 `json:"id,omitempty"`
	// Group name.
	GroupName []CmdbsecretdatabaseidGroupname `json:"group-name,omitempty"`
	// Secret permission option.    none:No Permissions.    list:List permission.    view:View permission.    edit:Edit permission.    owner:Owner permission.
	Permission string `json:"permission,omitempty"`
}
