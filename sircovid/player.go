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
	vidas      int
	v          int
	señalador  int
	a, b, c, d int
}

var chica, viejo humanos
var player1, player2 player

func initPlayer() {
	//////////////   Imangen VIEJO  //////////////////////////////
	viejo.img, _, err = ebitenutil.NewImageFromFile(`sircovid\data\viejo.png`, ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
	//imagen chica
	chica.img, _, err = ebitenutil.NewImageFromFile(`sircovid\data\segundoPlayer.png`, ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
	// viejo
	viejo.FrameOX = 32
	viejo.FrameOY = 96
	viejo.FrameNum = 1
	viejo.X = float64(25)
	viejo.Y = float64(375)
	viejo.FrameWidth = 32
	viejo.FrameHeight = 48
	//player
	player1.humanos = viejo
	player1.vidas = 3
	player1.v = 0
	player1.señalador = 0
	player1.a, player1.b, player1.c, player1.d = 0, 0, 0, 0

	//player2
	chica.FrameOX = 32
	chica.FrameOY = 96
	chica.FrameNum = 1
	chica.X = float64(25)
	chica.Y = float64(415)
	chica.FrameWidth = 32
	chica.FrameHeight = 48

	player2.humanos = chica
	player2.vidas = 3
	player2.v = 0
	player2.señalador = 1
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

	// trasladar viejo
	if ModeGame == 1 && p.posicionInicial != 1 {
		p.X = float64(25)
		p.Y = float64(375)
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
	p.X, p.Y, ok = obstaculos(p.X, p.Y+32, X1, Y1)
	switch {
	case ok:
		p.FrameOX = 32
		p.FrameNum = 1
	case !ok:
		p.Y -= 32
		p.FrameOX = 0
	}
	return p
}

func vida(h humanos, p player, b sumVidas) player {
	//pierde vidas con la nube
	nubX := nube1.X * scale
	nubY := nube1.Y * scale
	if nube1.Alpha > .3 {
		//pierde vidas con nube
		if p.X > nubX && p.X < nubX+120 && p.Y > nubY && p.Y < nubY+120 {
			p.v++
		}
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
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(.7, .7)
	op.GeoM.Translate(P.X, P.Y)
	i := (count1 / 7) % P.FrameNum
	sx, sy := P.FrameOX+i*P.FrameWidth, P.FrameOY
	screen.DrawImage(P.img.SubImage(image.Rect(sx, sy, sx+P.FrameWidth, sy+P.FrameHeight)).(*ebiten.Image), op)
}
