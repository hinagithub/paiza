// 問題
// 配列 A の要素数 N と配列 A の各要素 A_1, A_2, ..., A_N が与えられるので、
// 配列 A には何種類の値が含まれているかを求めてください。

package main
import (
    "bufio"
    "fmt"
    "os"
    "strconv"
)
func main(){
    scanner := bufio.NewScanner(os.Stdin)
    length := a2i(readline(scanner))
    
    // 重複のない値を記録するマップ
    seen := make(map[int]bool, length)
    for i:= 0; i<length; i++ {
        num := a2i(readline(scanner))
        seen[num] = true
    }
    fmt.Println(len(seen))
}

func a2i(text string) int {
    num, err := strconv.Atoi(text)
    if err != nil {
        fmt.Fprintf(os.Stderr, "failed to atoi: %v", err)
        os.Exit(1)
    }
    return num
}

func readline(scanner *bufio.Scanner) string {
    if !scanner.Scan() {
        fmt.Fprintln(os.Stderr, "failed to scan input")
        os.Exit(1)
    }
    return scanner.Text()
}
