package main

import (
	"fmt"
	"image"
	"image/color"
	_ "image/png"
	"log"
	"math/rand"
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

type humanos struct {
	FrameOX     int
	FrameOY     int
	FrameNum    int
	X           float64
	Y           float64
	MovX        int
	MovY        int
	FrameWidth  int
	FrameHeight int
	img         *ebiten.Image
}
type player struct {
	humanos
	posicionInicial int
}

//hombre
var hombre humanos
var mujer humanos
var humano humanos

//viejo
var viejo player
var player1 player

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

var (
	ModeGame     int
	ModeTitle    int
	ModeGameOver int
)

var (
	// imágenes
	imgTiles   *ebiten.Image
	imgNube    *ebiten.Image
	imgBarbijo *ebiten.Image

	// sonido
	audioContext   *audio.Context
	ragtimeContext *audio.Player
	deadSound      *audio.Player
	deadSound2     *audio.Player

	//para mover humanos
	a, a1, a2, a3, a4, a5, a6 int
	a7, a8, a9, a10           int

	// nube
	nubeX       = float64(rand.Intn(screenWidth) + 300)
	nubeY       = float64(rand.Intn(200) + 600)
	nubeAlpha   float64
	nubeAlphaUp bool

	//barbijo
	barbijoFrameOX  = 0
	barbijoFrameOY  = 74
	barbijoFrameNum = 1
	barbijoX        = float64(630)
	barbijoY        = float64(150)

	// vidas
	vidas   = 3
	vidtime int
	v       int

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
	imgNube, _, err = ebitenutil.NewImageFromFile(`sircovid\data\smoke.png`, ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
	imgBarbijo, _, err = ebitenutil.NewImageFromFile(`sircovid\data\barbijo.png`, ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}

	////////////// SONIDOS //////////////
	audioContext, _ = audio.NewContext(44100)

	ragtimeD, err := vorbis.Decode(audioContext, audio.BytesReadSeekCloser(raudio.Ragtime_ogg))
	if err != nil {
		log.Fatal(err)
	}
	ragtimeContext, err = audio.NewPlayer(audioContext, ragtimeD)
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
}

// Game es la estructura del juego
type Game struct {
	count int
}

////////////////////////////
// Update
////////////////////////////

// Update se llama 60 veces por segundo
func (g *Game) Update(screen *ebiten.Image) error {

	if ModeTitle == 0 {
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
		player1 = viejo

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
	case ModeGame == 0 && vidas != 0:
		ragtimeContext.SetVolume(.5)
		ragtimeContext.Play()

		// nube
		moverNube()

		// viejo
		player1 = moverPlayer(player1)

		// vida
		vida(hombre, player1)

		//hombre
		hombre = moverHumanos(hombre)
		humano = hombre

		//pasar de nivel
		siguienteNivel(player1)

	case ModeGame == 1 && vidas != 0:
		// nube
		moverNube()

		// viejo
		player1 = moverPlayer(player1)

		// vida
		vida(mujer, viejo)

		//hombre
		mujer = moverHumanos(mujer)
		humano = mujer

	case ModeGameOver == 0:
		ragtimeContext.Pause()
		// deadSound.SetVolume(1)

		time.Sleep(time.Millisecond * 150)
		deadSound2.SetVolume(.4)
		deadSound2.Play()

		ragtimeContext.Rewind()
		if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
			ModeTitle = 0
			vidas = 3
			// nubeAlpha -= 1
			// viejoX = 25
			// viejoY = 375
		}
	}
	return nil
}

// moverViejo recorte de imagen segun la direccion de movimiento del viejo
func moverPlayer(p player) player {
	// leer tecla
	switch {
	case inpututil.IsKeyJustPressed(ebiten.KeyRight) && p.MovY != 1 && p.MovY != 2:
		p.FrameOY = 96
		p.FrameNum = 3
		p.MovX = 1
		return p
	case inpututil.IsKeyJustReleased(ebiten.KeyRight) && p.MovY != 1 && p.MovY != 2:
		p.FrameNum = 1
		p.MovX = 0
		return p

	case inpututil.IsKeyJustPressed(ebiten.KeyLeft) && p.MovY != 1 && p.MovY != 2:
		p.FrameOY = 48
		p.FrameNum = 3
		p.MovX = 2
		return p

	case inpututil.IsKeyJustReleased(ebiten.KeyLeft) && p.MovY != 1 && p.MovY != 2:
		p.FrameNum = 1
		p.MovX = 0
		return p

	case inpututil.IsKeyJustPressed(ebiten.KeyUp) && p.MovX != 1 && p.MovX != 2:
		p.FrameOY = 144
		p.FrameNum = 3
		p.MovY = 1
		return p

	case inpututil.IsKeyJustReleased(ebiten.KeyUp) && p.MovX != 1 && p.MovX != 2:
		p.FrameNum = 1
		p.MovY = 0
		return p

	case inpututil.IsKeyJustPressed(ebiten.KeyDown) && p.MovX != 1 && p.MovX != 2:
		p.FrameOY = 0
		p.FrameNum = 3
		p.MovY = 2
		return p

	case inpututil.IsKeyJustReleased(ebiten.KeyDown) && p.MovX != 1 && p.MovX != 2:
		p.FrameNum = 1
		p.MovY = 0
		return p

	}

	// transladar viejo

	if ModeGame == 1 && p.posicionInicial != 1 {
		p.X = float64(25)
		p.Y = float64(375)
		p.posicionInicial = 1
	}

	var X1 = p.X
	var Y1 = p.Y
	switch {
	case p.MovX == 1:
		p.X++
		return p

	case p.MovX == 2:
		p.X--
		return p

	case p.MovY == 1:
		p.Y--
		return p

	case p.MovY == 2:
		p.Y++
		return p

	}
	// restringir viejo
	switch {
	case p.Y < 300 && p.X > 20 && p.X < 214:
		p.X = X1
		p.Y = Y1
		p.FrameNum = 1
		return p

	case p.Y < 130 && p.X > 214 && p.X < 768:
		p.X = X1
		p.Y = Y1
		p.FrameNum = 1
		return p

	case p.Y < 270 && p.X > 240 && p.X < 610:
		p.X = X1
		p.Y = Y1
		p.FrameNum = 1
		return p

	case p.Y < 270 && p.X > 675 && p.X < 768:
		p.X = X1
		p.Y = Y1
		p.FrameNum = 1
		return p

	case p.Y > 335 && p.Y < 528 && p.X > 40 && p.X < 350:
		p.X = X1
		p.Y = Y1
		p.FrameNum = 1
		return p

	case p.Y > 310 && p.Y < 450 && p.X > 390 && p.X < 630:
		p.X = X1
		p.Y = Y1
		p.FrameNum = 1
		return p

	}
	return p

}
func vida(h humanos, p player) player {
	//pierde vidas con la nuve
	nubeX := float64(nubeX * .4)
	nubeY := float64(nubeY * .4)
	if p.X > nubeX && p.X < nubeX+120 && p.Y > nubeY && p.Y < nubeY+120 {
		v++
	}
	if p.X > h.X && p.X < h.X+32 && p.Y+48 > h.Y && p.Y < h.Y+48 {
		v++
	}
	if p.X > barbijoX && p.X < barbijoX+32 && p.Y+48 > barbijoY && p.Y < barbijoY+48 {
		vidas++
		barbijoX = 1000
	}
	if v == 1 {
		vidas--
	}
	if v == 30 {
		v = 0
	}
	return p
}

// nubeCovid aumenta y disminuye transparencia de la nube (alpha)
func moverNube() {
	// creacion de nuevas nubes
	if nubeAlpha <= 0 {
		nubeX = float64(rand.Intn(1500))
		nubeY = float64(rand.Intn(500) + 600)
		nubeAlphaUp = true
	} else if nubeAlpha > 1 {
		time.Sleep(10000 * time.Microsecond)
		nubeAlphaUp = false
	}

	// movimiento nube
	if nubeAlpha >= 0 {
		nubeX--
	}

	// actualizar alpha
	if nubeAlphaUp {
		nubeAlpha += .009
	} else {
		nubeAlpha -= .003
	}
}

func moverHumanos(h humanos) humanos {
	h.FrameNum = 4
	switch {
	case a != 1:
		h.FrameOY = 48
		h.Y = 290
		h.X--
		if h.X < 228 {
			a = 1
		}
		return h
	case a == 1 && a1 != 1:
		h.FrameOY = 144
		h.Y--
		if h.Y == 137 {
			a1 = 1
		}
		return h
	case a1 == 1 && a2 != 1:
		h.FrameOY = 0
		h.Y++
		if h.Y == 310 {
			a2 = 1
		}
		return h
	case a2 == 1 && a3 != 1:
		h.FrameOY = 48
		h.X--
		if h.X == -100 {
			a3 = 1
		}
		return h
	case a3 == 1 && a4 != 1:
		h.FrameOY = 96
		h.Y = 460
		h.X++
		if h.X == 20 {
			a4 = 1
		}
		return h
	case a4 == 1 && a5 != 1:
		h.FrameOY = 144
		h.Y--
		if h.Y == 310 {
			a5 = 1
		}
		return h
	case a5 == 1 && a6 != 1:
		h.FrameOY = 96
		h.X++
		if h.X == 228 {
			a6 = 1
		}
		return h
	case a6 == 1 && a7 != 1:
		h.FrameOY = 144
		h.Y--
		if h.Y == 280 {
			a7 = 1
		}
		return h
	case a7 == 1 && a8 != 1:
		h.FrameOY = 96
		h.X++
		if h.X == 370 {
			a8 = 1
		}
		return h
	case a8 == 1 && a9 != 1:
		h.FrameOY = 0
		h.Y++
		if h.Y == 470 {
			a9 = 1
		}
		return h
	case a9 == 1:
		h.FrameOY = 96
		h.X++
		if h.X == 800 {
			a, a1, a2, a3, a4, a5, a6, a7, a8, a9 = 0, 0, 0, 0, 0, 0, 0, 0, 0, 0
		}
		return h
	}
	return h

}
func siguienteNivel(p player) {
	if p.X > 750 && p.Y > 450 {
		ModeTitle = 2
		ModeGame = 1
	}
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
	op.GeoM.Translate(humano.X, humano.Y)
	j := (g.count / 7) % humano.FrameNum
	hx, hy := humano.FrameOX+j*humano.FrameWidth, humano.FrameOY
	screen.DrawImage(humano.img.SubImage(image.Rect(hx, hy, hx+humano.FrameWidth, hy+humano.FrameHeight)).(*ebiten.Image), op)

	// dibujar nube
	op = &ebiten.DrawImageOptions{}
	op.GeoM.Translate(nubeX, nubeY)
	op.ColorM.Scale(1, 3, 2, nubeAlpha)
	op.GeoM.Scale(.4, .4)
	screen.DrawImage(imgNube, op)

	// dibujar barbijo
	op = &ebiten.DrawImageOptions{}
	op.GeoM.Scale(.5, .5)
	op.GeoM.Translate(barbijoX, barbijoY)
	bx, by := barbijoFrameOX+barbijoFrameWidth, barbijoFrameOY
	screen.DrawImage(imgBarbijo.SubImage(image.Rect(bx, by, bx+barbijoFrameWidth, by+barbijoFrameHeight)).(*ebiten.Image), op)

	// dibujar texto
	lifes := fmt.Sprintf("Vidas:%02d", vidas)
	text.Draw(screen, lifes, smallArcadeFont, fontSize, 40, color.RGBA{35, 27, 190, 0xff})

	switch {
	case ModeTitle == 0:
		texts = []string{"SIR-COVID", "", "PRIMER NIVEL", "", "", "PRESS SPACE KEY"}
	case ModeTitle == 1 && vidas != 0:
		texts = []string{}

	case ModeTitle == 2:
		texts = []string{"", "", "SEGUNDO NIVEL", "", "", "PRESS SPACE KEY"}
	case ModeTitle == 3 && vidas != 0:
		texts = []string{}

	case vidas == 0:
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
