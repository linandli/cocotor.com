package common

import (
	"net/http"
	"log"
	"github.com/gin-gonic/gin"
	"cocotor.com/model"
)

// Register 登录
func Register(c *gin.Context) {
	var registerInfo model.RegistInfo

	if c.BindJSON(&registerInfo) == nil {
		err := model.Register(registerInfo.Name, registerInfo.Phone, 
			registerInfo.Pwd, registerInfo.Gender, registerInfo.Permission)
		if err == nil {
			log.Println("用户：", registerInfo.Name, " 注册号码：", registerInfo.Phone, " 注册成功！")
			c.JSON(http.StatusOK, gin.H{
				"status": 0,
				"msg": "注册成功！",
			})
		} else {
			log.Println("用户：", registerInfo.Name, " 注册号码：", registerInfo.Phone, " 注册失败！")
			c.JSON(http.StatusOK, gin.H{
				"status": -1,
				"msg": "注册失败 " + err.Error(),
			})
		}
	} else {
		log.Println("注册使用数据格式异常！")
		c.JSON(http.StatusOK, gin.H{
			"status": -1,
			"msg": "解析数据失败！",
		})
	}
}
