package common

import (
	"net/http"
	"time"
	"log"

	jwtgo "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"

	cocotorjwt "cocotor.com/middleware/jwt"
	"cocotor.com/model"
)

// Login 登录
func Login(c *gin.Context) {
	var loginReq model.LoginReq

	if c.BindJSON(&loginReq) == nil {
		log.Println("用户：", loginReq.Phone, "开始登陆！")
		isPass, user, err := model.LoginCheck(loginReq)

		if isPass {
			log.Println("用户：", loginReq.Phone, "登陆成功！")	
			generateToken(c, user)
		} else {
			c.JSON(http.StatusOK, gin.H{
				"status": -1,
				"msg": "验证失败," + err.Error(),
			})
		}
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": -1,
			"msg":    "登录数据格式有误",
		})
	}
}


// 生成令牌
func generateToken(c *gin.Context, user model.User) {
	j := &cocotorjwt.JWT{
		[]byte("newtrekWang"),
	}
	claims := cocotorjwt.CustomClaims{
		user.Id,
		user.Name,
		user.Phone,
		jwtgo.StandardClaims{
			NotBefore: int64(time.Now().Unix() - 1000), // 签名生效时间
			ExpiresAt: int64(time.Now().Unix() + 3600), // 过期时间 一小时
			Issuer: "newtrekWang", //签名的发行者
		},
	}

	token, err := j.CreateToken(claims)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": -1,
			"msg":    err.Error(),
		})
		return
	}

	data := model.LoginResult{
		Token: token,
	}
	c.JSON(http.StatusOK, gin.H{
		"status": 0,
		"msg":    "登录成功！",
		"data":   data,
	})
	return
}