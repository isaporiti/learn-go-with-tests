package hello

type language string

const English language = "English"
const Spanish language = "Spanish"
const French language = "French"

var greetings = map[language]string{
	English: "Hello, ",
	Spanish: "Hola, ",
	French:  "Bonjour, ",
}

var defaultGreeting = greetings[English]

const defaultName = "World"

func Hello(name string, language language) string {
	return getGreetingPrefix(language) + getName(name)
}

func getGreetingPrefix(language language) string {
	if len(language) == 0 {
		return defaultGreeting
	}
	return greetings[language]
}

func getName(name string) string {
	if len(name) == 0 {
		return defaultName
	}
	return name
}
