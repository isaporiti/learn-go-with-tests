package hello

const greeting = "Hello, "

func Hello(name string) string {
	if len(name) == 0 {
		name = "World"
	}
	return greeting + name
}
