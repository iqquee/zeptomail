package zeptomail

type (

	// ErrorResponse is an object for errors
	ErrorResponse struct {
		Code string `json:"code"`
		// It consists of code, message and target parameters
		Details []struct {
			Code    string `json:"code"`
			Message string `json:"message"`
			Target  string `json:"target"`
		} `json:"details"`
		// Reason for the error
		Message string `json:"message"`
		// The field that caused the error
		Target string `json:"target"`
		// A unique id which is generated for every request
		RequestId string `json:"request_id"`
	}

	// SendHTMLEmailReq is the SendHTMLEmail() request object
	SendHTMLEmailReq struct {
		From     EmailAddress  `json:"from" validate:"required"`
		To       []SendEmailTo `json:"to" validate:"required"`
		Subject  string        `json:"subject" validate:"required"`
		Htmlbody string        `json:"htmlbody" validate:"required"`
	}

	// SendHTMLEmailRes is the SendHTMLEmail() response object
	SendHTMLEmailRes struct {
		Data      []SendEmailResData `json:"data"`
		Message   string             `json:"message"`
		RequestId string             `json:"request_id"`
		Object    string             `json:"object"`
		Error     ErrorResponse      `json:"error"`
	}

	// SendEmailResData is the Data object for the SendHTMLEmailRes object
	SendEmailResData struct {
		Code           string        `json:"code"`
		AdditionalInfo []interface{} `json:"additional_info"`
		Message        string        `json:"message"`
	}

	// SendTemplatedEmailReq is the SendTemplatedEmail() request object
	SendTemplatedEmailReq struct {
		/*
			Unique key identifier for your template.
			template_alias is the alias name given to the template key. It can be used in the place of template_key.​
		*/
		TemplateKey string `json:"template_key" validate:"required"`
		/*
			The email address to which bounced emails will be sent.

			Allowed value - A valid bounce email address as configured in your Mail Agent.
		*/
		BounceAddress string `json:"bounce_address"`
		// Allowed value - A valid sender email address with "address" and "name" key-value pairs
		From EmailAddress `json:"from" validate:"required"`
		// Allowed value - JSON object of email_address.
		To []SendEmailTo `json:"to" validate:"required"`

		CC  []SendEmailTo `json:"cc"`
		BCC []SendEmailTo `json:"bcc"`
		/*
			You can use merge tags to replace the placeholders with multiple values for different recipients. The merge tags of the mail to be sent. Merge Tags can be set in HTML body, text body, subject and client reference.

			Merge tags are tags that are set in the string.

			These tags will be replaced by the value provided in this param for that tag.

			In the case of single email API, the merge tag is common for all recipients. Whereas in the case of batch emails, the merge tag is specific to each recipient.
		*/
		MergeInfo map[string]interface{} `json:"merge_info" validate:"required"`
		// Allowed value - JSON object of reply_to email addresses.
		ReplyTo []EmailAddress `json:"reply_to"`
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

	// SendBatchTemplateEmailTo is the object for send batch recipient emails
	SendBatchTemplateEmailTo struct {
		// A valid recipient email address with "address" and "name" key-value pairs.
		EmailAddress EmailAddress `json:"email_address" validate:"required"`

		/*
			You can use merge tags to replace the placeholders with multiple values for different recipients. The merge tags of the mail to be sent. Merge Tags can be set in HTML body, text body, subject and client reference.

			Merge tags are tags that are set in the string.

			These tags will be replaced by the value provided in this param for that tag.

			In the case of single email API, the merge tag is common for all recipients. Whereas in the case of batch emails, the merge tag is specific to each recipient.
		*/
		MergeInfo map[string]interface{} `json:"merge_info" validate:"required"`
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

	// SendTemplatedEmailRes is the SendTemplatedEmail() response object
	SendTemplatedEmailRes struct {
		Data []struct {
			Code           string        `json:"code"`
			AdditionalInfo []interface{} `json:"additional_info"`
			Message        string        `json:"message"`
		} `json:"data"`
		Message   string        `json:"message"`
		RequestId string        `json:"request_id"`
		Object    string        `json:"object"`
		Error     ErrorResponse `json:"error"`
	}

	// SendBatchTemplatedEmailReq is the SendBatchTemplatedEmail() request object
	SendBatchTemplatedEmailReq struct {
		/*
			Unique key identifier for your template.
			template_alias is the alias name given to the template key. It can be used in the place of template_key.​
		*/
		TemplateKey string `json:"template_key" validate:"required"`
		/*
			The email address to which bounced emails will be sent.

			Allowed value - A valid bounce email address as configured in your Mail Agent.
		*/
		BounceAddress string `json:"bounce_address"`
		// Allowed value - A valid sender email address with "address" and "name" key-value pairs
		From EmailAddress `json:"from" validate:"required"`
		// Allowed value - JSON object of email_address.
		To []SendBatchTemplateEmailTo `json:"to" validate:"required"`
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

	// AddEmailTemplateReq is the AddEmailTemplate() response object
	AddEmailTemplateReq struct {
		// 	Name of the template
		TemplateName string `json:"template_name" validate:"required"`
		// Subject to be added in the email template.
		Subject string `json:"subject" validate:"required"`
		// The corresponding HTML content of the body in the email template.
		HtmlBody string `json:"htmlbody"`
		// 	Plain text body of the email in the template.
		TextBody string `json:"textbody"`
		// 	Unique alias for the template. It can be used instead of the template key. It can be obtained from the Edit template section in your ZeptoMail account.
		TemplateAlias string `json:"template_alias"`
		// Unique alias value given to the Mail Agent. It is available in the Setup info section of your Mail Agent.
		MailagentAlias string `json:"_" validate:"required"`
	}

	// AddEmailTemplateRes is the AddEmailTemplate() request object
	AddEmailTemplateRes struct {
		Data []struct {
			HtmlBody      string `json:"htmlbody"`
			UploadTime    string `json:"upload_time"`
			Template_name string `json:"template_name"`
			Template_key  string `json:"template_key"`
			Template_size int    `json:"template_size"`
			Modified_time string `json:"modified_time"`
			Subject       string `json:"subject"`
		} `json:"data"`
		Message string        `json:"message"`
		Object  string        `json:"object"`
		Error   ErrorResponse `json:"error"`
	}

	// UpdateEmailTemplateReq is the UpdateEmailTemplate() response object
	UpdateEmailTemplateReq struct {
		// 	Name of the template
		TemplateName string `json:"template_name" validate:"required"`
		// Subject to be added in the email template.
		Subject string `json:"subject" validate:"required"`
		// The corresponding HTML content of the body in the email template.
		HtmlBody string `json:"htmlbody"`
		// Plain text body of the email in the template.
		TextBody string `json:"textbody"`
		// TemplateKey is the template key.
		TemplateKey string `json:"" validate:"required"`
		// Unique alias value given to the Mail Agent. It is available in the Setup info section of your Mail Agent.
		MailagentAlias string `json:"" validate:"required"`
	}

	// GetEmailTemplateRes is the GetEmailTemplate() response object
	GetEmailTemplateRes struct {
		Data struct {
			HtmlBody     string `json:"htmlbody"`
			CreatedTime  string `json:"created_time"`
			TemplateName string `json:"template_name"`
			TemplateKey  string `json:"template_key"`
			ModifiedTime string `json:"modified_time"`
			Attachments  []struct {
				FileCacheKey string `json:"file_cache_key"`
				ContentType  string `json:"content_type"`
				FileName     string `json:"file_name"`
			} `json:"attachments"`
			Subject         string `json:"subject"`
			TemplateAlias   string `json:"template_alias"`
			SampleMergeInfo struct {
				Name string `json:"name"`
				Link string `json:"link"`
			} `json:"sample_merge_info"`
		} `json:"data"`
		Message string        `json:"message"`
		Object  string        `json:"object"`
		Error   ErrorResponse `json:"error"`
	}

	// SendBatchHTMLEmailReq is the SendBatchHTMLEmail() request object
	SendBatchHTMLEmailReq struct {
		From      EmailAddress               `json:"from" validate:"required"`
		To        []SendBatchTemplateEmailTo `json:"to" validate:"required"`
		MergeInfo map[string]interface{}     `json:"merge_info" validate:"required"`
		Subject   string                     `json:"subject" validate:"required"`
		Htmlbody  string                     `json:"htmlbody" validate:"required"`
	}

	// SendHTMLEmailRes is the SendHTMLEmail() response object
	SendBatchHTMLEmailRes struct {
		Data      []SendBatchResData `json:"data"`
		Message   string             `json:"message"`
		RequestId string             `json:"request_id"`
		Object    string             `json:"object"`
		Error     ErrorResponse      `json:"error"`
	}

	// SendBatchResData is the Data object for the SendHTMLEmailRes object
	SendBatchResData struct {
		Code           string        `json:"code"`
		AdditionalInfo []interface{} `json:"additional_info"`
		Message        string        `json:"message"`
	}

	// FileCacheUploadAPIReq is the FileCacheUploadAPI() response
	FileCacheUploadAPIReq struct {
		// name of file uploaded
		FileName string `json:"name" validate:"required"`
		// file path of the file uploaded6
		FileContent []byte `json:"file" validate:"required"`
	}

	// FileCacheUploadAPIRes is the FileCacheUploadAPI() response object
	FileCacheUploadAPIRes struct {
		FileCacheKey string                      `json:"file_cache_key"`
		Data         []FileCacheUploadAPIResData `json:"data"`
		Message      string                      `json:"message"`
		Object       string                      `json:"object"`
		Error        ErrorResponse               `json:"error"`
	}

	// FileCacheUploadAPIResData is the Data object for the FileCacheUploadAPIRes object
	FileCacheUploadAPIResData struct {
		Code           string        `json:"code"`
		AdditionalInfo []interface{} `json:"additional_info"`
		Message        string        `json:"message"`
	}

	// ListEmailTemplatesRes is Data object for the ListEmailTemplates() response object
	ListEmailTemplatesRes struct {
		Metadata struct {
			Count  int `json:"count"`
			Offset int `json:"offset"`
			Limit  int `json:"limit"`
		}
		Data struct {
			CreatedTime   string `json:"created_time"`
			TemplateName  string `json:"template_name"`
			TemplateKey   string `json:"template_key"`
			ModifiedTime  string `json:"modified_time"`
			Subject       string `json:"subject"`
			TemplateAlias string `json:"template_alias"`
		} `json:"data"`
		Message string        `json:"message"`
		Error   ErrorResponse `json:"error"`
	}
)
