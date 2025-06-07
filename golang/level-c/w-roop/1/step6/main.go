// 整数 N , K と N 行 K 列 の二次元配列 A が与えられます。 
// A の要素のうち、1 要素だけ 1 になっている要素があるので、その要素の行と列を出力してください。

package main
import(
    "bufio"
    "fmt"
    "os"
    "strings"
    "strconv"
)
func main(){
    scanner := bufio.NewScanner(os.Stdin)
    inputs := strings.Fields(readLine(scanner))
    length := atoi(inputs[0])

    // 2重配列を走査し1を見つけたら行例を出力
    for i := 0; i < length; i ++ {
        row := strings.Fields(readLine(scanner))
        for j, text := range row {
            if text == "1" {
                fmt.Println(i+1, j+1)
            }
        }
    }
 
}


func readLine(scanner *bufio.Scanner) string {
	if !scanner.Scan() {
		fmt.Fprintln(os.Stderr, "failed to scan input")
		os.Exit(1)
	}
	return scanner.Text()
}

func atoi(text string) int {
	num, err := strconv.Atoi(text)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to atoi: %v \n", err)
		os.Exit(1)
	}
	return num
}
