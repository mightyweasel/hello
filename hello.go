package main

import "fmt"
import "time"

func main() {
	fmt.Printf("hello, world\n")

	t := time.Now()

	fmt.Println(t.Format("2006-01-02 15:04:05"))
	fmt.Println(t.UTC())
	fmt.Println(t.Format(time.RFC850))
}
