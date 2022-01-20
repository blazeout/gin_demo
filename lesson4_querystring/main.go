package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// querystring
func main() {
	r := gin.Default()
	// GET请求URL?后面的是querystring参数
	// 使用key=value的格式,多个key-value用&连接
	r.GET("/web_query", func(c *gin.Context) {
		// 获取浏览器那边发请求携带的query string 参数
		name := c.Query("query") // 通过Query获取请求中携带的querystring参数
		age := c.Query("age")
		//name := c.DefaultQuery("query","somebody") //取不到就使用默认值
		//name, ok := c.GetQuery("query")
		//if !ok {
		//	// 取不到name就等于somebody
		//	name = "somebody"
		//}
		c.JSON(http.StatusOK, gin.H{
			"age":  age,
			"name": name,
		})
	})

	r.Run(":9090")
}
