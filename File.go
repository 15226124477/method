package method

import (
	"bufio"
	log "github.com/sirupsen/logrus"
	"os"
	"strings"
)

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
	}(file) // 在main函数结束时关闭文件
	// 准备写入器，这里使用bufio来提高写入效率
	writer := bufio.NewWriter(file)
	// 写入字符串到文件
	_, err = writer.WriteString(strings.Join(lines, "\r\n"))
	if err != nil {
		log.Warning(err)
		return
	}
	// 刷新缓冲区，确保所有数据都被写入到文件中
	err = writer.Flush()
	if err != nil {
		log.Warning(err)
		return
	}

}
