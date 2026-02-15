# Telegram String Session Generator (Gogram)

Eh tool Telegram string session generate karda hai using Gogram library. Tusi eh session apne music bot ya kise vi Telegram automation project vich use kar sakde ho.

## Prerequisites

1. **Go installed hona chahida** (version 1.21 ya newer)
   - Download: https://go.dev/download/

2. **Telegram API Credentials**
   - my.telegram.org te jao
   - Login karo apne phone number naal
   - "API development tools" te click karo
   - Ek navi app create karo
   - API ID te API Hash save kar lo

## Installation & Setup

### Step 1: Dependencies Install karo

```bash
go mod download
```

### Step 2: Program Run karo

```bash
go run session_generator.go
```

## Usage

Program run karan te eh info maangi jayegi:

1. **API ID**: Tusada API ID (number)
2. **API Hash**: Tusada API Hash (string)
3. **Phone Number**: Tusada phone number with country code (e.g., +911234567890)
4. **OTP**: Jo code tusanu SMS ya Telegram app vich milega
5. **2FA Password** (agar enabled hai): Tusada two-factor authentication password

## Output

Program tusanu ek string session generate kar ke dega jo kuch aisa hoga:

```
1BVtsOIABu6...longstring...xyz
```

Eh session tusade account di complete access provide karda hai, so **iski safe rakho aur kise naal share nahi karo!**

## Apne Music Bot vich Use karo

Generated session string nu apne bot di config vich add karo:

```python
# Python (Pyrogram/Telethon) example
STRING_SESSION = "1BVtsOIABu6...your_session_here..."
```

```go
// Go (Gogram) example
client, err := telegram.NewClient(telegram.ClientConfig{
    AppID:   apiID,
    AppHash: apiHash,
    Session: telegram.NewSessionFromString(stringSession),
})
```

## Security Warning ⚠️

- **String session KABHI share nahi karna**
- Eh session kise nu vi tusade account di full access de sakda hai
- Session nu private rakho aur secure jagah store karo
- Git repository vich commit nahi karna!

## Troubleshooting

### "Error connecting"
- Apna internet connection check karo
- API ID aur API Hash sahi check karo

### "Error with 2FA password"
- Password sahi type karo (case-sensitive hai)
- Agar bhul gaye ho, Telegram app vich jaa ke reset kar sako

### "Error sending code"
- Phone number format check karo (country code zaroori hai, e.g., +91)
- Koi special characters ya spaces nahi hone chahide

## Dependencies

- [Gogram](https://github.com/AmarnathCJD/gogram) - Telegram MTProto API library for Go

## License

This tool is for educational purposes. Use responsibly!
