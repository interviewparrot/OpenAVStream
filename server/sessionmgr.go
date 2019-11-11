package server

import (
	"github.com/gorilla/websocket"
	"github.com/speps/go-hashids"
	"log"
	"time"
)

const(
	SESSION_STARTED = "STARTED"
	SESSION_INPROGRESS = "INPROGRESS"
	SESSION_ENDED = "ENDED"

)

var SessionStore = map[string]*Session {
}

// Session is started when first user connects to it the server
// a unique session Id is given to it.
type Session struct {
	SessionId string
	AuthToken string
    ConnGroup ConnectionGroup
	State string
}

// Holds the socket connection and a unique id for it.
type Connection struct {
	Id string
	ClientAddr string
	Conn *websocket.Conn
}

// Holds a group of connection. Group is defined by one user connection
// and multiple Admin connection.
type ConnectionGroup struct {
	UserConnection Connection
	AdminConnection []Connection
}
const UserConnType = 0
const AdminConnType = 1

func makeTimestamp() int64 {
	return time.Now().UnixNano() / int64(time.Second)
}

func IsSessionExist(sessionId string) bool {
	if _, ok:= SessionStore[sessionId]; ok {
		return true
	}
	return false
}

// Creates a new session
func CreateNewSession(conn *websocket.Conn) *Session {
	id := newHashId()
	userConnect := Connection{id, conn.RemoteAddr().String(), conn}
	conngroup := ConnectionGroup{userConnect, nil}
	session := Session { id, "",conngroup , "STARTED"}
	SessionStore[id] = &session
	return &session

}

// write the binary data to the socket
func (s *Session) WriteBinary(data []byte, connType int) {
	log.Println("Start Writing to the connection")
	if connType == 0 {
		s.ConnGroup.UserConnection.Conn.WriteMessage(2, data)
		log.Println("finished writing to the connection")
	} else {
		for _, adminConn := range s.ConnGroup.AdminConnection {
			adminConn.Conn.WriteMessage(2, data)
		}
	}
}

// Write the text data to the socket
func (s *Session) WriteText(data []byte, connType int) {
	log.Println("Writing text data: "+ string(data))
	if connType == 0 {
		s.ConnGroup.UserConnection.Conn.WriteMessage(1, data)
		log.Println("finished writing to the connection")
	} else {
		for _, adminConn := range s.ConnGroup.AdminConnection {
			adminConn.Conn.WriteMessage(1, data)
		}
	}
}

func newHashId() string {
	var hd = hashids.NewData()
	hd.Salt = "Parrot Salt"
	h, err := hashids.NewWithData(hd)
	handleError(err)
	now := time.Now()
	year := now.Year()
	month := int(now.Month())
	day := now.Day()
	hour := now.Hour()
	minute := now.Minute()
	second := now.Second()
	id, _ := h.Encode([]int{year, month, day, hour, minute, second})
	return id
}

func handleError(err error) {
	if err != nil {
		log.Println("handling error::::", err)

	}
}
