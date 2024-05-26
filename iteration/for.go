package iteration

const repeatCount = 10

func Repeat(character string) string {
	fullStr := ""
	for i := 0; i <= repeatCount; i++ {
		fullStr = fullStr + character
	}

	return fullStr
}
