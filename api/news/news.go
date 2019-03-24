package news

import (
	"net/http"
	"log"

	// jwtgo "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"	
	"github.com/globalsign/mgo/bson"
	db "cocotor.com/utils/mongodb"
	model_news "cocotor.com/model/news"
)

const(
	database   = "telespy"
	collection = "clues"
)

func News(c *gin.Context) {
	var resultWithPage [] model_news.NewData
	// querys := bson.M{"$or":[]bson.M{bson.M{"website": "中国政府网", "group": "政务联播-地方"}, bson.M{"website": "中国政府网", "group": "政务联播-部门"}, bson.M{"website": "中国政府网", "group": "新闻滚动"}, bson.M{"website": "中国政府网", "group": "新闻要闻"}, bson.M{"website": "云南 国家能源局云南监管办", "group": "监管信息"}, bson.M{"website": "云南国家能源局云南监管办", "group": "监管动态"}, bson.M{"website": "四川国家能源局四川监管办", "group": "新闻中心"}, bson.M{"website": "四川国家能源局四川监管办", "group": "监管动态"}, bson.M{"website": "国家发改委", "group": "发展改革委令"}, bson.M{"website": "国家发改委", "group": "规划文本"}, bson.M{"website": "国家发改委", "group": "规范性文件"}, bson.M{"website": "国家发改委", "group": "解读"}, bson.M{"website": "国家发改委", "group": "通知"}, bson.M{"website": "国家能源局", "group": "最新文件"}, bson.M{"website": "国家能源局", "group": "解读"}, bson.M{"website": "山东国家能源局山东监管办", "group": "重要通知"}, bson.M{"website": "广东国家能源局南方监督局", "group": "要闻动态"}, bson.M{"website": "广东国家能源局南方监督局", "group": "通知公告"}, bson.M{"website": "江苏国家能源局江苏监管办", "group": "通知"}, bson.M{"website": "河南国家能源局河南监管办", "group": "监管公告"}, bson.M{"website": "河南国家能源局河南监管办", "group": "通知"}, bson.M{"website": "浙江国家能源局浙江监管办", "group": "通知公告"}, bson.M{"website": "湖南国家能源局湖南监管办公室", "group": "公告公示"}, bson.M{"website": "湖南国家能源局湖南监管办公室", "group": "通知"}, bson.M{"website": "甘肃国家能源局甘肃监管办", "group": "监管动态"}, bson.M{"website": "福建国家能源局福建监管办", "group": "政策发布"}, bson.M{"website": "福建国家能源局福建监管办", "group": "通知公告"}}}
	querys := bson.M{"$or":[]bson.M{bson.M{"website": "中国政府网", "group": "政务联播-地方"},bson.M{"website": "福建国家能源局福建监管办", "group": "政策发布"}}}
	err := db.FindPageSort(database, collection, 0, 10, querys, nil,  &resultWithPage, "-publish_time")

	if err != nil {
		log.Fatal(err)
		c.JSON(200, gin.H{
			"status": -1,
			"message": "error!",
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"status": 1,
		"msg": resultWithPage,
	})
}