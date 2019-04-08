package provider

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/plugintest"
	"github.com/hashicorp/terraform/states"
)

type TestHTTPMock struct {
	server *httptest.Server
}

const testDataSourceConfig_basic = `
data "http" "http_test" {
  url = "%s/meta_%d.txt"
}

output "body" {
  value = "${data.http.http_test.body}"
}
`

func TestDataSource_http200(t *testing.T) {
	testHTTPMock := setUpMockHTTPServer()

	defer testHTTPMock.server.Close()

	plugintest.UnitTest(t, plugintest.TestCase{
		Providers: testProviders,
		Steps: []plugintest.TestStep{
			{
				Config: fmt.Sprintf(testDataSourceConfig_basic, testHTTPMock.server.URL, 200),
				Check: func(s *states.State) error {
					_, ok := s.RootModule().Resources["data.http.http_test"]
					if !ok {
						return fmt.Errorf("missing data resource")
					}

					outputs := s.RootModule().OutputValues

					if outputs["body"].Value.AsString() != "1.0.0" {
						return fmt.Errorf(
							`'body' output is %s; want '1.0.0'`,
							outputs["body"].Value,
						)
					}

					return nil
				},
			},
		},
	})
}

func TestDataSource_http404(t *testing.T) {
	testHTTPMock := setUpMockHTTPServer()

	defer testHTTPMock.server.Close()

	plugintest.UnitTest(t, plugintest.TestCase{
		Providers: testProviders,
		Steps: []plugintest.TestStep{
			{
				Config:      fmt.Sprintf(testDataSourceConfig_basic, testHTTPMock.server.URL, 404),
				ExpectError: regexp.MustCompile("HTTP request error. Response code: 404"),
			},
		},
	})
}

const testDataSourceConfig_withHeaders = `
data "http" "http_test" {
  url = "%s/restricted/meta_%d.txt"

  request_headers = {
    "Authorization" = "Zm9vOmJhcg=="
  }
}

output "body" {
  value = data.http.http_test.body
}
`

func TestDataSource_withHeaders200(t *testing.T) {
	testHTTPMock := setUpMockHTTPServer()

	defer testHTTPMock.server.Close()

	plugintest.UnitTest(t, plugintest.TestCase{
		Providers: testProviders,
		Steps: []plugintest.TestStep{
			{
				Config: fmt.Sprintf(testDataSourceConfig_withHeaders, testHTTPMock.server.URL, 200),
				Check: func(s *states.State) error {
					_, ok := s.RootModule().Resources["data.http.http_test"]
					if !ok {
						return fmt.Errorf("missing data resource")
					}

					outputs := s.RootModule().OutputValues

					if outputs["body"].Value.AsString() != "1.0.0" {
						return fmt.Errorf(
							`'body' output is %s; want '1.0.0'`,
							outputs["body"].Value,
						)
					}

					return nil
				},
			},
		},
	})
}

const testDataSourceConfig_utf8 = `
data "http" "http_test" {
  url = "%s/utf-8/meta_%d.txt"
}

output "body" {
  value = data.http.http_test.body
}
`

func TestDataSource_utf8(t *testing.T) {
	testHTTPMock := setUpMockHTTPServer()

	defer testHTTPMock.server.Close()

	plugintest.UnitTest(t, plugintest.TestCase{
		Providers: testProviders,
		Steps: []plugintest.TestStep{
			{
				Config: fmt.Sprintf(testDataSourceConfig_utf8, testHTTPMock.server.URL, 200),
				Check: func(s *states.State) error {
					_, ok := s.RootModule().Resources["data.http.http_test"]
					if !ok {
						return fmt.Errorf("missing data resource")
					}

					outputs := s.RootModule().OutputValues

					if outputs["body"].Value.AsString() != "1.0.0" {
						return fmt.Errorf(
							`'body' output is %s; want '1.0.0'`,
							outputs["body"].Value,
						)
					}

					return nil
				},
			},
		},
	})
}

const testDataSourceConfig_utf16 = `
data "http" "http_test" {
  url = "%s/utf-16/meta_%d.txt"
}

output "body" {
  value = data.http.http_test.body
}
`

func TestDataSource_utf16(t *testing.T) {
	testHTTPMock := setUpMockHTTPServer()

	defer testHTTPMock.server.Close()

	plugintest.UnitTest(t, plugintest.TestCase{
		Providers: testProviders,
		Steps: []plugintest.TestStep{
			{
				Config:      fmt.Sprintf(testDataSourceConfig_utf16, testHTTPMock.server.URL, 200),
				ExpectError: regexp.MustCompile("Content-Type is not a text type. Got: application/json; charset=UTF-16"),
			},
		},
	})
}

const testDataSourceConfig_error = `
data "http" "http_test" {

}
`

func TestDataSource_compileError(t *testing.T) {
	plugintest.UnitTest(t, plugintest.TestCase{
		Providers: testProviders,
		Steps: []plugintest.TestStep{
			{
				Config:      testDataSourceConfig_error,
				ExpectError: regexp.MustCompile("The argument \"url\" is required, but no definition was found."),
			},
		},
	})
}

func setUpMockHTTPServer() *TestHTTPMock {
	Server := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			w.Header().Set("Content-Type", "text/plain")
			if r.URL.Path == "/meta_200.txt" {
				w.WriteHeader(http.StatusOK)
				w.Write([]byte("1.0.0"))
			} else if r.URL.Path == "/restricted/meta_200.txt" {
				if r.Header.Get("Authorization") == "Zm9vOmJhcg==" {
					w.WriteHeader(http.StatusOK)
					w.Write([]byte("1.0.0"))
				} else {
					w.WriteHeader(http.StatusForbidden)
				}
			} else if r.URL.Path == "/utf-8/meta_200.txt" {
				w.Header().Set("Content-Type", "text/plain; charset=UTF-8")
				w.WriteHeader(http.StatusOK)
				w.Write([]byte("1.0.0"))
			} else if r.URL.Path == "/utf-16/meta_200.txt" {
				w.Header().Set("Content-Type", "application/json; charset=UTF-16")
				w.WriteHeader(http.StatusOK)
				w.Write([]byte("\"1.0.0\""))
			} else if r.URL.Path == "/meta_404.txt" {
				w.WriteHeader(http.StatusNotFound)
			} else {
				w.WriteHeader(http.StatusNotFound)
			}
		}),
	)

	return &TestHTTPMock{
		server: Server,
	}
}
