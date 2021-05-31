package linebot

import (
	"encoding/json"
)

// QuickReplyItems struct
type QuickReplyItems struct {
	Items []*QuickReplyButton `json:"items"`
}

// NewQuickReplyItems function
func NewQuickReplyItems(buttons ...*QuickReplyButton) *QuickReplyItems {
	return &QuickReplyItems{
		Items: buttons,
	}
}

// QuickReplyButton type
type QuickReplyButton struct {
	ImageURL string
	Action   QuickReplyAction
}

// MarshalJSON method of QuickReplyButton
func (b *QuickReplyButton) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Type     string           `json:"type"`
		ImageURL string           `json:"imageUrl,omitempty"`
		Action   QuickReplyAction `json:"action"`
	}{
		Type:     "action",
		ImageURL: b.ImageURL,
		Action:   b.Action,
	})
}

// NewQuickReplyButton function
func NewQuickReplyButton(imageURL string, action QuickReplyAction) *QuickReplyButton {
	return &QuickReplyButton{
		ImageURL: imageURL,
		Action:   action,
	}
}
