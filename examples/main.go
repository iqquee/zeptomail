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

	// SendBatchTemplatedEmail sends a batch templated email
	SendBatchTemplatedEmail()

	// AddEmailTemplate can be used to add an email template.
	AddEmailTemplate()

	// UpdateEmailTemplate can be used to update an email template.
	UpdateEmailTemplate()

	// GetEmailTemplate is used to fetch a particular email template.
	GetEmailTemplate()
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

// AddEmailTemplate can be used to add an email template.
func AddEmailTemplate() {
	zeptomailToken := "your zeptomail authorization token"

	client := zeptomail.New(*http.DefaultClient, zeptomailToken)

	req := zeptomail.AddEmailTemplateReq{
		TemplateName:   "E-invite",
		Subject:        "Invitation to the event",
		HtmlBody:       "<h1> Hi {{name}}</h1>, invitation link {{link}}", // for a html template
		TemplateAlias:  "e-invite emails",
		MailagentAlias: "mail agent alias",
	}

	res, err := client.AddEmailTemplate(req)
	if err != nil {
		fmt.Printf("This is the error: %v", err.Error())
	}

	fmt.Printf("response message: %v\n", res.Message)
}

// UpdateEmailTemplate can be used to update an email template.
func UpdateEmailTemplate() {
	zeptomailToken := "your zeptomail authorization token"

	client := zeptomail.New(*http.DefaultClient, zeptomailToken)

	req := zeptomail.UpdateEmailTemplateReq{
		TemplateName:   "E-invite",
		Subject:        "Invitation to the event",
		HtmlBody:       "<h1> Hi {{name}}</h1>, invitation link {{link}}", // for a html template
		TemplateKey:    "your template key",
		MailagentAlias: "mail agent alias",
	}

	res, err := client.UpdateEmailTemplate(req)
	if err != nil {
		fmt.Printf("This is the error: %v", err.Error())
	}

	fmt.Printf("response message: %v\n", res.Message)
}

// GetEmailTemplate is used to fetch a particular email template.
func GetEmailTemplate() {
	zeptomailToken := "your zeptomail authorization token"

	client := zeptomail.New(*http.DefaultClient, zeptomailToken)

	var TemplateKey, MailagentAlias string = "your template key", "mail agent alias"

	res, err := client.GetEmailTemplate(MailagentAlias, TemplateKey)
	if err != nil {
		fmt.Printf("This is the error: %v", err.Error())
	}

	fmt.Printf("response message: %v\n", res.Message)
}
