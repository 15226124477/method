package method

import (
	"fmt"
	"github.com/15226124477/define"
	"github.com/hpcloud/tail"
	log "github.com/sirupsen/logrus"
	"net"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"
)

// TCP Server端测试
func process(conn net.Conn, line chan *tail.Line) {
	defer conn.Close() //
	for l := range line {
		conn.Write([]byte(l.Text + "\n")) // 发送数据
		fmt.Println(l.Time.Format("2006-01-02 15:04:05.000"), "LOG: Send Log Length:", len(l.Text))
	}
}

func LiveServer() {
	port := ""
	for i := define.LogTcpServerPortStart; i < define.LogTcpServerPortEnd; i++ {
		tmp, err := net.Listen("tcp", fmt.Sprintf("%s:%d", WorkIP(), i))
		if err != nil {
			log.Error("Listen() failed, err: ", err)
			continue
		} else {
			port = fmt.Sprintf("%s:%d", WorkIP(), i)
			tmp.Close()
			break
		}
	}
	log.Warning("Log Live Server is :", port)
	listen, err := net.Listen("tcp", port)
	if err != nil {
		log.Info("Listen() failed, err: ", err)
		return
	}
	runExe, _ := os.Executable()
	_, exec := filepath.Split(runExe)
	filenameWithSuffix := path.Base(exec)
	fileSuffix := path.Ext(filenameWithSuffix)
	filenameOnly := strings.TrimSuffix(filenameWithSuffix, fileSuffix)
	filePath := fmt.Sprintf("./log/%s%s.log", filenameOnly, time.Now().Format("2006-01-02"))
	log.Warning(filePath)
	// 配置tail
	config := tail.Config{
		ReOpen:    true,                                 // 文件被截断后重新打开
		Follow:    true,                                 // 跟随文件，监控新增内容
		Location:  &tail.SeekInfo{Offset: 0, Whence: 2}, // 从文件末尾开始读取
		MustExist: false,                                // 文件不存在不报错
		Poll:      true,                                 // 使用轮询模式
	}
	// 使用tail.TailFile创建一个Tail对象
	tails, err := tail.TailFile(filePath, config)
	if err != nil {
		log.Error(err)
	}

	for {
		conn, err := listen.Accept() // 监听客户端的连接请求
		if err != nil {
			log.Error("Accept() failed, err: ", err)
			continue
		}
		log.Warning("client connect...")
		go process(conn, tails.Lines) // 启动一个goroutine来处理客户端的连接请求
	}

}
