package main

import (
	"fmt"
	"image"
	"image/color"
	_ "image/png"
	"log"
	"time"

	"github.com/golang/freetype/truetype"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/inpututil"
	"github.com/hajimehoshi/ebiten/text"
	"golang.org/x/image/font"
)

const (
	// game
	screenWidth  = 768
	screenHeight = 528

	// viejo
	viejoFrameWidth  = 32
	viejoFrameHeight = 48

	// tiles
	tileSize = 16
	tileXNum = 48

	// texto
	fontSize = 20
)

var (
	// imágenes
	imgTiles *ebiten.Image
	imgNube  *ebiten.Image
	imgViejo *ebiten.Image

	// viejo
	viejoFrameOX  = 0
	viejoFrameOY  = 96
	viejoFrameNum = 1
	viejoX        = float64(25)
	viejoY        = float64(375)
	viejoMovX     int
	viejoMovY     int

	// nube
	nubeX       = float64(screenWidth + 1200)
	nubeY       = float64(800)
	nubeAlpha   float64
	nubeAlphaUp bool

	// vidas
	vidas   = 3
	vidtime int

	// texto
	arcadeFont font.Face

	err error
)

// init carga los datos
func init() {

	imgTiles, _, err = ebitenutil.NewImageFromFile(`C:\Aprendiendo Go\Games\sircovid\sircovid\data\city2.png`, ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
	imgViejo, _, err = ebitenutil.NewImageFromFile(`C:\Aprendiendo Go\Games\sircovid\sircovid\data\viejo.png`, ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
	imgNube, _, err = ebitenutil.NewImageFromFile(`C:\Aprendiendo Go\Games\sircovid\sircovid\data\smoke.png`, ebiten.FilterDefault)
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

	// game counter
	g.count++

	// nube
	moverNube()

	// viejo
	moverViejo()

	// vida
	// vida()

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
}

// vida estoy trabajando en esta funcion de abajo
func vida() {
	if vidas == vidtime {
		time.Sleep(1 * time.Second)
	}
	if int(viejoX) > int(nubeX) && int(viejoX) < int(nubeX+50) && int(viejoY) < 428 && int(viejoY) > 298 {
		vidas = vidas - 1
		vidtime = vidas
		fmt.Println(vidas)
	}

}

// nubeCovid aumenta y disminuye transparencia de la nube (alpha)
func moverNube() {
	nubeX--

	// actualizar alpha
	if nubeAlphaUp {
		nubeAlpha += .003
	} else {
		nubeAlpha -= .003
	}

	// cambiar dirección alpha
	if nubeAlpha < 1 {
		nubeAlphaUp = true
	}
	if nubeAlpha > 1 {
		nubeAlphaUp = false
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

	// dibujar nube
	op = &ebiten.DrawImageOptions{}
	op.GeoM.Translate(nubeX, nubeY)
	op.ColorM.Scale(1, 3, 2, nubeAlpha)
	op.GeoM.Scale(.4, .4)
	screen.DrawImage(imgNube, op)

	// dibujar texto
	lifes := fmt.Sprintf("Vidas:%02d", vidas)
	text.Draw(screen, lifes, arcadeFont, fontSize, 40, color.Black)
}

// Layout manja las dimensiones de pantalla
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
