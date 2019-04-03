package main

import (
	sdk "github.com/hashicorp/terraform-plugin-sdk"

	"github.com/terraform-providers/terraform-provider-http/provider"
)

func main() {
	sdk.ServeProvider(provider.New())
}
