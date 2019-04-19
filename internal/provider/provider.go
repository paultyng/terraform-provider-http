package provider

import (
	"context"
	"net/http"

	"github.com/hashicorp/terraform/helper/logging"

	sdk "github.com/hashicorp/terraform-plugin-sdk"
)

//go:generate tfplugingen -gen provider -type provider
type provider struct {
	RequestHeaders map[string]string `tf:"request_headers,optional"`
}

func (p *provider) NewClient() *http.Client {
	cli := &http.Client{
		Transport: logging.NewTransport("HTTP", http.DefaultTransport),
	}
	return cli
}

func New() sdk.Provider {
	return &provider{}
}

func (p *provider) Configure(ctx context.Context, tfVersion string) error {
	//nothing to do here
	return nil
}

func (p *provider) Stop(ctx context.Context) error {
	//nothing to do here
	return nil
}
