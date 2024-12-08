# zeptomail
A golang sdk for interacting with zeptomail api 

# Get Started
In order to use this package, you need to first head over to https://zoho.com to create an account to get your `Authorization Token Key`.
# Documentation
See the [ZeptoMail API Docs](https://www.zoho.com/zeptomail/help/introduction.html)

# Installation
This package can in installed using the go command below.
```sh
go get github.com/iqquee/zeptomail@latest
```
# Quick start
```sh
# assume the following codes in example.go file
$ touch example.go
# open the just created example.go file in the text editor of your choice
```

# SendHTMLEmail
`SendHTMLEmail()` sends a HTML template email.

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

	data := "<div><b> Kindly click on Verify Account to confirm your account </b></div>"
	req := zeptomail.SendHTMLEmailReq{
		To: []zeptomail.SendEmailTo{
			{
				EmailAddress: zeptomail.EmailAddress{
					Address: "rudra.d@zylker.com",
					Name:    "Rudra",
				},
			},
		},
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


# SendTemplatedEmail
`SendTemplatedEmail()` sends a templated email set from your ZeptoMail dashboard.

This method takes in the `zeptomail.SendTemplatedEmailReq{}` struct as a parameter.
### List of all the fields available in the zeptomail.SendTemplatedEmailReq{} struct.
```go
type ( 
    // SendHTMLEmailReq is the SendHTMLEmail() request object
   SendTemplatedEmailReq struct {
		/*
			Unique key identifier for your template.
			template_alias is the alias name given to the template key. It can be used in the place of template_key.​
		*/
		TemplateKey string `json:"template_key"`
		/*
			The email address to which bounced emails will be sent.

			Allowed value - A valid bounce email address as configured in your Mail Agent.
		*/
		BounceAddress string `json:"bounce_address"`
		// Allowed value - A valid sender email address with "address" and "name" key-value pairs
		From EmailAddress `json:"from" validate:"required"`
		// Allowed value - JSON object of email_address.
		To []SendEmailTo `json:"to" validate:"required"`
		/*
			You can use merge tags to replace the placeholders with multiple values for different recipients. The merge tags of the mail to be sent. Merge Tags can be set in HTML body, text body, subject and client reference.

			Merge tags are tags that are set in the string.

			These tags will be replaced by the value provided in this param for that tag.

			In the case of single email API, the merge tag is common for all recipients. Whereas in the case of batch emails, the merge tag is specific to each recipient.
		*/
		MergeInfo map[string]interface{} `json:"merge_info" validate:"required"`
		// Allowed value - JSON object of reply_to email addresses.
		ReplyTo EmailAddress `json:"reply_to"`
		/*
			You can enable or disable email click tracking here.

			You can also enable email click tracking in your Mail Agent under Email Tracking section.

			Note: The API setting will override the Mail Agent settings in your ZeptoMail account.

			Allowed value

			True - Enable email click tracking.

			False - Disable email click tracking.
		*/
		TrackClicks bool `json:"track_clicks"`
		/*
					You can enable or disable email open tracking.

			You can also enable email open tracking in your Mail Agent under Email Tracking section.

			Note: The API setting will override the Mail Agent settings in your ZeptoMail account.

			Allowed value

			True - Enable email open tracking.

			False - Disable email open tracking.
		*/
		TrackOpens bool `json:"track_opens"`
		// An identifier set by the user to track a particular transaction.
		ClientReference string `json:"client_reference"`
		// The additional headers to be sent in the email for your reference purposes.
		MimeHeaders map[string]interface{} `json:"mime_headers"`
		/*
			The attachments you want to add to your transactional emails. Visit [https://www.zoho.com/zeptomail/help/file-cache.html#alink-un-sup-for] to view the list of unsupported formats.

			Allowed value - JSON object of attachments.

			It can either be a base64 encoded content or file_cache_key or both.
		*/
		Attachments []struct {
			/*
				The content of your attachment.

				Allowed value - Base64 encoded value of a file.
			*/
			Content string `json:"content"`
			/*
				Indicates the content type in your attachment.

				Allowed value

				simple text message - plain / text

				image file - image / jpg
			*/
			MimeType string `json:"mime_type"`
			/*
				he file name of your attachment as in the File Cache section in your Mail Agent.

				Obtain file name from the File Cache section in your Mail Agent.
			*/
			Name string `json:"name"`
			/*
				The unique key for your attached files in a Mail Agent.

				Obtain file_cache_key from the File Cache section in your Mail Agent.
			*/
			FileCacheKey string `json:"file_cache_key"`
			/*
				the content id used by html body for content lookup. Each content should be given a separate cid value.

				Allowed value

				It can either a base64 encoded content or a file_cache_key or both.
			*/
			Cid string `json:"cid"`
		} `json:"attachments"`
	}

	// SendEmailTo is the object for recipient email
	SendEmailTo struct {
		// A valid recipient email address with "address" and "name" key-value pairs.
		EmailAddress EmailAddress `json:"email_address" validate:"required"`
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
	tempKey := "your zeptomail template key"
	client := zeptomail.New(*http.DefaultClient, zeptomailToken)

	req := zeptomail.SendTemplatedEmailReq{
		To: []zeptomail.SendEmailTo{
			{
				EmailAddress: zeptomail.EmailAddress{
					Address: "rudra.d@zylker.com",
					Name:    "Rudra",
				},
			},
		},
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
```
