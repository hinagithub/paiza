// 問題
// 自然数 N と N 個の要素の数列 M が与えられます。
// 1 ≦ i ≦ N の各 i について、i 行目には以下の数列を出力してください。
// * 1 以上 M_i 以下のすべての自然数を昇順、半角スペース区切りで出力してください。

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
	stringNumbers := strings.Split(sc.Text(), " ")
	intNumbers := make([]int, len(stringNumbers))
	for i, strNumber := range stringNumbers {
		intNumbers[i], _ = strconv.Atoi(strNumber)
	}

	for i := 0; i < n; i++ {
		results := []string{}
		for j := 1; j <= intNumbers[i]; j++ {
			results = append(results, strconv.Itoa(j))
		}
		fmt.Println(strings.Join(results, " "))
	}

}
