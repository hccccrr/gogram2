package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/amarnathcjd/gogram/telegram"
)

var (
	appID   int32  = 25742938                           
	appHash string = "b35b715fe8dc0a58e8048988286fc5b6" 
	token   string = "7623679464:AAGqdslPgtOzrAtycf6iuuDGPAJZCw4vJR0" 
)

func main() {
	// Gogram client initialization (Updated Syntax)
	client, err := telegram.NewClient(telegram.ClientConfig{
		AppID:    appID,
		AppHash:  appHash,
		LogLevel: telegram.LogInfo,
	})
	if err != nil {
		log.Fatal("Failed to create client:", err)
	}

	// Login as Bot
	err = client.LoginBot(token)
	if err != nil {
		log.Fatal("Bot login failed:", err)
	}

	fmt.Println(">> Bot is online! Send /start to generate session.")

	client.OnMessage(func(ctx *telegram.NewMessage) error {
		if ctx.Text() == "/start" {
			ctx.Reply("Gogram String Session generate karan layi apna Phone Number bhejo (e.g., +919876543210):")
			return nil
		}

		// Agar message + naal start hunda hai (Phone Number)
		if strings.HasPrefix(ctx.Text(), "+") {
			phone := ctx.Text()
			
			// User client setup for session generation
			userClient, _ := telegram.NewClient(telegram.ClientConfig{
				AppID:   appID,
				AppHash: appHash,
			})

			sentCode, err := userClient.AuthSendCode(phone)
			if err != nil {
				ctx.Reply("Error: " + err.Error())
				return nil
			}

			ctx.Reply("OTP mil gaya? Kripya OTP bhej dyo (Format -> otp:12345):")
			
			// OTP Handler
			client.OnMessage(func(otpCtx *telegram.NewMessage) error {
				if strings.HasPrefix(otpCtx.Text(), "otp:") {
					code := strings.TrimPrefix(otpCtx.Text(), "otp:")
					_, err := userClient.AuthSignIn(phone, sentCode.PhoneCodeHash, code)
					if err != nil {
						otpCtx.Reply("Invalid OTP: " + err.Error())
						return nil
					}

					// Export session for your bot
					session, _ := userClient.ExportSession()
					otpCtx.Reply("✅ Tuhadi Gogram Session String:\n\n`" + session + "`")
				}
				return nil
			})
		}
		return nil
	})

	client.Idle()
}
