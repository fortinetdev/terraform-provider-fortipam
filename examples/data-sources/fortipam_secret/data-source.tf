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
