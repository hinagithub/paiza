package main

import(
  "fmt"
  "bufio"
  "os"
)

func main(){
  sc := bufio.NewScanner(os.Stdin)
  sc.Scan()
  word := sc.Text()
  fmt.Println(word)
}
