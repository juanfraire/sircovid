package main

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
	"github.com/hajimehoshi/ebiten/text"
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
	aspirina        bool
	plasma1         bool
	medicina        bool
	farmacia        bool
	cafe            bool
	pan             bool
	bakery          bool
	completeLevel   bool

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
	return p
}
func compar(p player) player {
	//compras en farmacia

	switch {
	case inpututil.IsKeyJustPressed(ebiten.KeyUp) && p.Coins >= 2 && plasma1 && farmacia:
		aspirina = true
		plasma1 = false
		medicina = false
	case (inpututil.IsKeyJustPressed(ebiten.KeyDown) && p.Coins >= 3 && !aspirina && !medicina && !plasma1 && farmacia) || (inpututil.IsKeyJustPressed(ebiten.KeyDown) && p.Coins >= 3 && aspirina && !plasma1 && farmacia) || (inpututil.IsKeyJustPressed(ebiten.KeyUp) && p.Coins >= 2 && medicina && !plasma1 && farmacia):
		plasma1 = true
		aspirina = false
		medicina = false
	case inpututil.IsKeyJustPressed(ebiten.KeyDown) && p.Coins >= 3 && plasma1 && farmacia:
		plasma1 = false
		aspirina = false
		medicina = true
	}
	//compras en bakery
	if inpututil.IsKeyJustPressed(ebiten.KeyUp) && p.Coins >= 2 && bakery {
		cafe = true
		pan = false
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyDown) && p.Coins >= 3 && bakery {
		pan = true
		cafe = false
	}

	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		if plasma1 || (!aspirina && !medicina) {
			p.Coins = p.Coins - 3
			p.vidas++
		}
		if aspirina || cafe {
			p.Coins = p.Coins - 2
			p.Fast = true
			p.CountPoder = 600
		}
		if (medicina && ModeGame == 0) || (pan && ModeGame == 1) {
			p.Coins = p.Coins - 2
			completeLevel = true
		}
		farmacia = false
		p.Compras = false
	}
	return p
}
func dibujarTextoCompras(p player, screen *ebiten.Image) {
	if p.Compras {
		if p.Coins < 2 {
			jugadores := fmt.Sprintf("YOU DONT HAVE MONEY\n  COME BACK SOON")
			text.Draw(screen, jugadores, arcadeFont, 230, 200, color.White)
		}
		//EN FARMACIA
		if p.Coins >= 2 && farmacia && !aspirina && !plasma1 && !medicina {
			jugadores := fmt.Sprintf(">$2-ASPIRIN -GO FAST-\n $3-PLASMA -GET LIFE-\n $2 MEDICINE")
			text.Draw(screen, jugadores, arcadeFont, 300, 250, color.White)
			fmt.Println("esta aca")

		}
		if aspirina {
			jugadores := fmt.Sprintf(">$2-ASPIRIN -GO FAST-\n $3-PLASMA -GET LIFE-\n $2 MEDICINE")
			text.Draw(screen, jugadores, arcadeFont, 300, 250, color.White)
			fmt.Println("1esta aca")
		}
		if plasma1 {
			jugadores := fmt.Sprintf(" $2-ASPIRIN -GO FAST-\n>$3-PLASMA -GET LIFE-\n $2 MEDICINE")
			text.Draw(screen, jugadores, arcadeFont, 300, 250, color.White)
			fmt.Println("2esta aca")

		}
		if medicina {
			jugadores := fmt.Sprintf(" $2-ASPIRIN -GO FAST-\n $3-PLASMA -GET LIFE-\n>$2 MEDICINE")
			text.Draw(screen, jugadores, arcadeFont, 300, 250, color.White)
			fmt.Println("3esta aca")

		}
		//EN BAKERY
		if p.Coins >= 2 && bakery && !pan && !cafe {
			jugadores := fmt.Sprintf(">$2-CAFE -GO FAST-\n $2-BREAD")
			text.Draw(screen, jugadores, arcadeFont, 300, 250, color.White)
		}
		if cafe {
			jugadores := fmt.Sprintf(">$2-CAFE -GO FAST-\n $2-BREAD")
			text.Draw(screen, jugadores, arcadeFont, 300, 250, color.White)
		}
		if pan {
			jugadores := fmt.Sprintf(" $2-CAFE -GO FAST-\n>$2-BREAD")
			text.Draw(screen, jugadores, arcadeFont, 300, 250, color.White)

		}
	}
}
