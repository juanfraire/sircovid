package main

var nivel = int(1)

func iniciarVariables() {
	Relato = true
	count = 0
	nivel = 1
	Level = 1
	//cosas de Game
	ModeTitle = true
	ElectNumPlayers = 0
	ElectPlayer = 0
	Game1.numPlayers = 1

}
func pasarNivel() {
	Level++
	ModeTitleLevel = true
	nivel = Level * 2
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
