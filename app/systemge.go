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
			app.pingsReceived++
			println("received ping-sync; sending pong-sync")
			return "pong", nil
		},
	}
}

/* var startedAt = time.Time{} */

func (app *App) GetAsyncMessageHandlers() map[string]Node.AsyncMessageHandler {
	return map[string]Node.AsyncMessageHandler{
		topics.PING: func(node *Node.Node, message *Message.Message) error {
			app.pingsReceived++
			/* if app.pingsReceived == 2 {
				startedAt = time.Now()
			}
			if app.pingsReceived == 100001 {
				println("100000 pings received in " + time.Since(startedAt).String())
				app.pingsReceived = 0
			} */
			err := node.AsyncMessage("pong", "pong")
			if err != nil {
				return Error.New("error sending pong message", err)
			}
			return nil
		},
	}
}
