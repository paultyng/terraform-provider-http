module github.com/terraform-providers/terraform-provider-http

require (
	github.com/hashicorp/terraform v0.12.0-alpha4.0.20190423184126-64f419ac76f3
	github.com/hashicorp/terraform-plugin-sdk v1.0.0
	github.com/pkg/errors v0.8.1
	github.com/zclconf/go-cty v0.0.0-20190430221426-d36a6f0dbffd
)

replace (
	github.com/hashicorp/terraform-plugin-sdk => ../../hashicorp/terraform-plugin-sdk
	github.com/zclconf/go-cty => ../../zclconf/go-cty
)
