package main

import (
	"image"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten"
)

var (
	num    = rand.Intn(5)
	tmp    int
	cambio = rand.Intn(50) + 50
	count  int
	ok     bool
)

func init() {
	rand.Seed(time.Now().UnixNano())
	//hombre
	hombre.FrameOX = 0
	hombre.FrameOY = 48
	hombre.FrameNum = 1
	hombre.X = float64(350)
	hombre.Y = float64(290)
	hombre.FrameWidth = 32
	hombre.FrameHeight = 48
	//mujer
	mujer.FrameOX = 0
	mujer.FrameOY = 48
	mujer.FrameNum = 1
	mujer.X = float64(screenWidth - 50)
	mujer.Y = float64(290)
	mujer.FrameWidth = 32
	mujer.FrameHeight = 48

}

func moverEnemigos(h humanos) humanos {
	// var rand int
	h.FrameNum = 4
	count++

	x1, y1 := h.X, h.Y
	_, _, ok = obstaculos(h.X+20, h.Y+32, x1, y1)
	if count == cambio {
		for tmp = num; tmp == num || tmp == 0; tmp = rand.Intn(5) {
		}
		cambio += rand.Intn(100) + 50
		num = tmp

	} else if ok {
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
	}

	switch num {
	case 1:
		h.FrameOY = 48
		h.X--
	case 2:
		h.FrameOY = 96
		h.X++
	case 3:
		h.FrameOY = 144
		h.Y--
	case 4:
		h.FrameOY = 0
		h.Y++
	}
	return h
}

func dibujarEnemigos(E humanos, screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(.8, .8)
	op.GeoM.Translate(E.X, E.Y)
	j := (count1 / 7) % E.FrameNum
	hx, hy := E.FrameOX+j*E.FrameWidth, E.FrameOY
	screen.DrawImage(E.img.SubImage(image.Rect(hx, hy, hx+E.FrameWidth, hy+E.FrameHeight)).(*ebiten.Image), op)
}
