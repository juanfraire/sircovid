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
	text.Draw(screen, lifesP1, smallArcadeFont, fontSize, 40, color.White)
	if Game1.numPlayers == 2 {
		lifesP2 := fmt.Sprintf("Life:%02d", player2.vidas)
		text.Draw(screen, lifesP2, smallArcadeFont, 900, 40, color.White)
	}
	//dibujar monedas
	CoinsP1 := fmt.Sprintf("Coins:%d", player1.Coins)
	text.Draw(screen, CoinsP1, smallArcadeFont, fontSize, 70, color.White)
	//cartel para llamar a comandos
	Comandos := fmt.Sprintf("Commans: Key C")
	text.Draw(screen, Comandos, smallArcadeFont, 10, screenHeight-10, color.White)

	if Game1.numPlayers == 2 {
		CoinsP1 := fmt.Sprintf("Coins:%d", player2.Coins)
		text.Draw(screen, CoinsP1, smallArcadeFont, 900, 70, color.White)
	}
	//dibujar inmunidad
	if player1.Inmune {
		Inm := fmt.Sprintf("Inmune for:%02d", player1.CountPoder/60)
		text.Draw(screen, Inm, smallArcadeFont, fontSize, 100, color.White)
	}
	if player2.Inmune {
		Inm := fmt.Sprintf("Inmune for:%02d", player2.CountPoder/60)
		text.Draw(screen, Inm, smallArcadeFont, 840, 100, color.White)
	}
	//dibujar Fast
	if player1.Fast {
		Inm := fmt.Sprintf("Fast for:%02d", player1.CountPoder/60)
		text.Draw(screen, Inm, smallArcadeFont, fontSize, 130, color.White)
	}
	if player2.Fast {
		Inm := fmt.Sprintf("Fast for:%02d", player2.CountPoder/60)
		text.Draw(screen, Inm, smallArcadeFont, 840, 130, color.White)
	}
	switch {
	case ModeTitle:
		// intro draw
		intro1.drawIntro(screen, screenWidth, screenHeight)

	case ModeTitleLevel:
		nivel := fmt.Sprintf("    LEVEL %d\n\nPRESS SPACE KEY", Level)
		text.Draw(screen, nivel, arcadeFont, 300, 280, color.White)
	case ModeGameOver:
		lost := fmt.Sprintf("  GAME OVER!\n\n  TRAY AGAIN?\n\nPRESS SPACE KEY")
		text.Draw(screen, lost, arcadeFont, 310, 200, color.White)
		if player1.Coins < 2 && player2.Coins < 2 && monedas.X == 1500 {
			noMoney := fmt.Sprintf("(NOT COMPLETE LEVEL)")
			text.Draw(screen, noMoney, arcadeFont, 230, 370, color.White)
		}
	case ModeWin == true:
		win := fmt.Sprintf("YOU WIN")
		text.Draw(screen, win, arcadeFont, 400, 300, color.White)
	}
	switch {
	case ModePause && count1 < 40:
		jugadores := fmt.Sprintf("PAUSE")
		text.Draw(screen, jugadores, arcadeFont, 450, 270, color.White)

		//elegir numero de jugadores
	case ElectNumPlayers == 0 && Game1.numPlayers == 1:
		jugadores := fmt.Sprintf(">1 PLAYER\n 2 PLAYERS")
		text.Draw(screen, jugadores, arcadeFont, 350, 300, color.White)
	case ElectNumPlayers == 0 && Game1.numPlayers == 2:
		jugadores := fmt.Sprintf(" 1 PLAYER\n>2 PLAYERS")
		text.Draw(screen, jugadores, arcadeFont, 350, 300, color.White)

		//elegir jugador
	case ElectPlayer == 0:
		jugadores := fmt.Sprintf("CHOOSE PLAYER")
		text.Draw(screen, jugadores, arcadeFont, 320, 300, color.White)

	case ElectNumPlayers == 1:
		for i, l := range texts {
			x := (screenWidth - len(l)*fontSize) / 2
			text.Draw(screen, l, arcadeFont, x, (i+5)*fontSize, color.White)
		}
	}
	//dibujar Comandos
	if Commands {
		plasmaVida := fmt.Sprintf("Player 1 Keys:\nUp: Key Up\nDow: Key Down\nRigth: Key Rigth\nLeft: Key Left")
		text.Draw(screen, plasmaVida, smallArcadeFont, 325, 150, color.White)
		alcholInmune := fmt.Sprintf("Player 2 Keys:\nUp: Key W\nDow: Key S\nRigth: Key D\nLeft: Key A")
		text.Draw(screen, alcholInmune, smallArcadeFont, 325, 250, color.White)
		barbijoInmune := fmt.Sprintf("Pause: P\nSilence: X\nVolume Up: Ctrl Key Up\nVolume Down: Ctrl Key Down")
		text.Draw(screen, barbijoInmune, smallArcadeFont, 325, 350, color.White)

	}
}
