package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// gin框架返回json格式的数据
func main() {
	r := gin.Default()
	r.GET("/json", func(c *gin.Context) {
		// 方法1 : 使用map
		//data := map[string]interface{}{
		//	"name" : "小王子",
		//	"message" : "hello world",
		//	"age" : 18,
		//}
		data2 := gin.H{
			"name":    "小王子",
			"message": "hello world",
			"age":     18,
		}
		c.JSON(http.StatusOK, data2)
	})
	// 方法2 : 结构体, 灵活使用tag来对结构体字段做定制化操作

	type msg struct {
		Name    string `json:"name"` // 返回json格式的文件是通过json调用反射来获取字段的 所以首字母大写
		Age     int    `json:"age"`
		Message string `json:"message"`
	}
	r.GET("/struct/json", func(c *gin.Context) {
		data := msg{
			"小王子",
			18,
			"hello golang",
		}
		c.JSON(http.StatusOK, data)
	})
	r.Run(":9000")
}
