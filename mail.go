package zeptomail

import (
	"fmt"
	"net/http"
)

// SendHTMLEmail sends a HTML email
func (c *Client) SendHTMLEmail(req SendHTMLEmailReq) (*SendHTMLEmailRes, error) {
	url := "email"

	if err := validate.Struct(&req); err != nil {
		return nil, err
	}

	var response SendHTMLEmailRes
	if err := c.newRequest(http.MethodPost, url, req, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

// SendTemplatedEmail sends a templated email
func (c *Client) SendTemplatedEmail(req SendTemplatedEmailReq) (*SendTemplatedEmailRes, error) {
	url := "email/template"

	if err := validate.Struct(&req); err != nil {
		return nil, err
	}

	var response SendTemplatedEmailRes
	if err := c.newRequest(http.MethodPost, url, req, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

// SendBatchTemplatedEmail sends a batch templated email
func (c *Client) SendBatchTemplatedEmail(req SendBatchTemplatedEmailReq) (*SendTemplatedEmailRes, error) {
	url := "email/template/batch"

	if err := validate.Struct(&req); err != nil {
		return nil, err
	}

	var response SendTemplatedEmailRes
	if err := c.newRequest(http.MethodPost, url, req, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

// AddEmailTemplate is used to add an email template.
func (c *Client) AddEmailTemplate(req AddEmailTemplateReq) (*AddEmailTemplateRes, error) {
	url := fmt.Sprintf("mailagents/%s/templates", req.MailagentAlias)

	if err := validate.Struct(&req); err != nil {
		return nil, err
	}

	var response AddEmailTemplateRes
	if err := c.newRequest(http.MethodPost, url, req, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

// UpdateEmailTemplate is used to update an email template.
func (c *Client) UpdateEmailTemplate(req UpdateEmailTemplateReq) (*AddEmailTemplateRes, error) {
	url := fmt.Sprintf("mailagents/%s/templates/%s", req.MailagentAlias, req.TemplateKey)

	if err := validate.Struct(&req); err != nil {
		return nil, err
	}

	var response AddEmailTemplateRes
	if err := c.newRequest(http.MethodPut, url, req, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

// GetEmailTemplate is used to fetch a particular email template.
func (c *Client) GetEmailTemplate(MailAgentAlias, TemplateKey string) (*GetEmailTemplateRes, error) {
	url := fmt.Sprintf("mailagents/%s/templates/%s", MailAgentAlias, TemplateKey)

	var response GetEmailTemplateRes
	if err := c.newRequest(http.MethodGet, url, nil, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

// DeleteEmailTemplate is used to delete a template using template key.
func (c *Client) DeleteEmailTemplate(MailAgentAlias, TemplateKey string) (*interface{}, error) {
	url := fmt.Sprintf("mailagents/%s/templates/%s", MailAgentAlias, TemplateKey)

	var response interface{}
	if err := c.newRequest(http.MethodDelete, url, nil, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

// SendBatchHTMLEmail The API is used to send a batch of transactional HTML emails.
func (c *Client) SendBatchHTMLEmail(req SendBatchHTMLEmailReq) (res *SendBatchHTMLEmailRes, err error) {
	url := "email/batch"

	if err := validate.Struct(&req); err != nil {
		return nil, err
	}

	var response SendBatchHTMLEmailRes
	if err := c.newRequest(http.MethodPost, url, req, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

// FileCacheUploadAPI The API is used to upload files to File Cache
func (c *Client) FileCacheUploadAPI(req FileCacheUploadAPIReq) (res *FileCacheUploadAPIRes, err error) {
	url := fmt.Sprintf("%sfiles?name=%s", c.BaseUrl, req.FileName)

	if err := validate.Struct(&req); err != nil {
		return nil, err
	}

	var response FileCacheUploadAPIRes
	if err := c.newRequest(http.MethodPost, url, req, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

// ListEmailTemplates You can use this API to list the required number of email templates in your ZeptoMail account.
func (c *Client) ListEmailTemplates(MailAgentAlias string, Offset, limit int) (req *ListEmailTemplatesRes, err error) {
	url := fmt.Sprintf("mailagents/%s/templates?offset=%d&limit=%d", MailAgentAlias, Offset, limit)

	if err := validate.Struct(&req); err != nil {
		return nil, err
	}

	var response ListEmailTemplatesRes
	if err := c.newRequest(http.MethodGet, url, nil, &response); err != nil {
		return nil, err
	}

	return &response, nil
}
