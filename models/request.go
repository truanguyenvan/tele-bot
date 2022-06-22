package models

type SendMessage struct {
	Text    string `json:"text" validate:"required"`
	Header  string `json:"header" validate:"required"`
	Chat_id string `json:"chat_id" validate:"required"`
	Emoji   string `json:"emoji" validate:"required"`
}
