#!/bin/bash

# Gogram Session Generator using Python
# Generates session that works directly with Gogram!

echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
echo "🔐 Gogram Session Generator (via Python)"
echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
echo ""
echo "This will generate a session that works with Gogram!"
echo ""

# Your credentials
API_ID="25742938"
API_HASH="b35b715fe8dc0a58e8048988286fc5b6"

echo "Choose method:"
echo "  1) Telethon (Recommended - Most compatible)"
echo "  2) Pyrogram (Also works)"
echo ""
read -p "Enter choice (1 or 2): " choice

if [ "$choice" = "1" ]; then
    echo ""
    echo "📦 Installing Telethon..."
    pip3 install telethon --break-system-packages --quiet 2>/dev/null || pip3 install telethon --quiet
    
    echo "✅ Telethon installed!"
    echo ""
    echo "📱 Generating Gogram-compatible session..."
    echo ""
    
    python3 << 'EOFTELETHON'
from telethon.sync import TelegramClient
from telethon.sessions import StringSession

# Credentials
api_id = 25742938
api_hash = "b35b715fe8dc0a58e8048988286fc5b6"

print("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
print("⏳ Connecting to Telegram...")
print("   Enter phone with country code: +91...")
print()

try:
    with TelegramClient(StringSession(), api_id, api_hash) as client:
        session_string = client.session.save()
        
        print()
        print("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
        print("✅ GOGRAM SESSION GENERATED!")
        print("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
        print()
        print("📝 Your STRING_SESSION (Gogram compatible):")
        print()
        print(session_string)
        print()
        print("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
        print()
        
        # Save to file
        with open('gogram_session.txt', 'w') as f:
            f.write(session_string)
        print("✅ Saved to: gogram_session.txt")
        print()
        
        # Update .env
        import os, re
        
        update = input("💾 Update .env file? (y/n): ").strip().lower()
        if update in ['y', 'yes']:
            if os.path.exists('.env'):
                with open('.env', 'r') as f:
                    content = f.read()
                
                if 'STRING_SESSION=' in content:
                    content = re.sub(
                        r'STRING_SESSION=.*',
                        f'STRING_SESSION={session_string}',
                        content
                    )
                else:
                    content += f'\nSTRING_SESSION={session_string}\n'
                
                with open('.env', 'w') as f:
                    f.write(content)
                
                print("✅ .env updated!")
            else:
                print("⚠️  .env not found!")
        
        print()
        print("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
        print("✅ Done! Session works with Gogram!")
        print("   Run: ./shizumusic")
        print("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")

except Exception as e:
    print(f"❌ Error: {e}")
EOFTELETHON

elif [ "$choice" = "2" ]; then
    echo ""
    echo "📦 Installing Pyrogram..."
    pip3 install pyrogram tgcrypto --break-system-packages --quiet 2>/dev/null || pip3 install pyrogram tgcrypto --quiet
    
    echo "✅ Pyrogram installed!"
    echo ""
    echo "📱 Generating Gogram-compatible session..."
    echo ""
    
    python3 << 'EOFPYROGRAM'
from pyrogram import Client
import os, re

# Credentials
api_id = 25742938
api_hash = "b35b715fe8dc0a58e8048988286fc5b6"

print("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
print("⏳ Connecting to Telegram...")
print("   Enter phone with country code: +91...")
print()

try:
    with Client("temp", api_id, api_hash) as app:
        session_string = app.export_session_string()
        
        print()
        print("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
        print("✅ GOGRAM SESSION GENERATED!")
        print("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
        print()
        print("📝 Your STRING_SESSION (Gogram compatible):")
        print()
        print(session_string)
        print()
        print("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
        print()
        
        # Save to file
        with open('gogram_session.txt', 'w') as f:
            f.write(session_string)
        print("✅ Saved to: gogram_session.txt")
        print()
        
        # Update .env
        update = input("💾 Update .env file? (y/n): ").strip().lower()
        if update in ['y', 'yes']:
            if os.path.exists('.env'):
                with open('.env', 'r') as f:
                    content = f.read()
                
                if 'STRING_SESSION=' in content:
                    content = re.sub(
                        r'STRING_SESSION=.*',
                        f'STRING_SESSION={session_string}',
                        content
                    )
                else:
                    content += f'\nSTRING_SESSION={session_string}\n'
                
                with open('.env', 'w') as f:
                    f.write(content)
                
                print("✅ .env updated!")
            else:
                print("⚠️  .env not found!")
        
        print()
        print("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
        print("✅ Done! Session works with Gogram!")
        print("   Run: ./shizumusic")
        print("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
    
    # Cleanup
    if os.path.exists("temp.session"):
        os.remove("temp.session")

except Exception as e:
    print(f"❌ Error: {e}")
    if os.path.exists("temp.session"):
        os.remove("temp.session")
EOFPYROGRAM

else
    echo "❌ Invalid choice!"
    exit 1
fi

echo ""
