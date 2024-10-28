package iteration

func Repeat(character string, repeatCount int) string {
	var repeat string
	for i := 0; i < repeatCount; i++ {
		repeat += character
	}
	return repeat
}
