package main

import (
	"math/rand"
	"time"
)

//Game es la estructura del juego
type Game struct {
	count int
	nube
	moverNube      (nube)
	siguienteNivel (humanos)
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
	n.nubeX, n.nubeY, _ = obstaculos(n.nubeX, n.nubeY, nubeX1, nubeY1)

	// actualizar alpha
	if n.nubeAlphaUp {
		n.nubeAlpha += .009
	} else {
		n.nubeAlpha -= .003
	}
	return n
}
func siguienteNivel(p humanos) humanos {
	if p.X >= 748 && p.Y > 450 {
		ModeTitle = 2
		ModeGame = 1
		enemigos1.humanos = mujer
		a, b, c, d = 0, 0, 0, 0
		fondo.Pause()
		fondo.Rewind()
	}
	return p
}