package main

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func iniciarVariables() {
	count = 0

	//cosas de Game
	ModeGame = 0
	ModeTitle = 0
	ModeGameOver = 0
	ElectNumPlayers = 0
	ElectPlayer = 0
	Game1.numPlayers = 1

}
func pasarNivel() {
	// initSonido()
	count = 0
	initEnemigos()

	//reinciar enemigos
	count1 = 0
	initPlayer()

	//nube
	initNube()

	//paso de nivel
	ModeTitle = 2
	ModeGame = 1

	// para movimiento de player
	// player1.a, player1.b, player1.c, player1.d = 0, 0, 0, 0
	// player2.a, player2.b, player2.c, player2.d = 0, 0, 0, 0
}
