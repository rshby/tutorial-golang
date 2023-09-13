package main

import (
	"fmt"
	"strings"
)

// membuat function login -> ada pengecekan apakah user 'admin' atau bukan
func Login(inputUser string, inputFilter func(string) bool) {
	if inputFilter(inputUser) {
		fmt.Println("failed to login. you are blocked because your role is admin")
	} else {
		fmt.Println("success to login")
	}
}

func main() {
	// == 1. create function ==

	// == 2. create variabel dengan value anonymous function ==
	blacklist := func(inputUser string) bool {
		return strings.ToLower(inputUser) == "admin"
	}

	fmt.Println("- 2. create variabel dengan value anonymout function -")

	userId := "usr_dev_team"
	fmt.Println("hasil dari filter blacklist:", blacklist(userId))
	fmt.Println("hasil login:")
	Login(userId, blacklist)

	// == 3. create anonymous function langsung di input parameter ==
	fmt.Println("\n- 3. create anonymous function langsung di input parameter -")
	Login("root", func(s string) bool {
		if strings.ToLower(s) == "admin" || strings.ToLower(s) == "root" {
			return true
		} else {
			return false
		}
	})
}
