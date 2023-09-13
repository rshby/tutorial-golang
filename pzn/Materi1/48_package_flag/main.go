package main

import (
	"flag"
	"fmt"
)

func main() {
	var host *string = flag.String("host", "localhost", "Put Your Host")
	var user *string = flag.String("user", "root", "Put yout database user")
	var password *string = flag.String("password", "", "put yout database password")

	flag.Parse()

	fmt.Println("Host", *host)
	fmt.Println("User", *user)
	fmt.Println("Password", *password)
}
