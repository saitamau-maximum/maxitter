package external

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type DiscordImg struct {
	URL string `json:"url"`
	H   int    `json:"height"`
	W   int    `json:"width"`
}
type DiscordAuthor struct {
	Name string `json:"name"`
	URL  string `json:"url"`
	Icon string `json:"icon_url"`
}
type DiscordField struct {
	Name   string `json:"name"`
	Value  string `json:"value"`
	Inline bool   `json:"inline"`
}
type DiscordEmbed struct {
	Title     string         `json:"title"`
	Desc      string         `json:"description"`
	URL       string         `json:"url"`
	Color     int            `json:"color"`
	Image     DiscordImg     `json:"image"`
	Thum      DiscordImg     `json:"thumbnail"`
	Author    DiscordAuthor  `json:"author"`
	Fields    []DiscordField `json:"fields"`
	TimeStamp string         `json:"timestamp"`
}

type DiscordWebhook struct {
	UserName  string         `json:"username"`
	AvatarURL string         `json:"avatar_url"`
	Content   string         `json:"content"`
	Embeds    []DiscordEmbed `json:"embeds"`
	TTS       bool           `json:"tts"`
}

func SendWebhookDiscord(whurl, username, avater_url string, dw *DiscordWebhook) error {
	j, err := json.Marshal(dw)
	if err != nil {
		return fmt.Errorf("JSON error: %v", err)
	}

	req, err := http.NewRequest("POST", whurl, bytes.NewBuffer(j))
	if err != nil {
		return fmt.Errorf("new request error: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("http client error: %v", err)
	}
	if resp.StatusCode != 204 {
		return fmt.Errorf("failed to send a message to Webhook in Discord (%#v)", resp)
	}

	return nil
}
