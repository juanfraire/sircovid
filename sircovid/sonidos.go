package main

import (
	"log"
	"os"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/audio"
	"github.com/hajimehoshi/ebiten/audio/mp3"
	"github.com/hajimehoshi/ebiten/inpututil"
)

var (
	vol  float64
	temp float64
	up   bool
	down bool
	fade bool
	tmpS = 1
	// sonido
	audioContext *audio.Context
	audiomp3     *audio.Context
	deadSound    *audio.Player
	deadSound2   *audio.Player
	sonidoFondo  *audio.InfiniteLoop
	fondo        *audio.Player
	sonidoIntro  *audio.InfiniteLoop
	sIntro       *audio.Player
	sPuerta      *audio.Player
	sDinero      *audio.Player
	sNube        *audio.Player
	sFast        *audio.Player
	sBarbijo     *audio.Player
	sLevelUp     *audio.Player
)

// Inicio valores de sonido del juego al final

// sonidos
func sonido(p player) {

	switch {
	case ModePause:
		fondo.Pause()

	case Level > tmpS:
		fondo.Pause()
		fondo.Rewind()
		sLevelUp.Play()
		sLevelUp.Rewind()
		tmpS = Level

	case ModeTitle:
		fadeIn()
		fondo.Pause()
		sIntro.Play()

	case ModeGame:
		sIntro.Pause()
		sIntro.Rewind()
		deadSound2.Rewind()
		fondo.Play()

	case ModeGameOver:
		fondo.Pause()
		fondo.Rewind()
		deadSound.Pause()
		deadSound2.Play()
		sIntro.Play()
	}

	// volumen on/off
	if inpututil.IsKeyJustPressed(ebiten.KeyX) {
		fade = false
		switch {
		case vol != .001:
			vol = .001
		case vol == .001:
			vol = temp
		}
	}
	// volumen +/-
	if inpututil.IsKeyJustPressed(ebiten.KeyKPAdd) || inpututil.IsKeyJustPressed(ebiten.Key9) {
		up = true
	} else if inpututil.IsKeyJustReleased(ebiten.KeyKPAdd) || inpututil.IsKeyJustReleased(ebiten.Key9) {
		up = false
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyKPSubtract) || inpututil.IsKeyJustPressed(ebiten.Key8) {
		down = true
	} else if inpututil.IsKeyJustReleased(ebiten.KeyKPSubtract) || inpututil.IsKeyJustReleased(ebiten.Key8) {
		down = false
	}
	switch {
	case vol < .99 && up:
		vol += .01
		temp = vol
	case vol > .01 && down:
		vol -= .01
		temp = vol
	}
	// set volume
	fondo.SetVolume(vol)
	deadSound.SetVolume(vol)
	deadSound2.SetVolume(vol)
	sBarbijo.SetVolume(vol)
	sDinero.SetVolume(vol)
	sFast.SetVolume(vol)
	sIntro.SetVolume(vol)
	sLevelUp.SetVolume(vol)
	sPuerta.SetVolume(vol)
}

func sonidoVidas() {
	deadSound.Play()
	deadSound.Rewind()
}
func sonidomonedas() {
	sPuerta.Pause()
	sDinero.Play()
	sDinero.Rewind()
}

func fadeIn() {
	if vol == 0 {
		fade = true
	} else if vol > .99 {
		fade = false
	}
	if fade {
		vol += .01
	}
}

