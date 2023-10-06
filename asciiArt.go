package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

func ReadData(banner string) string {
	data, err := os.ReadFile(banner)
	if err != nil {
		log.Fatal(err)
	}

	return string(data)
}

// w := http.ResponseWriter (lets code Fprint to html); banner := font text file, as chosen by radio buttons; textbox := input in textbox on html
func AsciiArt(w http.ResponseWriter, banner string, textbox string) {
	// standard.txt (readData), split by '\n'
	splitStr := strings.Split(string(ReadData(banner+".txt")), "\n")

	// replaces new lines (\n) in textbox with '\\n'
	replaceNewline := strings.ReplaceAll(textbox, "\r\n", "\\n")
	// Textbox, split by '\\n', literally the symbols \ and n together
	splitText := strings.Split(string(replaceNewline), "\\n")
	// for each slice of arg...
	for a := 0; a < len(splitText); a++ {

		runeArgs := []rune(splitText[a])
		// if the slice of arg contains nothing then print a new line

		for i := 1; i <= 8; i++ {
			for j := 0; j <= len(runeArgs)-1; j++ {
				letterArgs := runeArgs[j]
				// rune to line number
				lineNumber := (int(letterArgs)-32)*9 + i
				// prints from line number to the next 8 lines
				//	fmt.Println(splitStr[lineNumber])
				fmt.Fprint(w, splitStr[lineNumber])
			}
			fmt.Fprintln(w)
		}
	}
}
