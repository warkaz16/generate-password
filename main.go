package main

import (
	"fmt"
	"math/rand"
)

func generatePassword(length int, difficulty string) string {

	letters := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z", "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}

	digits := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}

	symbols := []string{"!", "#", "$", "%", "&", "(", ")", "*", "+", ",", "-", ".", "/", ":", ";", "<", "=", ">", "?", "@", "[", "]", "^", "_", "{", "|", "}", "~"}

	newPassword := ""

	types := [][]string{letters, digits, symbols}

	if difficulty == "easy" {

		for i := 0; i < length; i++ {

			randomLetter := letters[rand.Intn(len(letters))]
			newPassword += randomLetter

		}
	}

	if difficulty == "medium" {

		for i := 0; i < length; i++ {

			randomFull := types[rand.Intn(len(types)-1)]

			randomLetter := randomFull[rand.Intn(len(randomFull)-1)]
			newPassword += randomLetter
		}
	}

	if difficulty == "hard" {

		for i := 0; i < length; i++ {

			randomFull := types[rand.Intn(len(types))]

			randomLetter := randomFull[rand.Intn(len(randomFull)-1)]
			newPassword += randomLetter

		}
	}

	return newPassword
}

func main() {
	difficulty := ""
	fmt.Print("Choose difficulty (easy, medium, hard): ")
	fmt.Scan(&difficulty)

	if difficulty == "easy" {
		fmt.Println(generatePassword(10, "easy"))
	} else if difficulty == "medium" {
		fmt.Println(generatePassword(10, "medium"))
	} else if difficulty == "hard" {
		fmt.Println(generatePassword(10, "hard"))
	} else {
		fmt.Println("ERROR: WRONG DIFFICULTY")
	}

}
