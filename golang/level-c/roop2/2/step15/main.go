// 整数 N が与えられます。
// N の約数を小さい方から順に改行区切りで出力してください。

package main
import "fmt"
func main(){
   var n int;
   fmt.Scan(&n)
   
   for i := 1; i <= n; i++ {
       if n % i == 0 {
           fmt.Println(i)
       }
   }
}
