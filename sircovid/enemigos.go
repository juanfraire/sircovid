package main

import (
	"image"
	"log"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

var (
	enemigo  humanos
	enemigo1 humanos
	enemigo2 humanos
	enemigo3 humanos
	enemigo4 humanos
	// num      = rand.Intn(5)
	count int
	tmp   int
	ok    bool
	x     float64
	y     float64
	m     []int
	arr   [][]int
)

func randXY() (x float64, y float64) {
	rand.Seed(time.Now().UnixNano())
	_, _, obs := obstaculos(x, y, x, y)
	for obs {
		x = float64(rand.Intn(screenWidth))
		y = float64(rand.Intn(screenHeight))
		_, _, obs = obstaculos(x, y, x, y)
	}
	// fmt.Println(x, y, obs)
	return
}

func initEnemigos() {

	rand.Seed(time.Now().UnixNano())

	//enemigo1
	enemigo1.FrameOX = 48
	enemigo1.FrameOY = 72 * rand.Intn(4)
	enemigo1.FrameNum = 1
	enemigo1.X, enemigo1.Y = randXY()
	enemigo1.FrameWidth = 48
	enemigo1.FrameHeight = 72
	enemigo1.num = rand.Intn(5)
	enemigo1.cambio = rand.Intn(50) + 100

	en := `sircovid\data\HERO-Jessica-Poses.png`

	enemigo1.img, _, err = ebitenutil.NewImageFromFile(en, ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}

	// fmt.Println(enemigo1)

	rand.Seed(time.Now().UnixNano())
	//enemigo2
	enemigo2.FrameOX = 48
	enemigo2.FrameOY = 72 * rand.Intn(4)
	enemigo2.FrameNum = 1
	enemigo2.X, enemigo2.Y = randXY()
	enemigo2.FrameWidth = 48
	enemigo2.FrameHeight = 72
	enemigo2.num = rand.Intn(5)
	enemigo2.cambio = rand.Intn(50) + 200

	enemigo2.img, _, err = ebitenutil.NewImageFromFile(`sircovid\data\cobani.png`, ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}

	rand.Seed(time.Now().UnixNano())

	enemigo3.FrameOX = 48
	enemigo3.FrameOY = 72 * rand.Intn(4)
	enemigo3.FrameNum = 1
	enemigo3.X, enemigo3.Y = randXY()
	enemigo3.FrameWidth = 48
	enemigo3.FrameHeight = 72
	enemigo3.num = rand.Intn(5)
	enemigo3.cambio = rand.Intn(50) + 100

	enemigo3.img, _, err = ebitenutil.NewImageFromFile(`sircovid\data\mujer-con-sombrero.png`, ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}

	rand.Seed(time.Now().UnixNano())

	enemigo4.FrameOX = 48
	enemigo4.FrameOY = 72 * rand.Intn(4)
	enemigo4.FrameNum = 1
	enemigo4.X, enemigo4.Y = randXY()
	enemigo4.FrameWidth = 48
	enemigo4.FrameHeight = 72
	enemigo4.num = rand.Intn(5)
	enemigo4.cambio = rand.Intn(50) + 100

	enemigo4.img, _, err = ebitenutil.NewImageFromFile(`sircovid\data\mujer-pelo-largo.png`, ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
}

func moverHumanos(E humanos) humanos {

	// initEnemigos()

	// arr = append(arr, [1]float64)
	// for i := 0; i < 10; i++ {
	// arr = enemigos()

	// fmt.Println(arr)

	if ModeGame >= 0 {
		count++
	}

	E.FrameNum = 3
	E.FrameOX = 0

	switch E.num {
	case 0:
		E.FrameNum = 1
		E.FrameOX = 48
	case 1:
		E.FrameOY = 72
		E.X--
	case 2:
		E.FrameOY = 144
		E.X++
	case 3:
		E.FrameOY = 216
		E.Y--
	case 4:
		E.FrameOY = 0
		E.Y++
	}

	if count >= E.cambio {
		for tmp = E.num; tmp == E.num; tmp = rand.Intn(5) {
		}
		E.cambio += rand.Intn(200) + 200
		E.num = tmp
	}

	x1, y1 := E.X, E.Y
	_, _, ok = obstaculos(E.X, E.Y, x1, y1)

	if ok {
		switch E.num {
		case 1:
			E.num = 2
		case 2:
			E.num = 1
		case 3:
			E.num = 4
		case 4:
			E.num = 3
		}
		E.cambio = count + 5
	}
	return E
}

func dibujarEnemigos(E humanos, screen *ebiten.Image) {
	if ModePause {
		E.FrameNum = 1
		E.FrameOX = 0
	}
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(.65, .51)
	op.GeoM.Translate(E.X, E.Y)
	j := (count1 / 7) % E.FrameNum
	hx, hy := E.FrameOX+j*E.FrameWidth, E.FrameOY
	screen.DrawImage(E.img.SubImage(image.Rect(hx, hy, hx+E.FrameWidth, hy+E.FrameHeight)).(*ebiten.Image), op)
}

// func enemigos() [][]int {
// 	var m []int
// 	x, y := randXY()
// 	m = append(m, 48, 72*rand.Intn(4), 1, int(x), int(y), 48, 72, rand.Intn(5), rand.Intn(50)+100)
// 	arr = append(arr, m)

// 	return arr
// }
