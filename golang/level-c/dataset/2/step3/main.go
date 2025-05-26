package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const wall = '#' // 壁を表す定数

func main() {
	sc := bufio.NewScanner(os.Stdin)

	// 最初の行を読み取り
	sc.Scan()
	tokens := strings.Fields(sc.Text())

	height, _ := strconv.Atoi(tokens[0])
	r, _ := strconv.Atoi(tokens[2])
	c, _ := strconv.Atoi(tokens[3])

	// 迷路データを作成
	maze := make([][]rune, height)
	for i := 0; i < height; i++ {
		sc.Scan()
		maze[i] = []rune(sc.Text())
	}

	// インデックス補正（0始まり）
	if maze[r-1][c-1] == wall {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
