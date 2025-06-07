// 正の整数 N が与えられるので、1 〜 N の整数を 1 から順に半角スペース区切りで 1 行で出力してください。

package main
import (
    "fmt"
    "strings"
    "strconv"
)
func main(){
    var n int
    fmt.Scan(&n)
    strNumbers := []string{}
    for i := 1; i <= n; i++ {
        strNumbers = append(strNumbers, strconv.Itoa(i))
    }
    fmt.Println(strings.Join(strNumbers, " "))
}
