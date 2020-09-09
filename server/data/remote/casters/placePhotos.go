package casters

import (
	"fmt"
	"github.com/imakiri/playground/server/core"
	"github.com/imakiri/playground/server/data/remote"
	"sync"
)

var PlacePhotos placePhotos0

type placePhotos0 bool

func (placePhotos0) Cast(group *sync.WaitGroup, c chan core.ThingImp) {
	defer group.Done()

	resp, err := remote.Client.Get("http://google.com")
	if err != nil {
		fmt.Println(err)
		return
	}
	//_, _ = io.Copy(os.Stdout, resp.Body)
	resp.Body.Close()
}
