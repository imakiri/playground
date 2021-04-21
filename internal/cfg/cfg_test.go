package cfg

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"os"
	"testing"
)

func TestNewService(t *testing.T) {
	const file = "main.yaml"
	fmt.Println(os.Getwd())
	var service, err = NewService(file)
	if err != nil {
		t.Error(err)
	}

	spew.Dump(service.config)
}
