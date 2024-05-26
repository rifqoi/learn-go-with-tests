package hello

import "fmt"

func Hello(name string, language string) string {

	const englishHelloPrefix = "Hello"
	helloLocales := map[string]string{
		"English":   "Hello",
		"Spanish":   "Hola",
		"Indonesia": "Assalamualaikum",
		"":          "Hello",
	}
	if name == "" {
		name = "World"
	}

	return fmt.Sprintf("%s, %s", helloLocales[language], name)
}

func main() {

}
