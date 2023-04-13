<p align="center">
	<a href="https://github.com/boz3r/Hellcat">
		<img src="assets/Hellcat.png" alt="fatebot" width="400" height="400">
	</a>
	<br>
	<a href="https://github.com/boz3r/Hellcat/blob/master/LICENSE">
		<img src="https://img.shields.io/badge/license-MIT license-orange?style=plastic">
	</a>
	<a href="https://github.com/boz3r/Hellcat">
    		<img src="https://img.shields.io/badge/version-v1.1.0-black?style=plastic">
	</a>
	<a href="https://go.dev/">
    		<img src="https://img.shields.io/badge/language-Go-orange?style=plastic">
	</a>
	<a href="https://www.microsoft.com/en-gb/software-download/windows10ISO">
    		<img src="https://img.shields.io/badge/platform-windows-black?style=plastic">
	</a>
  	</br>
</p>

<p align="center">
	<b><ins>⚠️ DISCLAIMER ⚠️</ins></b>
	<br>
		I create this project for education purposes only, the use of this software is your responsibility!!!
	<br>
</p>

---

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

Hellcat is a backdoor built to infect Windows machines. It uses Telegram as a C2 server to communicate between the attacker and the client. However, it can only set one Telegram bot API per payload. This means that if you want to infect another machine, you need to build a new payload with a new Telegram bot API. It is inspired by the <a href="https://github.com/byt3bl33d3r/gcat">gcat</a> and <a href="https://github.com/maldevel/gdog">gdog</a> backdoors.

## Build
To build the payload, you can just <b>double-click</b> the batch file. In case you want to pack it, you can move your <a href="https://upx.github.io/">UPX</a> to the path that's already set in the batch file or just specify the path yourself. If you don't want to pack the payload, then just comment it out.

<img src="assets/build.gif" alt="fatebot" width="700" height="410">

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
	<li>UPX Packer</li>
</ul>

<ul>
	<li>Code\Text Editor</li>
</ul>
