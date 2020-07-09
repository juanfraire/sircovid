package main

import (
	"math/rand"
	"time"
)

type enemigos struct {
	humanos
	moverHumanos (humanos)
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

var (
	num    = rand.Intn(5)
	tmp    int
	cambio = rand.Intn(50) + 50
	count  int
	ok     bool
)

func moverHumanos(h humanos) humanos {
	// var rand int
	h.FrameNum = 4
	count++

	x1, y1 := h.X, h.Y
	_, _, ok = obstaculos(h.X, h.Y, x1, y1)
	if count == cambio {
		for tmp = num; tmp == num || tmp == 0; tmp = rand.Intn(5) {
		}
		cambio += rand.Intn(100) + 50
		num = tmp

	} else if ok {
		switch num {
		case 1:
			num = 2
		case 2:
			num = 1
		case 3:
			num = 4
		case 4:
			num = 3
		}
	}

	switch num {
	case 1:
		h.FrameOY = 48
		h.X--
	case 2:
		h.FrameOY = 96
		h.X++
	case 3:
		h.FrameOY = 144
		h.Y--
	case 4:
		h.FrameOY = 0
		h.Y++
	}
	return h
}
