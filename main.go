package main

import (
	"fmt"
	"runtime"
)

func main() {
	if runtime.GOOS == "windows" {
		fmt.Println("Belum tersedia untuk sistem operasi Windows")
	} else {
		Exec()
	}
}