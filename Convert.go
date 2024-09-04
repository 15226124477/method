package method

import (
	"math"
	"strconv"
)

const (
	DEGREES = 1
	RADIAN  = 2

	// GGASOL = 3
	// POSSOL = 4
)

// Decimal 浮点数保留N位小数
func Decimal(value float64, prec int) float64 {
	value, _ = strconv.ParseFloat(strconv.FormatFloat(value, 'f', prec, 64), 64)
	return value
}

/*
	// Degrees2Radians 度转弧度
	func Degrees2Radians(degrees float64) float64 {
		return degrees * math.Pi / 180
	}
*/
// Radians2Degrees 弧度转度
func Radians2Degrees(radians float64) float64 {
	return radians * 180 / math.Pi
}

/*
	func RemoveRepeatedElement(s []string) []string {
		result := make([]string, 0)
		m := make(map[string]bool) //map的值不重要
		for _, v := range s {
			if _, ok := m[v]; !ok {
				result = append(result, v)
				m[v] = true
			}
		}
		return result
	}
*/

/*
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

	func AbleExists(mpath string) bool {
		start.Info(fmt.Sprintf("确保路径存在:%s", mpath))
		err := os.MkdirAll(mpath, 0755)
		if err != nil {
			start.Warning(fmt.Sprintf("路径异常:%s", err))
			return false
		}
		return true
	}


	func PathExists(path string) bool {
		_, err := os.Stat(path)
		if err == nil {
			return true
		}
		//is not exist来判断，是不是不存在的错误
		if os.IsNotExist(err) { //如果返回的错误类型使用os.isNotExist()判断为true，说明文件或者文件夹不存在
			return false
		}
		return false //如果有错误了，但是不是不存在的错误，所以把这个错误原封不动的返回
	}



	func IsContain(items []string, item string) bool {
		for _, eachItem := range items {
			if eachItem == item {
				return true
			}
		}
		return false
	}


	func Copy(from, to string) error {
		f, e := os.Stat(from)
		if e != nil {
			start.Debug("没有找到要拷贝的路径：", from)
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
					start.Error(err)
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
					start.Error(err)
				}
			}(out)
			// 然后将文件流和文件流对接起来
			_, e = io.Copy(out, bufReader)
		}
		return e
	}


	func Move(from, to string) error {
		err := Copy(from, to)
		if err == nil {
			err1 := os.RemoveAll(from)
			if err1 != nil {
				start.Debug("复制完成，删除失败")
				return err1
			} else {
				start.Debug("移动成功：", from, " => ", to)
			}
		} else {
			start.Debug("移动失败：", from, " => ", to)
		}

		return err
	}
*/
