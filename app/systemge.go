package app

import (
	"SystemgeSamplePingPong/topics"

	"github.com/neutralusername/Systemge/Error"
	"github.com/neutralusername/Systemge/Message"
	"github.com/neutralusername/Systemge/Node"
)

func (app *App) GetSyncMessageHandlers() map[string]Node.SyncMessageHandler {
	return map[string]Node.SyncMessageHandler{
		topics.PINGPONG: func(node *Node.Node, message *Message.Message) (string, error) {
			println("received ping-sync; sending pong-sync")
			return "pong", nil
		},
	}
}

/* var startedAt = time.Time{}
var counter = atomic.Uint32{} */

func (app *App) GetAsyncMessageHandlers() map[string]Node.AsyncMessageHandler {
	return map[string]Node.AsyncMessageHandler{
		topics.PING: func(node *Node.Node, message *Message.Message) error {
			/* 	val := counter.Add(1)
			if val == 1 {
				startedAt = time.Now()
			}
			if val == 100000 {
				println("100000 pings received in " + time.Since(startedAt).String())
				counter.Store(0)
			} */
			err := node.AsyncMessage("pong", "pong")
			if err != nil {
				return Error.New("error sending pong message", err)
			}
			return nil
		},
	}
}
