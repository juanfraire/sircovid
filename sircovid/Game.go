package main

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/audio"
	"github.com/hajimehoshi/ebiten/inpututil"
	"golang.org/x/image/font"
)

//Game es la estructura del juego
type Game struct {
	count int
	nube
	numPlayers     int
	electPlayer    int
	siguienteNivel (player)
}

// Game1 es el juego
var Game1 Game
var intro1 intro

var (
	ModeGame        int
	ModeWin         bool
	ModeTitle       int
	ElectNumPlayers int
	ElectPlayer     int
	ModeGameOver    int
	count1          int
	ModePause       bool
	pulse           bool
	pulso           int

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

//introduccion es la eleccion de los players, etc
func introduccion() {
	// intro update
	intro1.updateIntro(screenWidth, screenHeight)

	switch {
	case ElectNumPlayers == 0:
		if inpututil.IsKeyJustPressed(ebiten.KeyDown) {
			Game1.numPlayers = 2
		}
		if inpututil.IsKeyJustPressed(ebiten.KeyUp) {
			Game1.numPlayers = 1
		}
		if inpututil.IsKeyJustPressed(ebiten.KeySpace) && (Game1.numPlayers == 1 || Game1.numPlayers == 2) {
			ElectNumPlayers = 1
		}
	case ElectPlayer == 0 && Game1.numPlayers == 1 || Game1.numPlayers == 2:
		if Game1.numPlayers == 1 || Game1.numPlayers == 2 {
			if inpututil.IsKeyJustPressed(ebiten.KeyUp) {
				player1.humanos.img = humano1.img
			}
			if inpututil.IsKeyJustPressed(ebiten.KeyDown) {
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
}

func siguienteNivel(p player) player {
	if p.X[0] >= 746 && p.Y[0] > 450 {

		pasarNivel()
		fondo.Pause()
		fondo.Rewind()
	}
	if p.X[0] > 350 && p.X[0] < 370 && p.Y[0] < -30 {
		ModeWin = true
		fondo.Pause()
	}
	return p
}
