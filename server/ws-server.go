package wsServer

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
	"github.com/alchemist7991/scalable-chat-service/constant"
	redisHelper "github.com/alchemist7991/scalable-chat-service/redisHelper"
	"golang.org/x/net/websocket"
)

type Server struct {
	conns map[*websocket.Conn]string
}

func RegisterSocketHandlers(s *Server) {
	http.Handle("/init", websocket.Handler(s.HandleNewWS))
	fmt.Printf("Starting server on%s", constant.WS_PORT)
	http.ListenAndServe(constant.WS_PORT, nil)
}

func NewServer() *Server {
	return &Server{
		conns: make(map[*websocket.Conn]string),
	}
}

func (s *Server) HandleNewWS(ws *websocket.Conn) {
	defer ws.Close()
	log.Println("New connection received from ", ws.RemoteAddr())
	s.conns[ws] = GenerateSocketId()
	s.ReadMessage(ws)
}

func (s *Server) ReadMessage(ws *websocket.Conn) {
	buf := make([]byte, 2048)
	for {
		n, err := ws.Read(buf)
		if err != nil {
			log.Fatalln("Unable to read message", err)
		}
		msg := string(buf[:n])
		go redisHelper.StoreMessage(msg, s.conns[ws], ws.RemoteAddr().String())
		log.Printf("%s (%s): %s",s.conns[ws], ws.RemoteAddr(), msg)
	}
}

func GenerateSocketId() string {
	uniqId := strconv.FormatInt(time.Now().UTC().UnixNano(), 10)
	return uniqId
}

func StartServer() {
	server := NewServer()
	redisHelper.SetClientInstance()
	RegisterSocketHandlers(server)
}