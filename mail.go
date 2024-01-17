package zeptomail

import "fmt"

// TODO create an object for the error response from the server
// SendHTMLEmail sends a HTML template email
func (c *Client) SendHTMLEmail(req SendHTMLEmailReq) (*SendHTMLEmailRes, error) {
	url := "email"
	method := MethodPOST

	if err := validate.Struct(&req); err != nil {
		return nil, err
	}

	var response SendHTMLEmailRes
	if err := c.newRequest(method, url, req, response); err != nil {
		return nil, err
	}

	return &response, nil
}

// SendTemplatedEmail sends a templated email
func (c *Client) SendTemplatedEmail(req SendTemplatedEmailReq) (*SendTemplatedEmailRes, error) {
	url := "email/template"
	method := MethodPOST

	if err := validate.Struct(&req); err != nil {
		return nil, err
	}

	var response SendTemplatedEmailRes
	if err := c.newRequest(method, url, req, response); err != nil {
		return nil, err
	}

	return &response, nil
}

// SendBatchTemplatedEmail sends a batch templated email
func (c *Client) SendBatchTemplatedEmail(req SendBatchTemplatedEmailReq) (*SendTemplatedEmailRes, error) {
	url := "email/template/batch"
	method := MethodPOST

	if err := validate.Struct(&req); err != nil {
		return nil, err
	}

	var response SendTemplatedEmailRes
	if err := c.newRequest(method, url, req, response); err != nil {
		return nil, err
	}

	return &response, nil
}

// AddEmailTemplate can be used to add an email template.
func (c *Client) AddEmailTemplate(req AddEmailTemplateReq) (*AddEmailTemplateRes, error) {
	url := fmt.Sprintf("mailagents/%s/templates", req.MailagentAlias)
	method := MethodPOST

	if err := validate.Struct(&req); err != nil {
		return nil, err
	}

	var response AddEmailTemplateRes
	if err := c.newRequest(method, url, req, response); err != nil {
		return nil, err
	}

	return &response, nil
}
