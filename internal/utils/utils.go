package utils

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"net/http"
)

func IsNil(l ...interface{}) bool {
	for i := 0; i < len(l); i++ {
		if l[i] == nil {
			return true
		}
	}
	return false
}

func IsNilEx(l ...interface{}) (b []bool) {
	for i := 0; i < len(l); i++ {
		if l[i] == nil {
			b = append(b, true)
		} else {
			b = append(b, false)
		}
	}
	return
}

func ReadYAML(path string, dest interface{}) error {
	var raw, err = ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	return yaml.Unmarshal(raw, dest)
}

func SendBytes(data []byte, w http.ResponseWriter, r *http.Request) error {
	var _, err = w.Write(data)
	return err
}
