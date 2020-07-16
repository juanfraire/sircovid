package main

import (
	"fmt"
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
	num      = rand.Intn(5)
	count    int
	ok       bool
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

	enemigo1.img, _, err = ebitenutil.NewImageFromFile(`sircovid\data\enemigo1.png`, ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}

	rand.Seed(time.Now().UnixNano())
	//enemigo2
	enemigo2.FrameOX = 0
	enemigo2.FrameOY = 48
	enemigo2.FrameNum = 1
	enemigo2.X = float64(screenWidth - 50)
	enemigo2.Y = float64(290)
	enemigo2.FrameWidth = 32
	enemigo2.FrameHeight = 48
	enemigo2.num = rand.Intn(5)
	enemigo2.cambio = rand.Intn(50) + 50

	enemigo2.img, _, err = ebitenutil.NewImageFromFile(`sircovid\data\enemigo2.png`, ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
}

func moverHumanos(E humanos) humanos {

	E.FrameNum = 4

	switch E.num {
	case 0:
		E.FrameNum = 1
	case 1:
		E.FrameOY = 48
		E.X--
	case 2:
		E.FrameOY = 96
		E.X++
	case 3:
		E.FrameOY = 144
		E.Y--
	case 4:
		E.FrameOY = 0
		E.Y++
	}

	return E
}

func cambioDireccion(E humanos, count int) humanos {

	var tmp int
	//hay un problema con countH vs count
	if count >= E.cambio {
		for tmp = E.num; tmp == E.num; tmp = rand.Intn(5) {
		}
		E.cambio += rand.Intn(100) + 20
		E.num = tmp
	}
	fmt.Println(count)
	return E

}

func obstEnemigo(E humanos) humanos {
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
		E.cambio = count + 5
	}

	return E
}

func dibujarEnemigos(E humanos, screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(.8, .8)
	op.GeoM.Translate(E.X, E.Y)
	j := (count1 / 7) % E.FrameNum
	hx, hy := E.FrameOX+j*E.FrameWidth, E.FrameOY
	screen.DrawImage(E.img.SubImage(image.Rect(hx, hy, hx+E.FrameWidth, hy+E.FrameHeight)).(*ebiten.Image), op)
}
