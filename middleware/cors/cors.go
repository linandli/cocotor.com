package cors

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lestrrat/go-file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"

	cocotorjwt "cocotor.com/middleware/jwt"
)

type conf struct {
	Logpath  string `yaml:"logpath"`
	Logname  string `yaml:"logname"`
	Loglevel string `yaml:"loglevel"`
}

func (c *conf) getConf() *conf {
	yamlFile, err := ioutil.ReadFile("conf.yaml")
	if err != nil {
		fmt.Println(err.Error())
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		fmt.Println(err.Error())
	}
	return c
}

func Cors() gin.HandlerFunc {
	var c conf
	myConf := c.getConf()
	logClient := logrus.New()

	//禁止logrus的输出  
	src, err := os.OpenFile(os.DevNull, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		fmt.Println("err", err)
	}
	logClient.Out = src
	logClient.SetLevel(logrus.DebugLevel)
	baseLogPath := path.Join(myConf.Logpath, myConf.Logname)
	logWriter, err := rotatelogs.New(
		baseLogPath+".%Y-%m-%d-%H-%M.log",
		rotatelogs.WithLinkName(baseLogPath),      // 生成软链，指向最新日志文件
		rotatelogs.WithMaxAge(7*24*time.Hour),     // 文件最大保存时间
		rotatelogs.WithRotationTime(24*time.Hour), // 日志切割时间间隔
	)
	writeMap := lfshook.WriterMap{
		logrus.InfoLevel:  logWriter,
		logrus.FatalLevel: logWriter,
		logrus.WarnLevel:  logWriter,
		logrus.ErrorLevel: logWriter,
		logrus.PanicLevel: logWriter,
		logrus.TraceLevel: logWriter,
	}
	lfHook := lfshook.NewHook(writeMap, &logrus.JSONFormatter{})
	logClient.AddHook(lfHook)

	return func(c *gin.Context) {
		method := c.Request.Method               //请求方法
		origin := c.Request.Header.Get("Origin") //请求头部
		var headerKeys []string                  // 声明请求头keys
		for k, _ := range c.Request.Header {
			headerKeys = append(headerKeys, k)
		}
		headerStr := strings.Join(headerKeys, ", ")
		if headerStr != "" {
			headerStr = fmt.Sprintf("access-control-allow-origin, access-control-allow-headers, %s", headerStr)
		} else {
			headerStr = "access-control-allow-origin, access-control-allow-headers"
		}
		if origin != "" {
			c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
			c.Header("Access-Control-Allow-Origin", "*")                                       // 这是允许访问所有域
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE") //服务器支持的所有跨域请求的方法,为了避免浏览次请求的多次'预检'请求
			//  header的类型
			c.Header("Access-Control-Allow-Headers", "Authorization, Content-Length, X-CSRF-Token, Token,session,X_Requested_With,Accept, Origin, Host, Connection, Accept-Encoding, Accept-Language,DNT, X-CustomHeader, Keep-Alive, User-Agent, X-Requested-With, If-Modified-Since, Cache-Control, Content-Type, Pragma")
			// 允许跨域设置   可以返回其他字段
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers,Cache-Control,Content-Language,Content-Type,Expires,Last-Modified,Pragma,FooBar") // 跨域关键设置 让浏览器可以解析
			c.Header("Access-Control-Max-Age", "172800")                                                                                                                                                           // 缓存请求信息 单位为秒
			c.Header("Access-Control-Allow-Credentials", "false")                                                                                                                                                  //  跨域请求是否需要带cookie信息 默认设置为true
			c.Set("content-type", "application/json")                                                                                                                                                              // 设置返回格式是json
		}

		data, err := c.GetRawData()
		if err != nil {
			fmt.Println(err.Error())
		}

		postData := string(data)
		//把读过的字节流重新放到body
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(data))

		//放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.JSON(http.StatusOK, "Options Request!")
		}
		// 处理请求
		// 开始时间
		start := time.Now()
		// 处理请求
		c.Next()
		// 结束时间
		end := time.Now()
		//执行时间
		latency := end.Sub(start)
		path := c.Request.URL.Path
		clientIP := c.ClientIP()
		statusCode := c.Writer.Status()
		rawQuery := c.Request.URL.RawQuery

		if c.Request.Header.Get("token") != "" {
			j := &cocotorjwt.JWT{
				[]byte("newtrekTorch"),
			}

			claims, err := j.ParseToken(c.Request.Header.Get("token"))
			if err == nil {
				phone := claims.Phone
				name := claims.Name
				logClient.Infof("| %3d | %13v | %15s | %s  %s ? %s| %s : %s | data: %s |",
					statusCode,
					latency,
					clientIP,
					method,
					path,
					rawQuery,
					name,
					phone,
					postData,
				)
			} else {
				logClient.Infof("| %3d | %13v | %15s | %s  %s ? %s| token: %s |",
					statusCode,
					latency,
					clientIP,
					method,
					path,
					rawQuery,
					c.Request.Header.Get("token"),
				)
			}
		} else {
			logClient.Infof("| %3d | %13v | %15s | %s  %s ? %s|  data: %s |",
				statusCode,
				latency,
				clientIP,
				method,
				path,
				rawQuery,
				postData,
			)
		}
	}
}
