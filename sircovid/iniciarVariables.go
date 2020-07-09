package main

import "math/rand"

func iniciarVariables() {
	hombre.FrameOX = 0
	hombre.FrameOY = 48
	hombre.FrameNum = 1
	hombre.X = float64(350)
	hombre.Y = float64(290)
	hombre.FrameWidth = 32
	hombre.FrameHeight = 48

	mujer.FrameOX = 0
	mujer.FrameOY = 48
	mujer.FrameNum = 1
	mujer.X = float64(750)
	mujer.Y = float64(290)
	mujer.FrameWidth = 32
	mujer.FrameHeight = 48

	// viejo
	viejo.FrameOX = 0
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

	//enemigos
	enemigos1.humanos = hombre

	//nube
	nube1.nubeX = float64(rand.Intn(screenWidth) + 300)
	nube1.nubeY = float64(rand.Intn(1500))

	//defino el juego
	Game1.nube = nube1
	Game1.siguienteNivel = player1.humanos

	//paso de nivel
	ModeGame = 0
	ModeTitle = 0
	ModeGameOver = 0

	//para movimiento de player
	a, b, c, d = 0, 0, 0, 0

}
