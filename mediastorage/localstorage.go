package mediastorage

import (
	"io/ioutil"
	"os"
)

type LocalStorage struct {
	dirName     string

}

var LocalStorageInstance *LocalStorage


// put data
func (instance *LocalStorage) PutData(key string, data []byte) {
	ioutil.WriteFile(instance.dirName+"/"+ key, data, os.ModePerm);
}

func (instance *LocalStorage) GetData(key string) []byte {
	data, err := ioutil.ReadFile(instance.dirName+"/"+ key)
	handleError(err)
	return data
}
