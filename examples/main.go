package main

import (
	"fmt"
	"net/http"

	"github.com/iqquee/zeptomail"
)

func main() {
	// sendHTMLEmail sends a HTML Email
	sendHTMLEmail()

	// sendHTMLEmail sends a templated email
	sendTemplatedEmail()
}

// sendHTMLEmail sends a HTML Email
func sendHTMLEmail() {
	zeptomailToken := "your zeptomail authorization token"

	client := zeptomail.New(*http.DefaultClient, zeptomailToken)
	sendTo := []zeptomail.SendEmailTo{}
	sendTo = append(sendTo, zeptomail.SendEmailTo{
		EmailAddress: zeptomail.EmailAddress{
			Address: "rudra.d@zylker.com",
			Name:    "Rudra",
		},
	})

	data := "<div><b> Kindly click on Verify Account to confirm your account </b></div>"

	req := zeptomail.SendHTMLEmailReq{
		To: sendTo,
		From: zeptomail.EmailAddress{
			Address: "accounts@info.zylker.com",
			Name:    "Paula",
		},
		Subject:  "Account Confirmation",
		Htmlbody: data,
	}
	res, err := client.SendHTMLEmail(req)
	if err != nil {
		fmt.Printf("This is the error: %v", err.Error())
	}

	for _, e := range res.Data {
		fmt.Printf("response message: %v\n", e.Message)
	}
}

// sendHTMLEmail sends a templated email
func sendTemplatedEmail() {
	zeptomailToken := "your zeptomail authorization token"
	tempKey := "your zeptomail template key"

	client := zeptomail.New(*http.DefaultClient, zeptomailToken)
	sendTo := []zeptomail.SendEmailTo{}
	sendTo = append(sendTo, zeptomail.SendEmailTo{
		EmailAddress: zeptomail.EmailAddress{
			Address: "rudra.d@zylker.com",
			Name:    "Rudra",
		},
	})

	req := zeptomail.SendTemplatedEmailReq{
		To: sendTo,
		From: zeptomail.EmailAddress{
			Address: "accounts@info.zylker.com",
			Name:    "Paula",
		},
		TemplateKey: tempKey,
	}

	res, err := client.SendTemplatedEmail(req)
	if err != nil {
		fmt.Printf("This is the error: %v", err.Error())
	}

	for _, e := range res.Data {
		fmt.Printf("response message: %v\n", e.Message)
	}
}

// SendBatchTemplatedEmail sends a batch templated email
func SendBatchTemplatedEmail() {
	zeptomailToken := "your zeptomail authorization token"
	tempKey := "your zeptomail template key"

	client := zeptomail.New(*http.DefaultClient, zeptomailToken)
	sendTo := []zeptomail.SendBatchTemplateEmailTo{}
	sendTo = append(sendTo, zeptomail.SendBatchTemplateEmailTo{
		EmailAddress: zeptomail.EmailAddress{
			Address: "rudra.d@zylker.com",
			Name:    "Rudra",
		},
		MergeInfo: map[string]interface{}{
			"userID": "my userID",
		},
	})

	req := zeptomail.SendBatchTemplatedEmailReq{
		To: sendTo,
		From: zeptomail.EmailAddress{
			Address: "accounts@info.zylker.com",
			Name:    "Paula",
		},
		TemplateKey: tempKey,
	}

	res, err := client.SendBatchTemplatedEmail(req)
	if err != nil {
		fmt.Printf("This is the error: %v", err.Error())
	}

	for _, e := range res.Data {
		fmt.Printf("response message: %v\n", e.Message)
	}
}
