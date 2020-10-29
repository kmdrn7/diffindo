package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

var (
	allPcaps []string
	allOutput []string
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

func listPcap() {
	reader := bufio.NewScanner(os.Stdin)
	fmt.Println("Menampilkan file pcap yang ada pada direktori saat ini.")
	fmt.Println("===========================================================================")
	notfound := true
	pcaps, err := ioutil.ReadDir(".")
	if err != nil {
		fmt.Println("Terjadi kesalahan saat menampilkan list pcap pada direktori saat ini.")
		os.Exit(1)
	}
	for idx, pcap := range pcaps {
		if filepath.Ext(pcap.Name()) == ".pcap" {
			notfound = false
			fmt.Printf("[%d] %s\n", idx, pcap.Name())
			fmt.Print("Apakah anda ingin memproses pcap ini? [y/n]: ")
			reader.Scan()
			response := reader.Text()
			response = strings.ToLower(strings.TrimSpace(response))
			if response == "y" || response == "yes" {
				fmt.Print("Masukkan nama output untuk pcap ini: ")
				reader.Scan()
				response := reader.Text()
				allPcaps = append(allPcaps, pcap.Name())
				allOutput = append(allOutput, response)
			}
			fmt.Println("===========================================================================")
		}
	}
	if notfound == true {
		fmt.Println("Tidak ada file pcap yang terdeteksi pada direktori saat ini.")
	}
}

func splitPcap() {
	fmt.Printf("\nMemproses semua pcap.\n")
	for idx, pcap := range allPcaps {
		fmt.Printf("[%d] %s -> %s\n", idx, pcap, allOutput[idx])
		output, err := exec.Command("editcap", "-i", "3600", pcap, "split/" + allOutput[idx]).CombinedOutput()
		if err != nil {
			fmt.Println(err)
			fmt.Println("Terjadi kesalahan ketika membagi file pcap.")
			fmt.Println(string(output[:]))
			os.Exit(1)
		}
	}
	fmt.Println("Semua pcap telah berhasil dibagi.")
}

// Exec digunakan untuk menjalankan program utama
func Exec() {

	// Cek apakah direktori untuk hasil output sudah ada
	checkDirectory()

	// Menampilkan semua file pcap yang ada pada direktori saat ini
	listPcap()

	// Membagi file pcap kedalam beberapa bagian
	splitPcap()
}