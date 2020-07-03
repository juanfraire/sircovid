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

//hombre
var hombre humanos
var mujer humanos
var humano humanos

const (
	// game
	screenWidth  = 768
	screenHeight = 528

	// viejo
	viejoFrameWidth  = 32
	viejoFrameHeight = 48

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
	ModeGame int

	ModeTitle    int
	ModeGameOver int
)

var (
	// imágenes
	imgTiles   *ebiten.Image
	imgNube    *ebiten.Image
	imgViejo   *ebiten.Image
	imgBarbijo *ebiten.Image

	// viejo
	viejoFrameOX    = 0
	viejoFrameOY    = 96
	viejoFrameNum   = 1
	viejoX          = float64(25)
	viejoY          = float64(375)
	posicionInicial int
	viejoMovX       int
	viejoMovY       int

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
	// Imagen city
	imgTiles, _, err = ebitenutil.NewImageFromFile(`sircovid\data\city2.png`, ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
	// Imangen Viejo
	imgViejo, _, err = ebitenutil.NewImageFromFile(`sircovid\data\viejo.png`, ebiten.FilterDefault)
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
	// Imagen nube Covid
	imgNube, _, err = ebitenutil.NewImageFromFile(`sircovid\data\smoke.png`, ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
	imgBarbijo, _, err = ebitenutil.NewImageFromFile(`sircovid\data\barbijo.png`, ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
	// Textos
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

//////////////////////////////////////////////////////////////////////////////////////////////
//////////////////////// AUDIO (separado variables y func init momentaneamente)///////////////
//////////////////////////////////////////////////////////////////////////////////////////////
var (
	audioContext   *audio.Context
	ragtimeContext *audio.Player
	deadSound      *audio.Player
	deadSound2     *audio.Player
)

func init() {
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
		moverViejo()

		// vida
		vida(hombre)

		//hombre
		hombre = moverHumanos(hombre)
		humano = hombre

		//pasar de nivel
		siguienteNivel()

	case ModeGame == 1 && vidas != 0:
		// nube
		moverNube()

		// viejo
		moverViejo()

		// vida
		vida(mujer)

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
func moverViejo() {
	// leer tecla
	switch {
	case inpututil.IsKeyJustPressed(ebiten.KeyRight) && viejoMovY != 1 && viejoMovY != 2:
		viejoFrameOY = 96
		viejoFrameNum = 3
		viejoMovX = 1
	case inpututil.IsKeyJustReleased(ebiten.KeyRight) && viejoMovY != 1 && viejoMovY != 2:
		viejoFrameNum = 1
		viejoMovX = 0
	case inpututil.IsKeyJustPressed(ebiten.KeyLeft) && viejoMovY != 1 && viejoMovY != 2:
		viejoFrameOY = 48
		viejoFrameNum = 3
		viejoMovX = 2
	case inpututil.IsKeyJustReleased(ebiten.KeyLeft) && viejoMovY != 1 && viejoMovY != 2:
		viejoFrameNum = 1
		viejoMovX = 0
	case inpututil.IsKeyJustPressed(ebiten.KeyUp) && viejoMovX != 1 && viejoMovX != 2:
		viejoFrameOY = 144
		viejoFrameNum = 3
		viejoMovY = 1
	case inpututil.IsKeyJustReleased(ebiten.KeyUp) && viejoMovX != 1 && viejoMovX != 2:
		viejoFrameNum = 1
		viejoMovY = 0
	case inpututil.IsKeyJustPressed(ebiten.KeyDown) && viejoMovX != 1 && viejoMovX != 2:
		viejoFrameOY = 0
		viejoFrameNum = 3
		viejoMovY = 2
	case inpututil.IsKeyJustReleased(ebiten.KeyDown) && viejoMovX != 1 && viejoMovX != 2:
		viejoFrameNum = 1
		viejoMovY = 0
	}

	// transladar viejo

	if ModeGame == 1 && posicionInicial != 1 {
		viejoX = float64(25)
		viejoY = float64(375)
		posicionInicial = 1
	}

	var viejoX1 = viejoX
	var viejoY1 = viejoY
	switch {
	case viejoMovX == 1:
		viejoX++
	case viejoMovX == 2:
		viejoX--
	case viejoMovY == 1:
		viejoY--
	case viejoMovY == 2:
		viejoY++
	}
	// restringir viejo
	switch {
	case viejoY < 300 && viejoX > 20 && viejoX < 214:
		viejoX = viejoX1
		viejoY = viejoY1
		viejoFrameNum = 1
	case viejoY < 130 && viejoX > 214 && viejoX < 768:
		viejoX = viejoX1
		viejoY = viejoY1
		viejoFrameNum = 1
	case viejoY < 270 && viejoX > 240 && viejoX < 610:
		viejoX = viejoX1
		viejoY = viejoY1
		viejoFrameNum = 1
	case viejoY < 270 && viejoX > 675 && viejoX < 768:
		viejoX = viejoX1
		viejoY = viejoY1
		viejoFrameNum = 1
	case viejoY > 335 && viejoY < 528 && viejoX > 40 && viejoX < 350:
		viejoX = viejoX1
		viejoY = viejoY1
		viejoFrameNum = 1
	case viejoY > 310 && viejoY < 450 && viejoX > 390 && viejoX < 630:
		viejoX = viejoX1
		viejoY = viejoY1
		viejoFrameNum = 1
	}
	fmt.Println(viejoX, viejoY)

}
func vida(h humanos) {
	//pierde vidas con la nuve
	nubeX := float64(nubeX * .4)
	nubeY := float64(nubeY * .4)
	if viejoX > nubeX && viejoX < nubeX+120 && viejoY > nubeY && viejoY < nubeY+120 {
		v++
	}
	if viejoX > h.X && viejoX < h.X+32 && viejoY+48 > h.Y && viejoY < h.Y+48 {
		v++
	}
	if viejoX > barbijoX && viejoX < barbijoX+32 && viejoY+48 > barbijoY && viejoY < barbijoY+48 {
		vidas++
		barbijoX = 1000
	}
	if v == 1 {
		vidas--
	}
	if v == 30 {
		v = 0
	}
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
func siguienteNivel() {
	if viejoX > 750 && viejoY > 450 {
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

	// dibujar viejo
	op.GeoM.Scale(.7, .7)
	op.GeoM.Translate(viejoX, viejoY)
	i := (g.count / 7) % viejoFrameNum
	sx, sy := viejoFrameOX+i*viejoFrameWidth, viejoFrameOY
	screen.DrawImage(imgViejo.SubImage(image.Rect(sx, sy, sx+viejoFrameWidth, sy+viejoFrameHeight)).(*ebiten.Image), op)

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
