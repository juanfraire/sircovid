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
	barbijo, alchol, plasma, home1, home, monedas, relato, ciudad Objetos
	objScale                                                      = .3
	barHScale                                                     float64
	barWscale                                                     float64
	coinHScale                                                    float64
	coinWscale                                                    float64
	plasmaHScale                                                  float64
	plasmaWScale                                                  float64
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
	barbijo.X = float64(300)
	barbijo.Y = float64(500)
	barbijo.img, _, err = ebitenutil.NewImageFromFile(`sircovid\data\objetos.png`, ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
	plasma.FrameOX = 0
	plasma.FrameOY = 200
	plasma.FrameNum = 1
	plasma.FrameWidth = 60
	plasma.FrameHeight = 120
	plasma.X = float64(90)
	plasma.Y = float64(250)
	plasma.img = barbijo.img

	alchol.FrameOX = 0
	alchol.FrameOY = 320
	alchol.FrameNum = 1
	alchol.FrameWidth = 65
	alchol.FrameHeight = 120
	alchol.X = float64(1000)
	alchol.Y = float64(30)
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
	home.FrameOX = 0
	home.FrameOY = 195
	home.FrameNum = 1
	home.FrameWidth = 190
	home.FrameHeight = 120
	home.X = float64(0)
	home.Y = float64(260)
	home.img, _, err = ebitenutil.NewImageFromFile(`sircovid\data\carteles.png`, ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
	home1.FrameOX = 0
	home1.FrameOY = 195
	home1.FrameNum = 1
	home1.FrameWidth = 190
	home1.FrameHeight = 120
	home1.X = float64(110)
	home1.Y = float64(85)
	home1.img = home.img

	//realto inicio juego
	relato.FrameOX = 0
	relato.FrameOY = 0
	relato.FrameNum = 1
	relato.FrameWidth = 1500
	relato.FrameHeight = 2500
	relato.X = float64(30)
	relato.Y = float64(100)
	relato.img, _, err = ebitenutil.NewImageFromFile(`sircovid\data\relato1.png`, ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}

	ciudad.FrameOX = 0
	ciudad.FrameOY = 0
	ciudad.FrameNum = 1
	ciudad.FrameWidth = 1500
	ciudad.FrameHeight = 2500
	ciudad.X = float64(30)
	ciudad.Y = float64(100)
	ciudad.img, _, err = ebitenutil.NewImageFromFile(`sircovid\data\relato1.png`, ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}

}

func dibujarObjetos(B Objetos, screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	if B != relato {
		op.GeoM.Scale(objScale, objScale)
	}
	op.GeoM.Translate(B.X, B.Y)
	bx, by := B.FrameOX, B.FrameOY
	screen.DrawImage(B.img.SubImage(image.Rect(bx, by, bx+B.FrameWidth, by+B.FrameHeight)).(*ebiten.Image), op)
}
