package plugin

func encrypt(buchstabe rune) rune {
	switch {
	case buchstabe == 'a':
		return 'b'
	case buchstabe == 'b':
		return 'c'
	case buchstabe == 'c':
		return 'd'
	}
	return buchstabe
}

func decrypt(buchstabe rune) rune {
	switch {
	case buchstabe == 'b':
		return 'a'
	case buchstabe == 'c':
		return 'b'
	case buchstabe == 'd':
		return 'c'
	}
	return buchstabe
}
