package main

import "github.com/hajimehoshi/ebiten"

type humanos struct {
	FrameOX         [20]int
	FrameOY         [20]int
	FrameNum        [20]int
	X               [20]float64
	Y               [20]float64
	posicionInicial [20]int
	FrameWidth      [20]int
	FrameHeight     [20]int
	img             [20]*ebiten.Image
	num             [20]int
	cambio          [20]int
	// contagio        [20]bool
}
