package mydict

import (
	"errors"
)

// Dictionary type
type Dictionary map[string]string

var (
	errorNotFound   = errors.New("Error: value not found")
	errCantUpdate   = errors.New("Error: update non-existing word")
	errorWordExists = errors.New("Error: word already exist")
)

// search dictionary_ key: true,false
func (d Dictionary) Search(word string) (string, error) {
	value, exists := d[word]
	if exists {
		return value, nil
	}
	return "", errorNotFound
}

// add dictionry key:value
func (d Dictionary) Add(word, def string) error {
	_, err := d.Search(word)
	if err == errorNotFound {
		d[word] = def
		return nil
	}
	return errorWordExists
}

// update dictionary
func (d Dictionary) Update(word, def string) error {
	_, err := d.Search(word)
	switch err {
	case nil:
		d[word] = def
	case errorNotFound:
		return errCantUpdate
	}
	var tte string
	tte = "text"
	var text string = "text"
	return nil
}

func (d Dictionary) Delete(word string) {
	delete(d, word)
}
