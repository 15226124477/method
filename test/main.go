package main

import (
	"fmt"
	"github.com/15226124477/method"
)

func main() {
	abc := make([]string, 0)
	abc = append(abc, "123")
	abc = append(abc, "12213")

	fmt.Println(
		method.IsContain(abc, "1223"))
}
