package method

import (
	"bufio"
	"fmt"
	log "github.com/sirupsen/logrus"
	"io"
	"os"
	"path/filepath"
	"strings"
)

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

func PathExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

func Move(from, to string) error {
	err := Copy(from, to)
	if err == nil {
		err1 := os.RemoveAll(from)
		if err1 != nil {
			log.Debug("复制完成，删除失败")
			return err1
		} else {
			log.Debug("移动成功：", from, " => ", to)
		}
	} else {
		log.Debug("移动失败：", from, " => ", to)
	}

	return err
}

func AbleExists(mpath string) bool {
	log.Info(fmt.Sprintf("确保路径存在:%s", mpath))
	err := os.MkdirAll(mpath, 0755)
	if err != nil {
		log.Warning(fmt.Sprintf("路径异常:%s", err))
		return false
	}
	return true
}

func Copy(from, to string) error {
	f, e := os.Stat(from)
	if e != nil {
		log.Debug("没有找到要拷贝的路径：", from)
		return e
	}
	if f.IsDir() {
		//from是文件夹，那么定义to也是文件夹
		if list, e := os.ReadDir(from); e == nil {
			for _, item := range list {
				if e = Copy(filepath.Join(from, item.Name()), filepath.Join(to, item.Name())); e != nil {
					return e
				}
			}
		}
	} else {
		//from是文件，那么创建to的文件夹
		p := filepath.Dir(to)
		if _, e = os.Stat(p); e != nil {
			if e = os.MkdirAll(p, 0777); e != nil {
				return e
			}
		}
		//读取源文件
		file, e := os.Open(from)
		if e != nil {
			return e
		}
		defer func(file *os.File) {
			err := file.Close()
			if err != nil {
				log.Error(err)
			}
		}(file)
		bufReader := bufio.NewReader(file)
		// 创建一个文件用于保存
		out, e := os.Create(to)
		if e != nil {
			return e
		}
		defer func(out *os.File) {
			err := out.Close()
			if err != nil {
				log.Error(err)
			}
		}(out)
		_, e = io.Copy(out, bufReader)
	}
	return e
}

func CreatFolders(folders []string) {
	for i := 0; i < len(folders); i++ {
		err := os.MkdirAll(folders[i], 0755) // 创建多层文件夹
		go func() {
			log.Warning("Watching Folder:", folders[i])
			Watch(folders[i])
		}()
		if err != nil {
			log.Error(err)
			return
		}
	}
}
