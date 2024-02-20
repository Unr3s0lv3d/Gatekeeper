package main

import (
	"fmt"
	"time"
)

func main() {

	user_input := ""

	fmt.Print("Kenteken: ")
	fmt.Scan(&user_input)

	currentTime := time.Now()
	fmt.Println(currentTime.Hour())

	tijd := currentTime.Hour()
	var found bool
	Groet := ""
	test_slice := []string{"AB-123-CD", "XY-456-ZW", "GH-789-IJ", "KL-012-MN", "PQ-345-RS"}
	switch {
	case tijd >= 7 && tijd < 12:
		Groet = ("Goedenmorgen, welkom bij Fonteyn Vakantieparken.")

	case tijd >= 12 && tijd < 18:
		Groet = ("Goedenmiddag, welkom bij Fonteyn Vakantieparken.")

	case tijd >= 18 && tijd < 23:
		Groet = ("Goedenavond, welkom bij Fonteyn Vakantieparken.")
	default:
		Groet = ("Sorry, de parkeerplaats is 's nachts gelosten.")

	}

	for i := 0; i < len(test_slice); i++ {
		if user_input == (test_slice[i]) {
			fmt.Print("Ewa")
			found = true
			break
		}

	}
	if !found {
		fmt.Print("Ewa2\n")
	}

	fmt.Print(Groet)
	fmt.Print("\n" + test_slice[1])

}
