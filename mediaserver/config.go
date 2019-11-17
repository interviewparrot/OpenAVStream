package mediaserver

import (
	"github.com/interviewparrot/OpenAVStream/mediastorage"
	"github.com/interviewparrot/ParrotServer/server"
	"github.com/magiconair/properties"
	"log"
	"os"
)

var BUCKET_NAME string
var VIDEO_PREFIX = "video-data"
var AUDIO_PREFIX = "audio-data"

const envVar = "APP_PROFILE"
var PROFILE = "dev"
var PROPERTIES *properties.Properties

func init() {
	if profile := os.Getenv(envVar); profile != "" {
		PROFILE = profile
	}
	log.Println("Loading the properties for profile: " + PROFILE)
	PROPERTIES = properties.MustLoadFile("resources/application-"+PROFILE+".properties", properties.UTF8)
	server.BUCKET_NAME = PROPERTIES.MustGet("video-mediastorage-bucket")
	mediastorage.StorageBucketInstance = mediastorage.CreateStorageBucket(server.BUCKET_NAME)
}

func GetProperty(key string) string {
	return PROPERTIES.MustGet(key)
}