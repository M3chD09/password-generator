package core

func IslowCase(text string) bool {
	l := len(text)
	for i := 0; i <= l-1; i++ {
		if text[i] < 97 || text[i] > 122 {
			return false
		}
	}
	return true
}
func Changeletter(input byte, mode int) (output byte, omode int) {
	d1 := map[byte]byte{'a': '4', 'b': '8', 's': '$'}
	d2 := map[byte]byte{'a': '@', 'b': '6', 'c': '#', 'e': '3', 'g': '9', 'i': '!', 'j': '7', 'l': '1', 'o': '0', 'q': '?', 's': '5', 'x': '*', 'z': '2'}
	if mode == 1 {
		output, ok := d1[input]
		if !ok {
			mode = 2
		} else {
			return output, 1
		}
	}
	if mode == 2 {
		output, ok := d2[input]
		if !ok {
			mode = 3
		} else {
			return output, 2
		}
	}
	if mode == 3 {
		return input, 3
	}
	if mode == 4 {
		return (input - 32), 4
	}
	return
}
func count(x *int) {
	if *x == 4 {
		*x = 1
	} else {
		*x = *x + 1
	}
}
func Generator(text string) string {
	var mode [26]int
	for i := 0; i <= 25; i++ {
		mode[i] = 1
	}
	l := len(text)
	otext := []byte(text)
	for i := 0; i <= l-1; i++ {
		n := int(byte(text[i] - 97))
		otext[i], mode[n] = Changeletter(text[i], mode[n])
		count(&mode[n])
	}
	return string(otext)
}
