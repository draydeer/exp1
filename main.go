package main

import (
	"fmt"
	"envd/src/drivers"
	"envd/src"
)

func main() {
	var s1 drivers.Driver

	s1 = drivers.DriverInstance{}

	fmt.Println("map:", s1.Get("1", 2))

	var core = src.NewCore()

	fmt.Println("map:", core)
}
