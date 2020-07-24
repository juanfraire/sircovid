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

var barbijo, nextLevel, cartFarmacy, cartSupermarket, cartStore sumVidas

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
	nextLevel.FrameOX = 0
	nextLevel.FrameOY = 0
	nextLevel.FrameNum = 1
	nextLevel.FrameWidth = 240
	nextLevel.FrameHeight = 240
	nextLevel.X = float64(695)
	nextLevel.Y = float64(450)
	nextLevel.img, _, err = ebitenutil.NewImageFromFile(`sircovid\data\NextLevel.png`, ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
	//carteles
	cartFarmacy.FrameOX = 0
	cartFarmacy.FrameOY = 0
	cartFarmacy.FrameNum = 1
	cartFarmacy.FrameWidth = 310
	cartFarmacy.FrameHeight = 135
	cartFarmacy.X = float64(380)
	cartFarmacy.Y = float64(20)
	cartFarmacy.img, _, err = ebitenutil.NewImageFromFile(`sircovid\data\carteles.png`, ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}

	cartSupermarket.FrameOX = 0
	cartSupermarket.FrameOY = 260
	cartSupermarket.FrameNum = 1
	cartSupermarket.FrameWidth = 500
	cartSupermarket.FrameHeight = 135
	cartSupermarket.X = float64(160)
	cartSupermarket.Y = float64(70)
	cartSupermarket.img = cartFarmacy.img

	cartStore.FrameOX = 0
	cartStore.FrameOY = 135
	cartStore.FrameNum = 1
	cartStore.FrameWidth = 400
	cartStore.FrameHeight = 135
	cartStore.X = float64(-8)
	cartStore.Y = float64(250)
	cartStore.img = cartFarmacy.img
}

func dibujarSumVidas(B sumVidas, screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(.3, .3)
	op.GeoM.Translate(B.X, B.Y)
	bx, by := B.FrameOX, B.FrameOY
	screen.DrawImage(B.img.SubImage(image.Rect(bx, by, bx+B.FrameWidth, by+B.FrameHeight)).(*ebiten.Image), op)
}
