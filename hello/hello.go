package hello

const englishGreetingPrefix = "Hello, "
const spanishGreetingPrefix = "Hola, "
const frenchGreetingPrefix = "Bonjour, "

func Hello(name string, language string) string {
	if len(name) == 0 {
		name = "World"
	}
	greetingPrefix := englishGreetingPrefix
	if language == "Spanish" {
		greetingPrefix = spanishGreetingPrefix
	}
	if language == "French" {
		greetingPrefix = frenchGreetingPrefix
	}
	return greetingPrefix + name
}
