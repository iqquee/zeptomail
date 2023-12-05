package zeptomail

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/pkg/errors"
)

// validate runs a validation on the incoming json payload
var validate = validator.New(validator.WithRequiredStructEnabled())

const (
	// MethodPOST for HTTP POST request method
	MethodPOST string = "POST"
	// MethodGET for HTTP GET request method
	MethodGET string = "GET"
)

// Client is an object for the configs
type Client struct {
	Http    *http.Client
	BaseUrl string
	Token   string
}

// New is the zeptomail config initializer
func New(h http.Client, token string) *Client {
	return &Client{
		BaseUrl: "https://api.zeptomail.com/v1.1/",
		Http:    &h,
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
		req.Header.Set("Authorization", fmt.Sprintf("Zoho-enczapikey %v", c.Token))
		fmt.Printf("This is the Token: %v\n", c.Token)
	}

	if err != nil {
		return errors.Wrap(err, "http client ::: unable to create request body")
	}

	res, err := c.Http.Do(req)
	if err != nil {
		return errors.Wrap(err, "http client ::: client failed to execute request")
	}

	defer res.Body.Close()

	if err := json.NewDecoder(res.Body).Decode(&resp); err != nil {
		return errors.Wrap(err, "http client ::: unable to unmarshal response body")
	}

	if err := c.error(res.StatusCode); err != nil {
		return err
	}

	return nil
}

// error checks for errors in respect to the status code and return accordingly
func (c *Client) error(statusCode int) error {
	code := fmt.Sprint(statusCode)
	if strings.HasPrefix(code, "4") {
		errors.Wrap(nil, "Failure === Kindly check your reqest as there is an error.")
	} else if strings.HasPrefix(code, "5") {
		errors.Wrap(nil, "Server error")
	}

	return nil
}
