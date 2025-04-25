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
	inputs := strings.Split(sc.Text(), " ")

	height, _ := strconv.Atoi(inputs[0])
	width, _ := strconv.Atoi(inputs[1])
	a, _ := strconv.Atoi(inputs[2])
	b, _ := strconv.Atoi(inputs[3])

	var rowWidth int

	for i := 0; i < height; i++ {
		rowCells := make([]string, width)
		for j := 0; j < width; j++ {
			cellText := fmt.Sprintf("(%d, %d)", a, b)
			rowCells[j] = cellText
		}

		rowText := strings.Join(rowCells, " | ")
		fmt.Println(rowText)

		// 1行目で横幅を測定
		if i == 0 {
			rowWidth = len(rowText)
		}

		if i != height-1 {
			fmt.Println(strings.Repeat("=", rowWidth))
		}
	}
}
