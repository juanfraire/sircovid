package main

import "github.com/hajimehoshi/ebiten"

type nube struct {
	nubeX       float64
	nubeY       float64
	nubeAlpha   float64
	nubeAlphaUp bool
	img         *ebiten.Image
}
