// 問題
// 自然数 N, M と N 個の自然数からなる数列 A と M 個の自然数からなる数列 B が与えられます。
// 1 行目には数列 A の最初の B_1 個の値を出力し、 2 行目にはその次から B_2 個の値を出力します。
// このように、i 行目には数列 A の 1 + B_1 + B_2 + ... + B_{i - 1} 番目の値から B_i 個の値を出力してください。
// 言い換えると、数列 A の値を B_1 個、B_2個、... B_M 個で分割し、それぞれの数列を改行区切りで出力してください。

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
	m, _ := strconv.Atoi(inputs[1])
	sc.Scan()
	arrA := strings.Split(sc.Text(), " ")
	sc.Scan()
	arrB := strings.Split(sc.Text(), " ")

	start := 0
	end := 0
	for i := 0; i < m; i++ {
		start = end
		end, _ = strconv.Atoi(arrB[i])
		end += start
		resutls := arrA[start:end]
		fmt.Println(strings.Join(resutls, " "))
	}
}
