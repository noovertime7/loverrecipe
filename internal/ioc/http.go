package ioc

import (
	"github.com/gin-gonic/gin"
	"github.com/gotomicro/ego/server/egin"
)

func InitHTTP() *egin.Component {
	server := egin.Load("server.http").Build()

	server.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	return server
}
