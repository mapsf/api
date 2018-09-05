package ws

import (
	"net/http"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"encoding/json"
	"github.com/mapsf/api/app/common"
	"github.com/mapsf/api/app/models"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type ServerEvent struct {
	Type string      `json:"type"`
	Data interface{} `json:"data"`
}

type ClientEvent struct {
	User *models.Character `json:"-"`
	Type string            `json:"type"`
	Data interface{}       `json:"data"`
}

type hub struct {
	Connections map[uint]*websocket.Conn
}

var Hub = hub{
	Connections: make(map[uint]*websocket.Conn, 0),
}

func logError(err error) {
	if err != nil {
		log.Println(err)
	}
}

func debug(format string, args ... interface{}) {
	log.Printf("[WebScokets] "+format, args...)
}

func Emit(event ServerEvent, ids ... uint) error {
	// отправить одному подключенному клиенту
	if len(ids) == 1 {
		id := ids[0]
		conn, ok := Hub.Connections[id]
		if !ok {
			return fmt.Errorf(`Socket connection for userID "%v" does not exist`, id)
		}
		debug(`[SEND] to single userID -> %v, event -> "%v", data -> %v`, id, event.Type, event.Data)
		logError(conn.WriteJSON(event))
	} else {
		// отправить всем
		debug(`[SEND] to all clients, event -> "%v", data %v`)
		for userID, conn := range Hub.Connections {
			debug("[SEND_TO] user_id %v", userID)
			logError(conn.WriteJSON(event))
		}
	}
	return nil
}

func Handler(p common.Params) common.ResponseRenderer {

	debug("Подключение клиента...")

	conn, err := upgrader.Upgrade(p.ResponseWriter, p.Request, nil)
	if err != nil {
		fmt.Printf("Не удалось установить рукопожатие, причина %v", err)
		return nil
	}

	Hub.Connections[p.User.ID] = conn

	debug("Клиент подключен. Всего клиентов %v", len(Hub.Connections))

	Emit(ServerEvent{Type: "user", Data: p.User}, p.User.ID)

	//go func() {
	//	time.Sleep(5 * time.Second);
	//	Emit(ServerEvent{Type: "ping", Data: ":)"}, p.User.ID)
	//}()

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			debug("Клиент отключен, причина: %v", err)
			delete(Hub.Connections, p.User.ID)
			break
		}

		readMessage(msg, p.User)
	}

	return nil
}

func readMessage(message []byte, character *models.Character) {

	event := &ClientEvent{}

	err := json.Unmarshal(message, event)
	if err != nil {
		return
	}

	event.User = character

	debug(`[RECEIVED] event -> "%v", data -> %v`, event.Type, event.Data)

	processClientEvent(event)
}
