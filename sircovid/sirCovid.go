package main

import (
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
)

// init carga los datos
func init() {
	//inicia los textos
	initTextos()
	//inicializa a players
	initPlayer()
	//inicia sumarVidas
	initSumarVidas()
	//inicia nube
	initNube()
	//inicia enemigos
	initEnemigos()
	//iniciar otra variables
	iniciarVariables()
	//iniciar sonidos
	initSonido()

}

////////////////////////////
//////// Update ////////////
////////////////////////////

// Update se llama 60 veces por segundo
func (g *Game) Update(screen *ebiten.Image) error {
	sonido()
	// game counter
	g.count++
	count1++
	//funcion pausa
	if inpututil.IsKeyJustPressed(ebiten.KeyP) {
		ModePause = !ModePause
	}
	switch {

	case ModePause:

	case ModeTitle == 0:
		// intro update
		intro1.updateIntro(screenWidth, screenHeight)

		switch {
		case ElectNumPlayers == 0:
			if inpututil.IsKeyJustPressed(ebiten.KeyRight) {
				Game1.numPlayers = 2
			}
			if inpututil.IsKeyJustPressed(ebiten.KeyLeft) {
				Game1.numPlayers = 1
			}
			if inpututil.IsKeyJustPressed(ebiten.KeySpace) && (Game1.numPlayers == 1 || Game1.numPlayers == 2) {
				ElectNumPlayers = 1
			}
		case ElectPlayer == 0 && Game1.numPlayers == 1 || Game1.numPlayers == 2:
			if Game1.numPlayers == 1 || Game1.numPlayers == 2 {
				if inpututil.IsKeyJustPressed(ebiten.KeyRight) {
					player1.humanos.img = humano1.img
				}
				if inpututil.IsKeyJustPressed(ebiten.KeyLeft) {
					player1.humanos.img = humano2.img
				}
			}
			if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
				ElectPlayer = 1
			}
		}
		if inpututil.IsKeyJustPressed(ebiten.KeySpace) && ElectPlayer == 1 {
			ModeTitle = 1
		}
	case ModeTitle == 2:
		if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
			ModeTitle = 3
		}
	case ModeGame == 0 && player1.vidas != 0 && player2.vidas != 0:

		//// sonido ////
		sonidoGame()
		// nube
		nube1 = moverNube(nube1)

		// palyer
		player1 = moverPlayer(player1)
		player2 = moverPlayer(player2)

		// vida
		player1 = vida(enemigo1, player1, barbijo)
		player1 = vida(enemigo2, player1, barbijo)
		player1 = vida(enemigo3, player1, barbijo)
		player1 = vida(enemigo4, player1, barbijo)

		if Game1.numPlayers == 2 {
			player2 = vida(enemigo1, player2, barbijo)
			player2 = vida(enemigo2, player2, barbijo)
			player2 = vida(enemigo3, player2, barbijo)
			player2 = vida(enemigo4, player2, barbijo)
		}
		//enemigos
		enemigo1 = moverHumanos(enemigo1)
		enemigo2 = moverHumanos(enemigo2)
		enemigo3 = moverHumanos(enemigo3)
		enemigo4 = moverHumanos(enemigo4)

		//siguiente nivel
		Game1.siguienteNivel = siguienteNivel(player1)
		Game1.siguienteNivel = siguienteNivel(player2)

	case ModeGame == 1 && player1.vidas != 0 && player2.vidas != 0:

		// sonido
		sonidoGame()

		// nube
		nube1 = moverNube(nube1)

		// player
		player1 = moverPlayer(player1)
		player2 = moverPlayer(player2)

		// vida
		player1 = vida(enemigo1, player1, barbijo)
		player1 = vida(enemigo2, player1, barbijo)
		player1 = vida(enemigo3, player1, barbijo)
		player1 = vida(enemigo4, player1, barbijo)

		if Game1.numPlayers == 2 {
			player2 = vida(enemigo1, player2, barbijo)
			player2 = vida(enemigo2, player2, barbijo)
			player2 = vida(enemigo3, player2, barbijo)
			player2 = vida(enemigo4, player2, barbijo)
		}
		//enemigos
		enemigo1 = moverHumanos(enemigo1)
		enemigo2 = moverHumanos(enemigo2)
		enemigo3 = moverHumanos(enemigo3)
		enemigo4 = moverHumanos(enemigo4)

	case ModeGameOver == 0:
		sonidoGameover()
		if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
			iniciarVariables()
			initPlayer()
			initNube()
			initSumarVidas()
			initEnemigos()
		}

	}
	return nil
}

////////////////////////////
///////// Draw /////////////
////////////////////////////

// Draw dibuja la pantalla 60 veces por segundo
func (g *Game) Draw(screen *ebiten.Image) {

	// dubujar fondo
	op := &ebiten.DrawImageOptions{}
	screen.DrawImage(imgTiles, op)

	//dibujar palyers
	dibujarPlayer(player1, screen)
	if Game1.numPlayers == 2 {
		dibujarPlayer(player2, screen)
	}

	//dibuja al enemigo
	if ModeGame == 0 {
		dibujarEnemigos(enemigo1, screen)
		dibujarEnemigos(enemigo2, screen)
		dibujarEnemigos(enemigo3, screen)
		dibujarEnemigos(enemigo4, screen)
	}
	if ModeGame == 1 {
		dibujarEnemigos(enemigo1, screen)
		dibujarEnemigos(enemigo2, screen)
		dibujarEnemigos(enemigo3, screen)
		dibujarEnemigos(enemigo4, screen)
	}

	// dibujar nube
	dibujarNube(nube1, screen)

	// dibujar barbijo
	dibujarSumVidas(barbijo, screen)

	// dibujar texto
	dibujarTextos(screen)
}

// Layout maneja las dimensiones de pantalla
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return int(screenWidth), int(screenHeight)
}

////////////////////////////
// Main
////////////////////////////

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Sircovid")
	ebiten.SetWindowResizable(true)

	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
