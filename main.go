package main

import (
	"fmt"
	"envd/src/drivers"
	"envd/src"
	"envd/src/ads"
	"envd/src/router"
	"envd/src/drivers/memory"
	"envd/src/cache"
)

func main() {
	var s1 = memory.NewMemoryDriverWithKeys(map[string]interface{}{
		"a": map[string]interface{}{
			"b": 2,
			"c": []interface{}{1,2,3},
		},
	})

	var newd = drivers.NewDriverManager()
	var newr = router.NewRouter()
	var news = cache.NewCache()

	var core = src.NewCore(&newd, &newr, &news)

	core.GetRouter().AddMatchByTrimmedPrefix(&s1, "memory.")

	fmt.Println("map:", core)

	fmt.Println(core.GetKey("memory.a.c.2.4", 0))
	fmt.Println(core.GetKey("memory.a.b", 0))

	var t = ads.NewAdsMap(map[string]interface{}{"a": map[string]interface{}{"b": 2, "c": []interface{}{1,2,3}}})

	fmt.Println("ads:", ads.GetKey(t, "a.c.1", 6))
}
