# zeptomail
A golang sdk for interacting with zeptomail api 

# Get Started
Inother to use this package, you need to first head over to https://zoho.com to create an account to get your `Authorization Token Key`.
# Documentation
See the [ZeptoMail API Docs](https://www.zoho.com/zeptomail/help/introduction.html)

# Installation
This package can in installed using the go command below.
```sh
go get github.com/iqquee/zeptomail
```
# Quick start
```sh
# assume the following codes in example.go file
$ touch example.go
# open the just created example.go file in the text editor of your choice
```

# SendHTMLEmail
`SendHTMLEmail()` sends a HTML template email

This method takes in the `zeptomail.SendHTMLEmailReq{}` struct as a parameter.
### List of all the fields available in the zeptomail.SendHTMLEmailReq{} struct.
```go
type ( 
    // SendHTMLEmailReq is the SendHTMLEmail() request object
    SendHTMLEmailReq struct {
		From     EmailAddress  `json:"from" validate:"required"`
		To       []SendEmailTo `json:"to" validate:"required"`
		Subject  string        `json:"subject" validate:"required"`
		Htmlbody string        `json:"htmlbody" validate:"required"`
	}

    // EmailAddress is an email address object
    EmailAddress struct {
		/*
			The email address to which the recipient's email responses will be addressed.

			Allowed value - A valid email address containing a domain that is verified in your Mail Agent.
		*/
		Address string `json:"address" validate:"required"`
		// Recipient's name.
		Name string `json:"name" validate:"required"`
	}
)
```

```go
package main

import (
	"fmt"
	"net/http"

	"github.com/iqquee/zeptomail"
)

func main() {
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
```
