package main

import (
	"crypto/md5"
	"fmt"
	"imakiteki/playground/server/store"
	"io"
	"strings"
)

func PassHash(str string) string {
	h := md5.New()
	_, _ = io.WriteString(h, str)
	return fmt.Sprintf("%x", h.Sum(nil))
}

func strFormat(str ...string) string {
	if len(str) == 0 || str[0] == "" {
		return "Hello, World!"
	}

	return fmt.Sprintf("Hello, %s!", strings.Title(strings.ToLower(str[0])))
}

func main() {
	store.Run()

	//server.Run()
}
