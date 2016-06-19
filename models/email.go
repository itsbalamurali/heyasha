package models

type Email struct {
	Recepient string `json:"recepient"`
	Sender string `json:"sender"`
	From	string `json:"from"`
	Subject string `json:"subject"`
	BodyPlain string `json:"body-plain"`
	StrippedText string `json:"stripped-text"`
	StrippedSignature string `json:"stripped-signature"`
	BodyHtml string `json:"body-html"`
	StrippedHtml string `json:"stripped-html"`
	MessageUrl string `json:"message-url"`
	ContentIdMap string `json:"content-id-map"`
	Timestamp int `json:"timestamp"`
	Token string `json:"token"`
	Signature string `json:"signature"`
	MessageHeaders string `json:"message-headers"`
	Attachments []EmailAttachment `json:"attachments"`
}

type EmailAttachment struct {
	Size int `json:"size"`
	Url string `json:"url"`
	Name string `json:"name"`
	ContentType string `json:"content-type"`
}