/**
 * Random tools
 */

package xrandom

import (
	"math/rand"
	"strings"
	"time"
)

const (
	UpperCase        = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	LowerCase        = "abcdefghijklmnopqrstuvwxyz"
	Number           = "0123456789"
	SpecialCharacter = "~`!@#$%^&*()_+=-;':/?.>,<\"\\"
)

/*
 * Random for string
 * case, number, special characters, length
 * "x" for lowercase
 * "X" for uppercase
 * "7" for number
 * ";" for special character
 */
func String(source string, length int) string {
	if strings.Contains(source, "X") {
		source += UpperCase
	}

	if strings.Contains(source, "x") {
		source += LowerCase
	}

	if strings.Contains(source, "7") {
		source += Number
	}

	if strings.Contains(source, ";") {
		source += SpecialCharacter
	}

	return GetRandomChar(source, length)
}

func GetRandomChar(source string, length int) string {
	bytes := []byte(source)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}
