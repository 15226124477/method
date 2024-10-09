package method

import (
	"bufio"
	log "github.com/sirupsen/logrus"
	"io"
	"os"
	"path/filepath"
)

func Move(from, to string) error {
	err := Copy(from, to)
	if err == nil {
		err1 := os.RemoveAll(from)
		if err1 != nil {
			log.Debug("Copy Success,DeL File Error!!!")
			return err1
		} else {
			log.Debug("Move Success：", from, " => ", to)
		}
	} else {
		log.Debug("Move Error：", from, " => ", to)
	}

	return err
}

func Copy(from, to string) error {
	f, e := os.Stat(from)
	if e != nil {
		log.Debug("No Found the path：", from)
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
