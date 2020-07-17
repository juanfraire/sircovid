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
	} else if n.Alpha > 1 {
		time.Sleep(10000 * time.Microsecond)
		n.AlphaUp = false
	}

	// nubeX1 := n.X
	// nubeY1 := n.Y
	// movimiento nube
	if n.Alpha >= 0 {
		n.X--
	}
	// n.X, n.Y, ok = obstaculos(n.X+70, n.Y+70, nubeX1, nubeY1)
	// if !ok {
	// 	n.X -= 70
	// 	n.Y -= 70
	// }
	// actualizar alpha
	if n.AlphaUp {
		n.Alpha += .009
	} else {
		n.Alpha -= .003
	}
	return n
}
func dibujarNube(n nube, screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(n.X, n.Y+384)
	op.ColorM.Scale(1, 3, 2, n.Alpha)
	op.GeoM.Scale(.4, .4)
	screen.DrawImage(n.img, op)
}
