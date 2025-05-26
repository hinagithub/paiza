// 問題
// 1 行目に整数 N が与えられます。
// N 番目までのフィボナッチ数を出力してください。
// フィボナッチ数は
// F_0 = 0
// F_1 = 1
// F_(n+2) = F_n + F_(n+1) (n は 0 以上)
// とし、F_0 を 1 番目とします。

package main
import "fmt"

func main() {
    var limit int 
    fmt.Scan(&limit)
    
    var a, b int = 0, 1
    for i := 0; i < limit; i++ {
        fmt.Println(a)
        a, b = b, a + b
    }
}
