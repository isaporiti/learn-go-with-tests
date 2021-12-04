package hello

const englishGreetingPrefix = "Hello, "
const spanishGreetingPrefix = "Hola, "

func Hello(name string, language string) string {
	if len(name) == 0 {
		name = "World"
	}
	greetingPrefix := englishGreetingPrefix
	if language == "Spanish" {
		greetingPrefix = spanishGreetingPrefix
	}
	return greetingPrefix + name
}
