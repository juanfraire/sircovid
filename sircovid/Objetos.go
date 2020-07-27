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

var (
	barbijo, alchol, plasma, nextLevel, monedas Objetos
	objScale                                    = .3
	barHScale                                   float64
	barWscale                                   float64
	coinHScale                                  float64
	coinWscale                                  float64
	plasmaHScale                                float64
	plasmaWScale                                float64
	// alcholHScale := float64(alchol.FrameHeight) * objScale
	// alcholWScale := float64(alchol.FrameWidth) * objScale
)

//cartFarmacy, cartSupermarket, cartStore, cartBanck

func initObjetos() {
	//objetos
	barbijo.FrameOX = 0
	barbijo.FrameOY = 160
	barbijo.FrameNum = 1
	barbijo.FrameWidth = 105
	barbijo.FrameHeight = 40
	barbijo.X = float64(635)
	barbijo.Y = float64(150)
	barbijo.img, _, err = ebitenutil.NewImageFromFile(`sircovid\data\objetos.png`, ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
	plasma.FrameOX = 0
	plasma.FrameOY = 200
	plasma.FrameNum = 1
	plasma.FrameWidth = 60
	plasma.FrameHeight = 120
	plasma.X = float64(30)
	plasma.Y = float64(450)
	plasma.img = barbijo.img

	alchol.FrameOX = 0
	alchol.FrameOY = 320
	alchol.FrameNum = 1
	alchol.FrameWidth = 65
	alchol.FrameHeight = 120
	alchol.X = float64(90)
	alchol.Y = float64(250)
	alchol.img = barbijo.img

	monedas.FrameOX = 0
	monedas.FrameOY = 440
	monedas.FrameNum = 1
	monedas.FrameWidth = 65
	monedas.FrameHeight = 80
	monedas.X = float64(640)
	monedas.Y = float64(70)
	monedas.img = barbijo.img

	//carteles
	nextLevel.FrameOX = 0
	nextLevel.FrameOY = 77
	nextLevel.FrameNum = 1
	nextLevel.FrameWidth = 170
	nextLevel.FrameHeight = 120
	nextLevel.X = float64(screenWidth) - float64(nextLevel.FrameWidth)*.3
	nextLevel.Y = float64(440)
	nextLevel.img, _, err = ebitenutil.NewImageFromFile(`sircovid\data\carteles.png`, ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
	//carteles
	// cartFarmacy.FrameOX = 0
	// cartFarmacy.FrameOY = 0
	// cartFarmacy.FrameNum = 1
	// cartFarmacy.FrameWidth = 290
	// cartFarmacy.FrameHeight = 78
	// cartFarmacy.X = float64(380)
	// cartFarmacy.Y = float64(40)
	// cartFarmacy.img = nextLevel.img

	// cartSupermarket.FrameOX = 175
	// cartSupermarket.FrameOY = 77
	// cartSupermarket.FrameNum = 1
	// cartSupermarket.FrameWidth = 440
	// cartSupermarket.FrameHeight = 78
	// cartSupermarket.X = float64(165)
	// cartSupermarket.Y = float64(80)
	// cartSupermarket.img = cartFarmacy.img

	// cartStore.FrameOX = 300
	// cartStore.FrameOY = 0
	// cartStore.FrameNum = 1
	// cartStore.FrameWidth = 190
	// cartStore.FrameHeight = 78
	// cartStore.X = float64(0)
	// cartStore.Y = float64(260)
	// cartStore.img = cartFarmacy.img

	// cartBanck.FrameOX = 500
	// cartBanck.FrameOY = 0
	// cartBanck.FrameNum = 1
	// cartBanck.FrameWidth = 210
	// cartBanck.FrameHeight = 78
	// cartBanck.X = float64(620)
	// cartBanck.Y = float64(20)
	// cartBanck.img = cartFarmacy.img
}

func dibujarObjetos(B Objetos, screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(objScale, objScale)
	op.GeoM.Translate(B.X, B.Y)
	bx, by := B.FrameOX, B.FrameOY
	screen.DrawImage(B.img.SubImage(image.Rect(bx, by, bx+B.FrameWidth, by+B.FrameHeight)).(*ebiten.Image), op)
}
