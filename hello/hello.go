package hello

import (
	"fmt"
)

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

func Hello(name string, language language) (string, error) {
	greetingPrefix, err := getGreetingPrefix(language)
	if err != nil {
		return "", err
	}
	return greetingPrefix + getName(name), nil
}

func getGreetingPrefix(language language) (string, error) {
	if len(language) == 0 {
		return defaultGreeting, nil
	}
	if greeting, ok := greetings[language]; ok {
		return greeting, nil
	}
	return "", fmt.Errorf("'%s' language is unsupported", language)
}

func getName(name string) string {
	if len(name) == 0 {
		return defaultName
	}
	return name
}
