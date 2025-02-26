package method

import (
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
				if result[0:6] == "10.100" {
					req = result
				}
				if result[0:7] == "192.168" {
					req = result
				}
			}
		}
	}
	return req
}
