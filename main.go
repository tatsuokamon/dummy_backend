package main;

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type HelloWorld struct {
	Target string `json:"target"`;
	Now int64 `json:"now"`;
}

func main() {
	router := gin.Default();
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, HelloWorld {
			Target: "TatsuoKamon",
			Now: time.Now().UnixMilli(),
		})
	})

	router.Run(":80")
}
