// 問題
// 直線上の道があり、最初は X = 0 の地点にいます。この道はX = T + 0.1 の地点で崖になっています。
// また、 1 歩でX軸方向にK進むことができます。崖に落ちずにN歩進むことはできるでしょうか。

package main

import "fmt"

func main() {
	// N 歩数, K 歩幅, T 崖の位置
	var stepCount, stride, cliff int
	fmt.Scan(&stepCount, &stride, &cliff)

	totalDistanceWaled := stride * stepCount
	if totalDistanceWaled <= cliff {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}

}
