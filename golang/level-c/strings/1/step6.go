package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	text := readline(scanner)
	target := readline(scanner)

	for i, char := range []rune(text) {
		if char == rune(target[0]) {
			fmt.Println(i + 1)
		}
	}

}

func readline(scanner *bufio.Scanner) string {
	if !scanner.Scan() {
		fmt.Fprintln(os.Stderr, "scan err")
		os.Exit(1)
	}
	return scanner.Text()
}
