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
	vidas       int
	v           float32
	señalador   int
	a, b, c, d  int
	MovX        int
	MovY        int
	Inmune      bool
	CountInmune int
	Coins       int
}

var player1, player2 player
var humano1, humano2 humanos
var nivel = int(1)

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
	player1.X[0] = float64(365)
	player1.Y[0] = float64(300)
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
	player2.X[0] = float64(365)
	player2.Y[0] = float64(350)
	player2.FrameWidth[0] = 32
	player2.FrameHeight[0] = 48
	player1.MovX = 0
	player1.MovY = 0

	// player2.humanos = chica
	player2.vidas = 3
	player2.v = 0
	player2.señalador = 1
	player2.a, player2.b, player2.c, player2.d = 0, 0, 0, 0
	player2.humanos.img = humano2.img

}
func pasarNivelPlayer() {
	nivel *= 2
	if nivel > 13 {
		nivel = 13
	}
	//player1
	player1.FrameNum[0] = 1
	player1.X[0] = float64(365)
	player1.Y[0] = float64(300)
	player1.MovX = 0
	player1.MovY = 0
	player1.a, player1.b, player1.c, player1.d = 0, 0, 0, 0

	//player2
	player2.FrameNum[0] = 1
	player2.X[0] = float64(365)
	player2.Y[0] = float64(350)
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
	if inpututil.IsKeyJustReleased(ebiten.KeyRight) || inpututil.IsKeyJustReleased(ebiten.KeyD) {
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
	if inpututil.IsKeyJustReleased(ebiten.KeyLeft) || inpututil.IsKeyJustReleased(ebiten.KeyA) {
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
	if inpututil.IsKeyJustReleased(ebiten.KeyUp) || inpututil.IsKeyJustReleased(ebiten.KeyW) {
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
	if inpututil.IsKeyJustReleased(ebiten.KeyDown) || inpututil.IsKeyJustReleased(ebiten.KeyS) {
		p.FrameOX[0] = 32
		p.FrameNum[0] = 1
		p.MovY = 0
		p.d = 0
	}

	// trasladar player1
	if ModeGame == 1 && p.posicionInicial[0] != 1 {
		p.posicionInicial[0] = 1
	}
	var X1 = p.X[0]
	var Y1 = p.Y[0]
	switch {
	case p.MovX == 1:
		p.X[0]++
	case p.MovX == 2:
		p.X[0]--
	case p.MovY == 1:
		p.Y[0]--
	case p.MovY == 2:
		p.Y[0]++
	}
	p.X[0], p.Y[0], obs = obstaculos(p.X[0], p.Y[0], X1, Y1)
	if obs {
		p.FrameOX[0] = 32
		p.FrameNum[0] = 1
	}
	switch {
	case (p.X[0] > 125 && p.X[0] < 135 && p.Y[0] < 97 && p.Y[0] > 95) || (p.X[0] > 10 && p.X[0] < 25 && p.Y[0] < 272 && p.Y[0] > 270) || (p.X[0] > 635 && p.X[0] < 645 && p.Y[0] < 47 && p.Y[0] > 44) || (p.X[0] > 415 && p.X[0] < 425 && p.Y[0] < 52 && p.Y[0] > 50) || (p.X[0] > 193 && p.X[0] < 258 && p.Y[0] < 110 && p.Y[0] > 78) || (p.X[0] > 749 && p.X[0] < 751 && p.Y[0] < 222 && p.Y[0] > 220):
		p.Y[0] = -40
	case p.Y[0] < -36 && p.Y[0] > -39 && p.X[0] > 125 && p.X[0] < 135:
		p.Y[0] = 98
	case p.Y[0] < -36 && p.Y[0] > -39 && p.X[0] > 10 && p.X[0] < 25:
		p.Y[0] = 272
	case p.Y[0] < -36 && p.Y[0] > -39 && p.X[0] > 635 && p.X[0] < 645:
		p.Y[0] = 46
	case p.Y[0] < -36 && p.Y[0] > -39 && p.X[0] > 415 && p.X[0] < 425:
		p.Y[0] = 52
	case p.Y[0] < -36 && p.Y[0] > -39 && p.X[0] > 193 && p.X[0] < 264:
		p.Y[0] = 110
	case p.Y[0] < -36 && p.Y[0] > -39 && p.X[0] > 749 && p.X[0] < 751:
		p.Y[0] = 223
	case p.Y[0] < -36:
		p.X[0] = X1
	}
	// fmt.Println(p.X, p.Y)
	return p
}

func vida(h humanos, p player, b Objetos, pl Objetos) (player, Objetos, Objetos) {
	if p.Inmune != true {
		//pierde vidas con la nube
		for i := 0; i < nivel; i++ {
			nubX := nube1.X[i] * scale
			nubY := nube1.Y[i] * scale
			//pierde vidas con nube
			if p.X[0] > nubX && p.X[0] < nubX+120 && p.Y[0] > nubY && p.Y[0] < nubY+120 && nube1.Alpha[i] > .3 {
				p.v += .1
			}
		}
		//pierde vidas con humanos
		for i := 0; i < nivel; i++ {
			if p.X[0]+20 > h.X[i] && p.X[0] < h.X[i]+20 && p.Y[0]+32 > h.Y[i] && p.Y[0] < h.Y[i]+32 {
				p.v++
			}
		}
	}
	//infmune con barbijo o alchol en gel
	if p.X[0]+32 > b.X && p.X[0] < b.X+20 && p.Y[0]+48 > b.Y && p.Y[0] < b.Y+32 {
		b.X = 1500
		p.Inmune = true
		p.CountInmune = 300
	}
	if p.Inmune == true {
		p.CountInmune--
	}
	if p.CountInmune == 0 {
		p.Inmune = false
	}

	//gana vida
	if p.X[0]+32 > pl.X && p.X[0] < pl.X+60 && p.Y[0]+48 > pl.Y && p.Y[0] < pl.Y+120 {
		p.vidas++
		pl.X = 1500
	}
	if p.v >= 30 {
		p.vidas--
		sonidoVidas()
		p.v = 0
	}
	//gana monedas
	if p.X[0]+32 > monedas.X && p.X[0] < monedas.X+20 && p.Y[0]+48 > monedas.Y && p.Y[0] < monedas.Y+32 {
		monedas.X = 1500
		p.Coins += 5
	}
	return p, b, pl
}

func dibujarPlayer(P player, screen *ebiten.Image) {
	if ModePause {
		P.FrameNum[0] = 1
		P.FrameOX[0] = 0
	}
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(.65, .65)
	op.GeoM.Translate(P.X[0], P.Y[0])
	i := (count1 / 7) % P.FrameNum[0]
	sx, sy := P.FrameOX[0]+i*P.FrameWidth[0], P.FrameOY[0]
	screen.DrawImage(P.img[0].SubImage(image.Rect(sx, sy, sx+P.FrameWidth[0], sy+P.FrameHeight[0])).(*ebiten.Image), op)
}
