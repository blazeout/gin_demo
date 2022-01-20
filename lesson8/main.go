package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.LoadHTMLFiles("lesson8/index.html")
	r.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	// 点击上传 数据会从前端往后端/upload的地方发送post的请求
	// 这里需要处理请求
	r.POST("/upload", func(c *gin.Context) {
		// 从请求中读取到文件,并且保存在服务端的本题
		file, err := c.FormFile("f1")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}
		filePath := fmt.Sprintf("D:\\GOfile\\src\\gin_demo\\lesson8\\%s", file.Filename)
		err = c.SaveUploadedFile(file, filePath)
		if err != nil {
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"statics": "ok",
		})
	})
	r.Run(":8080")
}
