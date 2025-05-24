// 問題
// 配列 A の要素数 N と整数 K , 配列 A の各要素 A_1, A_2, ..., A_N が与えられるので、
// A に K が 1 つでも含まれている場合は Yes を、含まれていない場合は No を出力してください。

package main
import(
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)
func main(){
    const (
        Yes = "Yes"
        No = "No"
    )
    
    // ターゲット値を取得
    scanner := bufio.NewScanner(os.Stdin)
    if !scanner.Scan(){
        fmt.Fprintln(os.Stderr, "failed to scan first line")
        os.Exit(1)
    }
    inputs := strings.Fields(scanner.Text())
    target, _ := strconv.Atoi(inputs[1])
    
    // ターゲット値が含まれていればYesを出力
    contains := false
    for scanner.Scan() {
        num, _ := strconv.Atoi(scanner.Text())
        if num == target {
            contains = true
        }
    }
    if contains {
        fmt.Println(Yes)
    } else {
        fmt.Println(No)
    }
}
