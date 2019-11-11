package main

import (
	"encoding/base64"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/go-redis/redis"
	"github.com/gorilla/websocket"
	"github.com/interviewparrot/ParrotServer/cloudstorage"
	"github.com/magiconair/properties"
	"log"
	"net/http"
	"os"
	"github.com/interviewparrot/ParrotServer/server"
	"time"
)


var upgrader = websocket.Upgrader{} // use default options
const envVar = "APP_PROFILE"

var PROFILE = "dev"
var PROPERTIES *properties.Properties

func init() {
	if profile := os.Getenv(envVar); profile != "" {
		PROFILE = profile
	}
	log.Println("Loading the properties for profile: " + PROFILE)
	PROPERTIES = properties.MustLoadFile("resources/application-"+PROFILE+".properties", properties.UTF8)
	server.BUCKET_NAME = PROPERTIES.MustGet("video-storage-bucket")
	cloudstorage.StorageBucketInstance = cloudstorage.CreateStorageBucket(server.BUCKET_NAME)
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

}

func makeTimestamp() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

func processIncomingMsg(session *server.Session, msg []byte) {
	log.Println("Writing video chunk")
	objectKey := server.AUDIO_PREFIX + "/" + session.SessionId +"/"+ GetCurrentTime()+".wav"
	cloudstorage.StorageBucketInstance.PutObject(objectKey, msg)
	//text := speechrec.SpeechToText([]byte(fileName), msg)
	log.Println("incoming data")
	//voicegraph.AddToIncomingData(text, currentTimeStr, session.SessionId)
}

func GetCurrentTime() string {
	return time.Now().Format("2006-01-02T15-04-05")
}

func processVideoMsg(session server.Session, msg []byte) {
	log.Println("Writing video chunk")
	objectKey := server.VIDEO_PREFIX + "/" + session.SessionId +"/"+ GetCurrentTime()+".webm"
	cloudstorage.StorageBucketInstance.PutObject(objectKey, msg)

}


func ProcessTextMessage(msg []byte) {
	clientMessage := server.ClientMsg{}
	json.Unmarshal(msg, &clientMessage)
	if server.IsSessionExist(clientMessage.SessionId) {
		session := server.SessionStore[clientMessage.SessionId]
		switch cmd := clientMessage.Command; cmd {
		case server.CMD_Auth:
			log.Println("Auth token: " + clientMessage.Data)
		case server.CMD_StartSession:
			log.Println("Starting the conversation..." + clientMessage.Data)
		case server.CMD_ReceiveChunk:
			data, err := base64.StdEncoding.DecodeString(clientMessage.Data)
			log.Println("receiving chunk for sessionID: "+ clientMessage.SessionId + " and session state is: " + session.State)
			if err != nil {
				fmt.Println("error:", err)
				return
			}
			if session.State == server.SESSION_ENDED {
				processIncomingMsg(server.SessionStore[clientMessage.SessionId], data)
				log.Println("Session has ended closing the connection")
				session.ConnGroup.UserConnection.Conn.Close()
			} else {
				processIncomingMsg(server.SessionStore[clientMessage.SessionId], data)
			}

		}
	}
}

func adminEcho(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()

	for {
		mt, message, err := c.ReadMessage()
		if err != nil {

			log.Println("read:", err)
			break
		}
		log.Printf("message type: %s", mt)
		if mt == 2 {
			log.Println("Cannot process binary message right now")
		} else {
			ProcessTextMessage(message)
		}
	}
}

func metrics(w http.ResponseWriter, r *http.Request) {

}

func sessionHandler(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	c, err := upgrader.Upgrade(w, r, nil)
	session := server.CreateNewSession(c)
	// Send the session id to the client
	msg := server.ServerMsg{server.CMD_ReceiveSessionId, session.SessionId, session.SessionId}
	msgBytes, err := json.Marshal(msg)
	if err != nil {
		fmt.Println(err)
	}
	c.WriteMessage(1, msgBytes);

	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()

	for {
		mt, message, err := c.ReadMessage()
		if err != nil {

			log.Println("read:", err)
			break
		}
		log.Printf("message type: %s", mt)
		if mt == 2 {
			log.Println("Cannot process binary message right now")
		} else {
			ProcessTextMessage(message)
		}
	}
}

func pingHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("healthy"))
}

func main() {
	flag.Parse()
	log.SetFlags(0)
	http.HandleFunc("/session", sessionHandler)
	http.HandleFunc("/ping", pingHandler)
	log.Fatal(http.ListenAndServe(":"+ PROPERTIES.MustGet("parrot.server.port"), nil))
}
