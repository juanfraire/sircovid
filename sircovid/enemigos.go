package main

import (
	"image"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten"
)

type enemigos struct {
	humanos
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

var (
	num = rand.Intn(5)

	count int
	ok    bool
)

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
		cambio += rand.Intn(50) + 20
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
		cambio = count + 15
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
