// 整数 N が与えられるので、 1 から 5 までの数字を半角スペース区切りしたもの
// "1 2 3 4 5" を N 行出力してください。

package main
import "fmt"
func main(){
    var n int
    fmt.Scan(&n)
    
    const Text = "1 2 3 4 5"
    for i := 1; i <= n; i++ {
        fmt.Println(Text)
    }
}
