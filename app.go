package main

import (
	"fmt"
	"os"
	"os/exec"
)

func checkDirectory() {
	_, err := os.Stat("split")
	if os.IsNotExist(err) {
		fmt.Println("Direktori split tidak ada.")
		fmt.Println("Membuat direktori baru.")
		_, err := exec.Command("mkdir", "split").Output()
		if err != nil {
			fmt.Println("Terjadi error ketika membuat direktori split.")
			fmt.Println(err)
			os.Exit(1)
		} else {
			fmt.Println("Direktori split berhasil dibuat.")
		}
	} else {
		fmt.Println("Direktori split sudah ada.")
	}
}

func splitPcap(in string, out string) {
	output, err := exec.Command("editcap", "-i", "3600", in, "split/"+out).Output()
	if err != nil {
		fmt.Println("Error when spliting pcap file.")
		fmt.Println(string(output[:]))
		os.Exit(1)
	}
	fmt.Println("Pcap is splited successfully.")
}

// Exec digunakan untuk menjalankan program utama
func Exec(pcapIn string, pcapOut string) {

	// Check directory for output location
	checkDirectory()

	// Split pcap file
	splitPcap(pcapIn, pcapOut)
}