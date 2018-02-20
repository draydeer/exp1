package main

import (
	"testing"
	"envd/pkg/drivers/memory"
	"envd/pkg/drivers"
	"envd/pkg/routes"
	"envd/pkg/local_cache"
	"envd/pkg"
	"envd/pkg/lib"
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
	var newr = routes.NewRouter()
	var news = local_cache.NewCache()

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
