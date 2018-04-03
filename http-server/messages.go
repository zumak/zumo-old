package server

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (server *Server) getMessage(c *gin.Context) {
	messages, err := server.backend.GetMessages(c.Param("channelID"), 100)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, messages)
}
func (server *Server) postMessage(c *gin.Context) {
	req := &struct {
		Text   string
		Detail json.RawMessage
	}{}
	if err := c.ShouldBind(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		c.Abort()
		return
	}
	username := c.MustGet(keyUsername).(string)
	msg, err := server.backend.AppendMessage(username, c.Param("channelID"), req.Text, req.Detail)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, msg)
}
