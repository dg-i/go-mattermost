package mattermost

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// Client holds the Mattermost client config
type Client struct {
	WebhookURL string
	User       string
	Channel    string
}

// Post holds information about a Mattermost message
type Post struct {
	Text        string       `json:"text,omitempty"`
	Username    string       `json:"username,omitempty"`
	IconURL     string       `json:"icon_url,omitempty"`
	Channel     string       `json:"channel,omitempty"`
	Attachments []Attachment `json:"attachments,omitempty"`
}

// Attachment contains a Slack attachment of a Mattermost chat message
type Attachment struct {
	Fallback   string            `json:"fallback,omitempty"`
	Color      string            `json:"color,omitempty"`
	Pretext    string            `json:"pretext,omitempty"`
	Text       string            `json:"text,omitempty"`
	Title      string            `json:"title,omitempty"`
	TitleLink  string            `json:"title_link,omitempty"`
	AuthorName string            `json:"author_name,omitempty"`
	AuthorIcon string            `json:"author_icon,omitempty"`
	AuthorLink string            `json:"author_link,omitempty"`
	Fields     []AttachmentField `json:"fields,omitempty"`
	ImageURL   string            `json:"image_url,omitempty"`
}

// AttachmentField contains attachment fields for usage in attachments
type AttachmentField struct {
	Short bool   `json:"short"`
	Title string `json:"title,omitempty"`
	Value string `json:"value,omitempty"`
}

// NewClient returns a Mattermost client config
func NewClient(url, user, channel string) *Client {
	return &Client{
		WebhookURL: url,
		User:       user,
		Channel:    channel,
	}
}

// PostString builds a message with the given string and posts it using PostMessage
func (c *Client) PostString(text string) error {
	return c.Post(Post{
		Username: c.User,
		Channel:  c.Channel,
		Text:     text,
	})
}

// Post posts a message to the configured channel
func (c *Client) Post(msg Post) error {
	msg.Username = c.User
	msg.Channel = c.Channel

	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(msg)
	resp, err := http.Post(c.WebhookURL, "application/json; charset=utf-8", b)
	if err != nil {
		return err
	}
	if (err == nil) && (resp.StatusCode != http.StatusOK) {
		err = fmt.Errorf("HTTP response was %d", resp.StatusCode)
	}
	defer resp.Body.Close()
	return nil
}
