package cloudstorage

import (
	"cloud.google.com/go/storage"
	"context"
	"google.golang.org/api/iterator"
	"io/ioutil"
	"log"
)

// This file is inspired from
// https://github.com/GoogleCloudPlatform/golang-samples/blob/8deb2909eadf32523007fd8fe9e8755a12c6d463/docs/appengine/storage/app.go

var client *storage.Client
var ctx = context.Background()
var StorageBucketInstance *StorageBucket


func init() {
	log.Println("Creating storage bucket instance")
	client, _ = storage.NewClient(ctx)
}



type StorageBucket struct {
	client     *storage.Client
	bucketName string
	bucket     *storage.BucketHandle
}


func CreateStorageBucket(bucketName string) *StorageBucket {
	return &StorageBucket{
		client:     client,
		bucketName: bucketName,
		bucket:     client.Bucket(bucketName),
	}
}

// put a new object
func (sb *StorageBucket) PutObject(key string, data []byte) {
	sb.bucket = client.Bucket(sb.bucketName)
	wc := sb.bucket.Object(key).NewWriter(ctx)
	code, err := wc.Write(data)
	handleError(err)
	log.Println(code)
	wc.Close()

}

// read the object bytes
func (sb *StorageBucket) ReadBytes( object string) []byte {
	log.Println("Reading from bucket: "+ sb.bucketName+" key: "+ object)
	sb.bucket = client.Bucket(sb.bucketName)
	rc, err := sb.bucket.Object(object).NewReader(ctx)
	handleError(err)
	defer rc.Close()

	data, err := ioutil.ReadAll(rc)
	handleError(err)
	return data
}

func (sb *StorageBucket) ListObjects() []byte {
	log.Println("Reading from bucket: "+ sb.bucketName)
	sb.bucket = client.Bucket(sb.bucketName)
	it := sb.bucket.Objects(ctx, nil)
	for {
		attrs, err := it.Next()
		if err == iterator.Done {
			break
		}
		handleError(err)
		log.Println(attrs.Name)
	}

	return nil
}

// Generates a composite object
func (sb *StorageBucket)  Compose(bucket string, srcObjects[] string, destObj string) {
	bkt := client.Bucket(bucket)
	src := []*storage.ObjectHandle{}
	for _, srcObj := range srcObjects {
		objHandle := bkt.Object(srcObj)
		src = append(src, objHandle)
	}
	dst := bkt.Object(destObj)
	attrs, err := dst.ComposerFrom(src...).Run(ctx)
	handleError(err)
	log.Println(attrs)

}

func handleError(err error) {
	if err != nil {
		log.Println("handling error::::", err)

	}
}

