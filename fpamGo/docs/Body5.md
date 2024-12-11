# Body5

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | Folder ID. | [default to null]
**Name** | **string** | Folder name. | [default to null]
**InheritPolicy** | **string** | Enable/disable inherit policy from parent.    enable:Enable setting.    disable:Disable setting. | [optional] [default to null]
**InheritPermission** | **string** | Enable/disable inherit permission from parent.    enable:Enable setting.    disable:Disable setting. | [default to null]
**InheritFolderZtna** | **string** | Enable/disable inherit ztna control from parent.    enable:Enable setting.    disable:Disable setting. | [optional] [default to null]
**FolderType** | **string** | Personal/Regular folder.    personal:Personal folder.    public:Public folder | [optional] [default to null]
**SecretPolicy** | **string** | Secret policy name. | [optional] [default to null]
**ParentFolder** | **int32** | ID of parent folder, Root &#x3D;&#x3D; 0 | [default to null]
**UserPermission** | [**[]CmdbsecretfolderidUserpermission**](cmdbsecretfolderid_userpermission.md) | Configure user permission. | [default to null]
**GroupPermission** | [**[]CmdbsecretfolderidGrouppermission**](cmdbsecretfolderid_grouppermission.md) | Configure group permission. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


