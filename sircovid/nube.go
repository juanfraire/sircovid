package main

import (
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten"
)

type nube struct {
	nubeX       float64
	nubeY       float64
	nubeAlpha   float64
	nubeAlphaUp bool
	img         *ebiten.Image
}

//// nubeCovid aumenta y disminuye transparencia de la nube (alpha)
func moverNube(n nube) nube {
	// creacion de nuevas nubes
	if n.nubeAlpha <= 0 {
		n.nubeX = float64(rand.Intn(1500))
		n.nubeY = float64(rand.Intn(1500))
		n.nubeAlphaUp = true
	} else if n.nubeAlpha > 1 {
		time.Sleep(10000 * time.Microsecond)
		n.nubeAlphaUp = false
	}

	nubeX1 := n.nubeX
	nubeY1 := n.nubeY
	// movimiento nube
	if n.nubeAlpha >= 0 {
		n.nubeX--
	}
	n.nubeX, n.nubeY, ok = obstaculos(n.nubeX+70, n.nubeY+70, nubeX1, nubeY1)
	if !ok {
		n.nubeX -= 70
		n.nubeY -= 70
	}
	// actualizar alpha
	if n.nubeAlphaUp {
		n.nubeAlpha += .009
	} else {
		n.nubeAlpha -= .003
	}
	return n
}
func dibujarNube(n nube, screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(n.nubeX, nube1.nubeY+384)
	op.ColorM.Scale(1, 3, 2, n.nubeAlpha)
	op.GeoM.Scale(.4, .4)
	screen.DrawImage(n.img, op)
}
