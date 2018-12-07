package line

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"
)

const accessTokenPath = "accesstoken"

type LineServer struct {
	baseURL      string
	auth2Path    string
	callbackPath string

	loginID     string
	loginSecret string

	Bot *LineBot
}

func NewLineAuthServer(baseURL, auth2Path, callbackPath string) (*LineServer, error) {
	const loginID = "1624507559"
	const loginSecret = "a9ddb1060860a4a26e9c27d43d8e3249"
	const botID = "1571865955"
	const botSecret = "c8fc4a0931b9a11f7b2179de442ae779"
	const botToken = "jFkLZLb7XZhd+/sE6a2noJ1QD6Qac8nr82Vw4KaeZd85iS+PRugvjhiUO2ei3GhFdvuBPXorPBWBxDemJiphzW9Oy0G2VQOsFzxSBBPfsabaDMuIi9Bq2dLKYL6mOBMOlZU6JvMr2WgX9uTlFNnBxAdB04t89/1O/w1cDnyilFU="

	bot := NewLineBot(botToken)

	return &LineServer{
		baseURL,
		auth2Path,
		callbackPath,

		loginID,
		loginSecret,

		bot,
	}, nil
}

func (server *LineServer) Start(port string) {
	http.HandleFunc("/"+string(server.auth2Path), server.authen)
	http.HandleFunc("/"+accessTokenPath, server.accessToken)

	log.Println("Start server listening to port", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}

func (server *LineServer) authen(w http.ResponseWriter, r *http.Request) {
	state := "abcde"
	url := fmt.Sprintf("https://access.line.me/oauth2/v2.1/authorize?response_type=code&client_id=%s&redirect_uri=%s&state=%s&scope=profile&prompt=consent&bot_prompt=aggressive", server.loginID, url.QueryEscape(server.baseURL+accessTokenPath), state)
	log.Println("Redirect to", url)
	http.Redirect(w, r, url, http.StatusSeeOther)
}

func (server *LineServer) accessToken(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	err := query.Get("error")
	if err != "" {
		reason := query.Get("error_description")
		log.Print(reason)
		return
	}

	state := query.Get("state")
	code := query.Get("code")
	friendshipChange := query.Get("friendship_status_changed")
	log.Print("state = ", state)
	log.Print("code = ", code)
	log.Print("friendship change = ", friendshipChange)

	body := url.Values{}
	body.Set("grant_type", "authorization_code")
	body.Set("code", code)
	body.Set("redirect_uri", server.baseURL+accessTokenPath)
	body.Set("client_id", server.loginID)
	body.Set("client_secret", server.loginSecret)

	resp, respErr := http.Post("https://api.line.me/oauth2/v2.1/token", "application/x-www-form-urlencoded", strings.NewReader(body.Encode()))
	if respErr != nil {
		log.Panic(respErr)
	}
	defer resp.Body.Close()

	respBody := map[string]interface{}{}
	dErr := json.NewDecoder(resp.Body).Decode(&respBody)
	if dErr != nil {
		log.Panic(dErr)
	}

	if resp.StatusCode != 200 {
		fmt.Println("response Status:", resp.Status)
		fmt.Println("response Body:", respBody)
		return
	}

	accessToken := respBody["access_token"].(string)
	log.Printf(accessToken)
	getProfileReq, _ := http.NewRequest("GET", "https://api.line.me/v2/profile", nil)
	getProfileReq.Header.Add("Authorization", "Bearer "+accessToken)

	respBody2 := map[string]interface{}{}
	client := &http.Client{}
	getProfileResp, _ := client.Do(getProfileReq)
	log.Println(getProfileResp.Status)
	json.NewDecoder(getProfileResp.Body).Decode(&respBody2)
	log.Println(respBody2)
	log.Println("userID =", respBody2["userId"].(string))

	userID := respBody2["userId"].(string)
	server.Bot.PushMessage(userID, "test")
}
