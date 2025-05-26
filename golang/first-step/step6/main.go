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
	n, _ := strconv.Atoi(sc.Text())
	sc.Scan()
	inputs := strings.Split(sc.Text(), " ")
	for i := 0; i < n; i++ {
		fmt.Println(inputs[i])
	}
}
