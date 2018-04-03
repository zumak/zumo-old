package server

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/zumak/zumo/datatypes"
	"golang.org/x/net/websocket"
)

type translater struct {
	encoder *json.Encoder
	decoder *json.Decoder
}

func (t *translater) OnMessage(channelID string, msg datatypes.Message) {
	t.encoder.Encode(struct {
		*datatypes.Message
		Type      string
		ChannelID string
	}{&msg, "message", channelID})
}
func (t *translater) Read() (string, datatypes.Message, error) {
	// Must block until end
	msg := struct {
		ChannelID string
		datatypes.Message
	}{}
	if err := t.decoder.Decode(&msg); err != nil {
		return "", datatypes.Message{}, err
	}
	return msg.ChannelID, msg.Message, nil
}
func (t *translater) channelChnaged(channelID string) {
	t.encoder.Encode(struct {
		Type string
		Name string
		Data map[string]string
	}{"event", "channel", map[string]string{"ID": channelID}})
}
func (t *translater) OnJoinChannel(channelID string) {
	// send to client
	t.channelChnaged(channelID)
}
func (t *translater) OnLeaveChannel(channelID string) {
	// send to client
	t.channelChnaged(channelID)
}

func (server *Server) ws(c *gin.Context) {
	username := c.MustGet(keyUsername).(string)
	websocket.Handler(func(conn *websocket.Conn) {
		defer conn.Close()

		encoder := json.NewEncoder(conn)
		decoder := json.NewDecoder(conn)

		agent := &translater{encoder, decoder}

		// Will wait until end
		server.backend.OpenSession(username, agent)
	}).ServeHTTP(c.Writer, c.Request)
}
