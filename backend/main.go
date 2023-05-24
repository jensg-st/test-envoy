package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/*dd", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("%s%s", c.Request.Host+c.Request.URL.Path),
		})
	})
	r.Run(":2555") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
