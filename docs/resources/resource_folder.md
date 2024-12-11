---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "fortipam_resource_folder Resource - fortipam"
subcategory: ""
description: |-
  
---

# fortipam_resource_folder (Resource)



## Example Usage

```terraform
resource "fortipam_resource_folder" "folder" {
  parent = "tf_test"
  name   = "terraform"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) folder name under parent folder
- `parent` (String) full folder path to the parent folder	e.g. public_folder/local

### Read-Only

- `id` (String) The ID of this resource.
- `path` (String) full path to the folder

## Import

Import is supported using the following syntax:

```shell
terraform import fortipam_resource_folder.example 123
```