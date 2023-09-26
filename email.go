package zeptomail

// SendEmail sends a HTML template email
func SendEmail() {

}

// SendTemplatedEmail sends a templated email
func (c *Client) SendTemplatedEmail(req SendTemplatedEmailReq) (interface{}, error) {
	url := "email/template"
	method := MethodPOST

	if err := validate.Struct(&req); err != nil {
		return nil, err
	}

	var response interface{}
	if err := c.newRequest(method, url, req, response); err != nil {
		return nil, err
	}

	return &response, nil
}
