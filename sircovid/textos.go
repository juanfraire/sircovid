package main

import (
	"fmt"
	"image/color"
	"log"

	"github.com/golang/freetype/truetype"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/text"
	"golang.org/x/image/font"
)

func initTextos() {

	intro1.initIntro(screenWidth, screenHeight)

	imgTiles, _, err = ebitenutil.NewImageFromFile(`sircovid\data\city.png`, ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}

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

func dibujarTextos(screen *ebiten.Image) {
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

		// fmt.Println(pulso)
		jugadores := fmt.Sprintf("PAUSE")
		text.Draw(screen, jugadores, arcadeFont, 300, 200, color.White)

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
