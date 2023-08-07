## Hellcat

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

Hellcat is a backdoor built to infect Windows machines. It uses Telegram as a C2 server to communicate between the attacker and the client. However, it can only set one Telegram bot API per payloaf. This means that if you want to infect another machine, you need to build a new payload with a new Telegram bot API. It is inspired by the <a href="https://github.com/byt3bl33d3r/gcat">gcat</a> and <a href="https://github.com/maldevel/gdog">gdog</a> backdoors.

## Build
To build the payload, you can simply <b>double-click</b> the batch file. The binary file that is used to be implanted in the target machine will be located in the "bin" folder.

## Bot API
To build the Telegram bot API, you can go to <a href="https://t.me/botfather">BotFather</a> and ask for a new bot. Then, name your bot and choose a username for it, and you're done! If you don't know the command, you can type <b>"/help"</b> and it will show you a lot of useful commands.

## Requirement

<ul>
	<li>Telegram</li>
</ul>

<ul>
	<li>Go Compiler</li>
</ul>

<ul>
	<li>Code\Text Editor</li>
</ul>
