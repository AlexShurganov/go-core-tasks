package main

import (
	"crypto/sha256"
	"fmt"
	"strings"
)

var (
	result strings.Builder
)

func ShowType(value any) error {
	fmt.Printf("Value: %v, Type: %T\n", value, value)
	_, err := result.Write([]byte(fmt.Sprintf("%v", value)))
	if err != nil {
		return err
	}
	result.WriteString(" ")

	return nil
}

func ParseToRuneSlice(st string) []rune {
	sliceOfRune := []rune(st)

	return sliceOfRune
}

func AddSalt(runes []rune) []byte {
	salt := "go-2024"

	data := []byte(string(runes))

	dataWithSalt := string(data[:len(data)/2]) + salt + string(data[len(data)/2:])

	return []byte(dataWithSalt)
}

func Hash(bytes []byte) string {
	h := sha256.New()
	h.Write(bytes)
	return fmt.Sprintf("%x", h.Sum(nil))
}

func main() {
	var numDecimal int = 42
	var numOctal int = 052
	var numHexadecimal int = 0x2A
	var pi float64 = 3.14
	var name string = "Golang"
	var isActive bool = true
	var complexNum complex64 = 1 + 2i

	fmt.Println("Values and types:")
	vars := []any{numDecimal, numOctal, numHexadecimal, pi, name, isActive, complexNum}
	for _, v := range vars {
		ShowType(v)
	}

	fmt.Println("\nHandle string:")
	runes := ParseToRuneSlice(name)
	saltedBytes := AddSalt(runes)
	hashedString := Hash(saltedBytes)

	fmt.Printf("Original: %s\n", name)
	fmt.Printf("With salt: %s\n", string(saltedBytes))
	fmt.Printf("SHA256 hash: %s\n", hashedString)

	fmt.Printf("\nBuilder: %s\n", result.String())
}
