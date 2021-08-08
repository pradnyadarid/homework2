package services

//jika nilai x <= 40 maka nilainya C
//jika nilai 40 < x <= 70 maka nilai B
//jika nilai x >70 maka nilainya A

func Nilai(x int) string {
	if x <= 40 {
		return "C"

	} else if 40 < x && x <= 70 {
		return "B"

	} else {
		return "A"

	}
}
