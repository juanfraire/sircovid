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

var vol = float64(1)
var temp = float64(vol)

// Inicio valores de sonido del juego
func initSonido() {

	audioContext, _ = audio.NewContext(44100)
	s, err := os.Open(`sircovid\data\SIR-COVID.wav`)
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
	if inpututil.IsKeyJustPressed(ebiten.KeyKPAdd) && vol < .9 {
		vol += .1
		temp = vol
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyKPSubtract) && vol > .1 {
		vol -= .1
		temp = vol
	}

	fondo.SetVolume(vol)
	deadSound.SetVolume(vol)
	deadSound2.SetVolume(vol * .4)

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
	// deadSound2.SetVolume(.4)
	deadSound2.Play()

}

func sonidoVidas() {
	deadSound.Play()
	deadSound.Rewind()
}
