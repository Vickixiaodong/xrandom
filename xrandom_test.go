package xrandom

import (
	"log"
	"testing"
)

func TestGetRandomChar(t *testing.T) {
	log.Println(GetRandomChar(UpperCase, 32))
	log.Println(GetRandomChar(LowerCase, 32))
	log.Println(GetRandomChar(Number, 32))
	log.Println(GetRandomChar(SpecialCharacter, 32))
	log.Println(GetRandomChar("{}", 32))
}

func TestString(t *testing.T) {
	log.Println(String("x", 32))
	log.Println(String("X", 32))
	log.Println(String("7", 32))
	log.Println(String(";", 32))
	log.Println(String("xX7;", 32))

	log.Println(String("xX", 32))
}
