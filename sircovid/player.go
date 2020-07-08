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

	objetos := make([][]int, 12)
	objetos[0] = []int{578, 189, 2, 148}
	objetos[1] = []int{386, 191, 0, 79}
	objetos[2] = []int{193, 159, -1, 112}
	objetos[3] = []int{211, 43, 109, 34}
	objetos[4] = []int{83, 106, 1, 124}
	objetos[5] = []int{1, 77, 1, 301}
	objetos[6] = []int{83, 188, 416, 113}
	objetos[7] = []int{458, 129, 337, 106}
	objetos[8] = []int{705, 61, 352, 108}
	objetos[9] = []int{723, 45, 149, 107}
	objetos[10] = []int{306, 268, 162, 111}
	objetos[11] = []int{306, 189, 148, 13}

	for i := 0; i < len(objetos); i++ {
		if (int(p.X+20) > objetos[i][0] && int(p.X) < objetos[i][0]+objetos[i][1]) && int(p.Y+35) > objetos[i][2] && int(p.Y+35) < objetos[i][2]+objetos[i][3] {
			p.X = X1
			p.Y = Y1
		}
	}
	return p

}
func vida(h humanos, p humanos) player {
	//pierde vidas con la nuve
	collisionX := float64(Game1.nubeX * .4)
	collisionY := float64(Game1.nubeY * .4)
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
		deadSound.Play()
		deadSound.Rewind()
	}
	if player1.v == 30 {
		player1.v = 0
	}
	return player1
}
