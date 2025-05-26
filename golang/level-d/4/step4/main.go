// 問題
// ある電車に a 人が乗っています。
// 駅に到着した時に b 人が降りて新たに c 人が乗車する時、
// 電車に乗っている乗客人数を求めてください。

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
	ridingCount, _ := strconv.Atoi(inputs[0])
	gettingOffCount, _ := strconv.Atoi(inputs[1])
	newRidingCount, _ := strconv.Atoi(inputs[2])
	answer := ridingCount - gettingOffCount + newRidingCount
	fmt.Println(answer)
}
