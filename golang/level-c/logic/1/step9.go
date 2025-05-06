package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	inputs := strings.Fields(scanner.Text())
	if len(inputs) != 3 {
		fmt.Println("3つのビット（0または1）を入力してください")
		return
	}

	a, _ := strconv.Atoi(inputs[0])
	b, _ := strconv.Atoi(inputs[1])
	cin, _ := strconv.Atoi(inputs[2])

	sum := a ^ b ^ cin                       // S = A XOR B XOR Cin
	carry := (a & b) | (b & cin) | (a & cin) // C = AB + BCin + ACin

	fmt.Println(carry, sum)
}
