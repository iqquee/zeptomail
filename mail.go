package zeptomail

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
