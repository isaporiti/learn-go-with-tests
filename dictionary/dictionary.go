package dictionary

import "errors"

type Dictionary map[string]string

var (
	wordNotFoundError       = errors.New("could not find the word you're looking for")
	wordAlreadyDefinedError = errors.New("cannot add word because it already exists")
	wordDoesNotExistError   = errors.New("cannot update word missing in dictionary")
)

// Search looks up a word in a given dictionary and returns its definition if found.
func (d Dictionary) Search(word string) (definition string, err error) {
	definition, ok := d[word]
	if !ok {
		return definition, wordNotFoundError
	}
	return definition, nil
}

// Add inserts a new word and definition to the Dictionary, if the word is not defined yet.
// If already defined, its definition won't be overwritten and a wordAlreadyDefinedError error will be returned.
func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)
	if err == nil {
		return wordAlreadyDefinedError
	}
	if err == wordNotFoundError {
		d[word] = definition
	}
	return nil
}

// Update overrides an existing word's definition with a new definition.
// If the word doesn't exist, a non-nil error is returned.
func (d Dictionary) Update(word, newDefinition string) error {
	_, wordExists := d[word]
	if wordExists {
		d[word] = newDefinition
		return nil
	}
	return wordDoesNotExistError
}

func (d Dictionary) Delete(word string) {
	delete(d, word)
}
