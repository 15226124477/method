package method

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"path"
	"runtime"
)

// getCallerInfo 获取回调信息
func getCallerInfo(skip int) (info string) {
	pc, file, lineNo, ok := runtime.Caller(skip)
	if !ok {
		info = "runtime.Caller() failed"
		return
	}
	funcName := runtime.FuncForPC(pc).Name()
	fileName := path.Base(file) // Base函数返回路径的最后一个元素
	return fmt.Sprintf("FuncName:%s, file:%s, line:%d ", funcName, fileName, lineNo)
}

// ReturnStruct 打印结构体
func ReturnStruct(data interface{}) {
	res, _ := json.MarshalIndent(data, "", "    ")
	key := getCallerInfo(2)
	log.Debug(key, "\n", string(res))
}

// PostReqConvert 打印请求
func PostReqConvert(raw []byte, js interface{}, result interface{}) {
	key := getCallerInfo(2)
	log.Debug(key)
	res1, _ := json.MarshalIndent(string(raw), "", "    ")
	res2, _ := json.MarshalIndent(js, "", "    ")
	res3, _ := json.MarshalIndent(result, "", "    ")
	log.Info("Raw Text|\n", string(res1))
	log.Debug("JS  Text|\n", string(res2))
	log.Warning("Go  Text|\n", string(res3))
}

// GinLog 设置Gin日志格式
func GinLog() gin.HandlerFunc {
	return gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		// 自定义格式
		// token, num := AESDecrypt(param.Request.Header.Get("X-Access-Token"))
		fmt.Println(fmt.Sprintf("%s WEB: http://%s%-22s | %-20s:%-5s | StatusCode:%d Latency:%s %s %s \n",
			param.TimeStamp.Format("2006-01-02 15:04:05.000"),
			param.Request.Host,
			param.Path,
			param.ClientIP,
			param.Method,
			param.StatusCode,
			param.Latency,
			param.Request.Header.Get("X-Access-Token"),
			param.ErrorMessage,
		))
		return fmt.Sprintf("%s WEB: http://%s%-22s | %-20s:%-5s | StatusCode:%d Latency:%s %s %s \n",
			param.TimeStamp.Format("2006-01-02 15:04:05.000"),
			param.Request.Host,
			param.Path,
			param.ClientIP,
			param.Method,
			param.StatusCode,
			param.Latency,
			param.Request.Header.Get("X-Access-Token"),
			param.ErrorMessage,
		)

	})

}

// Cors HTTP请求的CORS设置
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin")
		// account := user.UuidFind(c.Request.Header.Get("X-Access-Token"))
		// start.Info(fmt.Sprintf("User|真实姓名:%s 账号权限:%s 账号名称:%s", account.RealName, account.Role, account.UserName))
		if origin != "" {
			c.Header("Access-Control-Allow-Origin", origin)
			c.Header("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE, OPTIONS")
			c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length,Token,x-access-token")
			c.Header("Access-Control-Expose-Headers", "Access-Control-Allow-Headers, Token")
			// c.Header("Access-Control-Max-Age", "172800")
			c.Header("Access-Control-Allow-Credentials", "true")
			c.Set("content-type", "application/json")
		}
		if method == "OPTIONS" {
			//c.JSON(200, Controller.R(200, nil, "Options Request"))
			c.AbortWithStatus(http.StatusNoContent)
		}
		c.Next()

	}
}

// GetSelfLocal 获取程序当前位置
func GetSelfLocal() string {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	// start.Warning("当前程序所在位置:", ex)
	return ex
}
