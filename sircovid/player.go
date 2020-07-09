package main

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
)

type player struct {
	humanos
	vidas       int
	v           int
	moverPlayer (humanos)
	vida        (humanos)
}

var as, b, c, d int

func moverPlayer(p humanos) humanos {
	// leer tecla

	// fmt.Println(as, b, c, d)
	if inpututil.IsKeyJustPressed(ebiten.KeyRight) {
		as = 1
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyLeft) {
		b = 1
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyUp) {
		c = 1
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyDown) {
		d = 1
	}
	if as == 1 && p.MovY != 1 && p.MovY != 2 {
		p.FrameOY = 96
		p.FrameNum = 3
		p.MovX = 1
	}
	if inpututil.IsKeyJustReleased(ebiten.KeyRight) && p.MovY != 1 && p.MovY != 2 {
		p.FrameNum = 1
		p.MovX = 0
		as = 0
	}
	if b == 1 && p.MovY != 1 && p.MovY != 2 {
		p.FrameOY = 48
		p.FrameNum = 3
		p.MovX = 2
	}
	if inpututil.IsKeyJustReleased(ebiten.KeyLeft) && p.MovY != 1 && p.MovY != 2 {
		p.FrameNum = 1
		p.MovX = 0
		b = 0
	}
	if c == 1 && p.MovX != 1 && p.MovX != 2 {
		p.FrameOY = 144
		p.FrameNum = 3
		p.MovY = 1
	}
	if inpututil.IsKeyJustReleased(ebiten.KeyUp) && p.MovX != 1 && p.MovX != 2 {
		p.FrameNum = 1
		p.MovY = 0
		c = 0
	}
	if d == 1 && p.MovX != 1 && p.MovX != 2 {
		p.FrameOY = 0
		p.FrameNum = 3
		p.MovY = 2
	}
	if inpututil.IsKeyJustReleased(ebiten.KeyDown) && p.MovX != 1 && p.MovX != 2 {
		p.FrameNum = 1
		p.MovY = 0
		d = 0
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
	p.X, p.Y, ok = obstaculos(p.X+10, p.Y+32, X1, Y1)
	if !ok {
		p.X -= 10
		p.Y -= 32
	}
	return p
}

func vida(h humanos, p humanos) player {
	//pierde vidas con la nuve
	collisionX := float64(Game1.nubeX * .4)
	collisionY := float64(Game1.nubeY*.4) + 106
	if Game1.nubeAlpha < .3 {
		collisionX = screenWidth + 300
	}
	if p.X > collisionX && p.X < collisionX+120 && p.Y > collisionY && p.Y < collisionY+120 {
		player1.v++
	}
	if p.X > h.X && p.X < h.X+32 && p.Y+48 > h.Y && p.Y < h.Y+48 {
		player1.v++
	}
	if p.X > barbijoX && p.X < barbijoX+32 && p.Y+48 > barbijoY && p.Y < barbijoY+48 {
		player1.vidas++
		barbijoX = 1000
	}
	if player1.v == 1 {
		player1.vidas--
		//sonido
		deadSound.Play()
		deadSound.Rewind()
	}
	if player1.v == 30 {
		player1.v = 0
	}
	return player1
}
