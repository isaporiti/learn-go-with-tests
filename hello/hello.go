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

func Hello(name string, language language) string {
	if len(name) == 0 {
		name = "World"
	}
	return getGreetingPrefix(language) + name
}

func getGreetingPrefix(language language) string {
	greetingPrefix := defaultGreeting
	if len(language) > 0 {
		greetingPrefix = greetings[language]
	}
	return greetingPrefix
}