func initSonido() {

	audioContext, _ = audio.NewContext(32000)

	// sonido fondo
	s, err := os.Open(`sircovid\data\audio\SIR-COVID.mp3`)
	if err != nil {
		panic(err)
	}
	defer s.Close()
	data := make([]byte, 521901)
	c, err := s.Read(data)
	// fmt.Println(c)

	fondoD, err := mp3.Decode(audioContext, audio.BytesReadSeekCloser(data))
	if err != nil {
		log.Fatal(err)
	}
	sonidoFondo = audio.NewInfiniteLoop(fondoD, int64(c)*20)
	if err != nil {
		log.Fatal(err)
	}
	fondo, err = audio.NewPlayer(audioContext, sonidoFondo)
	if err != nil {
		log.Fatal(err)
	}
	// sonido intro
	s, err = os.Open(`sircovid\data\audio\introconteclas.mp3`)
	if err != nil {
		panic(err)
	}
	defer s.Close()
	data = make([]byte, 371565)
	c, err = s.Read(data)
	// fmt.Println(c)

	introD, err := mp3.Decode(audioContext, audio.BytesReadSeekCloser(data))
	if err != nil {
		log.Fatal(err)
	}
	sonidoFondo = audio.NewInfiniteLoop(introD, int64(c)*20)
	if err != nil {
		log.Fatal(err)
	}
	sIntro, err = audio.NewPlayer(audioContext, sonidoFondo)
	if err != nil {
		log.Fatal(err)
	}

	//sonido Puerta
	s, err = os.Open(`sircovid\data\audio\puertas ingresos.mp3`)
	if err != nil {
		panic(err)
	}
	defer s.Close()
	data = make([]byte, 4653)
	c, err = s.Read(data)
	// fmt.Println("puerta", c)

	puertaD, err := mp3.Decode(audioContext, audio.BytesReadSeekCloser(data))
	if err != nil {
		log.Fatal(err)
	}
	sPuerta, err = audio.NewPlayer(audioContext, puertaD)
	if err != nil {
		log.Fatal(err)
	}

	//sonido Monedas
	s, err = os.Open(`sircovid\data\audio\DINERO.mp3`)
	if err != nil {
		panic(err)
	}
	defer s.Close()
	data = make([]byte, 9549)
	c, err = s.Read(data)
	// fmt.Println("dinero", c)

	dineroD, err := mp3.Decode(audioContext, audio.BytesReadSeekCloser(data))
	if err != nil {
		log.Fatal(err)
	}
	sDinero, err = audio.NewPlayer(audioContext, dineroD)
	if err != nil {
		log.Fatal(err)
	}
	// sonido Fast
	s, err = os.Open(`sircovid\data\audio\ALRIGHT! COFFE.mp3`)
	if err != nil {
		panic(err)
	}
	defer s.Close()
	data = make([]byte, 7533)
	c, err = s.Read(data)
	// fmt.Println("cofee", c)

	fastD, err := mp3.Decode(audioContext, audio.BytesReadSeekCloser(data))
	if err != nil {
		log.Fatal(err)
	}
	sFast, err = audio.NewPlayer(audioContext, fastD)
	if err != nil {
		log.Fatal(err)
	}

	// sonido barbijo o alcohol

	s, err = os.Open(`sircovid\data\audio\ponerse barbijo.mp3`)
	if err != nil {
		panic(err)
	}
	defer s.Close()
	data = make([]byte, 4653)
	c, err = s.Read(data)
	// fmt.Println("barbijo", c)

	barbijoD, err := mp3.Decode(audioContext, audio.BytesReadSeekCloser(data))
	if err != nil {
		log.Fatal(err)
	}
	sBarbijo, err = audio.NewPlayer(audioContext, barbijoD)
	if err != nil {
		log.Fatal(err)
	}

	//sonido Pasar Nivel

	s, err = os.Open(`sircovid\data\audio\PASAR DE NIVEL.mp3`)
	if err != nil {
		panic(err)
	}
	defer s.Close()
	data = make([]byte, 9549)
	c, err = s.Read(data)
	// fmt.Println("pasar nivel", c)

	levelD, err := mp3.Decode(audioContext, audio.BytesReadSeekCloser(data))
	if err != nil {
		log.Fatal(err)
	}
	sLevelUp, err = audio.NewPlayer(audioContext, levelD)
	if err != nil {
		log.Fatal(err)
	}

	// sonido perder vida
	s, err = os.Open(`sircovid\data\audio\tos1.mp3`)
	if err != nil {
		panic(err)
	}
	defer s.Close()
	data = make([]byte, 16461)
	c, err = s.Read(data)
	// fmt.Println("tos", c)

	tosD, err := mp3.Decode(audioContext, audio.BytesReadSeekCloser(data))
	if err != nil {
		log.Fatal(err)
	}
	deadSound, err = audio.NewPlayer(audioContext, tosD)
	if err != nil {
		log.Fatal(err)
	}

	// sonido muerte
	s, err = os.Open(`sircovid\data\audio\sonido muerte o daño por nube.mp3`)
	if err != nil {
		panic(err)
	}
	defer s.Close()
	data = make([]byte, 16461)
	_, err = s.Read(data)
	// fmt.Println("nube", c)
	jabD, err := mp3.Decode(audioContext, audio.BytesReadSeekCloser(data))
	if err != nil {
		log.Fatal(err)
	}
	deadSound2, err = audio.NewPlayer(audioContext, jabD)
	if err != nil {
		log.Fatal(err)
	}

}
