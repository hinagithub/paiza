// 問題
// 日時が "yyyy/MM/dd/hh:mm" の形式で与えられるので、年・月・日・時・分をそれぞれ出力してください。

package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	const layout = "2006/01/02/15:04"
	scanner := bufio.NewScanner(os.Stdin)
	if !scanner.Scan() {
		fmt.Fprintf(os.Stderr, "scan err")
		os.Exit(1)
	}
	text := scanner.Text()
	t, err := time.Parse(layout, text)
	if err != nil {
		fmt.Fprintf(os.Stderr, "time parse err", err)
		os.Exit(1)
	}

	fmt.Println(t.Year())
	fmt.Printf("%02d\n", t.Month())
	fmt.Printf("%02d\n", t.Day())
	fmt.Printf("%02d\n", t.Hour())
	fmt.Printf("%02d\n", t.Minute())

	// yyyy := text[:4]
	// MM := text[5:7]
	// dd := text[8:10]
	// hh := text[11:13]
	// mm := text[14:16]

	// fmt.Println(yyyy)
	// fmt.Println(MM)
	// fmt.Println(dd)
	// fmt.Println(hh)
	// fmt.Println(mm)

}
