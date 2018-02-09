package main

import (
	"fmt"
	"envd/src/drivers"
	"envd/src"
	"envd/src/ads"
	"envd/src/router"
	"envd/src/drivers/memory"
)

func main() {
	var s1 = memory.NewMemoryDriver()

	var newd = drivers.NewDriverManager()
	var newr = router.NewRouter()

	var core = src.NewCore(&newd, &newr)

	core.GetRouter().AddMatchByPrefix(&s1, "memory.")

	fmt.Println("map:", core)

	var a, b, c = core.GetRouter().Test("memory.a.b.c")

	fmt.Println("test", a, b, c)

	var t = ads.NewAdsMap(map[string]interface{}{"a": map[string]interface{}{"b": 2, "c": []interface{}{1,2,3}}})

	fmt.Println("ads:", ads.Get(t, "a.c.1", 6))
}
