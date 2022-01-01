package dictionary

// Search looks up a word in a given dictionary and returns its definition if found.
func Search(dictionary map[string]string, word string) (definition string) {
	definition = dictionary[word]
	return definition
}
