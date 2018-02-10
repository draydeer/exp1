package main

import (
	"testing"
	"envd/src/drivers/memory"
	"envd/src/drivers"
	"envd/src/router"
	"envd/src/cache"
	"envd/src"
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

	for n := 0; n < b.N; n++ {
		core.GetValue("memory.a.b", 0)
	}
}
