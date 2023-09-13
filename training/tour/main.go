package main

import (
	"fmt"
	"golang.org/x/exp/slices"
	"math"
	"runtime"
	"training/tour/model"
	"training/tour/service"
)

func Add(x, y int) int {
	return x + y
}

func main() {
	fmt.Println("ok")

	nilai := Add(10, 2)
	fmt.Println(nilai)

	fmt.Println("\n- 2. for -")
	for i := 0; i < 10; i++ {
		fmt.Println("index ke-", i)
	}

	fmt.Println("-3. if -")
	if x := 10; x == 10 {
		fmt.Println("x adalah 10")
	}

	fmt.Println("\n- 4. switch -")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("macos")
	case "linux":
		fmt.Println("linux")
	default:
		fmt.Println(os)
	}

	defer func() {
		fmt.Println("\nakhir")
	}()

	// create object user
	person := &model.User{
		X: 10,
		Y: 20,
	}

	fmt.Println(*person)

	// == array ==
	var a [2]string
	a[0] = "reo"

	c := make(map[int]int)
	c[0] = 2

	fmt.Println(a[0])
	fmt.Println(c)

	// == slice ==
	fmt.Println("\n- 3. slice -")
	mySlice := []string{"satu", "dua"}
	mySlice = append(mySlice, "tiga")
	fmt.Println(mySlice)

	fmt.Println("  delete slice")
	mySlice = slices.Delete(mySlice, 1, 2)
	fmt.Println(mySlice)

	// == map ==
	fmt.Println("\n- 4. map -")
	myMap := map[string]*model.User{
		"satu": &model.User{
			X: 1,
			Y: 1,
		},
		"dua": &model.User{
			X: 2,
			Y: 2,
		},
	}
	fmt.Println(myMap)
	v, ok := myMap["empat"]
	fmt.Println("value:", v, "present:", ok)

	// == function values ==
	fmt.Println("\n- 5. function values -")
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}

	fmt.Println(hypot(3, 4))
	fmt.Println(Compute(hypot))

	// == 6. method ==
	fmt.Println("\n- 6. method -")
	userService := service.NewUserService("account")
	serviceName := userService.PrintName()
	fmt.Println("service name:", serviceName)

	// == 7. package stringer ==

	// == 8. channel ==
	ch := make(chan int, 4)
	ch <- 1
}

// create function
func Compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}
