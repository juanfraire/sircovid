package main

import (
	"image"
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/inpututil"
)

type player struct {
	humanos
	vidas         int
	v             float32
	señalador     int
	a, b, c, d    int
	MovX          int
	MovY          int
	Inmune        bool
	CountPoder    int
	Coins         int
	Fast          bool
	Compras       bool
	CompleteLevel bool
}

var (
	player1, player2 player
	humano1, humano2 humanos
	// nivel            = int(1)
	plyrScale = .65
	hgt       float64
	wth       float64
	hack      bool
)

func initPlayer() {

	//////////////   Imangen VIEJO  //////////////////////////////
	humano1.img[0], _, err = ebitenutil.NewImageFromFile(`sircovid\data\player1.png`, ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
	//imagen chica
	humano2.img[0], _, err = ebitenutil.NewImageFromFile(`sircovid\data\player2.png`, ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
	// player1

	player1.FrameOX[0] = 0
	player1.FrameOY[0] = 96
	player1.FrameNum[0] = 1
	player1.X[0] = 15
	player1.Y[0] = -40
	player1.FrameWidth[0] = 32
	player1.FrameHeight[0] = 48
	player1.MovX = 0
	player1.MovY = 0
	//player
	// player1.humanos = player1
	player1.vidas = 3
	player1.v = 0
	player1.señalador = 0
	player1.a, player1.b, player1.c, player1.d = 0, 0, 0, 0
	player1.humanos.img = humano1.img

	//player2
	player2.FrameOX[0] = 0
	player2.FrameOY[0] = 96
	player2.FrameNum[0] = 1
	player2.X[0] = 130
	player2.Y[0] = -40
	player2.FrameWidth[0] = 32
	player2.FrameHeight[0] = 48
	player1.MovX = 0
	player1.MovY = 0

	player2.vidas = 3
	player2.v = 0
	player2.señalador = 1
	player2.a, player2.b, player2.c, player2.d = 0, 0, 0, 0
	player2.humanos.img = humano2.img

}
func pasarNivelPlayer() {

	//player1
	player1.FrameNum[0] = 1
	player1.X[0] = 15
	player1.Y[0] = -40
	player1.MovX = 0
	player1.MovY = 0
	player1.a, player1.b, player1.c, player1.d = 0, 0, 0, 0

	//player2
	player2.FrameNum[0] = 1
	player2.X[0] = 130
	player2.Y[0] = -40
	player2.MovX = 0
	player2.MovY = 0
	player2.a, player2.b, player2.c, player2.d = 0, 0, 0, 0

}
func moverPlayer(p player) player {
	// leer tecla
	if inpututil.IsKeyJustPressed(ebiten.KeyRight) && p.señalador == 0 || inpututil.IsKeyJustPressed(ebiten.KeyD) && p.señalador == 1 {
		p.a = 1
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyLeft) && p.señalador == 0 || inpututil.IsKeyJustPressed(ebiten.KeyA) && p.señalador == 1 {
		p.b = 1
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyUp) && p.señalador == 0 || inpututil.IsKeyJustPressed(ebiten.KeyW) && p.señalador == 1 {
		p.c = 1
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyDown) && p.señalador == 0 || inpututil.IsKeyJustPressed(ebiten.KeyS) && p.señalador == 1 {
		p.d = 1
	}

	if p.a == 1 && p.MovY != 1 && p.MovY != 2 {
		p.FrameOX[0] = 0
		p.FrameOY[0] = 96
		p.FrameNum[0] = 3
		p.MovX = 1
	}
	if inpututil.IsKeyJustReleased(ebiten.KeyRight) && p.señalador == 0 || inpututil.IsKeyJustReleased(ebiten.KeyD) && p.señalador == 1 {
		p.FrameOX[0] = 32
		p.FrameNum[0] = 1
		p.MovX = 0
		p.a = 0
	}
	if p.b == 1 && p.MovY != 1 && p.MovY != 2 {
		p.FrameOX[0] = 0
		p.FrameOY[0] = 48
		p.FrameNum[0] = 3
		p.MovX = 2
	}
	if inpututil.IsKeyJustReleased(ebiten.KeyLeft) && p.señalador == 0 || inpututil.IsKeyJustReleased(ebiten.KeyA) && p.señalador == 1 {
		p.FrameOX[0] = 32
		p.FrameNum[0] = 1
		p.MovX = 0
		p.b = 0
	}
	if p.c == 1 && p.MovX != 1 && p.MovX != 2 {
		p.FrameOX[0] = 0
		p.FrameOY[0] = 144
		p.FrameNum[0] = 3
		p.MovY = 1
	}
	if inpututil.IsKeyJustReleased(ebiten.KeyUp) && p.señalador == 0 || inpututil.IsKeyJustReleased(ebiten.KeyW) && p.señalador == 1 {
		p.FrameOX[0] = 32
		p.FrameNum[0] = 1
		p.MovY = 0
		p.c = 0
	}
	if p.d == 1 && p.MovX != 1 && p.MovX != 2 {
		p.FrameOX[0] = 0
		p.FrameOY[0] = 0
		p.FrameNum[0] = 3
		p.MovY = 2
	}
	if inpututil.IsKeyJustReleased(ebiten.KeyDown) && p.señalador == 0 || inpututil.IsKeyJustReleased(ebiten.KeyS) && p.señalador == 1 {
		p.FrameOX[0] = 32
		p.FrameNum[0] = 1
		p.MovY = 0
		p.d = 0
	}

	// trasladar player1
	if Level == 2 && p.posicionInicial[0] != 1 {
		p.posicionInicial[0] = 1
	}
	var X1 = p.X[0]
	var Y1 = p.Y[0]
	switch {
	case p.MovX == 1 && p.Fast:
		p.X[0] = p.X[0] + 2.3
	case p.MovX == 2 && p.Fast:
		p.X[0] = p.X[0] - 2.3
	case p.MovY == 1 && p.Fast:
		p.Y[0] = p.Y[0] - 2.3
	case p.MovY == 2 && p.Fast:
		p.Y[0] = p.Y[0] + 2.3

	case p.MovX == 1:
		p.X[0] += 1.5
	case p.MovX == 2:
		p.X[0] -= 1.5
	case p.MovY == 1:
		p.Y[0] -= 1.5
	case p.MovY == 2:
		p.Y[0] += 1.5
	}
	p.X[0], p.Y[0], obs = obstaculos(p.X[0], p.Y[0], X1, Y1)
	if obs {
		p.FrameOX[0] = 32
		p.FrameNum[0] = 1
	}
	///ENTRADAS a puertas///
	switch {
	case p.c == 1 && ((p.X[0] > 125 && p.X[0] < 135 && p.Y[0] < 98 && p.Y[0] > 95) || (p.X[0] > 10 && p.X[0] < 25 && p.Y[0] < 275 && p.Y[0] > 270) || (p.X[0] > 635 && p.X[0] < 645 && p.Y[0] < 50 && p.Y[0] > 44) || (p.X[0] > 415 && p.X[0] < 425 && p.Y[0] < 56 && p.Y[0] > 50) || (p.X[0] > 193 && p.X[0] < 258 && p.Y[0] < 110 && p.Y[0] > 78)):
		p.Y[0] = -40
		fondo.Pause()
		sPuerta.Play()
		sPuerta.Rewind()
		//pharmacy
	case p.c == 1 && p.X[0] > 275 && p.X[0] < 285 && p.Y[0] < 81 && p.Y[0] > 77:
		p.Y[0] = -40
		p.a, p.b, p.c, p.d = 0, 0, 0, 0
		p.MovX = 0
		p.Compras = true
		farmacia = true
		//bakery
	case p.c == 1 && p.X[0] > 945 && p.X[0] < 955 && p.Y[0] < 228 && p.Y[0] > 224:
		p.a, p.b, p.c, p.d = 0, 0, 0, 0
		p.MovX = 0
		p.Y[0] = -40
		p.Compras = true
		bakery = true
		//para WIN
	case p.c == 1 && p.X[0] > 753 && p.X[0] < 763 && p.Y[0] < 439 && p.Y[0] > 430:
		p.Y[0] = -40
		ModeWin = true
		fondo.Pause()
		//PARA MART
	case p.c == 1 && p.X[0] > 810 && p.X[0] < 820 && p.Y[0] < 228 && p.Y[0] > 224:
		p.a, p.b, p.c, p.d = 0, 0, 0, 0
		p.MovX = 0
		p.Y[0] = -40
		p.Compras = true
		mart = true
		//SUPERMAKET
	case p.c == 1 && (p.X[0] > 845 && p.X[0] < 855 && p.Y[0] < 437 && p.Y[0] > 434):
		p.a, p.b, p.c, p.d = 0, 0, 0, 0
		p.MovX = 0
		p.Y[0] = -40
		p.Compras = true
		supermarket = true

		///SALIDAS///
		//edificio arriba izquierda
	case p.Y[0] < -36 && p.Y[0] > -39 && p.X[0] > 124 && p.X[0] < 136:
		p.Y[0] = 99
		//edificio abajo a la izquieda
	case p.Y[0] < -36 && p.Y[0] > -39 && p.X[0] > 9 && p.X[0] < 26:
		p.Y[0] = 276
		//banco
	case p.Y[0] < -36 && p.Y[0] > -39 && p.X[0] > 634 && p.X[0] < 646:
		p.Y[0] = 51
		//local a la izquieda del banco
	case p.Y[0] < -36 && p.Y[0] > -39 && p.X[0] > 414 && p.X[0] < 426:
		p.Y[0] = 57
		//tienda a la derecha del edeficio arriba a la izquierda
	case p.Y[0] < -36 && p.Y[0] > -39 && p.X[0] > 193 && p.X[0] < 265:
		p.Y[0] = 110

		//mart
	case p.Y[0] < -36 && p.Y[0] > -39 && p.X[0] > 812 && p.X[0] < 828:
		p.Y[0] = 230
		mart = false
		//Pharmacy
	case p.Y[0] < -36 && p.Y[0] > -39 && p.X[0] > 274 && p.X[0] < 286:
		p.Y[0] = 82
		farmacia = false
		//salida de Bakery
	case p.Y[0] < -36 && p.Y[0] > -39 && p.X[0] > 943 && p.X[0] < 957:
		p.Y[0] = 229
		bakery = false
	//supermarket
	case p.Y[0] < -36 && p.Y[0] > -39 && p.X[0] > 844 && p.X[0] < 856:
		p.Y[0] = 436
		supermarket = false
	}
	//Vuelta a la realidad
	if p.Y[0] < 0 {
		p.X[0] = X1
	}
	if p.Y[0] < -40 {
		p.Y[0] = -40
	}
	// coder shortcuts

	if inpututil.IsKeyJustPressed(ebiten.KeyControl) {
		hack = true
	}
	switch {
	case hack && inpututil.IsKeyJustPressed(ebiten.KeyF):
		p.Fast = !p.Fast
	case hack && inpututil.IsKeyJustPressed(ebiten.KeyN):
		pasarNivel()
	case hack && inpututil.IsKeyJustPressed(ebiten.KeyI):
		p.Inmune = !p.Inmune
	}
	return p
}

func vida(h humanos, p player, b Objetos, pl Objetos) (player, Objetos, Objetos) {
	barHScale = float64(barbijo.FrameHeight) * objScale
	barWscale = float64(barbijo.FrameWidth) * objScale
	coinHScale = float64(monedas.FrameHeight) * objScale
	coinWscale = float64(monedas.FrameWidth) * objScale
	// alcholHScale := float64(alchol.FrameHeight) * objScale
	// alcholWScale := float64(alchol.FrameWidth) * objScale
	plasmaHScale = float64(plasma.FrameHeight) * objScale
	plasmaWScale = float64(plasma.FrameWidth) * objScale
	hgt = float64(p.FrameHeight[0])*plyrScale - 2
	wth = float64(p.FrameWidth[0])*plyrScale - 6

	if p.Inmune != true {
		//pierde vidas con la nube
		for i := 0; i < nivel; i++ {
			nubX := nube1.X[i] * nubScale
			nubY := nube1.Y[i] * nubScale
			//pierde vidas con nube
			if p.X[0]+wth > nubX && p.X[0] < nubX+nubFrameWith && p.Y[0]+hgt > nubY && p.Y[0] < nubY+nubFrameHight && nube1.Alpha[i] > .3 {
				p.v += .1
			}
		}
		//pierde vidas con humanos
		for i := 0; i < nivel; i++ {
			if p.X[0]+20 > h.X[i] && p.X[0] < h.X[i]+20 && p.Y[0]+32 > h.Y[i] && p.Y[0] < h.Y[i]+32 {
				p.v += .5
			}
		}
	}
	//inmune con barbijo o alchol en gel
	if p.X[0]+wth > b.X && p.X[0] < b.X+barWscale && p.Y[0]+hgt > b.Y && p.Y[0]+hgt < b.Y+barHScale {
		sBarbijo.Play()
		sBarbijo.Rewind()
		b.X = 1500
		p.Inmune = true
		p.CountPoder = 600
	}
	if p.Inmune == true || p.Fast {

		p.CountPoder--
	}
	if p.CountPoder == 0 {
		p.Inmune = false
		p.Fast = false
	}

	//gana/pierde vida
	if p.X[0]+wth > pl.X && p.X[0] < pl.X+plasmaWScale && p.Y[0]+hgt > pl.Y && p.Y[0]+hgt < pl.Y+plasmaHScale {
		p.vidas++
		pl.X = 1500
	}
	if p.v >= 30 {
		p.vidas--
		sonidoVidas()
		p.v = 0
	}
	if p.vidas == 0 {
		ModeGameOver = true
		ModeGame = false
	}
	//gana monedas
	if p.X[0]+wth > monedas.X && p.X[0] < monedas.X+coinWscale && p.Y[0]+hgt > monedas.Y && p.Y[0]+hgt < monedas.Y+coinHScale {
		monedas.X = 1500
		p.Coins += 5
		sPuerta.Pause()
		sDinero.Play()
		sDinero.Rewind()
	}
	//PIERDE POR falta de plata
	if p.Coins < 2 && !p.CompleteLevel && monedas.X == 1500 {
		ModeGame = false
		ModeGameOver = true
	}

	return p, b, pl
}

func dibujarPlayer(P player, screen *ebiten.Image) {
	if ModePause {
		P.FrameNum[0] = 1
		P.FrameOX[0] = 0
	}
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(plyrScale, plyrScale)
	op.GeoM.Translate(P.X[0], P.Y[0])
	i := (count1 / 7) % P.FrameNum[0]
	sx, sy := P.FrameOX[0]+i*P.FrameWidth[0], P.FrameOY[0]
	screen.DrawImage(P.img[0].SubImage(image.Rect(sx, sy, sx+P.FrameWidth[0], sy+P.FrameHeight[0])).(*ebiten.Image), op)
}
