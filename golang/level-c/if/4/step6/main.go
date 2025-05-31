// 問題
// 京子ちゃんは 1 周がNメートルの円の外周上にいます。
// 京子ちゃんは 1 ターンで現在地点から K メートル時計回りに円周上を歩きます。
// スタート地点から開始してTターン歩いたとき、京子ちゃんが丁度スタート地点に戻ってきているかどうかを判定してください。
// ただし、京子ちゃんは各ターン必ず歩く必要があり、同じ場所にとどまることはできません。

package main

import "fmt"

func main() {

	var length, runDistance, times int
	fmt.Scan(&length, &runDistance, &times)

	// 1周歩いて最初の位置ちょうどに戻っているかを判定
	if (runDistance*times)%length == 0 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
