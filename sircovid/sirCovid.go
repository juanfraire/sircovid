package main

import (
	"fmt"
	"image"
	_ "image/png"
	"log"
	"time"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/inpututil"
)

const (
	screenWidth  = 999
	screenHeight = 687
	tileSize     = 16
	tileXNum     = 48
)

var (
	tilesImage  *ebiten.Image
	smokeImg    *ebiten.Image
	viejoImg    *ebiten.Image
	frameOX     = 0
	frameOY     = 96
	frameWidth  = 32
	frameHeight = 48
	frameNum    = 1
	x           = float64(25)
	y           = float64(375)
	x1          = x
	y1          = y
	nCount      = float64(screenWidth + 1200)
	smokeX      = float64(screenWidth + 1200)
	nube        = nCount / 2.2012
	smokeY      = float64(800)
	alpha       float64
	tmp         float64
	movx        int
	movy        int
	err         error
	vidas       = 3
	vidtime     int
)

func init() {

	tilesImage, _, err = ebitenutil.NewImageFromFile("data/city2.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
	viejoImg, _, err = ebitenutil.NewImageFromFile("data/viejo.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
	smokeImg, _, err = ebitenutil.NewImageFromFile("data/smoke.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
}

type Game struct {
	layers [][]int
	count  int
}

func (g *Game) Update(screen *ebiten.Image) error {

	g.count++
	nCount--
	nube = nCount / 2.57
	if nCount == 0 {
		nCount = smokeX
	}
	covid()
	mover()
	translado()
	restricciones()
	go vida()
	return nil
}

func mover() {
	// recorte de imagen segun la direccion de movimiento del viejo

	switch {
	case inpututil.IsKeyJustPressed(ebiten.KeyRight) && movy != 1 && movy != 2:
		frameOY = 96
		frameNum = 3
		movx = 1
	case inpututil.IsKeyJustReleased(ebiten.KeyRight) && movy != 1 && movy != 2:
		frameNum = 1
		movx = 0
	case inpututil.IsKeyJustPressed(ebiten.KeyLeft) && movy != 1 && movy != 2:
		frameOY = 48
		frameNum = 3
		movx = 2
	case inpututil.IsKeyJustReleased(ebiten.KeyLeft) && movy != 1 && movy != 2:
		frameNum = 1
		movx = 0
	case inpututil.IsKeyJustPressed(ebiten.KeyUp) && movx != 1 && movx != 2:
		frameOY = 144
		frameNum = 3
		movy = 1
	case inpututil.IsKeyJustReleased(ebiten.KeyUp) && movx != 1 && movx != 2:
		frameNum = 1
		movy = 0
	case inpututil.IsKeyJustPressed(ebiten.KeyDown) && movx != 1 && movx != 2:
		frameOY = 0
		frameNum = 3
		movy = 2
	case inpututil.IsKeyJustReleased(ebiten.KeyDown) && movx != 1 && movx != 2:
		frameNum = 1
		movy = 0
	}
}

func translado() {
	//traslado del viejo coordenadas en x-y
	x1 = x
	y1 = y
	switch {
	case movx == 1:
		x += 1.3
	case movx == 2:
		x -= 1.3
	case movy == 1:
		y -= 1.3
	case movy == 2:
		y += 1.3
	}
}
func restricciones() {
	switch {
	case y < 300*1.3 && x > 20*1.3 && x < 214*1.3:
		x = x1
		y = y1
		frameNum = 1
	case y < 130*1.3 && x > 214*1.3 && x < 768*1.3:
		x = x1
		y = y1
		frameNum = 1
	case y < 270*1.3 && x > 240*1.3 && x < 610*1.3:
		x = x1
		y = y1
		frameNum = 1
	case y < 270*1.3 && x > 675*1.3 && x < 768*1.3:
		x = x1
		y = y1
		frameNum = 1
	case y > 335*1.3 && y < 528*1.3 && x > 40*1.3 && x < 350*1.3:
		x = x1
		y = y1
		frameNum = 1
	case y > 310*1.3 && y < 450*1.3 && x > 390*1.3 && x < 630*1.3:
		x = x1
		y = y1
		frameNum = 1
	}
}

//Estoy trabajando en esta funcion de abajo
func vida() {
	if vidas == vidtime {
		time.Sleep(1 * time.Second)
	}
	if int(x) > int(nube) && int(x) < int(nube+50) && int(y) < 428 && int(y) > 298 {
		vidas = vidas - 1
		vidtime = vidas
		fmt.Println(vidas)
	}

}

func covid() {
	// aumenta y disminuye transparencia de la nube (alpha)
	switch {
	case tmp == 0 && alpha < 1:
		alpha += .003
	case alpha > 1 && tmp != 1:
		tmp = 1
	case tmp == 1 && alpha > 0:
		alpha -= .003
	case alpha <= 0:
		tmp = 0
	}
}

func (g *Game) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(1.3, 1.3)
	screen.DrawImage(tilesImage, op)

	op.GeoM.Scale(.7, .7)
	op.GeoM.Translate(x, y)
	i := (g.count / 7) % frameNum
	sx, sy := frameOX+i*frameWidth, frameOY
	screen.DrawImage(viejoImg.SubImage(image.Rect(sx, sy, sx+frameWidth, sy+frameHeight)).(*ebiten.Image), op)

	op = &ebiten.DrawImageOptions{}
	op.GeoM.Translate(nCount, smokeY)
	op.ColorM.Scale(1, 3, 2, alpha)
	op.GeoM.Scale(.4, .4)
	screen.DrawImage(smokeImg, op)

	ebitenutil.DebugPrint(screen, fmt.Sprintf("TPS: %0.2f", ebiten.CurrentTPS()))
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return int(screenWidth), int(screenHeight)
}

func main() {

	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Tiles (Ebiten Demo)")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
