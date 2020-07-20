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
	enemigo1 humanos
	enemigo2 humanos
	enemigo3 humanos
	// num      = rand.Intn(5)
	count int
	tmp   int
	ok    bool
)

func initEnemigos() {
	rand.Seed(time.Now().UnixNano())
	//enemigo1
	enemigo1.FrameOX = 48
	enemigo1.FrameOY = 72 * rand.Intn(4)
	enemigo1.FrameNum = 1
	enemigo1.X = float64(350)
	enemigo1.Y = float64(290)
	enemigo1.FrameWidth = 48
	enemigo1.FrameHeight = 72
	enemigo1.num = rand.Intn(5)
	enemigo1.cambio = rand.Intn(50) + 100

	enemigo1.img, _, err = ebitenutil.NewImageFromFile(`sircovid\data\HERO-Jessica-Poses.png`, ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}

	rand.Seed(time.Now().UnixNano())
	//enemigo2
	enemigo2.FrameOX = 48
	enemigo2.FrameOY = 72 * rand.Intn(4)
	enemigo2.FrameNum = 1
	enemigo2.X = float64(screenWidth - 100)
	enemigo2.Y = float64(290)
	enemigo2.FrameWidth = 48
	enemigo2.FrameHeight = 72
	enemigo2.num = rand.Intn(5)
	enemigo2.cambio = rand.Intn(50) + 100

	enemigo2.img, _, err = ebitenutil.NewImageFromFile(`sircovid\data\cobani.png`, ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}

	rand.Seed(time.Now().UnixNano())

	enemigo3.FrameOX = 48
	enemigo3.FrameOY = 72 * rand.Intn(4)
	enemigo3.FrameNum = 1
	enemigo3.X = float64(screenWidth - 100)
	enemigo3.Y = float64(290)
	enemigo3.FrameWidth = 48
	enemigo3.FrameHeight = 72
	enemigo3.num = rand.Intn(5)
	enemigo3.cambio = rand.Intn(50) + 100

	enemigo3.img, _, err = ebitenutil.NewImageFromFile(`sircovid\data\mujer-con-sombrero.png`, ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}

}

func moverHumanos(E humanos) humanos {

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
		E.cambio += rand.Intn(100) + 100
		E.num = tmp
	}

	x1, y1 := E.X, E.Y
	_, _, ok = obstaculos(E.X, E.Y+32, x1, y1)

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
		E.cambio = count + 10
	}
	return E
}

func dibujarEnemigos(E humanos, screen *ebiten.Image) {
	if ModePause {
		E.FrameNum = 1
		E.FrameOX = 0
	} else {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Scale(.65, .51)
		op.GeoM.Translate(E.X, E.Y)
		j := (count1 / 7) % E.FrameNum
		hx, hy := E.FrameOX+j*E.FrameWidth, E.FrameOY
		screen.DrawImage(E.img.SubImage(image.Rect(hx, hy, hx+E.FrameWidth, hy+E.FrameHeight)).(*ebiten.Image), op)
	}
}
