resource "fortipam_resource_secret" "secret" {
  name     = "terraform_sec"
  path     = "${resource.fortipam_resource_folder.folder.parent}/${resource.fortipam_resource_folder.folder.name}"
  template = "Unix Account (SSH Password)"
  dynamic "fields" {
    for_each = local.fields
    content {
      name  = fields.value.name
      value = fields.value.value
    }
  }
}