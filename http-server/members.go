package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (server *Server) joinnedChannel(c *gin.Context) {
	username := c.Param("username")
	if username == "me" {
		username = c.MustGet(keyUsername).(string)
	}
	channelIDs, err := server.backend.JoinnedChannel(username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, channelIDs)
		return
	}
	c.JSON(http.StatusOK, channelIDs)
}
func (server *Server) invite(c *gin.Context) {
	channelID := c.Param("channelID")
	username := c.Param("username")

	err := server.backend.Join(channelID, username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"msg": "ok"})
}

func (server *Server) kick(c *gin.Context) {
	channelID := c.Param("channelID")
	username := c.Param("username")

	err := server.backend.Leave(channelID, username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"msg": "ok"})
}

func (server *Server) joinChannel(c *gin.Context) {
	channelID := c.Param("channelID")
	username := c.MustGet(keyUsername).(string)

	err := server.backend.Join(channelID, username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"msg": "ok"})
}
