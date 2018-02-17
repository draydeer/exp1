package main

import (
	"fmt"
	"envd/src/core"
	"envd/src/router"
	"envd/src/drivers/memory"
	"envd/src/local_cache"
	"envd/src/drivers/consul"
	"envd/src/driver_manager"
	"envd/src/drivers/env"
)

func main() {
	var s1 = memory.NewMemoryDriverWithKeys(map[string]interface{}{
		"a": map[string]interface{}{
			"b": 2,
			"c": []interface{}{1,2,3},
		},
		"a.a": map[string]interface{}{
			"b": 3,
			"c": []interface{}{4,5,6},
		},
	})

	var s2 = consul.NewConsulDriver()
	var s3 = env.NewEnvDriver()

	var newd = driver_manager.NewDriverManager()
	var newr = router.NewRouter()
	var news = local_cache.NewCache()

	var core = core.NewCore(&newd, &newr, &news)

	core.GetRouter().AddTrimmedPrefixMatch(s1, "memory.")
	core.GetRouter().AddTrimmedPrefixMatch(s2, "consul.")
	core.GetRouter().AddTrimmedPrefixMatch(s3, "env.")

	fmt.Println("map:", core)

	fmt.Println(core.GetKey("memory.a.c.2.4", 0))
	fmt.Println(core.GetKey("memory.a.b", 0))
	fmt.Println(core.GetKey("memory.a\\.a.c", 0))
	fmt.Println(core.GetKey("consul.test.x", 0))
	fmt.Println(core.GetKey("env.abc", 0))

	//var t = ads.NewAdsMap(map[string]interface{}{"a": map[string]interface{}{"b": 2, "c": []interface{}{1,2,3}}})
	//
	//fmt.Println("ads:", ads.GetKey(t, "a.c.1", 6))
}
