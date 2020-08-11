package main

func obstaculos(X float64, Y float64, X1 float64, Y1 float64) (float64, float64, bool) {
	var objetos [][]int
	switch {
	case !banco:
		objetos = make([][]int, 60)
		objetos[0] = []int{237, 79, 125, 99}
		objetos[1] = []int{235, 5, 310, 33}
		objetos[2] = []int{147, 124, 26, 68}
		objetos[3] = []int{55, 89, 30, 80}
		objetos[4] = []int{0, 59, 0, 245}
		objetos[5] = []int{68, 29, 129, 116}
		objetos[6] = []int{115, 31, 134, 111}
		objetos[7] = []int{148, 9, 220, 35}
		objetos[8] = []int{15, 44, 329, 82}
		objetos[9] = []int{335, 67, 285, 151}
		objetos[10] = []int{357, 38, 450, 82}
		objetos[11] = []int{443, 70, 317, 69}
		objetos[12] = []int{456, 54, 303, 13}
		objetos[13] = []int{471, 20, 295, 23}
		objetos[14] = []int{520, 100, 292, 237}
		objetos[15] = []int{455, 64, 445, 85}
		objetos[16] = []int{355, 127, 145, 72}
		objetos[17] = []int{482, 122, 133, 80}
		objetos[18] = []int{522, 100, 28, 171}
		objetos[19] = []int{298, 62, 23, 42}
		objetos[20] = []int{316, 23, 61, 11}
		objetos[21] = []int{690, 49, 291, 240}
		objetos[22] = []int{795, 70, 309, 75}
		objetos[23] = []int{874, 101, 293, 237}
		objetos[24] = []int{810, 69, 441, 89}
		objetos[25] = []int{682, 3, 351, 28}
		objetos[26] = []int{241, 11, 353, 66}
		objetos[27] = []int{238, 23, 392, 20}
		objetos[28] = []int{16, 42, 458, 73}
		objetos[29] = []int{31, 14, 447, 11}
		objetos[30] = []int{96, 89, 320, 59}
		objetos[31] = []int{147, 38, 379, 146}
		objetos[32] = []int{92, 93, 459, 69}
		objetos[33] = []int{715, 341, 0, 177}
		objetos[34] = []int{698, 16, 1, 144}
		objetos[35] = []int{705, 11, 146, 33}
		objetos[36] = []int{710, 4, 181, 21}
		objetos[37] = []int{717, 77, 179, 26}
		objetos[38] = []int{800, 11, 200, 10}
		objetos[39] = []int{892, 33, 176, 33}
		objetos[40] = []int{795, 101, 179, 16}
		objetos[41] = []int{362, 162, 23, 35}
		objetos[42] = []int{603, 9, 200, 6}
		objetos[43] = []int{25, 613, 0, 28}
		objetos[44] = []int{638, 30, 0, 15}
		objetos[45] = []int{638, 18, 13, 11}
		objetos[46] = []int{925, 131, 176, 30}
		objetos[47] = []int{1039, 17, 301, 225}
		objetos[48] = []int{1012, 29, 330, 196}
		objetos[49] = []int{1023, 17, 312, 18}
		objetos[50] = []int{809, 40, 294, 13}
		objetos[51] = []int{891, 26, 283, 12}
		objetos[52] = []int{943, 21, 286, 13}
		objetos[53] = []int{699, 15, 283, 11}
		objetos[54] = []int{538, 25, 285, 11}
		objetos[55] = []int{370, 107, 132, 23}
		//abajo para que no salga de la pantalla
		objetos[56] = []int{0, 0, 0, screenHeight}
		objetos[57] = []int{screenWidth, 0, 0, screenHeight}
		objetos[58] = []int{0, screenWidth, screenHeight, 32}
		objetos[59] = []int{0, screenWidth, 30, 0}

	case banco:
		objetos = make([][]int, 35)
		objetos[0] = []int{455, 128, 212, 55}
		objetos[1] = []int{615, 66, 262, 4}
		objetos[2] = []int{555, 94, 212, 22}
		objetos[3] = []int{678, 2, 173, 92}
		objetos[4] = []int{616, 63, 159, 26}
		objetos[5] = []int{454, 115, 160, 59}
		objetos[6] = []int{563, 56, 291, 28}
		objetos[7] = []int{423, 2, 265, 31}
		objetos[8] = []int{359, 65, 264, 3}
		objetos[9] = []int{359, 1, 183, 82}
		objetos[10] = []int{358, 49, 184, 1}
		objetos[11] = []int{406, 3, 158, 26}
		objetos[12] = []int{391, 65, 211, 25}
		objetos[13] = []int{420, 54, 293, 23}
		objetos[14] = []int{614, 2, 265, 30}
		objetos[15] = []int{542, 146, 157, 1}
		//limite
		objetos[16] = []int{0, 1057, -1, 159}
		objetos[17] = []int{704, 353, 157, 398}
		objetos[18] = []int{0, 350, 157, 394}
		objetos[19] = []int{349, 117, 336, 215}
		objetos[20] = []int{464, 182, 401, 150}
		objetos[21] = []int{560, 225, 351, 203}
		objetos[22] = []int{575, 143, 335, 18}
		objetos[23] = []int{463, 17, 352, 61}
		objetos[24] = []int{567, 8, 344, 8}
		objetos[25] = []int{463, 8, 344, 10}
		objetos[26] = []int{673, 30, 164, 10}
		objetos[27] = []int{355, 25, 322, 13}
		objetos[28] = []int{690, 10, 310, 26}
		objetos[29] = []int{673, 28, 324, 9}
		objetos[30] = []int{353, 14, 308, 18}

		objetos[31] = []int{0, 0, 0, screenHeight}
		objetos[32] = []int{screenWidth, 0, 0, screenHeight}
		objetos[33] = []int{0, screenWidth, screenHeight, 32}
		objetos[34] = []int{0, screenWidth, 30, 0}
	}

	for i := 0; i < len(objetos); i++ {
		if int(X+wth) >= objetos[i][0] && int(X) <= objetos[i][0]+objetos[i][1] && int(Y+hgt) >= objetos[i][2] && int(Y+hgt) <= objetos[i][2]+objetos[i][3] {
			X = X1
			Y = Y1
			return X, Y, true
		}
	}

	for j := 0; j < numEnemigo; j++ {
		if i != j && X+wth > enemigo.X[j] && X < enemigo.X[j]+wth && Y+hgt > enemigo.Y[j]+hgt && Y+hgt < enemigo.Y[j]+hgt {
			X = X1
			Y = Y1
			return X, Y, true
		}
	}

	return X, Y, false
}
