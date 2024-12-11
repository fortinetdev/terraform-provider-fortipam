

terraform {
  //This part is required when running local binary with local_build.sh
  required_providers {
    fortipam = {
      source  = "fortipam.com/local/fortipam"
      version = "1.0.0"
    }
  }
}

provider "fortipam" {
  base_url     = "<FPAM_URL>"
  access_token = "<FPAM_TOKEN>"
  verify_ssl   = false
}


data "fortipam_secret" "host" {
  path  = "tf_test"
  name  = "terraform_sec"
  field = "Host"
}

data "fortipam_secret" "usr" {
  path  = "tf_test"
  name  = "terraform_sec"
  field = "Username"
}

data "fortipam_secret" "pwd" {
  path  = "tf_test"
  name  = "terraform_sec"
  field = "Password"
}

locals {
  fields = [
    {
      "name"  = "Host",
      "value" = data.fortipam_secret.host.value
    },
    {
      "name"  = "Username",
      "value" = data.fortipam_secret.usr.value
    },
    {
      "name"  = "Password",
      "value" = data.fortipam_secret.pwd.value
    },
  ]
}

import {
  to = fortipam_resource_folder.folder
  id = 9
}
resource "fortipam_resource_folder" "folder" {
  parent = "tf_test"
  name   = "terraform"
}

/*
1. Can also use `terraform plan -generate-config-out=generated.tf` to generate the resource
import {
 to = fortipam_resource_secret.importSec
 id = "21"
}

2. CLI command: terraform import fortipam_resource_secret.sec <id>
resource "fortipam_resource_secret" "sec" {
}
*/

import {
  to = fortipam_resource_secret.secret
  id = 29
}
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

