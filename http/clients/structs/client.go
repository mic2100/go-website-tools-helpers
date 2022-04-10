package HttpClientStructs

import (
	"errors"
	"fmt"
	HttpClientHelpers "github.com/mic2100/go-website-tools-helpers/http/clients/helpers"
	"net/http"
)

type Client struct {
	BaseUrl string
	Client  *http.Client
	Headers []Header
}

type Header struct {
	Key   string
	Value string
}

func (c *Client) AddHeader(header Header) *Client {
	c.Headers = append(c.Headers, header)

	return c
}

func (c *Client) Get(uri string, dataType struct{}) (struct{}, error) {
	req, err := HttpClientHelpers.NewRequest("GET", c.BaseUrl, uri)
	if err != nil {
		return dataType, errors.New(err.Error())
	}

	HttpClientHelpers.ApplyHeaders(req, c)

	resp, err := c.Client.Do(req)

	if err != nil {
		return dataType, errors.New(err.Error())
	}

	defer resp.Body.Close()

	return HttpClientHelpers.GetResponseStruct(resp, dataType)
}

func (c *Client) Post(url string, data struct{}, dataType struct{}) struct{} {
	fmt.Println(url)
	fmt.Println(dataType)

	return dataType
}

func (c *Client) Put(url string, dataType struct{}) struct{} {
	fmt.Println(url)
	fmt.Println(dataType)

	return dataType
}

func (c *Client) Delete(url string, dataType struct{}) struct{} {
	fmt.Println(url)
	fmt.Println(dataType)

	return dataType
}
