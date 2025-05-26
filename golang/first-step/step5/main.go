package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	text := sc.Text()
	inputs := strings.Split(text, " ")

	for _, input := range inputs {
		fmt.Println(input)
	}

}
