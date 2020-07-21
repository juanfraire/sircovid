package main

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/audio"
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

//// nubeCovid aumenta y disminuye transparencia de la nube (alpha)

func siguienteNivel(p player) player {
	if p.X >= 746 && p.Y > 450 {

		pasarNivel()
		fondo.Pause()
		fondo.Rewind()
	}
	return p
}
