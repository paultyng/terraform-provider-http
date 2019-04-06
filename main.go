package main

import (
	sdk "github.com/hashicorp/terraform-plugin-sdk"

	"github.com/terraform-providers/terraform-provider-http/internal/provider"
)

func main() {
	sdk.ServeProvider(provider.New())
}
