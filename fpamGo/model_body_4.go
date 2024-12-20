/*
 * FortiPAM
 *
 * REST API for configuring FortiPAM objects and settings
 *
 * API version: v1.1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package fpam_go

type Body4 struct {
	// Folder ID.
	Id int32 `json:"id"`
	// Folder name.
	Name string `json:"name"`
	// Enable/disable inherit policy from parent.    enable:Enable setting.    disable:Disable setting.
	InheritPolicy string `json:"inherit-policy,omitempty"`
	// Enable/disable inherit permission from parent.    enable:Enable setting.    disable:Disable setting.
	InheritPermission string `json:"inherit-permission"`
	// Enable/disable inherit ztna control from parent.    enable:Enable setting.    disable:Disable setting.
	InheritFolderZtna string `json:"inherit-folder-ztna,omitempty"`
	// Personal/Regular folder.    personal:Personal folder.    public:Public folder
	FolderType string `json:"folder-type,omitempty"`
	// Secret policy name.
	SecretPolicy string `json:"secret-policy,omitempty"`
	// ID of parent folder, Root == 0
	ParentFolder int32 `json:"parent-folder"`
	// Configure user permission.
	UserPermission []CmdbsecretfolderidUserpermission `json:"user-permission"`
	// Configure group permission.
	GroupPermission []CmdbsecretfolderidGrouppermission `json:"group-permission"`
}
