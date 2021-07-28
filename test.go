// hello60
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)


func main() {
    r := gin.Default()
	r.Use(gin.Logger())
    r.GET("/ping", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "pong",
        })
    })
	r.GET("/user/:name",func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Hello %s", name)
	})
	r.GET("/overall",func(c *gin.Context) {
		c.JSON(http.StatusOK,gin.H{
			"online":1200,
		})
	})
    r.Run(":9000") // 預設監聽本地8080埠，如果需要更改可以使用 r.Run(":9000")
}
