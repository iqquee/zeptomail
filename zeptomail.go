package zeptomail

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/pkg/errors"
)

// Client is an object for the configs
type Client struct {
	Http    http.Client
	BaseUrl string
	Token   string
}

// New is the zeptomail config initializer
func New(h http.Client, token string) *Client {
	return &Client{
		BaseUrl: "https://api.zeptomail.com/v1.1/",
		Http:    h,
		Token:   token,
	}
}

// newRequest makes a http request to the zeptomail server and decodes the server response into the reqBody parameter passed into the newRequest method
func (c *Client) newRequest(method, reqURL string, reqBody interface{}, resp interface{}) error {
	newURL := c.BaseUrl + reqURL
	var body io.Reader

	if reqBody != nil {
		bb, err := json.Marshal(reqBody)
		if err != nil {
			return errors.Wrap(err, "http client ::: unable to marshal request struct")
		}
		body = bytes.NewReader(bb)
	}

	req, err := http.NewRequest(method, newURL, body)
	if reqBody != nil {
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("zoho-enczapikey %v", c.Token))
	}

	if err != nil {
		return errors.Wrap(err, "http client ::: unable to create request body")
	}

	res, err := c.Http.Do(req)
	if err != nil {
		return errors.Wrap(err, "http client ::: client failed to execute request")
	}

	if err := json.NewDecoder(res.Body).Decode(&resp); err != nil {
		return errors.Wrap(err, "http client ::: unable to unmarshal response body")
	}

	return nil
}
