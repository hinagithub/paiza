// 問題
// 配列 A の要素数 N と配列 A の各要素 A_1, A_2, ..., A_N が整数として与えられるので、
// 配列 A の全ての要素の和を求めてください。

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
    var sum int = 0
    for scanner.Scan() {
        num, err := strconv.Atoi(scanner.Text())
        if err != nil {
            fmt.Fprintf(os.Stderr, "atoi err: %v", err)
            os.Exit(1)
        }
        sum += num
    }
    fmt.Println(sum)
}
