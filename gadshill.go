package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"unicode"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		someline := scanner.Text()

		for _, r := range someline {
			fmt.Printf("%c ", r)

			if unicode.IsControl(r) {
				fmt.Printf("Control\t")
			}

			if unicode.IsDigit(r) {
				fmt.Printf("Digit\t")
			}

			if unicode.IsGraphic(r) {
				fmt.Printf("Graphic\t")
			}

			if unicode.IsLower(r) {
				fmt.Printf("Lower\t")
			}

			if unicode.IsNumber(r) {
				fmt.Printf("Number\t")
			}

			if unicode.IsPrint(r) {
				fmt.Printf("Print\t")
			}

			if unicode.IsSpace(r) {
				fmt.Printf("Space\t")
			}

			if unicode.IsSymbol(r) {
				fmt.Printf("Symbol\t")
			}

			if unicode.IsTitle(r) {
				fmt.Printf("Title\t")
			}

			if unicode.IsUpper(r) {
				fmt.Printf("Upper\t")
			}

			if unicode.IsLetter(r) {
				fmt.Printf("Letter\t")
			}

			if unicode.Is(unicode.Han, r) {
				fmt.Printf("Han\t")
			}

			if unicode.Is(unicode.Hiragana, r) {
				fmt.Printf("Hiragana\t")
			}

			if unicode.Is(unicode.Katakana, r) {
				fmt.Printf("Katakana")
			}

			fmt.Println()
		}
	}
}
