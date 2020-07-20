package main

import (
	"fmt"
	"image/color"
	_ "image/png"
	"log"

	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"

	// "golang.ge/font"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/audio"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/inpututil"
	"github.com/hajimehoshi/ebiten/text"
)

//intro
var intro1 intro

// Game1 es el juego
var Game1 Game

const (
	// game
	screenWidth  = 768
	screenHeight = 528

	// tiles
	tileSize = 16
	tileXNum = 48

	//para start y game Over
	fontSize      = 32
	smallFontSize = fontSize / 2
)

//Mode Game es el en que parte del juego estamos
var (
	ModeGame        int
	ModeTitle       int
	ElectNumPlayers int
	ElectPlayer     int
	ModeGameOver    int
	count1          int
	ModePause       bool
	pulse           bool
	pulso           int
)

var (
	// imágenes
	imgTiles *ebiten.Image

	// sonido
	audioContext *audio.Context
	deadSound    *audio.Player
	deadSound2   *audio.Player
	sonidoFondo  *audio.InfiniteLoop
	fondo        *audio.Player

	//para start y game over
	arcadeFont      font.Face
	smallArcadeFont font.Face
	texts           = []string{}

	err error
)

// init carga los datos
func init() {

	// Intro Init
	intro1.initIntro(screenWidth, screenHeight)

	//////////////   Imagen CITY  ////////////////////////////////
	imgTiles, _, err = ebitenutil.NewImageFromFile(`sircovid\data\city.png`, ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
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

	////////////////  TEXTOS    ////////////////////////////
	tt, err := truetype.Parse(fonts.ArcadeN_ttf)
	if err != nil {
		log.Fatal(err)
	}
	const dpi = 72
	arcadeFont = truetype.NewFace(tt, &truetype.Options{
		Size:    fontSize,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	smallArcadeFont = truetype.NewFace(tt, &truetype.Options{
		Size:    smallFontSize,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})

}

////////////////////////////
// Update
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
					player1.humanos = chica
				}
				if inpututil.IsKeyJustPressed(ebiten.KeyLeft) {
					player1.humanos = viejo
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
			player2 = vida(enemigo2, player1, barbijo)
			player2 = vida(enemigo3, player1, barbijo)
			player2 = vida(enemigo4, player1, barbijo)
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
// Draw
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
	lifesP1 := fmt.Sprintf("Vidas:%02d", player1.vidas)
	text.Draw(screen, lifesP1, smallArcadeFont, fontSize, 40, color.RGBA{35, 27, 190, 0xff})

	if Game1.numPlayers == 2 {
		lifesP2 := fmt.Sprintf("Vidas:%02d", player2.vidas)
		text.Draw(screen, lifesP2, smallArcadeFont, 600, 40, color.RGBA{35, 27, 190, 0xff})
	}

	switch {
	case ModeTitle == 0:

		// intro draw
		intro1.drawIntro(screen, screenWidth, screenHeight)

		if ElectNumPlayers == 1 {
			texts = []string{"", "", "", "PRIMER NIVEL", "", "PRESS SPACE KEY"}
		}

	case ModeTitle == 1 && (player1.vidas != 0 && player2.vidas != 0):
		texts = []string{}

	case ModeTitle == 2:
		texts = []string{"", "", "SEGUNDO NIVEL", "", "", "PRESS SPACE KEY"}
	case ModeTitle == 3 && (player1.vidas != 0 && player2.vidas != 0):
		texts = []string{}

	case player2.vidas == 0 || player1.vidas == 0:
		texts = []string{"", "", "", "GAME OVER!", "", "TRY AGAIN?", "", "PRESS SPACE KEY"}
	}

	// fmt.Println(count)
	switch {
	case ModePause:

		fmt.Println(pulso)
		jugadores := fmt.Sprintf("PAUSE")
		// fmt.Println(pulse)
		if !pulse {
			text.Draw(screen, jugadores, arcadeFont, 300, 200, color.White)
		}

	case ElectNumPlayers == 0:
		jugadores := fmt.Sprintf(" Elija la cantidad de jugadores \n\ncon las flechas y presione la barra espaciadora\n\nJUGADORES:%02d", Game1.numPlayers)
		text.Draw(screen, jugadores, smallArcadeFont, 140, 250, color.White)

	case ElectPlayer == 0 && Game1.numPlayers == 1:
		jugadores := fmt.Sprintf("Elija el jugador \n\n con las flechas\n\ny presione la barra espaciadora")
		text.Draw(screen, jugadores, smallArcadeFont, 250, 250, color.White)

	case ElectNumPlayers == 1:
		for i, l := range texts {
			x := (screenWidth - len(l)*fontSize) / 2
			text.Draw(screen, l, arcadeFont, x, (i+5)*fontSize, color.White)
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
	ebiten.SetWindowTitle("Sircovid")
	ebiten.SetWindowResizable(true)

	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
