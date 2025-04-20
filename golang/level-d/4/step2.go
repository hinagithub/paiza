package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	text := sc.Text()
	inputs := strings.Split(text, " ")
	n1, _ := strconv.Atoi(inputs[0])
	n2, _ := strconv.Atoi(inputs[1])
	answer1 := n1 - n2
	answer2 := n1 * n2
	fmt.Println(answer1, answer2)
}
