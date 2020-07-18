package main

import (
	"log"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

type nube struct {
	X       float64
	Y       float64
	Alpha   float64
	AlphaUp bool
	img     *ebiten.Image
}

var nube1 nube

func initNube() {
	nube1.X = float64(rand.Intn(screenWidth / .4))
	nube1.Y = float64(rand.Intn(screenHeight / .4))

	nube1.img, _, err = ebitenutil.NewImageFromFile(`sircovid\data\smoke.png`, ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
}

//// nubeCovid aumenta y disminuye transparencia de la nube (alpha)
func moverNube(n nube) nube {
	// creacion de nuevas nubes
	if n.Alpha <= 0 {
		n.X = float64(rand.Intn(screenWidth / .4))
		n.Y = float64(rand.Intn(screenHeight / .4))
		n.AlphaUp = true
	} else if n.Alpha > .8 {
		time.Sleep(10000 * time.Microsecond)
		n.AlphaUp = false
	}

	// movimiento nube
	if n.Alpha >= 0 {
		n.X--
	}

	if n.AlphaUp {
		n.Alpha += .009
	} else {
		n.Alpha -= .003
	}
	return n
}

var scale = float64(.4)

func dibujarNube(n nube, screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(n.X, n.Y)
	op.ColorM.Scale(3, 2, 0, n.Alpha)
	op.GeoM.Scale(scale, scale)
	screen.DrawImage(n.img, op)
}
