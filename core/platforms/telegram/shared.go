package telegram

import "fmt"

// Recipient represents the recipient of a message
type Recipient struct {
	ChatID    *int
	ChannelID *string
}

// NewChatRecipient creates a new recipient for private or group chats
func NewChatRecipient(chatID int) Recipient {
	return Recipient{
		ChatID: &chatID,
	}
}

// NewChannelRecipient creates a new recipient for channels
func NewChannelRecipient(channelName string) Recipient {
	return Recipient{
		ChannelID: &channelName,
	}
}

// NewRecipientFromChat creates a recipient that addresses the given chat
func NewRecipientFromChat(chat Chat) Recipient {
	return NewChatRecipient(chat.ID) //No need to distinguish between channels and chats, bots cannot receive from channels
}

func (r Recipient) isChat() bool {
	return r.ChatID != nil
}

func (r Recipient) isChannel() bool {
	return r.ChannelID != nil
}

// MarshalJSON marshals the recipient to JSON
func (r Recipient) MarshalJSON() ([]byte, error) {
	toReturn := ""

	if r.isChannel() {
		toReturn = fmt.Sprintf("\"%s\"", *r.ChannelID)
	} else {
		toReturn = fmt.Sprintf("%d", *r.ChatID)
	}

	return []byte(toReturn), nil
}
