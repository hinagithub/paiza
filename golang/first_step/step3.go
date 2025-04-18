package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()

	for i := 0; i < 3; i++ {
		sc.Scan()
		text := sc.Text()
		fmt.Println(text)
	}
}
