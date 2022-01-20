package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// path参数 (URI)参数

func main() {
	r := gin.Default()
	// go 1.16.7 已经处理冲突了 所以不会冲突
	r.GET("/user/:name/:age", func(c *gin.Context) {
		// 获取路劲参数
		name := c.Param("name")
		age := c.Param("age")
		c.JSON(http.StatusOK, gin.H{
			"name": name,
			"age":  age,
		})
	})
	r.GET("/blog/:year/:month", func(c *gin.Context) {
		year := c.Param("year")
		month := c.Param("month")
		c.JSON(http.StatusOK, gin.H{
			"year":  year,
			"month": month,
		})
	})
	r.Run(":9090")
}
