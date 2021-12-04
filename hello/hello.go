package hello

func Hello(name string) string {
	if len(name) == 0 {
		return "Hello, World"
	}
	return "Hello, " + name
}
