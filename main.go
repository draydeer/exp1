package main

import (
	"fmt"
	"envd/pkg/core"
	"envd/pkg/router_manager"
	"envd/pkg/drivers/memory"
	"envd/pkg/local_cache"
	"envd/pkg/drivers/consul"
	"envd/pkg/driver_manager"
	"envd/pkg/drivers/environment"
	//"encoding/json"
	"envd/pkg/server/rest"
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
	var s3 = environment.NewEnvironmentDriver()

	var newd = driver_manager.NewDriverManager()
	var newr = router_manager.NewRouter()
	var news = local_cache.NewCache()

	var core = core.NewCore(&newd, &newr, &news)

	core.GetRouterManager().AddTrimmedPrefixMatch(s1, "", "memory:")
	core.GetRouterManager().AddTrimmedPrefixMatch(s2, "config/default","consul:")
	core.GetRouterManager().AddTrimmedPrefixMatch(s2, "","csl:")
	core.GetRouterManager().AddTrimmedPrefixMatch(s3, "","environment:")

	fmt.Println("map:", core)

	//fmt.Println(core.GetKey("memory:a.c.2.4", 0))
	//fmt.Println(core.GetKey("memory:a.b", 0))
	//fmt.Println(core.GetKey("memory:a\\.a.c", 0))
	//fmt.Println(core.GetKey("consul:config.ft-dev.application.ft\\.dcache\\.host", 0))
	//fmt.Println(core.GetKey("environment:abc", 0))
	//a, _ := core.GetKey("memory:a", 0)
	//b, _ := json.Marshal(a)
	//
	//fmt.Println(string(b))

	rest.RunServer(core)

	//var t = ads.NewAdsMap(map[string]interface{}{"a": map[string]interface{}{"b": 2, "c": []interface{}{1,2,3}}})
	//
	//fmt.Println("ads:", ads.GetKey(t, "a.c.1", 6))
}
