package algo_test

import (
	"algo/algo"
	"testing"
)

type subs map[string]int
var sMap map[string]subs
var benchStr []byte

func init() {
	sMap = map[string]subs{
		"":{"":1},
		"ababccabc": {"abc":1, "cab":1},
		"abcdefghi": {"abcdefghi":1},
		"aaaaaaaaa": {"a":1},
		"aba":{"ab":1,"ba":1},
		"ajsdjfjaiwjdsahsdahfjjhhajf87617813nndsanfadadjajdfasndfnffffdasdfsadewuowuthvqwaadjaggagagwadagag":{"hajf8761":1},
	}

	benchStr = []byte("ajsdjfjaiwjdsahsdahfjjhhajf87617813nndsanfadadjajdfasndfnffffdasdfsadewuowuthvqwaadjaggagagwadagag")
}

func TestLongestUnrepeatedSubStr1(t *testing.T) {
	for k, v := range sMap {
		rst := string(algo.LongestUnrepeatedSubStr1([]byte(k)))
		if _, ok := v[rst]; !ok {
			t.Errorf("str:%s,expect:%v,got:%s\n", k, v, rst)
		}
	}
}

func TestLongestUnrepeatedSubStr2(t *testing.T) {
	for k, v := range sMap {
		rst := string(algo.LongestUnrepeatedSubStr2([]byte(k)))
		if _, ok := v[rst]; !ok {
			t.Errorf("str:%s,expect:%v,got:%s\n", k, v, rst)
		}
	}
}

func BenchmarkLongestUnrepeatedSubStr1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		algo.LongestUnrepeatedSubStr1(benchStr)
	}
}

func BenchmarkLongestUnrepeatedSubStr2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		algo.LongestUnrepeatedSubStr2(benchStr)
	}
}