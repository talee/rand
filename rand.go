package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	prompt()

	var i int
	for setInput(&i) {
		fmt.Println(r.Intn(i))
		prompt()
	}
}

func prompt() {
	fmt.Print("rand < ")
}

func setInput(val ...interface{}) bool {
	n, err := fmt.Scanf("%d", val...)
	return n > 0 && err == nil
}
