package main

import (
	"image"
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

type Objetos struct {
	FrameOX     int
	FrameOY     int
	FrameNum    int
	FrameWidth  int
	FrameHeight int
	X           float64
	Y           float64
	img         *ebiten.Image
}

var barbijo, nextLevel, cartFarmacy, cartSupermarket, cartStore, cartBanck Objetos

func initObjetos() {
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
	nextLevel.FrameOY = 77
	nextLevel.FrameNum = 1
	nextLevel.FrameWidth = 170
	nextLevel.FrameHeight = 120
	nextLevel.X = float64(695)
	nextLevel.Y = float64(440)
	nextLevel.img, _, err = ebitenutil.NewImageFromFile(`sircovid\data\carteles.png`, ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
	//carteles
	cartFarmacy.FrameOX = 0
	cartFarmacy.FrameOY = 0
	cartFarmacy.FrameNum = 1
	cartFarmacy.FrameWidth = 290
	cartFarmacy.FrameHeight = 78
	cartFarmacy.X = float64(380)
	cartFarmacy.Y = float64(40)
	cartFarmacy.img = nextLevel.img

	cartSupermarket.FrameOX = 175
	cartSupermarket.FrameOY = 77
	cartSupermarket.FrameNum = 1
	cartSupermarket.FrameWidth = 440
	cartSupermarket.FrameHeight = 78
	cartSupermarket.X = float64(165)
	cartSupermarket.Y = float64(80)
	cartSupermarket.img = cartFarmacy.img

	cartStore.FrameOX = 300
	cartStore.FrameOY = 0
	cartStore.FrameNum = 1
	cartStore.FrameWidth = 190
	cartStore.FrameHeight = 78
	cartStore.X = float64(0)
	cartStore.Y = float64(260)
	cartStore.img = cartFarmacy.img

	cartBanck.FrameOX = 500
	cartBanck.FrameOY = 0
	cartBanck.FrameNum = 1
	cartBanck.FrameWidth = 210
	cartBanck.FrameHeight = 78
	cartBanck.X = float64(620)
	cartBanck.Y = float64(20)
	cartBanck.img = cartFarmacy.img
}

func dibujarObjetos(B Objetos, screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(.3, .3)
	op.GeoM.Translate(B.X, B.Y)
	bx, by := B.FrameOX, B.FrameOY
	screen.DrawImage(B.img.SubImage(image.Rect(bx, by, bx+B.FrameWidth, by+B.FrameHeight)).(*ebiten.Image), op)
}
