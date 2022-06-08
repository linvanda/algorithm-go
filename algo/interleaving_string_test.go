package algo_test

import (
	"algo/algo"
	"testing"
)

var interSamples = [][3]string{
	{"aabcc", "dbbca", "aadbbcbcac"},
	{"aabcc", "dbbca", "aadbbbaccc"},
	{"", "", ""},
	{"","b","b"},
}
var interResults = []bool{true, false, true, true}

func TestIsInterleavingString(t *testing.T) {
	for i, v := range interSamples {
		ok := algo.IsInterleavingString([]byte(v[0]), []byte(v[1]),[]byte(v[2]))
		if ok != interResults[i] {
			t.Errorf("sample:%v, expect:%v, got:%v\n", v, interResults[i], ok)
		}
	}
}
