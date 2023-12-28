package main

import (
	"fmt"
	"strings"
)

var morseCodeMap = map[string]string{
	".-": "A", "-...": "B", "-.-.": "C", "-..": "D", ".": "E",
	"..-.": "F", "--.": "G", "....": "H", "..": "I", ".---": "J",
	"-.-": "K", ".-..": "L", "--": "M", "-.": "N", "---": "O",
	".--.": "P", "--.-": "Q", ".-.": "R", "...": "S", "-": "T",
	"..-": "U", "...-": "V", ".--": "W", "-..-": "X", "-.--": "Y",
	"--..": "Z", ".----": "1", "..---": "2", "...--": "3", "....-": "4",
	".....": "5", "-....": "6", "--...": "7", "---..": "8", "----.": "9",
	"-----": "0", "--..--": ",", ".-.-.-": ".", "..--..": "?", "-..-.": "/",
	"-....-": "-", "-.--.": "(", "-.--.-": ")", "{": "{", "}": "}",
}

func reverseMap(inputMap map[string]string) map[string]string {
	reversedMap := make(map[string]string)
	for key, value := range inputMap {
		reversedMap[value] = key
	}
	return reversedMap
}
func Encrypt(message string) string {
	var cipher strings.Builder
	reversed := reverseMap(morseCodeMap)
	for _, letter := range message {
		if string(letter) != " " {
			morseCode, ok := reversed[string(letter)]
			if ok {
				cipher.WriteString(morseCode + " ")
			} else {
				cipher.WriteString(" ")
			}
		} else {
			cipher.WriteString(" ")
		}
	}

	return cipher.String()
}

func Decrypt(message string) string {
	words := strings.Split(message, " ")
	var decipher strings.Builder

	for _, word := range words {
		if word != "" {
			decipher.WriteString(morseCodeMap[word])
		} else {
			decipher.WriteString(" ")
		}
	}

	return decipher.String()
}

func main() {
	message := "HELLOWORLD"
	encrypted := Encrypt(strings.ToUpper(message))
	fmt.Println(encrypted)

	message = ".... . .-.. .-.. --- .-- --- .-. .-.. -.."
	fmt.Println(message)
	fmt.Println("Decrypt:")
	decrypted := Decrypt(message)
	fmt.Println(decrypted)
}
