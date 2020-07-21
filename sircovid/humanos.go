package main

import "github.com/hajimehoshi/ebiten"

type humanos struct {
	FrameOX         int
	FrameOY         int
	FrameNum        int
	X               float64
	Y               float64
	posicionInicial int
	FrameWidth      int
	FrameHeight     int
	img             *ebiten.Image
	num             int
	cambio          int
}
