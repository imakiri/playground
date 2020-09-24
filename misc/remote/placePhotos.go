package remote

import (
	"fmt"
	"github.com/imakiri/playground/core"
	"github.com/imakiri/playground/misc"
	"sync"
)

var PlacePhotos placePhotos0

type placePhotos0 bool

func (placePhotos0) Cast(group *sync.WaitGroup, c chan core.ThingImp) {
	defer group.Done()

	resp, err := misc.Client.Get("http://google.com")
	if err != nil {
		fmt.Println(err)
		return
	}
	//_, _ = io.Copy(os.Stdout, resp.Body)
	resp.Body.Close()
}
