package canvas

var denomW, denomH = 256, 256

func GetDenoms() (int, int) {
	return denomW, denomH
}

func SetDenoms(w, h int) {
	denomW = w
	denomH = h
}
