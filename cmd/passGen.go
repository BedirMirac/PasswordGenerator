package cmd

import (
	"errors"
	"math/rand/v2"
)

func PassGen() string {
	const specialChars = "*!@#$%^&*()-+=[]{}<>"
	const numbers = "0123456789"
	const characters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	passOfString := make([]byte, (Length - Numbers - Symbols))
	passOfNumber := make([]byte, Numbers)
	passOfSpeacialChars := make([]byte, Symbols)
	for i := range passOfString {
		passOfString[i] = characters[rand.IntN(len(characters))]
	}
	for j := range passOfNumber {
		passOfNumber[j] = numbers[rand.IntN(len(numbers))]
	}
	for k := range passOfSpeacialChars {
		passOfSpeacialChars[k] = specialChars[rand.IntN(len(specialChars))]
	}
	pass := string(passOfNumber) + string(passOfSpeacialChars) + string(passOfString)

	return shuffle(pass)
}

func LenOfPass() error {

	if Length < 0 {
		return errors.New("length of password cannot be negative")
	}
	return nil
}

func SizeOfNumber() error {
	if Numbers < 0 {
		return errors.New("number of numbers cannot be negative")
	}
	if Numbers > Length {
		return errors.New("number of characters cannot be more than length of password")
	}
	return nil

}

func SizeOfSpecialCharacter() error {

	if Symbols < 0 {
		return errors.New("number of characters cannot be negative")
	}
	if Symbols > Length {
		return errors.New("number of characters cannot be more than length of password")
	}
	return nil
}

func IsValid() error {
	if Numbers+Symbols > Length {
		return errors.New("sum of size of numbers and special characters cannot be more than length of password")
	}
	return nil
}

func shuffle(s string) string {
	pass := []rune(s)

	rand.Shuffle(len(pass), func(i, j int) {
		pass[i], pass[j] = pass[j], pass[i]
	})

	return string(pass)
}

func validateInputs() error {
	if err := LenOfPass(); err != nil {
		return err
	}
	if err := SizeOfNumber(); err != nil {
		return err
	}
	if err := SizeOfSpecialCharacter(); err != nil {
		return err
	}
	if err := IsValid(); err != nil {
		return err
	}
	return nil
}
