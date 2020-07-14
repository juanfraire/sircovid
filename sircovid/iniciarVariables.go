package main

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func iniciarVariables() {
<<<<<<< HEAD
	// hombre.FrameOX = 0
	// hombre.FrameOY = 48
	// hombre.FrameNum = 1
	// hombre.X = float64(350)
	// hombre.Y = float64(290)
	// hombre.FrameWidth = 32
	// hombre.FrameHeight = 48

	// mujer.FrameOX = 0
	// mujer.FrameOY = 48
	// mujer.FrameNum = 1
	// mujer.X = float64(screenWidth - 50)
	// mujer.Y = float64(290)
	// mujer.FrameWidth = 32
	// mujer.FrameHeight = 48
=======

	hombre.FrameOX = 0
	hombre.FrameOY = 48
	hombre.FrameNum = 1
	hombre.X = float64(350)
	hombre.Y = float64(290)
	hombre.FrameWidth = 32
	hombre.FrameHeight = 48
	hombre.num = rand.Intn(5)
	hombre.cambio = rand.Intn(100) + 100

	mujer.FrameOX = 0
	mujer.FrameOY = 48
	mujer.FrameNum = 1
	mujer.X = float64(screenWidth - 50)
	mujer.Y = float64(290)
	mujer.FrameWidth = 32
	mujer.FrameHeight = 48
	mujer.num = rand.Intn(5)
	mujer.cambio = rand.Intn(100) + 50

	//enemigos
	enemigos1.humanos = hombre
	enemigos2.humanos = mujer
>>>>>>> 3bb2d837a2760e6774f74a8c03e827626696b24f

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
	ElectNumPlayers = 0
	ElectPlayer = 0

}
func pasarNivel() {
	hombre.FrameOX = 0
	hombre.FrameOY = 48
	hombre.FrameNum = 1
	hombre.X = float64(200)
	hombre.Y = float64(200)
	hombre.FrameWidth = 32
	hombre.FrameHeight = 48

	mujer.FrameOX = 0
	mujer.FrameOY = 48
	mujer.FrameNum = 1
	mujer.X = float64(screenWidth - 50)
	mujer.Y = float64(290)
	mujer.FrameWidth = 32
	mujer.FrameHeight = 48

	//reinciar enemigos
	count1 = 0

	// viejo
	viejo.FrameOX = 0
	viejo.FrameOY = 96
	viejo.FrameNum = 1
	viejo.X = float64(25)
	viejo.Y = float64(375)
	viejo.FrameWidth = 32
	viejo.FrameHeight = 48

	//palyer 2
	chica.FrameOX = 0
	chica.FrameOY = 96
	chica.FrameNum = 1
	chica.X = float64(25)
	chica.Y = float64(415)
	chica.FrameWidth = 32
	chica.FrameHeight = 48

	//nube
	nube1.nubeX = float64(rand.Intn(screenWidth) + 300)
	nube1.nubeY = float64(rand.Intn(1500))

	//defino el juego
	Game1.nube = nube1
	Game1.siguienteNivel = player1.humanos

	//paso de nivel
	ModeTitle = 2
	ModeGame = 1

	//para movimiento de player
	player1.a, player1.b, player1.c, player1.d = 0, 0, 0, 0
	player2.a, player2.b, player2.c, player2.d = 0, 0, 0, 0
}
