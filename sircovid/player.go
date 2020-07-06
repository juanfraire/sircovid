package main

import (
	"math/rand"
	"time"

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

func moverPlayer(p humanos) humanos {
	// leer tecla
	switch {
	case inpututil.IsKeyJustPressed(ebiten.KeyRight) && p.MovY != 1 && p.MovY != 2:
		p.FrameOY = 96
		p.FrameNum = 3
		p.MovX = 1
	case inpututil.IsKeyJustReleased(ebiten.KeyRight) && p.MovY != 1 && p.MovY != 2:
		p.FrameNum = 1
		p.MovX = 0
	case inpututil.IsKeyJustPressed(ebiten.KeyLeft) && p.MovY != 1 && p.MovY != 2:
		p.FrameOY = 48
		p.FrameNum = 3
		p.MovX = 2
	case inpututil.IsKeyJustReleased(ebiten.KeyLeft) && p.MovY != 1 && p.MovY != 2:
		p.FrameNum = 1
		p.MovX = 0
	case inpututil.IsKeyJustPressed(ebiten.KeyUp) && p.MovX != 1 && p.MovX != 2:
		p.FrameOY = 144
		p.FrameNum = 3
		p.MovY = 1
	case inpututil.IsKeyJustReleased(ebiten.KeyUp) && p.MovX != 1 && p.MovX != 2:
		p.FrameNum = 1
		p.MovY = 0
	case inpututil.IsKeyJustPressed(ebiten.KeyDown) && p.MovX != 1 && p.MovX != 2:
		p.FrameOY = 0
		p.FrameNum = 3
		p.MovY = 2
	case inpututil.IsKeyJustReleased(ebiten.KeyDown) && p.MovX != 1 && p.MovX != 2:
		p.FrameNum = 1
		p.MovY = 0
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

	// restringir viejo
	switch {
	case p.Y < 300 && p.X > 20 && p.X < 214:
		p.X = X1
		p.Y = Y1
		p.FrameNum = 1
	case p.Y < 130 && p.X > 214 && p.X < 768:
		p.X = X1
		p.Y = Y1
		p.FrameNum = 1
	case p.Y < 270 && p.X > 240 && p.X < 610:
		p.X = X1
		p.Y = Y1
		p.FrameNum = 1
	case p.Y < 270 && p.X > 675 && p.X < 768:
		p.X = X1
		p.Y = Y1
		p.FrameNum = 1
	case p.Y > 335 && p.Y < 528 && p.X > 40 && p.X < 350:
		p.X = X1
		p.Y = Y1
		p.FrameNum = 1
	case p.Y > 310 && p.Y < 450 && p.X > 390 && p.X < 630:
		p.X = X1
		p.Y = Y1
		p.FrameNum = 1
	}
	return p

}
func vida(h humanos, p humanos) player {
	//pierde vidas con la nuve
	collisionX := float64(nubeX * .4)
	collisionY := float64(nubeY * .4)
	if nubeAlpha < .3 {
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
		deadSound.Play()
		deadSound.Rewind()
	}
	if player1.v == 30 {
		player1.v = 0
	}
	return player1
}

// nubeCovid aumenta y disminuye transparencia de la nube (alpha)
func moverNube() {
	// creacion de nuevas nubes
	if nubeAlpha <= 0 {
		nubeX = float64(rand.Intn(1500))
		nubeY = float64(rand.Intn(500) + 600)
		nubeAlphaUp = true
	} else if nubeAlpha > 1 {
		time.Sleep(10000 * time.Microsecond)
		nubeAlphaUp = false
	}

	// movimiento nube
	if nubeAlpha >= 0 {
		nubeX--
	}

	// actualizar alpha
	if nubeAlphaUp {
		nubeAlpha += .009
	} else {
		nubeAlpha -= .003
	}
}
