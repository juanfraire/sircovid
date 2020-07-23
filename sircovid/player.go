package main

import (
	"fmt"
	"image"
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/inpututil"
)

type player struct {
	humanos
	vidas      int
	v          int
	señalador  int
	a, b, c, d int
	MovX       int
	MovY       int
}

var player1, player2 player
var humano1, humano2 humanos

func initPlayer() {
	//////////////   Imangen VIEJO  //////////////////////////////
	humano1.img, _, err = ebitenutil.NewImageFromFile(`sircovid\data\player1.png`, ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
	//imagen chica
	humano2.img, _, err = ebitenutil.NewImageFromFile(`sircovid\data\player2.png`, ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
	// player1
	player1.FrameOX = 0
	player1.FrameOY = 96
	player1.FrameNum = 1
	player1.X = float64(365)
	player1.Y = float64(300)
	player1.FrameWidth = 32
	player1.FrameHeight = 48
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
	player2.FrameOX = 0
	player2.FrameOY = 96
	player2.FrameNum = 1
	player2.X = float64(365)
	player2.Y = float64(350)
	player2.FrameWidth = 32
	player2.FrameHeight = 48
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
	//player1
	player1.FrameNum = 1
	player1.X = float64(365)
	player1.Y = float64(300)
	player1.MovX = 0
	player1.MovY = 0
	player1.a, player1.b, player1.c, player1.d = 0, 0, 0, 0

	//player2
	player2.FrameNum = 1
	player2.X = float64(365)
	player2.Y = float64(350)
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
		p.FrameOX = 0
		p.FrameOY = 96
		p.FrameNum = 3
		p.MovX = 1
	}
	if inpututil.IsKeyJustReleased(ebiten.KeyRight) || inpututil.IsKeyJustReleased(ebiten.KeyD) {
		p.FrameOX = 32
		p.FrameNum = 1
		p.MovX = 0
		p.a = 0
	}
	if p.b == 1 && p.MovY != 1 && p.MovY != 2 {
		p.FrameOX = 0
		p.FrameOY = 48
		p.FrameNum = 3
		p.MovX = 2
	}
	if inpututil.IsKeyJustReleased(ebiten.KeyLeft) || inpututil.IsKeyJustReleased(ebiten.KeyA) {
		p.FrameOX = 32
		p.FrameNum = 1
		p.MovX = 0
		p.b = 0
	}
	if p.c == 1 && p.MovX != 1 && p.MovX != 2 {
		p.FrameOX = 0
		p.FrameOY = 144
		p.FrameNum = 3
		p.MovY = 1
	}
	if inpututil.IsKeyJustReleased(ebiten.KeyUp) || inpututil.IsKeyJustReleased(ebiten.KeyW) {
		p.FrameOX = 32
		p.FrameNum = 1
		p.MovY = 0
		p.c = 0
	}
	if p.d == 1 && p.MovX != 1 && p.MovX != 2 {
		p.FrameOX = 0
		p.FrameOY = 0
		p.FrameNum = 3
		p.MovY = 2
	}
	if inpututil.IsKeyJustReleased(ebiten.KeyDown) || inpututil.IsKeyJustReleased(ebiten.KeyS) {
		p.FrameOX = 32
		p.FrameNum = 1
		p.MovY = 0
		p.d = 0
	}

	// trasladar player1
	if ModeGame == 1 && p.posicionInicial != 1 {
		p.posicionInicial = 1
	}
	var X1 = p.X
	var Y1 = p.Y
	switch {
	case p.MovX == 1:
		p.X++
	case p.MovX == 2:
		p.X--
	case p.MovY == 1:
		p.Y--
	case p.MovY == 2:
		p.Y++
	}
	p.X, p.Y, ok = obstaculos(p.X, p.Y, X1, Y1)
	if ok {
		p.FrameOX = 32
		p.FrameNum = 1
	}
	switch {
	case (p.X > 125 && p.X < 135 && p.Y < 97 && p.Y > 95) || (p.X > 10 && p.X < 25 && p.Y < 272 && p.Y > 270) || (p.X > 635 && p.X < 645 && p.Y < 47 && p.Y > 44) || (p.X > 415 && p.X < 425 && p.Y < 52 && p.Y > 50) || (p.X > 193 && p.X < 258 && p.Y < 110 && p.Y > 78) || (p.X > 749 && p.X < 751 && p.Y < 222 && p.Y > 220):
		p.Y = -40
	case p.Y < -36 && p.Y > -39 && p.X > 125 && p.X < 135:
		p.Y = 98
	case p.Y < -36 && p.Y > -39 && p.X > 10 && p.X < 25:
		p.Y = 272
	case p.Y < -36 && p.Y > -39 && p.X > 635 && p.X < 645:
		p.Y = 46
	case p.Y < -36 && p.Y > -39 && p.X > 415 && p.X < 425:
		p.Y = 52
	case p.Y < -36 && p.Y > -39 && p.X > 193 && p.X < 264:
		p.Y = 110
	case p.Y < -36 && p.Y > -39 && p.X > 749 && p.X < 751:
		p.Y = 223
	case p.Y < -36:
		p.X = X1
	}
	fmt.Println(p.X, p.Y)
	return p
}

func vida(h humanos, p player, b sumVidas) player {
	//pierde vidas con la nube
	nubX := nube1.X * scale
	nubY := nube1.Y * scale

	//pierde vidas con nube
	if p.X > nubX && p.X < nubX+120 && p.Y > nubY && p.Y < nubY+120 && nube1.Alpha > .3 {
		p.v++
	}
	//pierde vidas con humanos
	if p.X+20 > h.X && p.X < h.X+20 && p.Y+32 > h.Y && p.Y < h.Y+32 {
		p.v++
	}
	//gana vida con barbijo
	if p.X > b.X && p.X < b.X+20 && p.Y+32 > b.Y && p.Y < b.Y+32 {
		p.vidas++
		barbijo.X = 1000
	}

	if p.v == 30 {
		p.vidas--
		p.v = 0
		sonidoVidas()
	}

	return p
}

func dibujarPlayer(P player, screen *ebiten.Image) {
	if ModePause {
		P.FrameNum = 1
		P.FrameOX = 0
	}
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(.65, .65)
	op.GeoM.Translate(P.X, P.Y)
	i := (count1 / 7) % P.FrameNum
	sx, sy := P.FrameOX+i*P.FrameWidth, P.FrameOY
	screen.DrawImage(P.img.SubImage(image.Rect(sx, sy, sx+P.FrameWidth, sy+P.FrameHeight)).(*ebiten.Image), op)
}
