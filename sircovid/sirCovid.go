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
	case ModeWin:
		//toda la introduccion con eleccion de players, etc
	case Relato:
		if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
			Relato = false
		}
	case ModeTitle:
		introduccion()

	case ModeTitleLevel:

		if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
			ModeTitleLevel = false
			ModeGame = true
		}

	case player1.Compras:
		if player1.Coins >= 2 {
			player1 = compar(player1)
		}
		if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
			player1.Compras = false
		}
	case player2.Compras:
		if player2.Coins >= 2 {
			player2 = compar(player2)
		}
		if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
			player2.Compras = false
		}
	case ModeGame:
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
			player2, barbijo, plasma = vida(enemigo, player2, barbijo, plasma)
			player2, alchol, plasma = vida(enemigo, player2, alchol, plasma)
		}
		//enemigos
		enemigo = moverHumanos(enemigo)

		//siguiente nivel
		Game1.siguienteNivel = siguienteNivel(player1)
		Game1.siguienteNivel = siguienteNivel(player2)

	case ModeGameOver:
		sonidoGameover()
		for i := 0; i < nivel; i++ {
			enemigo.FrameNum[i] = 1
			enemigo.FrameOX[i] = 0
		}
		if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
			iniciarVariables()
			initPlayer()
			initNube()
			initObjetos()
			initEnemigos()
			ModeGameOver = false
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
	// op.GeoM.Scale(2, 2)
	// op.GeoM.Translate(player1.X[0], player1.Y[0])
	screen.DrawImage(imgTiles, op)

	//dibujar objetos
	dibujarObjetos(barbijo, screen)
	dibujarObjetos(plasma, screen)
	dibujarObjetos(alchol, screen)
	dibujarObjetos(monedas, screen)

	//dibujar palyers
	dibujarPlayer(player1, screen)
	if Game1.numPlayers == 2 {
		dibujarPlayer(player2, screen)
	}
	//dibuja al enemigo
	dibujarEnemigos(enemigo, screen)

	// dibujar nube
	dibujarNube(nube1, screen)

	// dibujar texto
	if !Relato {
		dibujarTextos(screen)
	}

	//dibujar textos compras
	dibujarTextoCompras(player1, screen)
	dibujarTextoCompras(player2, screen)

	//Dibujar relato
	if Relato {
		dibujarObjetos(relato, screen)
		relato.Y = relato.Y - .3

	}
	if ModeTitleLevel {
		dibujarObjetos(mhome, screen)
		dibujarObjetos(money, screen)
		if Game1.numPlayers == 2 {
			dibujarObjetos(mhome1, screen)
		}
	}
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
