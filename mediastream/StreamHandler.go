package stream

import (
	"github.com/interviewparrot/OpenAVStream/mediaserver"
	"github.com/interviewparrot/ParrotServer/cloudstorage"
	"log"
	"time"
)

func ProcessIncomingMsg(session *mediaserver.Session, msg []byte) {
	log.Println("Writing video chunk")
	objectKey := mediaserver.AUDIO_PREFIX + "/" + session.SessionId +"/"+ GetCurrentTime()+".wav"
	cloudstorage.StorageBucketInstance.PutObject(objectKey, msg)
	log.Println("incoming data")
}


func GetCurrentTime() string {
	return time.Now().Format("2006-01-02T15-04-05")
}