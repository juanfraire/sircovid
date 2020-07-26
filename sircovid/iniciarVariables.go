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
	ModeGame++
	ModeTitle = 1

	pasarNivelPlayer()

	initObjetos()

	// initSonido()
	count = 0
	initEnemigos()
	//reinciar enemigos
	count1 = 0
	//nube
	initNube()
}
