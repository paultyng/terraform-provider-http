package provider

import (
	"context"
	"fmt"
	"io/ioutil"
	"mime"
	"net/http"
	"net/url"
	"regexp"
	"strings"

	sdk "github.com/hashicorp/terraform-plugin-sdk"
	cty "github.com/zclconf/go-cty/cty"
)

type urlAttribute string

func (u urlAttribute) Validate() error {
	_, err := url.ParseRequestURI(string(u))
	return err
}

//go:generate tfplugingen -gen datasource -type dataHTTP
type dataHTTP struct {
	provider *provider
	TFName   struct{} `tf:"http,datasource"`

	URL            urlAttribute      `tf:"url,required"`
	RequestHeaders map[string]string `tf:"request_headers,optional"`
	Body           string            `tf:"body,computed"`
}

func (d *dataHTTP) Read(ctx context.Context) error {
	client := d.provider.NewClient()

	req, err := http.NewRequest("GET", string(d.URL), nil)
	if err != nil {
		return fmt.Errorf("Error creating request: %s", err)
	}

	req = req.WithContext(ctx)

	for n, v := range d.provider.RequestHeaders {
		req.Header.Set(n, v)
	}

	for n, v := range d.RequestHeaders {
		req.Header.Set(n, v)
	}

	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("Error during making a request: %s", d.URL)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return fmt.Errorf("HTTP request error. Response code: %d", resp.StatusCode)
	}

	contentType := resp.Header.Get("Content-Type")
	if contentType == "" || isContentTypeAllowed(contentType) == false {
		return fmt.Errorf("Content-Type is not a text type. Got: %s", contentType)
	}

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("Error while reading response body. %s", err)
	}

	d.Body = string(bytes)

	return nil
}

// This is to prevent potential issues w/ binary files
// and generally unprintable characters
// See https://github.com/hashicorp/terraform/pull/3858#issuecomment-156856738
func isContentTypeAllowed(contentType string) bool {
	parsedType, params, err := mime.ParseMediaType(contentType)
	if err != nil {
		return false
	}

	allowedContentTypes := []*regexp.Regexp{
		regexp.MustCompile("^text/.+"),
		regexp.MustCompile("^application/json$"),
		regexp.MustCompile("^application/samlmetadata\\+xml"),
	}

	for _, r := range allowedContentTypes {
		if r.MatchString(parsedType) {
			charset := strings.ToLower(params["charset"])
			return charset == "" || charset == "utf-8"
		}
	}

	return false
}

// TODO: this could be generated...

func (r *dataHTTP) Validate() ([]sdk.Diagnostic, error) {
	diags := []sdk.Diagnostic{}
	var err error

	err = r.URL.Validate()
	if err != nil {
		diags = append(diags, sdk.Diagnostic{
			Path:     cty.Path{}.GetAttr("url"),
			Severity: sdk.SeverityError,
			Summary:  err.Error(),
		})
	}

	return diags, nil
}
