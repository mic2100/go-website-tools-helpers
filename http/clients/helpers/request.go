package HttpClientHelpers

import (
	"fmt"
	HttpClientStructs "github.com/mic2100/go-website-tools-helpers/http/clients/structs"
	"net/http"
)

func NewRequest(method string, baseUrl string, uri string) (*http.Request, error) {
	return http.NewRequest(method, fmt.Sprintf("%s/%s", baseUrl, uri), nil)
}

func ApplyHeaders(req *http.Request, c *HttpClientStructs.Client) {
	for _, h := range c.Headers {
		req.Header.Set(h.Key, h.Value)
	}
}
