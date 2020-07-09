package main

func obstaculos(X float64, Y float64, X1 float64, Y1 float64) (float64, float64, bool) {

	objetos := make([][]int, 16)
	objetos[0] = []int{578, 189, 2, 148}
	objetos[1] = []int{386, 191, 0, 79}
	objetos[2] = []int{193, 159, -1, 112}
	objetos[3] = []int{211, 43, 109, 34}
	objetos[4] = []int{83, 106, 1, 124}
	objetos[5] = []int{1, 77, 1, 301}
	objetos[6] = []int{83, 188, 416, 113}
	objetos[7] = []int{458, 129, 337, 106}
	objetos[8] = []int{705, 61, 352, 108}
	objetos[9] = []int{723, 45, 149, 107}
	objetos[10] = []int{306, 268, 162, 111}
	objetos[11] = []int{306, 189, 148, 13}
	//abajo para que no salga de la pantalla
	objetos[12] = []int{0, 0, 0, screenHeight}
	objetos[13] = []int{0, screenHeight, 0, 0}
	objetos[14] = []int{0, screenWidth, screenHeight, 0}
	objetos[15] = []int{screenWidth, 0, 0, screenHeight}

	for i := 0; i < len(objetos); i++ {
		if (int(X) > objetos[i][0] && int(X) < objetos[i][0]+objetos[i][1]) && int(Y) > objetos[i][2] && int(Y) < objetos[i][2]+objetos[i][3] {
			X = X1
			Y = Y1
			return X, Y, true
		}
	}
	return X, Y, false
}
