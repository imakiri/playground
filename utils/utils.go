package utils

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

func IsNilSafe(l ...interface{}) bool {
	for i := 0; i < len(l); i++ {
		if l[i] == nil {
			return false
		}
	}
	return true
}

func IsNilSafeEx(l ...interface{}) (b []bool) {
	for i := 0; i < len(l); i++ {
		if l[i] == nil {
			b = append(b, false)
		} else {
			b = append(b, true)
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
