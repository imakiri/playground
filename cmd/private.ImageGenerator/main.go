package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"os"
	"sync"
)

func pow(x, n uint16) (re uint64) {
	re = 1
	for n > 0 {
		re = re * uint64(x)
		n--
	}
	return
}

func base(x, b, dim uint16) (re []uint64) {
	re = make([]uint64, 0, dim)

	for i := dim - 1; i < dim; i-- {
		t := uint64(x) / pow(b, i)
		x = x % uint16(pow(b, i))
		re = append(re, t)
	}

	return
}

func count(x []uint64) (n uint8) {
	lookTable := make([]uint64, 0, len(x))

	for _, value := range x {
		isUnique := true

		for _, lookTableValue := range lookTable {
			if value == lookTableValue {
				isUnique = false
				break
			}
		}

		if isUnique {
			lookTable = append(lookTable, value)
		}
	}

	return uint8(len(lookTable))
}

func dim(x, b uint16) (dim uint16) {
	for {
		if uint64(x)/pow(b, dim) > 0 {
			dim++
		} else {
			break
		}
	}
	return
}

func draw(x, y, b uint16) *image.Paletted {
	rect := image.Rectangle{
		Min: image.Point{
			X: 0,
			Y: 0,
		},
		Max: image.Point{
			X: int(x),
			Y: int(y),
		},
	}
	pal := color.Palette([]color.Color{
		color.RGBA{243, 72, 72, 0},
		color.RGBA{242, 157, 73, 0},
		color.RGBA{242, 242, 73, 0},
		color.RGBA{157, 242, 73, 0},
		color.RGBA{73, 242, 157, 0},
		color.RGBA{73, 242, 242, 0},
		color.RGBA{73, 157, 242, 0},
		color.RGBA{73, 73, 242, 0},
		color.RGBA{157, 73, 242, 0},
		color.RGBA{242, 73, 242, 0},
	})
	img := image.NewPaletted(rect, pal)

	dX := dim(x, b) - 1
	dY := dim(y, b) - 1
	wg := &sync.WaitGroup{}

	for i := uint16(0); i < x; i++ {

		wg.Add(1)
		go func(i uint16, wg *sync.WaitGroup) {
			bX := base(i, b, dX)

			for j := uint16(0); j < y; j++ {
				bY := base(j, b, dY)
				t := append(bX, bY...)
				img.SetColorIndex(int(i), int(j), count(t))
			}
			wg.Done()
		}(i, wg)

	}

	wg.Wait()
	return img
}

func writeImg(x, y, b uint16) {
	f, err := os.Create("draw.gif")
	if err != nil {
		panic(err.Error())
	}

	img := []*image.Paletted{
		draw(x, y, b),
	}

	g := &gif.GIF{
		Image:           img,
		Delay:           []int{0},
		LoopCount:       0,
		Disposal:        nil,
		Config:          image.Config{},
		BackgroundIndex: 0,
	}

	err = gif.EncodeAll(f, g)
	if err != nil {
		panic(err.Error())
	}
}

func main() {
	var x, y, b int
	var err error

	for {
		fmt.Print("\nEnter x: ")
		_, err = fmt.Scanf("%d\n", &x)
		if err != nil {
			fmt.Println(err.Error())
		} else {
			break
		}
	}

	for {
		fmt.Print("\nEnter y: ")
		_, err = fmt.Scanf("%d\n", &y)
		if err != nil {
			fmt.Println(err.Error())
		} else {
			break
		}
	}

	for {
		fmt.Print("\nEnter base: ")
		_, err = fmt.Scanf("%d\n", &b)
		if err != nil {
			fmt.Println(err.Error())
		} else {
			break
		}
	}

	fmt.Printf("\nx: %d, y: %d, b: %d", x, y, b)
	fmt.Print("\n...\n")
	writeImg(uint16(x), uint16(y), uint16(b))
}
