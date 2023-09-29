package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/iqquee/zeptomail"
	"github.com/pkg/errors"
)

func main() {
	zeptomailToken := "wSsVR60lqx/wDat0zTKqc7s5mFtRVl2jQ0183wTw7iX+HqiR98c7nhXIAFOuSqJLQjFhHDoX8uh/mE1U0DRciNgtw14CXCiF9mqRe1U4J3x17qnvhDzNWmpakBOJLowAxg5inmhlFM8h+g=="
	// tempKey := "2d6f.51f8ea21ac70a66e.k1.349b6820-5cb5-11ee-bf8e-5254004d4100.18ad36f07a2"
	client := zeptomail.New(*http.DefaultClient, zeptomailToken)
	sendTo := []zeptomail.SendEmailTo{}
	sendTo = append(sendTo, zeptomail.SendEmailTo{
		EmailAddress: zeptomail.EmailAddress{
			Address: "hisyntax01@gmail.com",
			Name:    "matty",
		},
	})

	data := "<div>Welcome to Xnd, <b> Kindly click on Verify[https://github.com/iqquee] Account to confirm your account </b></div>"
	req := zeptomail.SendHTMLEmailReq{
		To: sendTo,
		From: zeptomail.EmailAddress{
			Address: "noreply@xnexd.io",
			Name:    "Xnd",
		},
		Subject:  "Welcome to Xnd",
		Htmlbody: data,
	}
	res, err := client.SendHTMLEmail(req)
	if err != nil {
		fmt.Printf("This is the error: %v", err.Error())
	}

	for _, e := range res.Data {
		fmt.Printf("response message: %v\n", e.Message)
	}

	// sendMail()
}

func sendMail() error {
	client := &http.Client{}
	token := "wSsVR60lqx/wDat0zTKqc7s5mFtRVl2jQ0183wTw7iX+HqiR98c7nhXIAFOuSqJLQjFhHDoX8uh/mE1U0DRciNgtw14CXCiF9mqRe1U4J3x17qnvhDzNWmpakBOJLowAxg5inmhlFM8h+g=="
	payload := []byte(`{
	"from": {"address": "noreply@xnexd.io"},
	"to": [{"email_address": {"address": "hisyntax0@gmail.com","name": "Purple"}}],
	"subject":"Welcome to xnd",
	"htmlbody":"<div><b> Test email sent successfully from Xnd. </b></div>"}`)

	req, err := http.NewRequest("POST", "https://api.zeptomail.com/v1.1/email", bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Zoho-enczapikey %v", token))

	if err != nil {
		return errors.Wrap(err, "http client ::: unable to create request body")
	}

	res, err := client.Do(req)
	if err != nil {
		return errors.Wrap(err, "http client ::: client failed to execute request")
	}

	defer res.Body.Close()

	var resp zeptomail.SendHTMLEmailRes

	if err := json.NewDecoder(res.Body).Decode(&resp); err != nil {
		return errors.Wrap(err, "http client ::: unable to unmarshal response body")
	}

	fmt.Printf("This is the status code: %v\n", res.Status)
	fmt.Printf("This is the response message: %v\n", resp.Message)
	fmt.Printf("This is the response requestID: %v\n", resp.RequestId)
	return nil
}
