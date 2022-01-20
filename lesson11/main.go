package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// lesson11 中间件

func indexHandler(c *gin.Context) {
	fmt.Println("index")
	name, ok := c.Get("name")
	if !ok {
		name = "不存在的用户"
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": name,
	})
}

// 定义一个中间件

func m1(c *gin.Context) {
	fmt.Println("m1 in...")
	// 计时
	start := time.Now()
	c.Next() // 调用后续的函数
	// c.Abort() 阻止调用后续的函数
	cost := time.Since(start)
	fmt.Printf("cost : %#v\n", cost)
	fmt.Println("m1 out ... ")
}
func m12(c *gin.Context) {
	fmt.Println("m12 in...")
	c.Set("name", "qimi")
	//c.Next() //调用后续的函数
	//c.Abort() // 阻止后续的函数
	//return
	fmt.Println("m12 out ... ")
}

func authMiddleware(doCheck bool) gin.HandlerFunc {
	// 链接数据库
	// 做一些其他的准备工作
	return func(c *gin.Context) {
		if doCheck {

		} else {
			c.Next()
		}
	}
	// 是否登录的判断
	// if是登录用户
	// c.Next()
	// else
	// c.Abort()
}

func main() {
	r := gin.Default() // 默认使用了logger和recovery中间件
	// 不想使用了的话 r := gin.New()
	r.Use(m1, m12, authMiddleware(true)) // 全局注册中间件函数 以后都不用在GET请求中写m1了

	// GET请求进来后先走m1中间件函数,再走indexHandler 比如这个视频只有vip才能查看,那么我就要先走这个检查你是否为vip身份的中间件函数
	r.GET("/index", indexHandler)
	r.Run(":8080")
}
