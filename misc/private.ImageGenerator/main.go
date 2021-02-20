package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"os"
	"sync"
)

// Compute n / (b ** p)
func div(n, b, p uint16) uint16 {
	for p > 0 {
		n = n / b
		p--
	}
	return n
}

// Compute ceil(log_b(n))
func dim(n, b uint16) (d uint16) {
	for {
		m := div(n, b, d)
		if m > 0 {
			d++
		} else {
			return
		}
	}
}

// Count unique numbers in array
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
	n = uint8(len(lookTable))
	return n
}

type task struct {
	size      [2]uint16
	base      [2]uint16
	dimension [2]uint16
	palette   []color.Color
}

func createTask(size, base [2]uint16, palette []color.Color) (t *task, err error) {
	t = new(task)
	t.size = size
	t.base = base
	t.palette = palette
	t.dimension[0] = dim(t.size[0], t.base[0])
	t.dimension[1] = dim(t.size[1], t.base[1])
	return
}

func (e task) compute() *image.Paletted {
	rect := image.Rectangle{
		Min: image.Point{
			X: 0,
			Y: 0,
		},
		Max: image.Point{
			X: int(e.size[0]),
			Y: int(e.size[1]),
		},
	}
	pal := color.Palette(e.palette)
	img := image.NewPaletted(rect, pal)
	wg := &sync.WaitGroup{}

	//re := make([][][]uint64, e.size[0])
	//for i := range re {
	//	re[i] = make([][]uint64, e.size[1])
	//}

	for i := uint16(0); i < e.size[0]; i++ {
		wg.Add(1)

		//bX := base(i, e.base[0], e.dimension[0])

		//for j := uint16(0); j < e.size[1]; j++ {
		//	bY := base(j, e.base[1], e.dimension[1])
		//	t := append(bX, bY...)
		//	re[i][j] = t
		//	img.SetColorIndex(int(i), int(j), count(t))
		//}

		go func(i uint16, wg *sync.WaitGroup) {
			bX := base(i, e.base[0], e.dimension[0])

			for j := uint16(0); j < e.size[1]; j++ {
				bY := base(j, e.base[1], e.dimension[1])
				t := append(bX, bY...)
				img.SetColorIndex(int(i), int(j), count(t))
			}

			wg.Done()
		}(i, wg)
	}

	wg.Wait()
	return img
}

func pow(x, n uint16) (re uint64) {
	re = 1
	for n > 0 {
		re = re * uint64(x)
		n--
	}
	return
}

func base(n, b, dim uint16) (re []uint64) {
	re = make([]uint64, 0, dim)

	for i := dim - 1; i < dim; i-- {
		t := uint64(div(n, b, i))
		//t := uint16(uint64(n) / pow(b, i))
		n = n % uint16(pow(b, i))
		re = append(re, t)
	}

	return
}

func writeImg(x, y, b uint16) {
	f, err := os.Create("draw.gif")
	if err != nil {
		panic(err.Error())
	}

	size := [2]uint16{x, y}
	base := [2]uint16{b, b}
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
	t, _ := createTask(size, base, pal)

	img := []*image.Paletted{
		t.compute(),
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

func generateImg() {
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

func main() {
	generateImg()
	//fmt.Println(div(2048, 8, 5))
}
