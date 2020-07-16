package main

import (
	"fmt"
	"image/color"
	_ "image/png"
	"log"
	"os"
	"time"

	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"

	// "golang.ge/font"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/audio"
	"github.com/hajimehoshi/ebiten/audio/vorbis"
	"github.com/hajimehoshi/ebiten/audio/wav"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	raudio "github.com/hajimehoshi/ebiten/examples/resources/audio"
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
)

var (
	// imágenes
	imgTiles *ebiten.Image
	//imgBarbijo *ebiten.Image

	// sonido
	audioContext *audio.Context
	deadSound    *audio.Player
	deadSound2   *audio.Player
	sonidoFondo  *audio.InfiniteLoop
	fondo        *audio.Player

	//barbijo
	// barbijoFrameOX  = 0
	// barbijoFrameOY  = 74
	// barbijoFrameNum = 1
	// barbijoX        = float64(630)
	// barbijoY        = float64(150)

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
	imgTiles, _, err = ebitenutil.NewImageFromFile(`sircovid\data\city2.png`, ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
	//inicializa a players
	initPlayer()
	//inicia sumarVidas
	initSumarVidas()
	//inicia nube
	initNube()
	///////////// Imagen NUBE COVID ///////////////////
	// nube1.img, _, err = ebitenutil.NewImageFromFile(`sircovid\data\smoke.png`, ebiten.FilterDefault)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	////////////// SONIDOS //////////////

	audioContext, _ = audio.NewContext(44100)
	s, err := os.Open(`sircovid\data\SIR-COVID.wav`)
	if err != nil {
		panic(err)
	}
	defer s.Close()
	data := make([]byte, 11491248)
	c, err := s.Read(data)
	fondoD, err := wav.Decode(audioContext, audio.BytesReadSeekCloser(data))
	//fmt.Println(c)
	if err != nil {
		log.Fatal(err)
	}
	sonidoFondo = audio.NewInfiniteLoop(fondoD, int64(c))
	if err != nil {
		log.Fatal(err)
	}
	fondo, err = audio.NewPlayer(audioContext, sonidoFondo)
	if err != nil {
		log.Fatal(err)
	}

	jumpD, err := vorbis.Decode(audioContext, audio.BytesReadSeekCloser(raudio.Jump_ogg))
	if err != nil {
		log.Fatal(err)
	}
	deadSound, err = audio.NewPlayer(audioContext, jumpD)
	if err != nil {
		log.Fatal(err)
	}
	jabD, err := wav.Decode(audioContext, audio.BytesReadSeekCloser(raudio.Jab_wav))
	if err != nil {
		log.Fatal(err)
	}
	deadSound2, err = audio.NewPlayer(audioContext, jabD)
	if err != nil {
		log.Fatal(err)

	}

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
	iniciarVariables()

}

////////////////////////////
// Update
////////////////////////////

// Update se llama 60 veces por segundo
func (g *Game) Update(screen *ebiten.Image) error {

	//  maneja sonido fondo (S = mute) en proceso (no sacar)
	if inpututil.IsKeyJustPressed(ebiten.KeyX) {
		if fondo.Volume() != 0 {
			fondo.SetVolume(0)
		} else if fondo.Volume() == 0 {
			fondo.SetVolume(1)
		}
	}

	// game counter
	g.count++
	count1++
	switch {
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
			if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
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
			if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
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
		deadSound2.Rewind()
		fondo.Play()

		// nube
		nube1 = moverNube(nube1)

		// palyer
		player1 = moverPlayer(player1)
		player2 = moverPlayer(player2)

		// vida
		player1 = vida(enemigo1, player1, barbijo)
		if Game1.numPlayers == 2 {
			player2 = vida(enemigo1, player2, barbijo)
		}
		//enemigo1
		enemigo1 = moverHumanos(enemigo1)

		//siguiente nivel
		Game1.siguienteNivel = siguienteNivel(player1)
		Game1.siguienteNivel = siguienteNivel(player2)

	case ModeGame == 1 && player1.vidas != 0 && player2.vidas != 0:

		// sonido
		deadSound2.Rewind()
		fondo.Play()

		// nube
		nube1 = moverNube(nube1)

		// player
		player1 = moverPlayer(player1)
		player2 = moverPlayer(player2)

		// vida
		player1 = vida(enemigo1, player1, barbijo)
		player1 = vida(enemigo2, player1, barbijo)
		if Game1.numPlayers == 2 {
			player2 = vida(enemigo1, player2, barbijo)
			player2 = vida(enemigo2, player2, barbijo)
		}
		//enemigos
		enemigo1 = moverHumanos(enemigo1)
		enemigo2 = moverHumanos(enemigo2)

	case ModeGameOver == 0:
		if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
			iniciarVariables()
			initPlayer()
		}

		// sonido
		fondo.Pause()
		fondo.Rewind()
		time.Sleep(time.Millisecond * 100)
		deadSound2.SetVolume(.4)
		deadSound2.Play()

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
	}
	if ModeGame == 1 {
		dibujarEnemigos(enemigo1, screen)
		dibujarEnemigos(enemigo2, screen)
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
	switch {
	case ElectNumPlayers == 0:
		jugadores := fmt.Sprintf(" Elija la cantidad de jugadores \n\ncon las flechas y presione enter\n\nJUGADORES:%02d", Game1.numPlayers)
		text.Draw(screen, jugadores, smallArcadeFont, 140, 250, color.White)

	case ElectPlayer == 0 && Game1.numPlayers == 1:
		jugadores := fmt.Sprintf("Elija el jugador \n\n con las flechas\n\ny presione enter")
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
