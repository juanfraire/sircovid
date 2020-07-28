package main

import (
	"log"
	"os"
	"time"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/audio"
	"github.com/hajimehoshi/ebiten/audio/vorbis"
	"github.com/hajimehoshi/ebiten/audio/wav"
	raudio "github.com/hajimehoshi/ebiten/examples/resources/audio"
	"github.com/hajimehoshi/ebiten/inpututil"
)

var (
	vol  = float64(1)
	temp = float64(vol)
	up   bool
	down bool
	// sonido
	audioContext *audio.Context
	deadSound    *audio.Player
	deadSound2   *audio.Player
	sonidoFondo  *audio.InfiniteLoop
	fondo        *audio.Player
	sonidoIntro  *audio.InfiniteLoop
	sIntro       *audio.Player
)

// Inicio valores de sonido del juego
func initSonido() {

	audioContext, _ = audio.NewContext(44100)
	// sonido fondo
	s, err := os.Open(`sircovid\data\audio\SIR-COVID sin moneditas (1).wav`)
	if err != nil {
		panic(err)
	}
	defer s.Close()
	data := make([]byte, 11491248)
	c, err := s.Read(data)
	fondoD, err := wav.Decode(audioContext, audio.BytesReadSeekCloser(data))
	if err != nil {
		log.Fatal(err)
	}
	sonidoFondo = audio.NewInfiniteLoop(fondoD, int64(c))
	if err != nil {
		log.Fatal(err)
	}
	fondo, err = audio.NewPlayer(audioContext, sonidoFondo)
	if err != nil {
		log.Fatal(err)
	}
	// sonido intro
	// sInt, err := os.Open(`sircovid\data\audio\introconteclas.wav`)
	// if err != nil {
	// 	panic(err)
	// }
	// defer sInt.Close()
	// dataInt := make([]byte, 11491250)
	// cInt, err := s.Read(dataInt)
	// introD, err := wav.Decode(audioContext, audio.BytesReadSeekCloser(dataInt))
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// sonidoIntro = audio.NewInfiniteLoop(introD, int64(cInt))
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// sIntro, err = audio.NewPlayer(audioContext, sonidoIntro)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	jumpD, err := vorbis.Decode(audioContext, audio.BytesReadSeekCloser(raudio.Jump_ogg))
	if err != nil {
		log.Fatal(err)
	}
	deadSound, err = audio.NewPlayer(audioContext, jumpD)
	if err != nil {
		log.Fatal(err)
	}
	jabD, err := wav.Decode(audioContext, audio.BytesReadSeekCloser(raudio.Jab_wav))
	if err != nil {
		log.Fatal(err)
	}
	deadSound2, err = audio.NewPlayer(audioContext, jabD)
	if err != nil {
		log.Fatal(err)
	}
}

func sonido() {

	// volumen on/off
	if inpututil.IsKeyJustPressed(ebiten.KeyX) {
		switch {
		case vol != 0:
			vol = 0
		case vol == 0:
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
	case vol > .01 && down:
		vol -= .01
	}

	fondo.SetVolume(vol)
	deadSound.SetVolume(vol)
	deadSound2.SetVolume(vol * .4)

	// sonido ModePause
	if ModePause {
		fondo.Pause()
	}

}
func sonidoGame() {
	// deadSound.Rewind()
	deadSound2.Rewind()
	fondo.Play()

}

func sonidoGameover() {
	fondo.Pause()
	fondo.Rewind()
	time.Sleep(time.Millisecond * 100)
	deadSound2.Play()

}

func sonidoVidas() {
	deadSound.Play()
	deadSound.Rewind()
}
