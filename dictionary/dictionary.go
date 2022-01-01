package dictionary

import "errors"

type Dictionary map[string]string

// Search looks up a word in a given dictionary and returns its definition if found.
func (d Dictionary) Search(word string) (definition string, err error) {
	definition, ok := d[word]
	if !ok {
		return "", errors.New("could not find the word you're looking for")
	}
	return definition, nil
}
