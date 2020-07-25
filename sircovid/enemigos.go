package main

import (
	"image"
	"log"
	"math/rand"
	"strconv"
	"time"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

var (
	enemigo humanos
	count   int
	tmp     int
	ok      bool
	x       float64
	y       float64
	en      string
)

func randXY() (x float64, y float64) {
	rand.Seed(time.Now().UnixNano())
	_, _, obs := obstaculos(x, y, x, y)
	for obs {
		x = float64(rand.Intn(screenWidth))
		y = float64(rand.Intn(screenHeight))
		_, _, obs = obstaculos(x, y, x, y)
	}
	return
}

func initEnemigos() {
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < nivel; i++ {
		en = `sircovid\data\enemigo` + strconv.Itoa(i+1) + `.png`
		enemigo.FrameOX[i] = 48
		enemigo.FrameOY[i] = 72 * rand.Intn(4)
		enemigo.FrameNum[i] = 1
		enemigo.X[i], enemigo.Y[i] = randXY()
		enemigo.FrameWidth[i] = 48
		enemigo.FrameHeight[i] = 72
		enemigo.num[i] = rand.Intn(5)
		enemigo.cambio[i] = rand.Intn(50) + 100
		enemigo.img[i], _, err = ebitenutil.NewImageFromFile(en, ebiten.FilterDefault)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func moverHumanos(E humanos) humanos {

	if ModeGame >= 0 {
		count++
	}
	for i := 0; i < nivel; i++ {

		E.FrameNum[i] = 3
		E.FrameOX[i] = 0

		switch E.num[i] {
		case 0:
			E.FrameNum[i] = 1
			E.FrameOX[i] = 48
		case 1:
			E.FrameOY[i] = 72
			E.X[i]--
		case 2:
			E.FrameOY[i] = 144
			E.X[i]++
		case 3:
			E.FrameOY[i] = 216
			E.Y[i]--
		case 4:
			E.FrameOY[i] = 0
			E.Y[i]++
		}

		if count >= E.cambio[i] {
			for tmp = E.num[i]; tmp == E.num[i]; tmp = rand.Intn(5) {
			}
			E.cambio[i] += rand.Intn(200) + 200
			E.num[i] = tmp
		}

		x1, y1 := E.X[i], E.Y[i]
		_, _, ok = obstaculos(E.X[i], E.Y[i], x1, y1)

		if ok {
			switch E.num[i] {
			case 1:
				E.num[i] = 2
			case 2:
				E.num[i] = 1
			case 3:
				E.num[i] = 4
			case 4:
				E.num[i] = 3
			}
			E.cambio[i] = count + 5
		}
	}
	return E
}

func dibujarEnemigos(E humanos, screen *ebiten.Image) {
	for i := 0; i < nivel; i++ {
		if ModePause || ModeWin {
			E.FrameNum[i] = 1
			E.FrameOX[i] = 0
		}
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Scale(.65, .51)
		op.GeoM.Translate(E.X[i], E.Y[i])
		j := (count1 / 7) % E.FrameNum[i]
		hx, hy := E.FrameOX[i]+j*E.FrameWidth[i], E.FrameOY[i]
		screen.DrawImage(E.img[i].SubImage(image.Rect(hx, hy, hx+E.FrameWidth[i], hy+E.FrameHeight[i])).(*ebiten.Image), op)
	}
}
