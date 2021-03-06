// Code generated by "tfplugingen -gen provider -type provider"; DO NOT EDIT.

package provider

import (
	"fmt"
	terraformpluginsdk "github.com/hashicorp/terraform-plugin-sdk"
	errors "github.com/pkg/errors"
	cty "github.com/zclconf/go-cty/cty"
	gocty "github.com/zclconf/go-cty/cty/gocty"
)

func (r *provider) Schema() terraformpluginsdk.Schema {
	return terraformpluginsdk.Schema{Block: terraformpluginsdk.Block{Attributes: []terraformpluginsdk.Attribute{terraformpluginsdk.Attribute{
		Computed:  false,
		ForceNew:  false,
		Name:      "request_headers",
		Optional:  true,
		Required:  false,
		Sensitive: false,
		Type:      cty.Map(cty.String),
	}}}}
}
func (r *provider) UnmarshalState(conf cty.Value) error {
	var err error
	_ = err
	if !conf.IsNull() && conf.IsKnown() {
		if !conf.GetAttr("request_headers").IsNull() && conf.GetAttr("request_headers").IsKnown() {
			err = gocty.FromCtyValue(conf.GetAttr("request_headers"), &r.RequestHeaders)
			if err != nil {
				return errors.WithStack(err)
			}
		}
	}
	return nil
}
func (r *provider) MarshalState() (cty.Value, error) {
	var err error
	_ = err
	var state cty.Value
	{
		state1 := map[string]cty.Value{}
		{
			state1["request_headers"], err = gocty.ToCtyValue(r.RequestHeaders, cty.Map(cty.String))
			if err != nil {
				return cty.NilVal, errors.WithStack(err)
			}
		}
		state = cty.ObjectVal(state1)
	}
	return state, nil
}
func (p *provider) DataSourceFactory(typeName string) terraformpluginsdk.DataSource {
	switch typeName {
	case "http":
		return &dataHTTP{provider: p}
	}
	panic(fmt.Sprintf("datasource %s unexpected", typeName))
}
func (p *provider) DataSourceSchemas() map[string]terraformpluginsdk.Schema {
	return map[string]terraformpluginsdk.Schema{"http": p.DataSourceFactory("http").Schema()}
}
func (p *provider) ResourceFactory(typeName string) terraformpluginsdk.Resource {
	panic(fmt.Sprintf("resource %s unexpected", typeName))
}
func (p *provider) ResourceSchemas() map[string]terraformpluginsdk.Schema {
	return map[string]terraformpluginsdk.Schema{}
}
