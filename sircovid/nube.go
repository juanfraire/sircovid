package main

import (
	"log"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

type nube struct {
	X       [20]float64
	Y       [20]float64
	Alpha   [20]float64
	AlphaUp [20]bool
	img     *ebiten.Image
}

var nube1 nube

func initNube() {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < nivel*2; i++ {
		nube1.X[i] = float64(rand.Intn(screenWidth / .4))
		nube1.Y[i] = float64(rand.Intn(screenHeight / .4))
		nube1.Alpha[i] = float64(rand.Intn(11)) * .1
	}
	nube1.img, _, err = ebitenutil.NewImageFromFile(`sircovid\data\smoke.png`, ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
}

//// nubeCovid aumenta y disminuye transparencia de la nube (alpha)
func moverNube(n nube) nube {
	rand.Seed(time.Now().UnixNano())
	// creacion de nuevas nubes
	for i := 0; i < nivel*2; i++ {
		if n.Alpha[i] >= 0 {
			n.X[i]--
		}

		if n.Alpha[i] <= 0 {
			n.X[i] = float64(rand.Intn(screenWidth / .4))
			n.Y[i] = float64(rand.Intn(screenHeight / .4))
			n.AlphaUp[i] = true
		} else if n.Alpha[i] > .9 {
			n.AlphaUp[i] = false
		}

		// movimiento nube

		if n.AlphaUp[i] {
			n.Alpha[i] += .009
		} else {
			n.Alpha[i] -= .003
		}
	}
	return n
}

var scale = float64(.4)

func dibujarNube(n nube, screen *ebiten.Image) {
	for i := 0; i < nivel*2; i++ {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(n.X[i], n.Y[i])
		op.ColorM.Scale(3, 2, 0, n.Alpha[i])
		op.GeoM.Scale(scale, scale)
		screen.DrawImage(n.img, op)
	}
}
