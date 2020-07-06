package main

type enemigos struct {
	humanos
	moverHumanos (humanos)
}

func moverHumanos(h humanos) humanos {
	h.FrameNum = 4
	switch {
	case a != 1:
		h.FrameOY = 48
		h.Y = 290
		h.X--
		if h.X < 228 {
			a = 1
		}
	case a == 1 && a1 != 1:
		h.FrameOY = 144
		h.Y--
		if h.Y == 137 {
			a1 = 1
		}
	case a1 == 1 && a2 != 1:
		h.FrameOY = 0
		h.Y++
		if h.Y == 310 {
			a2 = 1
		}
	case a2 == 1 && a3 != 1:
		h.FrameOY = 48
		h.X--
		if h.X == -100 {
			a3 = 1
		}
	case a3 == 1 && a4 != 1:
		h.FrameOY = 96
		h.Y = 460
		h.X++
		if h.X == 20 {
			a4 = 1
		}
	case a4 == 1 && a5 != 1:
		h.FrameOY = 144
		h.Y--
		if h.Y == 310 {
			a5 = 1
		}
	case a5 == 1 && a6 != 1:
		h.FrameOY = 96
		h.X++
		if h.X == 228 {
			a6 = 1
		}
	case a6 == 1 && a7 != 1:
		h.FrameOY = 144
		h.Y--
		if h.Y == 280 {
			a7 = 1
		}
	case a7 == 1 && a8 != 1:
		h.FrameOY = 96
		h.X++
		if h.X == 370 {
			a8 = 1
		}
	case a8 == 1 && a9 != 1:
		h.FrameOY = 0
		h.Y++
		if h.Y == 470 {
			a9 = 1
		}
	case a9 == 1:
		h.FrameOY = 96
		h.X++
		if h.X == 800 {
			a, a1, a2, a3, a4, a5, a6, a7, a8, a9 = 0, 0, 0, 0, 0, 0, 0, 0, 0, 0
		}
	}
	return h

}
