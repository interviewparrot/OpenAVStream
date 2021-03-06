package mediastream

import (
	"github.com/interviewparrot/OpenAVStream/pkg/mediaserver"
	"github.com/interviewparrot/OpenAVStream/pkg/mediastorage"
	"log"
	"time"
)

func ProcessIncomingMsg(session *mediaserver.Session, msg []byte) {
	log.Println("Writing media chunk")
	objectKey := mediaserver.VIDEO_PREFIX + "/" + session.SessionId +"/"+ GetCurrentTime()+".webm"
	mediastorage.LocalStorageInstance.PutData(objectKey, msg)
}


func GetCurrentTime() string {
	return time.Now().Format("2006-01-02T15-04-05")
}