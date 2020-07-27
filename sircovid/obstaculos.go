package main

func obstaculos(X float64, Y float64, X1 float64, Y1 float64) (float64, float64, bool) {

	objetos := make([][]int, 55)
	objetos[0] = []int{304, 192, 145, 125}
	objetos[1] = []int{496, 79, 162, 108}
	objetos[2] = []int{480, 66, 383, 17}
	objetos[3] = []int{496, 33, 366, 58}
	objetos[4] = []int{507, 12, 363, 8}
	objetos[5] = []int{488, 50, 374, 41}
	objetos[6] = []int{607, 17, 434, 25}
	objetos[7] = []int{403, 18, 438, 23}
	objetos[8] = []int{594, 15, 259, 16}
	objetos[9] = []int{716, 17, 265, 20}
	objetos[10] = []int{689, 14, 173, 13}
	objetos[11] = []int{257, 15, 265, 16}
	objetos[12] = []int{273, 15, 198, 48}
	objetos[13] = []int{178, 14, 198, 51}
	objetos[14] = []int{145, 15, 162, 138}
	objetos[15] = []int{113, 15, 161, 139}
	objetos[16] = []int{85, 42, 288, 12}
	objetos[17] = []int{145, 48, 288, 13}
	objetos[18] = []int{403, 13, 337, 18}
	objetos[19] = []int{609, 16, 339, 15}
	objetos[20] = []int{704, 192, 354, 110}
	objetos[21] = []int{915, 12, 374, 25}
	objetos[22] = []int{915, 12, 423, 22}
	objetos[23] = []int{578, 31, 66, 84}
	objetos[24] = []int{688, 32, 63, 86}
	objetos[25] = []int{668, 17, 56, 59}
	objetos[26] = []int{615, 16, 56, 58}
	objetos[27] = []int{382, 16, 62, 20}
	objetos[28] = []int{193, 158, 2, 106}
	objetos[29] = []int{211, 45, 100, 33}
	objetos[30] = []int{77, 114, 1, 124}
	objetos[31] = []int{1, 68, 3, 283}
	objetos[32] = []int{1, 63, 280, 22}
	objetos[33] = []int{450, 12, 391, 26}
	objetos[34] = []int{581, 11, 392, 23}
	objetos[35] = []int{454, 138, 434, 12}
	objetos[36] = []int{455, 134, 338, 13}
	objetos[37] = []int{384, 79, 0, 63}
	objetos[38] = []int{387, 71, 65, 7}
	objetos[39] = []int{410, 29, 59, 24}
	objetos[40] = []int{465, 494, 0, 77}
	objetos[41] = []int{955, 18, 1, 60}
	objetos[42] = []int{957, 10, 61, 11}
	objetos[43] = []int{998, 41, 17, 13}
	objetos[44] = []int{724, 283, 161, 94}
	objetos[45] = []int{1029, 10, 197, 26}
	objetos[46] = []int{964, 28, 339, 23}
	objetos[47] = []int{1027, 28, 340, 21}
	objetos[48] = []int{81, 191, 417, 111}
	objetos[49] = []int{924, 72, 149, 16}
	objetos[50] = []int{35, 42, 3, 271}
	//abajo para que no salga de la pantalla
	objetos[51] = []int{0, 0, 0, screenHeight}
	objetos[52] = []int{0, screenHeight, 0, 0}
	objetos[53] = []int{0, screenWidth, screenHeight, 32}
	objetos[54] = []int{screenWidth, 0, 0, screenHeight}

	for i := 0; i < len(objetos); i++ {
		if int(X)+17 >= objetos[i][0] && int(X)+2 <= objetos[i][0]+objetos[i][1] && int(Y)+30 >= objetos[i][2] && int(Y)+32 <= objetos[i][2]+objetos[i][3] {
			X = X1
			Y = Y1
			return X, Y, true
		}
	}
	return X, Y, false
}
