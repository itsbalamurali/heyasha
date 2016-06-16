package kik

import (
	"bytes"
	"encoding/json"
	"net/http"
)

// Action is used to determine what kind of message a webhook event is.
type Type int

const (
	// UnknownAction means that the event was not able to be classified.
	UnknownType Type = iota - 1
	TextType
	LinkType
	PictureType
	VideoType
	StartChatType
	StickerType
	IsTypingType
	ScanDataType
	DeliveryReceiptType
	ReadReceiptType
	//Kikbot Messaging Url
	SendMessageURL = "https://api.kik.com/v1/message"

)

type KikWebhook struct {
	Messages []MessageReceived `json:"messages"`
}

type KikResponse struct {
	Messages []SendMessage  `json:"messages"`
}

type SendMessage struct {
	ChatID string `json:"chatId"`
	To     string `json:"to"`
	Type   string `json:"type"`
	Body   string `json:"body"`
}

type MessageReceived struct {
	ChatID               string    `json:"chatId"`
	Type                 string    `json:"type"`
	From                 string    `json:"from"`
	Participants         []string  `json:"participants"`
	ID                   string    `json:"id"`
	MessageIDs           []string  `json:"messageIds,omitempty"`
	Timestamp            int64 `json:"timestamp"`
	ReadReceiptRequested bool      `json:"readReceiptRequested"`
	Mention              string    `json:"mention"`
	IsTyping             bool      `json:"isTyping,omitempty"`
	StickerPackId        string    `json:"stickerPackId,omitempty"`
	stickerUrl           string    `json:"stickerUrl,omitempty"`
	Data                 string    `json:"data,omitempty"`
	Body                 string    `json:"body,omitempty"`
	Url                  string    `json:"url,omitempty"`
	PicUrl               string    `json:"picUrl,omitempty"`
	Attribution          Attr      `json:"attribution,omitempty"`
	NoForward            bool      `json:"noForward,omitempty"`
	KikJsData            string    `json:"kikJsData,omitempty"`
	VideoUrl             string    `json:"videoUrl,omitempty"`
}

type Attr struct {
	Name    string `json:"name"`
	IconUrl string `json:"iconUrl"`
}

func Classify(msg MessageReceived) Type {

	switch msg.Type {
	case "text":
		return TextType
	case "link":
		return LinkType
	case "picture":
		return PictureType
	case "video":
		return VideoType
	case "start-chatting":
		return StartChatType
	case "scan-data":
		return ScanDataType
	case "sticker":
		return StickerType
	case "is-typing":
		return IsTypingType
	case "delivery-receipt":
		return DeliveryReceiptType
	case "read-receipt":
		return ReadReceiptType
	default:
		return UnknownType
	}

}

// Text sends a textual message.
func Text(to string,chatid string,message string) error {
	m := KikResponse{
		Messages: []SendMessage{
			{
			To: to,
			Type: "text",
			ChatID: chatid,
			Body: message,
			},
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
	req.SetBasicAuth("talkwithasha", "4ce43ba4-938e-4581-9989-0185c6b66ac2")
	client := &http.Client{}

	resp, err := client.Do(req)
	defer resp.Body.Close()

	return err
}