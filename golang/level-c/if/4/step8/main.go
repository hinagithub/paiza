// 問題
// 高さH、幅Wの箱( 0 ≦ H, W ≦ 10^9 )があります。
// この箱を 1 つ以上の高さ 2 、幅 2 の四角いタイルで敷き詰めます。
// 箱に隙間なくタイルを敷き詰めることはできますか？

package main

import (
	"fmt"
)

func main() {
	var height, width int
	fmt.Scan(&height, &width)

	// 高さまたは幅が2未満の場合はタイルを置けない
	if height < 2 || width < 2 {
		fmt.Println("NO")
		return
	}

	// 高さ・幅ともに2の倍数であれば、隙間なく敷き詰められる
	if height%2 == 0 && width%2 == 0 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
