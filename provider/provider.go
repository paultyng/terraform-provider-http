package provider

import (
	"net/http"

	sdk "github.com/hashicorp/terraform-plugin-sdk"
)

type provider struct {
	// TODO: global RequestHeaders map[string]string `tf:"request_headers,optional"`
}

func (p *provider) NewClient() *http.Client {
	cli := &http.Client{}
	return cli
}

func New() sdk.Provider {
	return &provider{}
}
