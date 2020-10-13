package utils

import (
	"errors"
	"io/ioutil"
)

func ReadFile(filename string) (string, error) {
	if IsEmpty(filename) {
		return "", errors.New("Verilen dosya boş olamaz!")
	}
	bytes, err := ioutil.ReadFile(filename)

	CheckError(err)
	return string(bytes), nil

}
