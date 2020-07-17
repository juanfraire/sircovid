package main

func obstaculos(X float64, Y float64, X1 float64, Y1 float64) (float64, float64, bool) {

	objetos := make([][]int, 61)
	objetos[0] = []int{0, 81, 0, 302}
	objetos[1] = []int{83, 110, -2, 129}
	objetos[2] = []int{212, 46, 96, 44}
	objetos[3] = []int{193, 160, 1, 108}
	objetos[4] = []int{384, 81, 0, 71}
	objetos[5] = []int{411, 28, 68, 14}
	objetos[6] = []int{467, 108, 1, 75}
	objetos[7] = []int{577, 33, 71, 80}
	objetos[8] = []int{688, 33, 75, 78}
	objetos[9] = []int{691, 11, 276, 9}
	objetos[10] = []int{703, 19, 274, 12}
	objetos[11] = []int{688, 18, 179, 12}
	objetos[12] = []int{595, 14, 273, 15}
	objetos[13] = []int{577, 15, 259, 16}
	objetos[14] = []int{595, 11, 392, 23}
	objetos[15] = []int{437, 10, 392, 24}
	objetos[16] = []int{452, 139, 439, 6}
	objetos[17] = []int{503, 34, 448, 11}
	objetos[18] = []int{400, 17, 453, 11}
	objetos[19] = []int{402, 16, 339, 13}
	objetos[20] = []int{611, 12, 337, 14}
	objetos[21] = []int{610, 17, 451, 14}
	objetos[22] = []int{675, 13, 467, 11}
	objetos[23] = []int{699, 29, 460, 17}
	objetos[24] = []int{707, 58, 354, 107}
	objetos[25] = []int{433, 16, 515, 13}
	objetos[26] = []int{496, 17, 514, 12}
	objetos[27] = []int{82, 190, 418, 107}
	objetos[28] = []int{454, 136, 344, 7}
	objetos[29] = []int{290, 14, 276, 9}
	objetos[30] = []int{255, 18, 274, 9}
	objetos[31] = []int{197, 11, 308, 9}
	objetos[32] = []int{177, 16, 203, 49}
	objetos[33] = []int{271, 19, 205, 47}
	objetos[34] = []int{148, 10, 161, 142}
	objetos[35] = []int{145, 50, 288, 15}
	objetos[36] = []int{114, 12, 160, 144}
	objetos[37] = []int{84, 43, 287, 14}
	objetos[38] = []int{89, 33, 304, 12}
	objetos[39] = []int{101, 12, 163, 13}
	objetos[40] = []int{371, 13, 99, 13}
	objetos[41] = []int{304, 273, 162, 106}
	objetos[42] = []int{721, 47, 17, 235}
	objetos[43] = []int{635, 0, 160, 0}
	objetos[44] = []int{290, 14, 341, 9}
	objetos[45] = []int{577, 141, 11, 65}
	objetos[46] = []int{596, 0, 407, 0}
	objetos[47] = []int{595, 13, 391, 25}
	objetos[48] = []int{437, 11, 393, 23}
	objetos[49] = []int{483, 77, 384, 16}
	objetos[50] = []int{492, 59, 375, 7}
	objetos[51] = []int{498, 45, 366, 8}
	objetos[52] = []int{489, 0, 400, 0}
	objetos[53] = []int{490, 64, 402, 6}
	objetos[54] = []int{495, 53, 410, 7}
	objetos[55] = []int{499, 47, 419, 8}
	objetos[56] = []int{303, 192, 144, 23}
	objetos[57] = []int{387, 0, 167, 15}
	objetos[58] = []int{486, 0, 187, 0}
	objetos[59] = []int{442, 0, 160, 0}
	objetos[60] = []int{436, 0, 156, 0}
	//abajo para que no salga de la pantalla
	objetos[12] = []int{0, 0, 0, screenHeight}
	objetos[13] = []int{0, screenHeight, 0, 0}
	objetos[14] = []int{0, screenWidth, screenHeight, 32}
	objetos[15] = []int{screenWidth, 0, 0, screenHeight}

	for i := 0; i < len(objetos); i++ {
		if int(X)+17 >= objetos[i][0] && int(X)+2 <= objetos[i][0]+objetos[i][1] && int(Y) >= objetos[i][2] && int(Y) <= objetos[i][2]+objetos[i][3] {
			X = X1
			Y = Y1
			return X, Y, true
		}
	}
	return X, Y, false
}
