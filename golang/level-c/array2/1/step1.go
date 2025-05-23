// 問題
// 配列 A の要素数 N と整数 K, 配列 A の各要素 A_1, A_2, ..., A_N が与えられるので、
// 配列 A に K がいくつ含まれるか数えてください。

package main
import (
    "bufio"
    "fmt"
    "os"
    "strings"
)
func main(){
    scanner := bufio.NewScanner(os.Stdin)
    
    // ターゲット文字を取得
    if !scanner.Scan() {
        fmt.Println("scan err")
        os.Exit(1)
    }
    inputs := strings.Fields(scanner.Text())
    target := inputs[1]
    
    // ターゲットと一致する要素の数をカウント
    count := 0
    for scanner.Scan() {
        if target == scanner.Text(){
            count++
        }
    }
    fmt.Println(count)

}
