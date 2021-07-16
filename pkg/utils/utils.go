package utils

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
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
