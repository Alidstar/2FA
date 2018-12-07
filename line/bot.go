package line

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type LineBot struct {
	accessToken string
}

func NewLineBot(accessToken string) *LineBot {
	return &LineBot{
		accessToken,
	}
}

type pushMessageReq struct {
	To       string    `json:"to"`
	Messages []message `json:"messages"`
}

func (bot *LineBot) PushMessage(userID string, text string) error {
	data, _ := json.Marshal(pushMessageReq{
		userID,
		[]message{
			marshalTextMessage(text),
		},
	})

	log.Println("data", string(data))

	client := http.Client{}
	req, reqErr := http.NewRequest("POST", "https://api.line.me/v2/bot/message/push", bytes.NewBuffer(data))
	if reqErr != nil {
		log.Panic(reqErr)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+bot.accessToken)

	resp, respErr := client.Do(req)
	if respErr != nil {
		log.Panic(respErr)
	}

	defer resp.Body.Close()

	log.Println("Push Status:", resp.Status)
	if resp.StatusCode != 200 {
		return fmt.Errorf("%v", resp.Status)
	}

	return nil
}
