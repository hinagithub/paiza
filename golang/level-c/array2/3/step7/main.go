// 問題
// あなたは集団行動のリーダーです。次のような指示を出すことで様々な列の操作ができます。

// ・swap A B
// 先頭から A 番目の人と、先頭から B 番目の人の位置を入れ替える。
// ・reverse
// 列の前後を入れ替える。
// ・resize C
// 先頭から C 人を列に残し、それ以外の人を全員列から離れさせる。ただし、列が既に C 人以下の場合、何も行わない。

// 初め、列には番号 1 〜 N の N 人がおり、先頭から番号の昇順に並んでいます。(1, 2 , 3, ..., N)
// あなたの出した指示の回数 Q とその指示の内容 S_i (1 ≦ i ≦ Q) が順に与えられるので、全ての操作を順に行った後の列を出力してください。

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
	inputs := strings.Fields(readLine(scanner))
	personCount := atoi(inputs[0])
	length := atoi(inputs[1])
	results := make([]int, personCount)
	for i := range results {
		results[i] = i + 1
	}

	for i := 0; i < length; i++ {
		orderInfo := strings.Fields(readLine(scanner))
		command := orderInfo[0]
		switch command {
		case "swap":
			indexA := atoi(orderInfo[1])
			indexB := atoi(orderInfo[2])
			results[indexA-1], results[indexB-1] = results[indexB-1], results[indexA-1]
		case "reverse":
			results = reverse(results)
		case "resize":
			newLength := atoi(orderInfo[1])
			if len(results) > newLength {
				results = results[:newLength]
			}
		default:
			fmt.Fprintf(os.Stderr, "unknown command: %s\n", command)
			os.Exit(1)
		}
	}
	for _, r := range results {
		fmt.Println(r)
	}
}

func reverse(arr []int) []int {
	results := make([]int, len(arr))
	for i, v := range arr {
		results[len(arr)-i-1] = v
	}
	return results
}

func readLine(scanner *bufio.Scanner) string {
	if !scanner.Scan() {
		fmt.Fprintln(os.Stderr, "failed to scan input")
		os.Exit(1)
	}
	return scanner.Text()
}

func atoi(text string) int {
	num, err := strconv.Atoi(text)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to convert to int: %v\n", err)
		os.Exit(1)
	}
	return num
}
