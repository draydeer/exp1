package main

import (
	"testing"
	"envd/pkg/drivers/memory"
	"envd/pkg/driver_manager"
	"envd/pkg/core"
	"envd/pkg/local_cache"
	"envd/pkg/router"
	//"envd/pkg/lib"
	"fmt"
)

func BenchmarkSelfConcatOperator1000(b *testing.B) {
	var s1 = memory.NewMemoryDriverWithKeys(map[string]interface{}{
		"a": map[string]interface{}{
			"b": 2,
			"c": []interface{}{1,2,3},
		},
	})

	var newd = driver_manager.NewDriverManager()
	var newr = router.NewRouter()
	var news = local_cache.NewCache()

	var core = core.NewCore(&newd, &newr, &news)

	core.GetRouter().AddTrimmedPrefixMatch(s1, "memory.")

	//var splitter = lib.SplitKeyFn(".")
	//var str = "aaasdf.bf.asdf.casd.fasd.dfasdf.eas.ffadsfa.sdf.gasdf.hfasfa.sdfadsf"
	//
	//fmt.Print(splitter(str))

	for n := 0; n < b.N; n++ {
		core.GetKey("memory.a.c.2", 0)
		//splitter(str)
	}

	fmt.Println(core.GetKey("memory.a.c.2", 0))
}
