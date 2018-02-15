package main

import (
	"testing"
	"envd/src/drivers/memory"
	"envd/src/drivers"
	"envd/src/router"
	"envd/src/cache"
	"envd/src"
	"envd/src/lib"
	"fmt"
	//"strings"
)

func BenchmarkSelfConcatOperator1000(b *testing.B) {
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

	var splitter = lib.SplitKeyFn(".")
	var str = "aaasdf.bf.asdf.casd.fasd.dfasdf.eas.ffadsfa.sdf.gasdf.hfasfa.sdfadsf"

	fmt.Print(splitter(str))

	for n := 0; n < b.N; n++ {
		core.GetKey("memory.a.b", 0)
		//splitter(str)
	}
}