package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"runtime"
	"strings"
	"sync"
)

func cURL(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func checkSite(username string, waitgroup *sync.WaitGroup) {
	var situs = [...]string{
		"https://instagram.com/%s/?__a=1",
	}
	for i := 0; i < len(situs); i++ {
		response, err := cURL(fmt.Sprintf(situs[i], username))
		if err != nil {
			fmt.Printf("[ERROR] %s %s", fmt.Sprintf(situs[i], username), err)
			waitgroup.Done()
		} else {
			if strings.Index(response, fmt.Sprintf(`"username":"%s","`, username)) > -1 {
				fmt.Println("[200 OK] Valid, sudah ada yang memakainya.")
				waitgroup.Done()
			} else {
				fmt.Println("[200 OK] Tidak valid, tidak ada yang memakainya.")
				waitgroup.Done()
			}
		}
	}
}

func main() {
	runtime.GOMAXPROCS(2)
	defer fmt.Printf("\nTerimakasih sudah memakai tools ini ^_^")

	var username string
	var waitgroup sync.WaitGroup
	waitgroup.Add(1)

	fmt.Println("[+] Username Check Avability.")
	fmt.Println("[!] Ditulis dengan Go")
	fmt.Printf("[+] Dibuat pada 26 May 2016\n\n")

	fmt.Print("[>] Masukkan username anda : ")
	fmt.Scanln(&username)

	if username == "" {
		fmt.Println("[!] Anda harus memasukkan username.")
		os.Exit(0)
	}
	fmt.Printf("[>] Memakai username %s untuk dicek.\n", username)
	fmt.Printf("[!] Proses..\n\n")

	go checkSite(username, &waitgroup)
	waitgroup.Wait()
}
