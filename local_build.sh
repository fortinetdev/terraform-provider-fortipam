BINARY=terraform-provider-fortipam
LOCAL_PLUGIN_REGISTRY=~/.terraform.d/plugins/fortipam.com/local/fortipam/1.0.0/linux_amd64

rm -rf ~/.terraform.d/plugins/fortipam/
mkdir -p $LOCAL_PLUGIN_REGISTRY
go build -o $LOCAL_PLUGIN_REGISTRY/$BINARY