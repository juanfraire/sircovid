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
	ciudad1, barbijo, alchol, plasma, fondoNegro, home1, home, monedas, relato, papiro, ciudad, tpaper, money, meds, mmeds, mhome, mhome1, bread, clothes, cruz Objetos
	objScale                                                                                                                                                    = .3
	barHScale                                                                                                                                                   float64
	barWscale                                                                                                                                                   float64
	coinHScale                                                                                                                                                  float64
	coinWscale                                                                                                                                                  float64
	plasmaHScale                                                                                                                                                float64
	plasmaWScale                                                                                                                                                float64
	// alcholHScale := float64(alchol.FrameHeight) * objScale
	// alcholWScale := float64(alchol.FrameWidth) * objScale
)

//cartFarmacy, cartSupermarket, cartStore, cartBanck

func initObjetos() {
	//objetos

	ciudad1.FrameOX = 0
	ciudad1.FrameOY = 0
	ciudad1.FrameNum = 1
	ciudad1.FrameWidth = 640
	ciudad1.FrameHeight = 320
	ciudad1.X = float64(0)
	ciudad1.Y = float64(0)
	ciudad1.img, _, err = ebitenutil.NewImageFromFile(`sircovid\data\ciudadFere.png`, ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}

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
	alchol.X = float64(1010)
	alchol.Y = float64(470)
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

	//realto inicio juego
	relato.FrameOX = 0
	relato.FrameOY = 0
	relato.FrameNum = 1
	relato.FrameWidth = 1500
	relato.FrameHeight = 2500
	relato.X = float64(230)
	relato.Y = float64(100)
	relato.img, _, err = ebitenutil.NewImageFromFile(`sircovid\data\relato2.png`, ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
	papiro.FrameOX = 0
	papiro.FrameOX = 0
	papiro.FrameOY = 0
	papiro.FrameNum = 1
	papiro.FrameWidth = 1288
	papiro.FrameHeight = 898
	papiro.X = float64(-90)
	papiro.Y = float64(-50)
	papiro.img, _, err = ebitenutil.NewImageFromFile(`sircovid\data\papiro.png`, ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
	cruz.FrameOX = 0
	cruz.FrameOX = 0
	cruz.FrameOY = 0
	cruz.FrameNum = 1
	cruz.FrameWidth = 862
	cruz.FrameHeight = 370
	cruz.X = float64(775)
	cruz.Y = float64(5)
	cruz.img, _, err = ebitenutil.NewImageFromFile(`sircovid\data\cruz.png`, ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}

	fondoNegro.FrameOX = 0
	fondoNegro.FrameOX = 0
	fondoNegro.FrameOY = 0
	fondoNegro.FrameNum = 1
	fondoNegro.FrameWidth = 862
	fondoNegro.FrameHeight = 250
	fondoNegro.X = float64(165)
	fondoNegro.Y = float64(170)
	fondoNegro.img, _, err = ebitenutil.NewImageFromFile(`sircovid\data\fondoNegro.png`, ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
	mhome.FrameOX = 0
	mhome.FrameOY = 0
	mhome.FrameNum = 1
	mhome.FrameWidth = 388
	mhome.FrameHeight = 757
	mhome.X = float64(0)
	mhome.Y = float64(250)
	mhome.img, _, err = ebitenutil.NewImageFromFile(`sircovid\data\mission-home.png`, ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
	mhome1.FrameOX = 0
	mhome1.FrameOY = 0
	mhome1.FrameNum = 1
	mhome1.FrameWidth = 388
	mhome1.FrameHeight = 757
	mhome1.X = float64(80)
	mhome1.Y = float64(50)
	mhome1.img, _, err = ebitenutil.NewImageFromFile(`sircovid\data\mission-home.png`, ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
	money.FrameOX = 0
	money.FrameOY = 0
	money.FrameNum = 1
	money.FrameWidth = 994
	money.FrameHeight = 538
	money.X = float64(410)
	money.Y = float64(10)
	money.img, _, err = ebitenutil.NewImageFromFile(`sircovid\data\mission-money.png`, ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
	meds.FrameOX = 0
	meds.FrameOY = 0
	meds.FrameNum = 1
	meds.FrameWidth = 945
	meds.FrameHeight = 759
	meds.X = float64(90)
	meds.Y = float64(20)
	meds.img, _, err = ebitenutil.NewImageFromFile(`sircovid\data\mission-meds.png`, ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
	bread.FrameOX = 0
	bread.FrameOY = 0
	bread.FrameNum = 1
	bread.FrameWidth = 1047
	bread.FrameHeight = 503
	bread.X = float64(710)
	bread.Y = float64(190)
	bread.img, _, err = ebitenutil.NewImageFromFile(`sircovid\data\mission-bread.png`, ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
	clothes.FrameOX = 0
	clothes.FrameOY = 0
	clothes.FrameNum = 1
	clothes.FrameWidth = 1079
	clothes.FrameHeight = 503
	clothes.X = float64(560)
	clothes.Y = float64(190)
	clothes.img, _, err = ebitenutil.NewImageFromFile(`sircovid\data\mission-clothes.png`, ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
	tpaper.FrameOX = 0
	tpaper.FrameOY = 0
	tpaper.FrameNum = 1
	tpaper.FrameWidth = 1151
	tpaper.FrameHeight = 388
	tpaper.X = float64(570)
	tpaper.Y = float64(400)
	tpaper.img, _, err = ebitenutil.NewImageFromFile(`sircovid\data\mission-tpaper.png`, ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
}

func dibujarNiveles(screen *ebiten.Image) {
	// dibujarObjetos(mhome, screen)
	// dibujarObjetos(money, screen)
	// if Game1.numPlayers == 2 {
	// 	dibujarObjetos(mhome1, screen)
	// }
	// if Level == 1 || Level == 5 || Level == 6 || Level == 8 {
	// 	dibujarObjetos(meds, screen)
	// }
	// if Level == 2 || Level == 5 || Level == 7 || Level == 10 {
	// 	dibujarObjetos(bread, screen)
	// }
	// if Level == 3 || Level == 7 || Level == 8 || Level == 9 {
	// 	dibujarObjetos(clothes, screen)
	// }
	// if Level == 4 || Level == 6 || Level == 9 || Level == 10 {
	// 	dibujarObjetos(tpaper, screen)
	// }
}

func dibujarObjetos(B Objetos, screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	if B != relato && B != papiro && B != fondoNegro && B != ciudad1 {
		op.GeoM.Scale(objScale, objScale)
	}
	op.GeoM.Translate(B.X, B.Y)
	bx, by := B.FrameOX, B.FrameOY
	screen.DrawImage(B.img.SubImage(image.Rect(bx, by, bx+B.FrameWidth, by+B.FrameHeight)).(*ebiten.Image), op)
}
