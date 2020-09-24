package data

import "sync"

func (e *external) placePhotos(queue string, wg *sync.WaitGroup, c chan Re) {
	defer wg.Done()

	//

	c <- re
	close(c)
}

func (e *external) placeSearch(queue string, wg *sync.WaitGroup, c chan Re) {
	defer wg.Done()

	//

	c <- re
	close(c)
}
