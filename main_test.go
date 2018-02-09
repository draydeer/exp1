package main

import (
	"testing"
	"envd/src/ads"
)

func BenchmarkSelfConcatOperator1000(b *testing.B) {
	var t = ads.NewAdsMap(map[string]interface{}{"a": map[string]interface{}{"b": 2, "c": []interface{}{1,2,3}}})

	for n := 0; n < b.N; n++ {
		ads.Get(t, "a..b", 6)
	}
}
