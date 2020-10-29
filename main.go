package main

import (
	"fmt"
	"runtime"

	flag "github.com/spf13/pflag"
)

var (
	pcapIn string
	pcapOut string
)

func main() {
	if runtime.GOOS == "windows" {
		fmt.Println("Belum tersedia untuk sistem operasi Windows")
	} else {
		flag.StringVar(&pcapIn, "pcap", "", "pcap input filename")
		flag.StringVar(&pcapOut, "out", "", "pcap output filename")
		flag.Parse()
		Exec(pcapIn, pcapOut)
	}
}