// 問題
// 配列 A の要素数 N と配列 A の各要素である整数 A_1, A_2, ..., A_N が与えられるので、
// 配列 A の要素の最小値 min を求めてください。

package main

import(
    "bufio"
    "fmt"
    "os"
    "strconv"
)

func main(){
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan() // 1行目の要素数は使用しないのでスキップ
    
    var minVal = *int
    for scanner.Scan(){
        num, err := strconv.Atoi(scanner.Text())
        if err != nil {
            fmt.Fprintf(os.Stderr, "atoi err: %v", err)
            os.Exit(1)
        }
        if minVal == nil || num < *minVal {
            minVal = num
        }
    }
    fmt.Println(*minVal)
}
