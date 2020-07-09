package main

import (
	"fmt"
	"image"
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

//hombre
var hombre humanos
var mujer humanos
var humano humanos
var viejo humanos

//intro
var intro1 intro

//jugador
var player1 player

//nube
var nube1 nube

// Game1 es el juego
var Game1 Game

//humanos enemigos
var enemigos1 enemigos

const (
	// game
	screenWidth  = 768
	screenHeight = 528

	//barbijo
	barbijoFrameWidth  = 105
	barbijoFrameHeight = 48

	// tiles
	tileSize = 16
	tileXNum = 48

	//para start y game Over
	fontSize      = 32
	smallFontSize = fontSize / 2
)

//Mode Game es el en que parte del juego estamos
var (
	ModeGame     int
	ModeTitle    int
	ModeGameOver int
)

var (
	// imágenes
	imgTiles   *ebiten.Image
	imgBarbijo *ebiten.Image

	// sonido
	audioContext *audio.Context
	deadSound    *audio.Player
	deadSound2   *audio.Player
	sonidoFondo  *audio.InfiniteLoop
	fondo        *audio.Player

	//para mover humanos
	// a, a1, a2, a3, a4, a5, a6 int
	// a7, a8, a9, a10           int

	//barbijo
	barbijoFrameOX  = 0
	barbijoFrameOY  = 74
	barbijoFrameNum = 1
	barbijoX        = float64(630)
	barbijoY        = float64(150)

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
	//////////////   Imangen VIEJO  //////////////////////////////
	viejo.img, _, err = ebitenutil.NewImageFromFile(`sircovid\data\viejo.png`, ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
	hombre.img, _, err = ebitenutil.NewImageFromFile(`sircovid\data\hombre.png`, ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
	mujer.img, _, err = ebitenutil.NewImageFromFile(`sircovid\data\mujer.png`, ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
	///////////// Imagen NUBE COVID ///////////////////
	nube1.img, _, err = ebitenutil.NewImageFromFile(`sircovid\data\smoke.png`, ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
	imgBarbijo, _, err = ebitenutil.NewImageFromFile(`sircovid\data\barbijo.png`, ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}

	////////////// SONIDOS //////////////

	// contexto de decodificacion
	audioContext, _ = audio.NewContext(44100)

	// abre sonido de fondo
	s, err := os.Open(`sircovid\data\SIR-COVID.wav`)
	if err != nil {
		panic(err)
	}
	defer s.Close()
	data := make([]byte, 11491248)
	c, err := s.Read(data)

	// decodifico sonido fondo
	fondoD, err := wav.Decode(audioContext, audio.BytesReadSeekCloser(data))
	if err != nil {
		log.Fatal(err)
	}

	// creo loop infinito de sonido
	sonidoFondo = audio.NewInfiniteLoop(fondoD, int64(c))
	if err != nil {
		log.Fatal(err)
	}

	// sonido de reproduccion
	fondo, err = audio.NewPlayer(audioContext, sonidoFondo)
	if err != nil {
		log.Fatal(err)
	}

	// decode y creacion sonido a reproducir - SONIDO PERDER VIDAS
	jumpD, err := vorbis.Decode(audioContext, audio.BytesReadSeekCloser(raudio.Jump_ogg))
	if err != nil {
		log.Fatal(err)
	}
	deadSound, err = audio.NewPlayer(audioContext, jumpD)
	if err != nil {
		log.Fatal(err)
	}

	// decode y creacion -SONIDO MUERTE FINAL
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
	if inpututil.IsKeyJustPressed(ebiten.KeyS) {
		if fondo.Volume() != 0 {
			fondo.SetVolume(0)
		} else if fondo.Volume() == 0 {
			fondo.SetVolume(1)
		}
	}

	// game counter
	g.count++

	switch {
	case ModeTitle == 0:

		// intro update
		intro1.updateIntro(screenWidth, screenHeight)

		if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
			ModeTitle = 1
		}
	case ModeTitle == 2:
		if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
			ModeTitle = 3
		}
	case ModeGame == 0 && player1.vidas != 0:
		//// sonido ////
		deadSound2.Rewind()
		fondo.Play()

		// nube
		Game1.nube = moverNube(Game1.nube)

		// palyer
		player1.humanos = moverPlayer(player1.humanos)

		// vida
		player1 = vida(enemigos1.humanos, player1.humanos)

		//hombre
		enemigos1.humanos = moverHumanos(enemigos1.humanos)

		//pasar de nivel
		Game1.siguienteNivel = siguienteNivel(player1.humanos)

	case ModeGame == 1 && player1.vidas != 0:
		// sonido
		deadSound2.Rewind()
		fondo.Play()

		// nube
		Game1.nube = moverNube(Game1.nube)

		// player
		player1.humanos = moverPlayer(player1.humanos)

		// vida
		player1 = vida(enemigos1.humanos, player1.humanos)

		//mujer
		enemigos1.humanos = moverHumanos(enemigos1.humanos)

	case ModeGameOver == 0:
		if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
			iniciarVariables()
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

	// dibujar player
	op.GeoM.Scale(.7, .7)
	op.GeoM.Translate(player1.X, player1.Y)
	i := (g.count / 7) % player1.FrameNum
	sx, sy := player1.FrameOX+i*player1.FrameWidth, player1.FrameOY
	screen.DrawImage(player1.img.SubImage(image.Rect(sx, sy, sx+player1.FrameWidth, sy+player1.FrameHeight)).(*ebiten.Image), op)

	//dibujar humano
	op = &ebiten.DrawImageOptions{}
	op.GeoM.Scale(.8, .8)
	op.GeoM.Translate(enemigos1.humanos.X, enemigos1.humanos.Y)
	j := (g.count / 7) % enemigos1.humanos.FrameNum
	hx, hy := enemigos1.humanos.FrameOX+j*enemigos1.humanos.FrameWidth, enemigos1.humanos.FrameOY
	screen.DrawImage(enemigos1.humanos.img.SubImage(image.Rect(hx, hy, hx+enemigos1.humanos.FrameWidth, hy+enemigos1.humanos.FrameHeight)).(*ebiten.Image), op)

	// dibujar nube
	op = &ebiten.DrawImageOptions{}
	op.GeoM.Translate(Game1.nubeX, Game1.nubeY+384)
	op.ColorM.Scale(1, 3, 2, Game1.nubeAlpha)
	op.GeoM.Scale(.4, .4)
	screen.DrawImage(Game1.nube.img, op)

	// dibujar barbijo
	op = &ebiten.DrawImageOptions{}
	op.GeoM.Scale(.3, .3)
	op.GeoM.Translate(barbijoX, barbijoY)
	bx, by := barbijoFrameOX+barbijoFrameWidth, barbijoFrameOY
	screen.DrawImage(imgBarbijo.SubImage(image.Rect(bx, by, bx+barbijoFrameWidth, by+barbijoFrameHeight)).(*ebiten.Image), op)

	// dibujar texto
	lifes := fmt.Sprintf("Vidas:%02d", player1.vidas)
	text.Draw(screen, lifes, smallArcadeFont, fontSize, 40, color.RGBA{35, 27, 190, 0xff})

	switch {
	case ModeTitle == 0:

		// intro draw
		intro1.drawIntro(screen, screenWidth, screenHeight)
		texts = []string{"", "", "", "", "", "", "PRIMER NIVEL", "", "", "PRESS SPACE KEY"}

	case ModeTitle == 1 && player1.vidas != 0:
		texts = []string{}

	case ModeTitle == 2:
		texts = []string{"", "", "SEGUNDO NIVEL", "", "", "PRESS SPACE KEY"}
	case ModeTitle == 3 && player1.vidas != 0:
		texts = []string{}

	case player1.vidas == 0:
		texts = []string{"", "", "", "GAME OVER!", "", "TRY AGAIN?", "", "PRESS SPACE KEY"}
	}
	for i, l := range texts {
		x := (screenWidth - len(l)*fontSize) / 2
		text.Draw(screen, l, arcadeFont, x, (i+5)*fontSize, color.White)
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
