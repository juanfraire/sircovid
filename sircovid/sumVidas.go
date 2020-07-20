package main

import (
	"image"
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

type sumVidas struct {
	FrameOX     int
	FrameOY     int
	FrameNum    int
	FrameWidth  int
	FrameHeight int
	X           float64
	Y           float64
	img         *ebiten.Image
}

var barbijo sumVidas

func initSumarVidas() {
	barbijo.FrameOX = 0
	barbijo.FrameOY = 74
	barbijo.FrameNum = 1
	barbijo.FrameWidth = 105
	barbijo.FrameHeight = 48
	barbijo.X = float64(630)
	barbijo.Y = float64(150)
	barbijo.img, _, err = ebitenutil.NewImageFromFile(`sircovid\data\barbijo.png`, ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
}

func dibujarSumVidas(B sumVidas, screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(.2, .2)
	op.GeoM.Translate(B.X, B.Y)
	bx, by := B.FrameOX+B.FrameWidth, B.FrameOY
	screen.DrawImage(B.img.SubImage(image.Rect(bx, by, bx+B.FrameWidth, by+B.FrameHeight)).(*ebiten.Image), op)
}
