// 整数 N , K が与えられるので、 1 から N までの数字を半角スペース区切りしたもの
// "1 2 ... (N-1) N" を K 行出力してください。

package main
import(
    "fmt"
    "strings"
    "strconv"
)
func main(){
    var n, k int
    fmt.Scan(&n, &k)
    
    strNumbers := []string{}
    for i := 1; i <= n; i++ {
        strNumbers = append(strNumbers, strconv.Itoa(i))
    }
    
    for i := 0; i < k; i++ {
        fmt.Println(strings.Join(strNumbers, " "))
    }
}
