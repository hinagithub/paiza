// 問題
// 配列 A の要素数 N と整数 K, 配列 A の各要素 A_1, A_2, ..., A_N が与えられるので、
// 配列 A の全ての要素を + K した後の A の各要素を出力してください。

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
    
    // 1行目の要素数と加算値を取得
    inputs := strings.Fields(readLine(scanner))
    if len(inputs) < 2 {
        fmt.Fprintln(os.Stderr, "invalid input: need 2 integers in first line")
        os.Exit(1)
    }
    length := parseInt(inputs[0])
    adding := parseInt(inputs[1])
    
    // 要素全てに加算して出力
    for i := 0; i < length; i++ {
        num := parseInt(readLine(scanner))
        fmt.Println(num + adding)
    }
}

func parseInt(text string) int {
    num, err := strconv.Atoi(text)
    if err != nil {
        fmt.Fprintf(os.Stderr, "failed to atoi: %v \n", err)
        os.Exit(1)
    }
    return num
}

func readLine(scanner *bufio.Scanner) string {
    if !scanner.Scan() {
        fmt.Fprintln(os.Stderr, "failed to scan input")
        os.Exit(1)
    }
    return scanner.Text()
}
