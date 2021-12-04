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
	greetingPrefix := defaultGreeting
	if len(language) > 0 {
		greetingPrefix = greetings[language]
	}
	return greetingPrefix
}

func getName(name string) string {
	if len(name) == 0 {
		name = defaultName
	}
	return name
}
