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

// AddEmailTemplate is used to add an email template.
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

// UpdateEmailTemplate is used to update an email template.
func (c *Client) UpdateEmailTemplate(req UpdateEmailTemplateReq) (*AddEmailTemplateRes, error) {
	url := fmt.Sprintf("mailagents/%s/templates/%s", req.MailagentAlias, req.TemplateKey)
	method := MethodPUT

	if err := validate.Struct(&req); err != nil {
		return nil, err
	}

	var response AddEmailTemplateRes
	if err := c.newRequest(method, url, req, response); err != nil {
		return nil, err
	}

	return &response, nil
}

// GetEmailTemplate is used to fetch a particular email template.
func (c *Client) GetEmailTemplate(MailAgentAlias, TemplateKey string) (*GetEmailTemplateReq, error) {
	url := fmt.Sprintf("mailagents/%s/templates/%s", MailAgentAlias, TemplateKey)
	method := MethodGET

	var response GetEmailTemplateReq
	if err := c.newRequest(method, url, nil, response); err != nil {
		return nil, err
	}

	return &response, nil
}

// DeleteEmailTemplate is used to delete a template using template key.
func (c *Client) DeleteEmailTemplate(MailAgentAlias, TemplateKey string) (*interface{}, error) {
	url := fmt.Sprintf("mailagents/%s/templates/%s", MailAgentAlias, TemplateKey)
	method := MethodDELETE

	var response interface{}
	if err := c.newRequest(method, url, nil, response); err != nil {
		return nil, err
	}

	return &response, nil
}


func (c *Client) SendBatchHTMLEmail(req SendBatchHTMLEmailReq) (res *SendBatchHTMLEmailRes, err error) {
	url := "email/batch"
	Method := MethodPOST

	if err := validate.Struct(&req); err != nil {
		return nil, err
	}

	var response SendBatchHTMLEmailRes

	if err := c.newRequest(url, Method, req, response); err != nil {
		return nil, err
	}

	return &response, nil
}

