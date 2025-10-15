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

	return "Your password: " + newPassword + "\n"
}

func main() {

	difficulty := ""
	length := 0

	fmt.Println("")
	fmt.Print("Choose password`s length (6-15): ")
	fmt.Scan(&length)

	fmt.Println("")

	fmt.Print("Choose difficulty (easy, medium, hard): ")
	fmt.Scan(&difficulty)
	fmt.Println("")

	if difficulty == "easy" && length >= 6 && length <= 15 {
		fmt.Println(generatePassword(length, "easy"))
	} else if difficulty == "medium" && length >= 6 && length <= 15 {
		fmt.Println(generatePassword(length, "medium"))
	} else if difficulty == "hard" && length >= 6 && length <= 15 {
		fmt.Println(generatePassword(length, "hard"))
	} else {
		fmt.Println("ERROR: WRONG DATA")
		fmt.Println("")
	}

}
