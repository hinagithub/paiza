// 問題
// 整数N, 2 つの数列A, B が与えられます。
//  1 ≦ i ≦ N を満たす整数 i のうち、A_i と B_i が等しくなるような i の個数を求めてください。

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
	length := atoi(readline(scanner))
	aStrs := strings.Fields(readline(scanner))
	bStrs := strings.Fields(readline(scanner))

	if len(aStrs) != length || len(bStrs) != length {
		fmt.Fprintln(os.Stderr, "input length does not match")
		os.Exit(1)
	}

	equalCount := 0
	for i := 0; i < length; i++ {
		if aStrs[i] == bStrs[i] {
			equalCount++
		}
	}
	fmt.Println(equalCount)

}

func readline(scanner *bufio.Scanner) string {
	if !scanner.Scan() {
		fmt.Fprintln(os.Stderr, "failed to scan input")
		os.Exit(1)
	}
	return scanner.Text()
}

func atoi(text string) int {
	num, err := strconv.Atoi(text)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to atoi: %v \n", err)
		os.Exit(1)
	}
	return num
}
