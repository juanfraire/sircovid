package main

import (
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

//bank ajuste interiores
func interior() {

	imgBanco, _, err = ebitenutil.NewImageFromFile(`sircovid\data\banco.png`, ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
	imgCintas, _, err = ebitenutil.NewImageFromFile(`sircovid\data\cintas.png`, ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
	initEnemigos()
	initNube()
	inicio = false
	// nivel = 6
	eScaleW = .6
	eScaleH = .45
	plyrScale = .75
}
func salida() {
	if banco {
		banco = false
		eScaleW = .5
		eScaleH = .4
		plyrScale = .6
		initNube()
		initEnemigos()
	}
}
