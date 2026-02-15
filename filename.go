package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	tg "github.com/amarnathcjd/gogram/telegram"
)

var (
	appID   int32  = 25742938
	appHash string = "b35b715fe8dc0a58e8048988286fc5b6"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("📞 Enter phone number (+91...): ")
	phone, _ := reader.ReadString('\n')
	phone = strings.TrimSpace(phone)

	// USER CLIENT
	user, err := tg.NewClient(tg.ClientConfig{
		AppID:   appID,
		AppHash: appHash,
	})
	if err != nil {
		log.Fatal(err)
	}

	if err := user.Connect(); err != nil {
		log.Fatal(err)
	}

	// ✅ SEND OTP
	sentCode, err := user.AuthSendCode(
		phone,
		appID,
		appHash,
		&tg.CodeSettings{},
	)
	if err != nil {
		log.Fatal("SendCode failed:", err)
	}

	fmt.Print("🔑 Enter OTP: ")
	otp, _ := reader.ReadString('\n')
	otp = strings.TrimSpace(otp)

	// ✅ SIGN IN — THIS IS THE KEY PART
	if err := sentCode.SignIn(user, otp); err != nil {
		log.Fatal("Login failed:", err)
	}

	// ✅ EXPORT SESSION
	session := user.ExportSession()

	fmt.Println("\n✅ GOGRAM STRING SESSION:\n")
	fmt.Println(session)
}
