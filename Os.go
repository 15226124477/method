package method

import (
	"bufio"
	"fmt"
	"github.com/fsnotify/fsnotify"
	log "github.com/sirupsen/logrus"
	"io"
	"os"
	"path/filepath"
	"strings"
)

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

// GetFilesPath 获取文件夹下的全部文件
func GetFilesPath(root string, suffix string) ([]string, error) {
	var files []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		if filepath.Ext(path) == suffix {
			files = append(files, path)
		}

		return nil
	})
	if err != nil {
		return nil, err
	}
	return files, nil
}

// WriteFile 写入文件
func WriteFile(fPath string, lines []string) {
	file, err := os.OpenFile(fPath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		log.Warning(err)
		return
	}
	defer func(file *os.File) {
		err = file.Close()
		if err != nil {
			log.Error(err)
		}
	}(file)
	writer := bufio.NewWriter(file)
	_, err = writer.WriteString(strings.Join(lines, "\r\n"))
	if err != nil {
		log.Warning(err)
		return
	}
	err = writer.Flush()
	if err != nil {
		log.Warning(err)
		return
	}
}

func WriteNewFile(fpath string, in io.Reader, fm os.FileMode) error {
	err := os.MkdirAll(filepath.Dir(fpath), 0755)
	if err != nil {
		return fmt.Errorf("%s: making directory for file: %v", fpath, err)
	}
	out, err := os.Create(fpath)
	if err != nil {
		return fmt.Errorf("%s: creating new file: %v", fpath, err)
	}
	defer func(out *os.File) {
		_ = out.Close()
	}(out)
	err = out.Chmod(fm)
	_, err = io.Copy(out, in)
	if err != nil {
		return fmt.Errorf("%s: writing file: %v", fpath, err)
	}
	return nil
}

// Mkdir 创建一个文件夹
func Mkdir(mpath string) bool {
	log.Info(fmt.Sprintf("Local Path is Exist:%s", mpath))
	err := os.MkdirAll(mpath, 0755)
	go func() {
		Watch(mpath)
	}()
	if err != nil {
		log.Warning(fmt.Sprintf("Local Path is Error:%s", err))
		return false
	}
	return true
}

// IsPathExist 判断路径是否存在
func IsPathExist(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

// Mkdirs 创建多个文件夹
func Mkdirs(folders []string) {
	for i := 0; i < len(folders); i++ {
		Mkdir(folders[i])
	}
}
