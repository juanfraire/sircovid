package main

//Game es la estructura del juego
type Game struct {
	count int
	nube
	numPlayers     int
	electPlayer    int
	siguienteNivel (player)
}

//// nubeCovid aumenta y disminuye transparencia de la nube (alpha)

func siguienteNivel(p player) player {

	if p.X >= 746 && p.Y > 450 {

		pasarNivel()

		fondo.Pause()
		fondo.Rewind()
	}
	return p
}
