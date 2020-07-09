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

// func moverHumanos(h humanos) humanos {
// 	h.FrameNum = 4
// 	switch {
// 	case a != 1:
// 		h.FrameOY = 48
// 		h.Y = 290
// 		h.X--
// 		if h.X < 228 {
// 			a = 1
// 		}
// 	case a == 1 && a1 != 1:
// 		h.FrameOY = 144
// 		h.Y--
// 		if h.Y == 137 {
// 			a1 = 1
// 		}
// 	case a1 == 1 && a2 != 1:
// 		h.FrameOY = 0
// 		h.Y++
// 		if h.Y == 310 {
// 			a2 = 1
// 		}
// 	case a2 == 1 && a3 != 1:
// 		h.FrameOY = 48
// 		h.X--
// 		if h.X == -100 {
// 			a3 = 1
// 		}
// 	case a3 == 1 && a4 != 1:
// 		h.FrameOY = 96
// 		h.Y = 460
// 		h.X++
// 		if h.X == 20 {
// 			a4 = 1
// 		}
// 	case a4 == 1 && a5 != 1:
// 		h.FrameOY = 144
// 		h.Y--
// 		if h.Y == 310 {
// 			a5 = 1
// 		}
// 	case a5 == 1 && a6 != 1:
// 		h.FrameOY = 96
// 		h.X++
// 		if h.X == 228 {
// 			a6 = 1
// 		}
// 	case a6 == 1 && a7 != 1:
// 		h.FrameOY = 144
// 		h.Y--
// 		if h.Y == 280 {
// 			a7 = 1
// 		}
// 	case a7 == 1 && a8 != 1:
// 		h.FrameOY = 96
// 		h.X++
// 		if h.X == 370 {
// 			a8 = 1
// 		}
// 	case a8 == 1 && a9 != 1:
// 		h.FrameOY = 0
// 		h.Y++
// 		if h.Y == 470 {
// 			a9 = 1
// 		}
// 	case a9 == 1:
// 		h.FrameOY = 96
// 		h.X++
// 		if h.X == 800 {
// 			a, a1, a2, a3, a4, a5, a6, a7, a8, a9 = 0, 0, 0, 0, 0, 0, 0, 0, 0, 0
// 		}
// 	}
// 	return h

// }
