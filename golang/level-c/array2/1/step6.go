// 問題
// 配列 A の要素数 N と整数 K , 配列 A の各要素 A_1, A_2, ..., A_N が与えられるので、
// K であるような A の要素のうち、要素番号が最小のものを出力してください。
// A に K が含まれない場合は -1 を出力してください。

package main
import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)
func main(){
    scanner := bufio.NewScanner(os.Stdin)
    
    // 1行目: 要素数とターゲット値
    scan(scanner)
    inputs := strings.Fields(scanner.Text())
    length := atoi(inputs[0])
    target := atoi(inputs[1])

    // 要素を1つずつ読み、ターゲットと一致する最初のインデックスを探す
    for i := 0; i < length; i++ {
        scan(scanner)
        num := atoi(scanner.Text())
        if num == target {
            fmt.Println(i + 1)
            os.Exit(0)
        }
    }
    
    // 見つからなかった場合
    fmt.Println(-1)
}

func scan(scanner *bufio.Scanner) {
    if !scanner.Scan(){
        fmt.Fprintln(os.Stderr, "failed to scan input")
        os.Exit(1)
    }
}

func atoi(text string) int {
    num, err := strconv.Atoi(text)
    if err != nil {
        fmt.Fprintf(os.Stderr, "failed to atoi: %v", err)
        os.Exit(1)
    }
    return num
}
