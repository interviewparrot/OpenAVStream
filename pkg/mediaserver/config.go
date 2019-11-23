package mediaserver

import (
	"github.com/interviewparrot/OpenAVStream/pkg/mediastorage"
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
	BUCKET_NAME = PROPERTIES.MustGet("openavstream.mediastorage.bucket")
	storageType := GetProperty("openavstream.mediastorage.type")
	if storageType == "cloud" {
		mediastorage.StorageBucketInstance = mediastorage.CreateStorageBucket(BUCKET_NAME)
	}
}

func GetProperty(key string) string {
	return PROPERTIES.MustGet(key)
}