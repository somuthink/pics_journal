package socket

import (
	"bytes"
	"context"
	"time"

	"github.com/a-h/templ"
	"github.com/charmbracelet/log"
	"github.com/gofiber/contrib/websocket"
	"github.com/golang-jwt/jwt/v5"
	"github.com/somuthink/pics_journal/core/internal/config"
	"github.com/somuthink/pics_journal/core/internal/crypto"
	"github.com/somuthink/pics_journal/core/internal/models"
	"github.com/somuthink/pics_journal/core/internal/queue"
)

type LlmOutput struct {
	Error  error
	UserID uint
	Result string
	models.LlmJob
}

type SocketClient struct {
	conn   *websocket.Conn
	userID uint
}

type SocketHub struct {
	clients              map[uint]*SocketClient
	register, unregister chan *SocketClient
	Output               chan LlmOutput
}

func NewSocketHub() *SocketHub {
	return &SocketHub{
		clients:    make(map[uint]*SocketClient),
		register:   make(chan *SocketClient),
		unregister: make(chan *SocketClient),
		Output:     make(chan LlmOutput),
	}
}

func templToBytes(t templ.Component) []byte {
	var buffer bytes.Buffer
	t.Render(context.Background(), &buffer)
	return buffer.Bytes()
}

func (hub *SocketHub) Run() {
	for {
		select {
		case client := <-hub.register:
			hub.clients[client.userID] = client
			log.Info("new client", "data", client)
		case client := <-hub.unregister:
			if _, ok := hub.clients[client.userID]; ok {
				delete(hub.clients, client.userID)
				client.conn.Close()
			}
		case output := <-hub.Output:
			if _, ok := hub.clients[output.UserID]; ok {
				hub.clients[output.UserID].conn.WriteMessage(websocket.TextMessage, []byte(output.Result))
			}

		}
	}
}

func generate(c *websocket.Conn) {
	token := c.Cookies("jwt")

	claims := jwt.MapClaims{}
	jwt.ParseWithClaims(token, &claims, func(token *jwt.Token) (any, error) {
		return []byte(config.Cfg.JWT_TOKEN), nil
	})

	userID := uint(claims["user_id"].(float64))

	client := &SocketClient{
		conn:   c,
		userID: userID,
	}

	HUB.register <- client
	// defer func() {
	// 	HUB.unregister <- client
	// }()

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	jobID := crypto.GenerateJobID(userID)

	ctx, cancel = context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	err := queue.QueueLlmJob(
		ctx,
		models.LlmJob{JobID: jobID, UserID: userID})
	if err != nil {
		log.Error("queue job", "err", err)
	}

	for {
		if _, _, err := c.ReadMessage(); err != nil {
			return
		}

		// ctx, cancel = context.WithTimeout(context.Background(), 2*time.Second)
		// defer cancel()
		// log.Info("canceling job", "jobID", jobID, "err", sessions.CancelJob(ctx, jobID))

		c.WriteMessage(websocket.TextMessage, []byte("думаю"))

	}
}
