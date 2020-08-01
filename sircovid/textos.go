package main

import (
	"fmt"
	"image/color"
	"log"

	"github.com/golang/freetype/truetype"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/text"
	"golang.org/x/image/font"
)

func initTextos() {

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
	// dibujar vidas
	lifesP1 := fmt.Sprintf("Life:%02d", player1.vidas)
	text.Draw(screen, lifesP1, smallArcadeFont, fontSize, 40, color.RGBA{35, 27, 190, 0xff})
	if Game1.numPlayers == 2 {
		lifesP2 := fmt.Sprintf("Life:%02d", player2.vidas)
		text.Draw(screen, lifesP2, smallArcadeFont, 900, 40, color.RGBA{35, 27, 190, 0xff})
	}
	//dibujar monedas
	CoinsP1 := fmt.Sprintf("Coins:%d", player1.Coins)
	text.Draw(screen, CoinsP1, smallArcadeFont, fontSize, 70, color.RGBA{35, 27, 190, 0xff})

	if Game1.numPlayers == 2 {
		CoinsP1 := fmt.Sprintf("Coins:%d", player2.Coins)
		text.Draw(screen, CoinsP1, smallArcadeFont, 900, 70, color.RGBA{35, 27, 190, 0xff})
	}
	//dibujar inmunidad
	if player1.Inmune {
		Inm := fmt.Sprintf("Inmune for:%02d", player1.CountPoder/60)
		text.Draw(screen, Inm, smallArcadeFont, fontSize, 100, color.RGBA{35, 27, 190, 0xff})
	}
	if player2.Inmune {
		Inm := fmt.Sprintf("Inmune for:%02d", player2.CountPoder/60)
		text.Draw(screen, Inm, smallArcadeFont, 840, 100, color.RGBA{35, 27, 190, 0xff})
	}
	//dibujar Fast
	if player1.Fast {
		Inm := fmt.Sprintf("Fast for:%02d", player1.CountPoder/60)
		text.Draw(screen, Inm, smallArcadeFont, fontSize, 130, color.RGBA{35, 27, 190, 0xff})
	}
	if player2.Fast {
		Inm := fmt.Sprintf("Fast for:%02d", player2.CountPoder/60)
		text.Draw(screen, Inm, smallArcadeFont, 840, 130, color.RGBA{35, 27, 190, 0xff})
	}
	switch {
	case ModeTitle:
		// intro draw
		intro1.drawIntro(screen, screenWidth, screenHeight)

	case ModeTitleLevel:
		nivel := fmt.Sprintf("    LEVEL %d\n\nPRESS SPACE KEY", Level)
		text.Draw(screen, nivel, arcadeFont, 310, 250, color.White)
	case ModeGameOver:
		lost := fmt.Sprintf("  GAME OVER!\n\n  TRAY AGAIN?\n\nPRESS SPACE KEY")
		text.Draw(screen, lost, arcadeFont, 310, 200, color.White)
	case ModeWin == true:
		win := fmt.Sprintf("YOU ARBITRARILY\n\n      WIN")
		text.Draw(screen, win, arcadeFont, 300, 200, color.White)
	}
	switch {
	case ModePause && count1 < 40:
		jugadores := fmt.Sprintf("PAUSE")
		text.Draw(screen, jugadores, arcadeFont, 450, 200, color.White)

		//elegir numero de jugadores
	case ElectNumPlayers == 0 && Game1.numPlayers == 1:
		jugadores := fmt.Sprintf(">1 PLAYER\n 2 PLAYERS")
		text.Draw(screen, jugadores, arcadeFont, 360, 250, color.White)
	case ElectNumPlayers == 0 && Game1.numPlayers == 2:
		jugadores := fmt.Sprintf(" 1 PLAYER\n>2 PLAYERS")
		text.Draw(screen, jugadores, arcadeFont, 360, 250, color.White)

		//elegir jugador
	case ElectPlayer == 0:
		jugadores := fmt.Sprintf("CHOOSE PLAYER")
		text.Draw(screen, jugadores, arcadeFont, 330, 250, color.White)

	case ElectNumPlayers == 1:
		for i, l := range texts {
			x := (screenWidth - len(l)*fontSize) / 2
			text.Draw(screen, l, arcadeFont, x, (i+5)*fontSize, color.White)
		}
	}
}
