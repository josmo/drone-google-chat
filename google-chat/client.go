package google_chat

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type Message struct {
	Text string `json:"text"`
}

type Client interface {
	SendMessage(*Message) error
}

type client struct {
	url string
}

func NewClient(url string, key string, token string, conversationKey string) Client {
	fullURL = url + "/messages?key=" + key + "&token=" + token
	return &client{fullURL}
}

func (c *client) SendMessage(msg *Message) error {

	body, _ := json.Marshal(msg)
	buf := bytes.NewReader(body)

	resp, err := http.Post(c.url, "application/json", buf)
	if err != nil {
		return err
	}
	//TODO: fix error and give better feedback
	if resp.StatusCode != 200 {
		//	//t, _ := ioutil.ReadAll(resp.Body)
		//	//return &Error{resp.StatusCode, string(t)}
	}

	return nil
}
