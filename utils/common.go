package utils

import (
	sid "github.com/ventu-io/go-shortid"
)

// GenerateID -
func GenerateID() (string, error) {
	return sid.Generate()
}
