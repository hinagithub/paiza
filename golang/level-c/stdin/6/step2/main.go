package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan() // 1行目（N）は読み飛ばす
	sc.Scan()
	for _, num := range strings.Fields(sc.Text()) {
		fmt.Println(num)
	}
}
