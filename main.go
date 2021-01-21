package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	resp, err := http.Get("example.com/auth/") // URL to webpage containing credentials
	if err != nil {
		fmt.Print("Error:", err.Error())
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Print("Error:", err.Error())
		return
	}
	allCombos := string(body)
	check(allCombos)
}

func check(allCombos string) {
	var (
		user string
		pass string
	)

	fmt.Print("User: ")
	fmt.Scan(&user)
	fmt.Print("Pass: ")
	fmt.Scan(&pass)

	creds := user + ":" + pass
	if strings.Contains(allCombos, creds) {
		fmt.Print("Correct info")
	} else {
		fmt.Print("Invalid info")
	}
}
