package main

import (
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/inpututil"
)

// init carga los datos
func init() {
	//init ciudad img
	imgTiles, _, err = ebitenutil.NewImageFromFile(`sircovid\data\cidade.png`, ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
	//init intro
	intro1.initIntro(screenWidth, screenHeight)
	//inicia los textos
	initTextos()
	//inicializa a players
	initPlayer()
	//inicia sumarVidas
	initObjetos()
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
	if count1 == 60 {
		count1 = 0
	}
	//funcion pausa
	if inpututil.IsKeyJustPressed(ebiten.KeyP) {
		ModePause = !ModePause
	}
	switch {
	//pausar el juego
	case ModePause:
	case ModeWin == true:

		//toda la introduccion con eleccion de players, etc
	case ModeTitle == 0:
		introduccion()

	case ModeTitle == 1:
		if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
			ModeTitle = 2
		}
	case player1.vidas != 0 && player2.vidas != 0 && ModeTitle == 2:

		//// sonido ////
		sonidoGame()
		// nube
		nube1 = moverNube(nube1)

		// palyer
		player1 = moverPlayer(player1)
		player2 = moverPlayer(player2)

		// vida
		player1, barbijo, plasma = vida(enemigo, player1, barbijo, plasma)
		player1, alchol, plasma = vida(enemigo, player1, alchol, plasma)

		if Game1.numPlayers == 2 {
			player2, barbijo, plasma = vida(enemigo, player1, barbijo, plasma)
			player2, alchol, plasma = vida(enemigo, player1, alchol, plasma)
		}
		//enemigos
		enemigo = moverHumanos(enemigo)

		//siguiente nivel
		Game1.siguienteNivel = siguienteNivel(player1)
		Game1.siguienteNivel = siguienteNivel(player2)

	case ModeGameOver == 0:
		sonidoGameover()
		if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
			iniciarVariables()
			initPlayer()
			initNube()
			initObjetos()
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

	//dibujar barbijo
	dibujarObjetos(barbijo, screen)
	dibujarObjetos(plasma, screen)
	dibujarObjetos(alchol, screen)
	dibujarObjetos(monedas, screen)

	//dibujar nextLevel
	dibujarObjetos(nextLevel, screen)
	//dibujar palyers
	dibujarPlayer(player1, screen)

	if Game1.numPlayers == 2 {
		dibujarPlayer(player2, screen)
	}
	//dibuja al enemigo
	dibujarEnemigos(enemigo, screen)

	// dibujar nube
	dibujarNube(nube1, screen)

	//dibujar carteles
	// dibujarObjetos(cartSupermarket, screen)
	// dibujarObjetos(cartFarmacy, screen)
	// dibujarObjetos(cartStore, screen)
	// dibujarObjetos(cartBanck, screen)

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
	ebiten.SetWindowTitle("Sir-covid")
	ebiten.SetWindowResizable(true)

	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
