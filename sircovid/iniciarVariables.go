package main

var nivel = int(1)

func iniciarVariables() {
	count = 0
	nivel = 1
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
	nivel = ModeGame * 2
	if nivel > 13 {
		nivel = 13
	}

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
