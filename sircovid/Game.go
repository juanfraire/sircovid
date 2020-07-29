package main

import (
	"github.com/hajimehoshi/ebiten"
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

	// imágenes
	imgTiles *ebiten.Image

	// sonido
	// audioContext *audio.Context
	// deadSound    *audio.Player
	// deadSound2   *audio.Player
	// sonidoFondo  *audio.InfiniteLoop
	// fondo        *audio.Player
	// sonidoIntro  *audio.InfiniteLoop
	// sIntro       *audio.Player

	//para start y game over
	arcadeFont      font.Face
	smallArcadeFont font.Face
	texts           = []string{}

	err error
)

const (
	// game
	screenWidth  = 66 * 16
	screenHeight = 33 * 16

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
			if inpututil.IsKeyJustPressed(ebiten.KeyW) {
				player2.humanos.img = humano1.img
			}
			if inpututil.IsKeyJustPressed(ebiten.KeyS) {
				player2.humanos.img = humano2.img
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
	if p.X[0] >= nextLevel.X && p.Y[0] > 450 {

		pasarNivel()
		fondo.Pause()
		fondo.Rewind()
		sLevelUp.Play()
		sLevelUp.Rewind()
	}
	if p.X[0] > 751 && p.X[0] < 763 && p.Y[0] < -40 {
		ModeWin = true
		fondo.Pause()
	}
	return p
}
