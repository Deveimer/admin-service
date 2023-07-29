package utils

import gonanoid "github.com/matoous/go-nanoid"

func GenerateAlphaNumericUniqueId(length int) (string, error) {
	id, err := gonanoid.Generate("ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789", length)
	if err != nil {
		return "", nil
	}
	return id, nil
}

func GenerateNumericUniqueId(length int) (string, error) {
	id, err := gonanoid.Generate("0123456789", length)
	if err != nil {
		return "", nil
	}
	return id, nil
}
