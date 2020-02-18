package websocket

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

// WsHandler -
type WsHandler struct {
	handler echo.HandlerFunc
}

func (h *WsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	e := echo.New()
	c := e.NewContext(r, w)

	forever := make(chan struct{})
	h.handler(c)
	<-forever
}

func TestWebsocketExample(t *testing.T) {
	h := WsHandler{handler: hello}
	server := httptest.NewServer(http.HandlerFunc(h.ServeHTTP))
	defer server.Close()

	wsURL := "ws" + strings.TrimPrefix(server.URL, "http") + "/ws"
	ws, _, err := websocket.DefaultDialer.Dial(wsURL, nil)
	assert.Nil(t, err, err)

	// write
	err = ws.WriteMessage(websocket.TextMessage, []byte("Hello, Server!"))
	assert.Nil(t, err, err)

	// read
	_, msg, err := ws.ReadMessage()
	assert.Nil(t, err, err)
	assert.Equal(t, "Hello, Client!", string(msg))
}
