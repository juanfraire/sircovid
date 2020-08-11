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
	barbijo, alchol, plasma, fondoNegro, fondoNegroCompras, fondoNegroPause, home1, home, monedas, relato, papiro, ciudad, tpaper, vaccine, money, meds, mmeds, mhome, mhome1, bread, clothes, cruz Objetos
	objScale                                                                                                                                                                                        = .3
	barHScale                                                                                                                                                                                       float64
	barWscale                                                                                                                                                                                       float64
	coinHScale                                                                                                                                                                                      float64
	coinWscale                                                                                                                                                                                      float64
	plasmaHScale                                                                                                                                                                                    float64
	plasmaWScale                                                                                                                                                                                    float64
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
	barbijo.X = float64(650)
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
	plasma.X = float64(489)
	plasma.Y = float64(73)
	plasma.img = barbijo.img

	alchol.FrameOX = 0
	alchol.FrameOY = 320
	alchol.FrameNum = 1
	alchol.FrameWidth = 65
	alchol.FrameHeight = 120
	alchol.X = float64(540)
	alchol.Y = float64(360)
	alchol.img = barbijo.img

	monedas.FrameOX = 0
	monedas.FrameOY = 440
	monedas.FrameNum = 1
	monedas.FrameWidth = 65
	monedas.FrameHeight = 80
	monedas.X = float64(99)
	monedas.Y = float64(99)
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
	cruz.X = float64(785)
	cruz.Y = float64(330)
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
	fondoNegroPause.FrameOX = 0
	fondoNegroPause.FrameOX = 0
	fondoNegroPause.FrameOY = 0
	fondoNegroPause.FrameNum = 1
	fondoNegroPause.FrameWidth = 862
	fondoNegroPause.FrameHeight = 250
	fondoNegroPause.X = float64(415)
	fondoNegroPause.Y = float64(215)
	fondoNegroPause.img, _, err = ebitenutil.NewImageFromFile(`sircovid\data\fondoNegro.png`, ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
	fondoNegroCompras.FrameOX = 0
	fondoNegroCompras.FrameOX = 0
	fondoNegroCompras.FrameOY = 0
	fondoNegroCompras.FrameNum = 1
	fondoNegroCompras.FrameWidth = 1100
	fondoNegroCompras.FrameHeight = 300
	fondoNegroCompras.X = float64(30)
	fondoNegroCompras.Y = float64(100)
	fondoNegroCompras.img, _, err = ebitenutil.NewImageFromFile(`sircovid\data\fondoNegro1.png`, ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
	mhome.FrameOX = 0
	mhome.FrameOY = 0
	mhome.FrameNum = 1
	mhome.FrameWidth = 388
	mhome.FrameHeight = 757
	mhome.X = float64(0)
	mhome.Y = float64(200)
	mhome.img, _, err = ebitenutil.NewImageFromFile(`sircovid\data\mission-home.png`, ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
	mhome1.FrameOX = 0
	mhome1.FrameOY = 0
	mhome1.FrameNum = 1
	mhome1.FrameWidth = 388
	mhome1.FrameHeight = 757
	mhome1.X = float64(220)
	mhome1.Y = float64(200)
	mhome1.img, _, err = ebitenutil.NewImageFromFile(`sircovid\data\mission-home.png`, ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
	money.FrameOX = 0
	money.FrameOY = 0
	money.FrameNum = 1
	money.FrameWidth = 1070
	money.FrameHeight = 423
	money.X = float64(50)
	money.Y = float64(50)
	money.img, _, err = ebitenutil.NewImageFromFile(`sircovid\data\mission-money.png`, ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
	meds.FrameOX = 0
	meds.FrameOY = 0
	meds.FrameNum = 1
	meds.FrameWidth = 945
	meds.FrameHeight = 759
	meds.X = float64(230)
	meds.Y = float64(310)
	meds.img, _, err = ebitenutil.NewImageFromFile(`sircovid\data\mission-meds.png`, ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
	bread.FrameOX = 0
	bread.FrameOY = 0
	bread.FrameNum = 1
	bread.FrameWidth = 1047
	bread.FrameHeight = 503
	bread.X = float64(750)
	bread.Y = float64(145)
	bread.img, _, err = ebitenutil.NewImageFromFile(`sircovid\data\mission-bread.png`, ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
	clothes.FrameOX = 0
	clothes.FrameOY = 0
	clothes.FrameNum = 1
	clothes.FrameWidth = 941
	clothes.FrameHeight = 694
	clothes.X = float64(270)
	clothes.Y = float64(0)
	clothes.img, _, err = ebitenutil.NewImageFromFile(`sircovid\data\mission-clothes.png`, ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
	tpaper.FrameOX = 0
	tpaper.FrameOY = 0
	tpaper.FrameNum = 1
	tpaper.FrameWidth = 1151
	tpaper.FrameHeight = 388
	tpaper.X = float64(560)
	tpaper.Y = float64(135)
	tpaper.img, _, err = ebitenutil.NewImageFromFile(`sircovid\data\mission-tpaper.png`, ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
	tpaper.FrameOX = 0
	tpaper.FrameOY = 0
	tpaper.FrameNum = 1
	tpaper.FrameWidth = 1151
	tpaper.FrameHeight = 388
	tpaper.X = float64(560)
	tpaper.Y = float64(135)
	tpaper.img, _, err = ebitenutil.NewImageFromFile(`sircovid\data\mission-tpaper.png`, ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
	vaccine.FrameOX = 0
	vaccine.FrameOY = 0
	vaccine.FrameNum = 1
	vaccine.FrameWidth = 1106
	vaccine.FrameHeight = 468
	vaccine.X = float64(560)
	vaccine.Y = float64(335)
	vaccine.img, _, err = ebitenutil.NewImageFromFile(`sircovid\data\mission-vaccine.png`, ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
}

func dibujarNiveles(screen *ebiten.Image) {
	dibujarObjetos(mhome, screen)
	dibujarObjetos(money, screen)
	if Game1.numPlayers == 2 {
		dibujarObjetos(mhome1, screen)
	}
	if Level == 1 || Level == 5 || Level == 6 || Level == 8 {
		dibujarObjetos(meds, screen)
	}
	if Level == 2 || Level == 5 || Level == 7 || Level == 10 {
		dibujarObjetos(bread, screen)
	}
	if Level == 3 || Level == 7 || Level == 8 || Level == 9 {
		dibujarObjetos(clothes, screen)
	}
	if Level == 4 || Level == 6 || Level == 9 || Level == 10 {
		dibujarObjetos(tpaper, screen)
	}
	if Level > 10 {
		dibujarObjetos(vaccine, screen)
	}
}

func dibujarObjetos(B Objetos, screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	if B != relato && B != papiro && B != fondoNegro && B != fondoNegroCompras {
		op.GeoM.Scale(objScale, objScale)
	}
	op.GeoM.Translate(B.X, B.Y)
	bx, by := B.FrameOX, B.FrameOY
	screen.DrawImage(B.img.SubImage(image.Rect(bx, by, bx+B.FrameWidth, by+B.FrameHeight)).(*ebiten.Image), op)
}
