package ejericios01

import "strconv"

func Ejercicios01(input string) (int, string) {

	num, err := strconv.Atoi(input)
	if err != nil {
		return 0, "Hubo un error: " + err.Error()
	}
	if num == 100 {
		return num, "Es igual a 100"
	} else if num >= 100 {
		return num, "Es mayor a 100"
	} else {
		return num, "Es menor a 100"
	}

}
