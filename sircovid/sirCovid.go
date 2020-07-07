package main

import (
	"fmt"
	"image"
	"image/color"
	_ "image/png"
	"log"
	"math/rand"
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
	audioContext   *audio.Context
	ragtimeContext *audio.Player
	deadSound      *audio.Player
	deadSound2     *audio.Player
	sonidoFondo    *audio.InfiniteLoop
	fondo          *audio.Player

	//para mover humanos
	a, a1, a2, a3, a4, a5, a6 int
	a7, a8, a9, a10           int

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

	audioContext, _ = audio.NewContext(44100)
	s, err := os.Open(`sircovid\data\SIR-COVID.wav`)
	if err != nil {
		panic(err)
	}
	defer s.Close()
	data := make([]byte, 11491248)
	c, err := s.Read(data)
	fondoD, err := wav.Decode(audioContext, audio.BytesReadSeekCloser(data))
	fmt.Println(c)
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

	// ragtimeD, err := vorbis.Decode(audioContext, audio.BytesReadSeekCloser(raudio.Ragtime_ogg))
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// ragtimeContext, err = audio.NewPlayer(audioContext, ragtimeD)
	// if err != nil {
	// 	log.Fatal(err)
	// }

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

	hombre.FrameOX = 0
	hombre.FrameOY = 48
	hombre.FrameNum = 1
	hombre.X = float64(750)
	hombre.Y = float64(290)
	hombre.FrameWidth = 32
	hombre.FrameHeight = 48

	mujer.FrameOX = 0
	mujer.FrameOY = 48
	mujer.FrameNum = 1
	mujer.X = float64(750)
	mujer.Y = float64(290)
	mujer.FrameWidth = 32
	mujer.FrameHeight = 48
	humano = hombre

	// viejo
	viejo.FrameOX = 0
	viejo.FrameOY = 96
	viejo.FrameNum = 1
	viejo.X = float64(25)
	viejo.Y = float64(375)
	viejo.FrameWidth = 32
	viejo.FrameHeight = 48

	//player
	player1.humanos = viejo
	player1.vidas = 3
	player1.v = 0

	//enemigos
	enemigos1.humanos = hombre

	//nube
	nube1.nubeX = float64(rand.Intn(screenWidth) + 300)
	nube1.nubeY = float64(rand.Intn(200) + 600)

	//defino el juego
	Game1.nube = nube1
	Game1.siguienteNivel = player1.humanos

}

////////////////////////////
// Update
////////////////////////////

// Update se llama 60 veces por segundo
func (g *Game) Update(screen *ebiten.Image) error {

	if ModeTitle == 0 {

	}

	// game counter
	g.count++

	switch {
	case ModeTitle == 0:
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
		// sonidoFondo.SetVolume(.)
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

		// nube
		Game1.nube = moverNube(Game1.nube)

		// player
		player1.humanos = moverPlayer(player1.humanos)

		// vida
		player1 = vida(enemigos1.humanos, player1.humanos)

		//mujer
		enemigos1.humanos = moverHumanos(enemigos1.humanos)

	case ModeGameOver == 0:

		fondo.Pause()
		fondo.Rewind()
		time.Sleep(time.Millisecond * 100)
		deadSound2.SetVolume(.4)
		deadSound2.Play()

		if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
			ModeGame = 0
			player1.vidas = 3
			// nubeAlpha -= 1
			// viejoX = 25
			// viejoY = 375
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
	op.GeoM.Translate(Game1.nubeX, Game1.nubeY)
	op.ColorM.Scale(1, 3, 2, Game1.nubeAlpha)
	op.GeoM.Scale(.4, .4)
	screen.DrawImage(Game1.nube.img, op)

	// dibujar barbijo
	op = &ebiten.DrawImageOptions{}
	op.GeoM.Scale(.5, .5)
	op.GeoM.Translate(barbijoX, barbijoY)
	bx, by := barbijoFrameOX+barbijoFrameWidth, barbijoFrameOY
	screen.DrawImage(imgBarbijo.SubImage(image.Rect(bx, by, bx+barbijoFrameWidth, by+barbijoFrameHeight)).(*ebiten.Image), op)

	// dibujar texto
	lifes := fmt.Sprintf("Vidas:%02d", player1.vidas)
	text.Draw(screen, lifes, smallArcadeFont, fontSize, 40, color.RGBA{35, 27, 190, 0xff})

	switch {
	case ModeTitle == 0:
		texts = []string{"SIR-COVID", "", "PRIMER NIVEL", "", "", "PRESS SPACE KEY"}
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
