package mediastorage

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

type LocalStorage struct {
	dirName     string

}

var LocalStorageInstance *LocalStorage

func init() {
	LocalStorageInstance = &LocalStorage{dirName:"/tmp"}
}


// put data
func (instance *LocalStorage) PutData(key string, data []byte) {

	dir := filepath.Dir(key)
	fullPath := instance.dirName +"/"+  dir
	log.Printf(fullPath)
	os.MkdirAll(fullPath, 0777)
	ioutil.WriteFile(instance.dirName+"/"+ key, data, os.ModePerm);
}

func (instance *LocalStorage) GetData(key string) []byte {
	data, err := ioutil.ReadFile(instance.dirName+"/"+ key)
	handleError(err)
	return data
}
