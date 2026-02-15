package main

import (
	"bufio"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/amarnathcjd/gogram/telegram"
)

type SessionData struct {
	AuthKey    []byte `json:"auth_key"`
	AuthKeyID  []byte `json:"auth_key_id"`
	ServerSalt []byte `json:"server_salt"`
	Addr       string `json:"addr"`
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("╔════════════════════════════════════════╗")
	fmt.Println("║   Telegram String Session Generator   ║")
	fmt.Println("║         Using Gogram Library          ║")
	fmt.Println("╚════════════════════════════════════════╝")
	fmt.Println()

	// Get API ID
	fmt.Print("Enter your API ID: ")
	var apiID int32
	fmt.Scanln(&apiID)

	// Get API Hash
	fmt.Print("Enter your API Hash: ")
	apiHash, _ := reader.ReadString('\n')
	apiHash = strings.TrimSpace(apiHash)

	// Get Phone Number
	fmt.Print("Enter your phone number (with country code, e.g., +911234567890): ")
	phoneNumber, _ := reader.ReadString('\n')
	phoneNumber = strings.TrimSpace(phoneNumber)

	fmt.Println("\n[*] Connecting to Telegram...")

	// Create client
	client, err := telegram.NewClient(telegram.ClientConfig{
		AppID:    apiID,
		AppHash:  apiHash,
		Session:  telegram.NewInMemorySession(),
		LogLevel: telegram.LogInfo,
	})

	if err != nil {
		fmt.Printf("[-] Error creating client: %v\n", err)
		return
	}

	// Connect to Telegram
	err = client.Connect()
	if err != nil {
		fmt.Printf("[-] Error connecting: %v\n", err)
		return
	}
	defer client.Disconnect()

	fmt.Println("[+] Connected successfully!")
	fmt.Println("[*] Sending OTP to your phone number...")

	// Send code
	sentCode, err := client.AuthSendCode(
		phoneNumber,
		apiID,
		apiHash,
		&telegram.CodeSettings{},
	)

	if err != nil {
		fmt.Printf("[-] Error sending code: %v\n", err)
		return
	}

	fmt.Println("[+] OTP sent successfully!")
	fmt.Print("Enter the OTP you received: ")
	code, _ := reader.ReadString('\n')
	code = strings.TrimSpace(code)

	// Sign in with code
	auth, err := client.AuthSignIn(
		phoneNumber,
		sentCode.PhoneCodeHash,
		code,
	)

	if err != nil {
		// Check if 2FA is enabled
		if strings.Contains(err.Error(), "SESSION_PASSWORD_NEEDED") || strings.Contains(err.Error(), "password") {
			fmt.Println("\n[!] 2FA is enabled on your account")
			fmt.Print("Enter your 2FA password: ")
			password, _ := reader.ReadString('\n')
			password = strings.TrimSpace(password)

			// Get password info
			accountPassword, err := client.AccountGetPassword()
			if err != nil {
				fmt.Printf("[-] Error getting password info: %v\n", err)
				return
			}

			// Sign in with password
			auth, err = client.AuthCheckPassword(accountPassword, password)
			if err != nil {
				fmt.Printf("[-] Error with 2FA password: %v\n", err)
				return
			}
		} else {
			fmt.Printf("[-] Error signing in: %v\n", err)
			return
		}
	}

	fmt.Println("\n[+] Login successful!")
	fmt.Printf("[+] Logged in as: %s\n", getUserName(auth))

	// Generate string session
	session := client.ExportStringSession()
	
	fmt.Println("\n╔════════════════════════════════════════╗")
	fmt.Println("║          STRING SESSION READY          ║")
	fmt.Println("╚════════════════════════════════════════╝")
	fmt.Println()
	fmt.Println("Your String Session:")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println(session)
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()
	fmt.Println("[!] Keep this session string SAFE and PRIVATE!")
	fmt.Println("[!] Anyone with this session can access your account!")
	fmt.Println()

	// Save to file
	fmt.Print("Do you want to save this to a file? (y/n): ")
	save, _ := reader.ReadString('\n')
	save = strings.TrimSpace(strings.ToLower(save))

	if save == "y" || save == "yes" {
		err = os.WriteFile("string_session.txt", []byte(session), 0600)
		if err != nil {
			fmt.Printf("[-] Error saving file: %v\n", err)
		} else {
			fmt.Println("[+] Session saved to string_session.txt")
		}
	}

	fmt.Println("\n[*] Session generation complete!")
}

func getUserName(auth telegram.AuthAuthorization) string {
	switch user := auth.(type) {
	case *telegram.AuthAuthorizationObj:
		if u, ok := user.User.(*telegram.UserObj); ok {
			name := u.FirstName
			if u.LastName != "" {
				name += " " + u.LastName
			}
			if u.Username != "" {
				name += fmt.Sprintf(" (@%s)", u.Username)
			}
			return name
		}
	}
	return "User"
}
