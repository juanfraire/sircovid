package main

import (
	"image"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

var (
	num = rand.Intn(5)

	count int
	ok    bool
)

func init() {
	rand.Seed(time.Now().UnixNano())
	//enemigo1
	enemigo1.FrameOX = 0
	enemigo1.FrameOY = 48
	enemigo1.FrameNum = 1
	enemigo1.X = float64(350)
	enemigo1.Y = float64(290)
	enemigo1.FrameWidth = 32
	enemigo1.FrameHeight = 48
	enemigo1.num = rand.Intn(5)
	enemigo1.cambio = rand.Intn(50) + 100

	rand.Seed(time.Now().UnixNano())
	//enemigo2
	enemigo2.FrameOX = 0
	enemigo2.FrameOY = 48
	enemigo2.FrameNum = 4
	enemigo2.X = float64(screenWidth - 50)
	enemigo2.Y = float64(290)
	enemigo2.FrameWidth = 32
	enemigo2.FrameHeight = 48
	enemigo2.num = rand.Intn(5)
	enemigo2.cambio = rand.Intn(50) + 50

	// rand.Seed(time.Now().UnixNano())
	// //enemigo2
	// enemigo3.FrameOX = 0
	// enemigo3.FrameOY = 32
	// enemigo3.FrameNum = 3
	// enemigo3.X = float64(20)
	// enemigo3.Y = float64(330)
	// enemigo3.FrameWidth = 32
	// enemigo3.FrameHeight = 32
	// enemigo3.num = rand.Intn(5)
	// enemigo3.cambio = rand.Intn(50) + 50

	// policia.FrameOX = 0
	// policia.FrameOY = 48
	// policia.FrameNum = 1
	// policia.X = float64(20)
	// policia.Y = float64(290)
	// policia.FrameWidth = 32
	// policia.FrameHeight = 48
	// policia.num = rand.Intn(5)
	// policia.cambio = rand.Intn(50) + 50

}

func moverHumanos(FrameOY int, FrameNum int, num int, X float64, Y float64) (int, int, float64, float64) {

	FrameNum = 4

	switch num {
	case 0:
		FrameNum = 1
	case 1:
		FrameOY = 48
		X--
	case 2:
		FrameOY = 96
		X++
	case 3:
		FrameOY = 144
		Y--
	case 4:
		FrameOY = 0
		Y++
	}

	return FrameOY, FrameNum, X, Y
}

func cambioDireccion(num int, cambio int, count int) (int, int) {
	var tmp int

	if count >= cambio {
		for tmp = num; tmp == num; tmp = rand.Intn(5) {
		}
		cambio += rand.Intn(100) + 20
		num = tmp
	}

	return num, cambio
}

func obstEnemigo(cambio int, count int, num int, X float64, Y float64) (int, int) {
	x1, y1 := X, Y
	_, _, ok = obstaculos(X, Y, x1, y1)

	if ok {
		switch num {
		case 1:
			num = 2
		case 2:
			num = 1
		case 3:
			num = 4
		case 4:
			num = 3
		}
		cambio = count + 5
	}
	return num, cambio
}

func dibujarEnemigos(E humanos, screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(.8, .8)
	op.GeoM.Translate(E.X, E.Y)
	j := (count1 / 7) % E.FrameNum
	hx, hy := E.FrameOX+j*E.FrameWidth, E.FrameOY
	screen.DrawImage(E.img.SubImage(image.Rect(hx, hy, hx+E.FrameWidth, hy+E.FrameHeight)).(*ebiten.Image), op)
}
