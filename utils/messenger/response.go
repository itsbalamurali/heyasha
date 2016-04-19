package messenger

import (
	"bytes"
	"encoding/json"
	"fmt"
	"image"
	"image/jpeg"
	"io"
	"mime/multipart"
	"net/http"
)

const (
	// SendMessageURL is API endpoint for sending messages.
	SendMessageURL = "https://graph.facebook.com/v2.6/me/messages"
)

// Response is used for responding to events with messages.
type Response struct {
	token string
	to    Recipient
}

// Text sends a textual message.
func (r *Response) Text(message string) error {
	m := SendMessage{
		Recipient: r.to,
		Message: MessageData{
			Text: message,
		},
	}

	data, err := json.Marshal(m)
	if err != nil {
		return nil
	}

	req, err := http.NewRequest("POST", SendMessageURL, bytes.NewBuffer(data))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.URL.RawQuery = "access_token=" + r.token

	client := &http.Client{}

	resp, err := client.Do(req)
	defer resp.Body.Close()

	return err
}

// Image sends an image.
func (r *Response) Image(im image.Image) error {
	var b bytes.Buffer
	w := multipart.NewWriter(&b)

	data, err := w.CreateFormFile("fielddata", "meme.jpg")
	if err != nil {
		return err
	}

	imageBytes := new(bytes.Buffer)
	err = jpeg.Encode(imageBytes, im, nil)
	if err != nil {
		return err
	}

	_, err = io.Copy(data, imageBytes)
	if err != nil {
		return err
	}

	w.WriteField("recipient", fmt.Sprintf(`{"id":"%v"}`, r.to.ID))
	w.WriteField("message", `{"attachment":{"type":"image", "payload":{}}}`)

	req, err := http.NewRequest("POST", SendMessageURL, &b)
	if err != nil {
		return err
	}

	req.URL.RawQuery = "access_token=" + r.token

	req.Header.Set("Content-Type", w.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	var res bytes.Buffer
	res.ReadFrom(resp.Body)
	fmt.Println(res.String(), "DONE!")
	return nil
}

// SendMessage is the information sent in an API request to Facebook.
type SendMessage struct {
	Recipient Recipient   `json:"recipient"`
	Message   MessageData `json:"message"`
}

// MessageData is a text message to be sent.
type MessageData struct {
	Text string `json:"text,omitempty"`
}
