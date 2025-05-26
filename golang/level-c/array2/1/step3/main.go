// 問題
// 配列 A の要素数 N と配列 A の各要素である整数 A_1, A_2, ..., 
// A_N が与えられるので、配列 A の要素の最大値 max を求めてください。

package main
import(
    "fmt"
    "os"
    "bufio"
    "strconv"
)
func main(){
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan() // 1行目の要素数は使用しないのでスキップ
    
    // nil でスタートするようポインタを使う
    var maxVal *int 
    for scanner.Scan() {
        num, err := strconv.Atoi(scanner.Text())
        if err != nil {
            fmt.Fprintf(os.Stderr, "atoi err:")
            os.Exit(1)
        }
        if maxVal == nil || num > *maxVal {
            maxVal = &num
        }
    }
    fmt.Println(*maxVal)
}
