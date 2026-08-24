package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type NodeController struct{}

var nodeHub = NewNodeHub()

func (NodeController) Health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "ok",
		"service": "openhlas-node",
	})
}

func (NodeController) Info(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"service":  "openhlas-node",
		"protocol": "websocket",
		"endpoint": "/ws",
	})
}

func (NodeController) WebSocket(c *gin.Context) {
	nodeHub.HandleConnection(c.Writer, c.Request)
}
