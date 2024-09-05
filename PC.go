package method

import (
	"github.com/fsnotify/fsnotify"
	log "github.com/sirupsen/logrus"
	"io"
	"net"
	"net/http"
	"os"
)

func IsURLAccessible(url string) bool {
	// 发送HTTP GET请求
	resp, err := http.Get(url)
	if err != nil {
		log.Warning(err)
		return false
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK {
		return true
	}
	return false
}

func HttpDownload(url string, filePath string) {
	resp, err := http.Get(url)
	if err != nil {
		log.Debug("Error downloading file:", err)
		return
	}
	defer resp.Body.Close()
	outFile, err := os.Create(filePath)
	if err != nil {
		log.Debug("Error creating file:", err)
		return
	}
	defer outFile.Close()
	_, err = io.Copy(outFile, resp.Body)
	if err != nil {
		log.Debug("Error saving file:", err)
		return
	}
	log.Info("File downloaded successfully:", filePath)
}

func WorkIP() string {
	ips, err := net.InterfaceAddrs()
	if err != nil {
		log.Error(err)
		return ""
	}
	req := ""
	for _, address := range ips {
		if pc, ok := address.(*net.IPNet); ok && !pc.IP.IsLoopback() {
			if pc.IP.To4() != nil {
				result := pc.IP.String()
				if result[0:6] == "172.16" {
					req = result
				}
			}
		}
	}
	return req
}

func Watch(folder string) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Error("NewWatcher failed: ", err)
	}
	defer func(watcher *fsnotify.Watcher) {
		err = watcher.Close()
		if err != nil {
			log.Error(err)
		}
	}(watcher)
	done := make(chan bool)
	go func() {
		defer close(done)
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				log.Warning(event.String())
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Error("error:", err)
			}
		}
	}()

	err = watcher.Add(folder)
	if err != nil {
		log.Fatal("Add failed:", err)
	}
	<-done
}
