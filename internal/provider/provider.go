package provider

import (
	"fmt"
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

// TODO: could be generated below here?

func (p *provider) DataSourceFactory(typeName string) sdk.DataSource {
	switch typeName {
	case "http":
		return &dataHTTP{
			provider: p,
		}
	}
	panic(fmt.Sprintf("data source %s unexpected", typeName))
}

func (p *provider) ResourceFactory(typeName string) sdk.Resource {
	panic(fmt.Sprintf("resource %s unexpected", typeName))
}

func (p *provider) ResourceSchemas() map[string]sdk.Schema {
	return map[string]sdk.Schema{}
}

func (p *provider) DataSourceSchemas() map[string]sdk.Schema {
	return map[string]sdk.Schema{
		"http": p.DataSourceFactory("http").Schema(),
	}
}
