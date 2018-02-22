package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (server *Server) getChannels(c *gin.Context) {
	channels, err := server.backend.GetChannels()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, channels)
}
func (server *Server) createChannel(c *gin.Context) {
	req := &struct {
		Name   string
		Labels map[string]string
	}{}
	if err := c.ShouldBind(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		c.Abort()
		return
	}
	channel, err := server.backend.CreateChannel(req.Name, req.Labels)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, channel)
}
