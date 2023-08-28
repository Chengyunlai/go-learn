package main

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"top.chengyunlai/go-learn/2-day/4-topic/res"
	"top.chengyunlai/go-learn/2-day/4-topic/service"
)

func main() {
	r := gin.Default()
	r.GET("/community/page/get/:id", func(c *gin.Context) {
		// topicId是路径中的参数，但是返回的是一个字符串
		topicId := c.Param("id")
		data := QueryPageInfo(topicId)
		c.JSON(data.Code, data)

	})
	r.Run() // 监听并在 0.0.0.0:8080 上启动服务
}

func QueryPageInfo(topicId string) *res.R {
	id, err := strconv.ParseInt(topicId, 10, 64)
	if err != nil {
		return res.Error()
	}
	info, err := service.QueryPageInfo(id)
	if err != nil {
		return res.Error()
	}
	return res.Success(info)
}
