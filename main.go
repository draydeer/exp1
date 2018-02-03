package main

import (
	"fmt"
	"envd/src/drivers"
	"envd/src/core"
	"envd/src/router"
)

func main() {
	var s1 drivers.Driver

	s1 = drivers.DriverInstance{A: 4}

	fmt.Println("map:", s1.Test())
	fmt.Println("map:", s1.Get("1", 2))

	var core = core.CoreInstance{
		DriverManager: drivers.NewDriverManager(make(map[string]interface{})),
		Router: router.NewRouter(),
	}

	fmt.Println("map:", core)
}
