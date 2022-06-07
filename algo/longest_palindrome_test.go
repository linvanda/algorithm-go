package algo_test

import (
	"algo/algo"
	"testing"
)

func TestLongestPalindrome1(t *testing.T) {
	s1 := "ab1238f321add"// 9
	s2 := "abc7cbfa"// 7

	l1 := algo.LongestPalindrome1([]byte(s1))
	if l1 != 9 {
		t.Errorf("s1:%s test wrong,real:%d,get:%d\n", s1, 9, l1)
	}

	l2 := algo.LongestPalindrome1([]byte(s2))
	if l2 != 7 {
		t.Errorf("s2:%s test wrong,real:%d,get:%d\n", s2, 7, l2)
	}
}

func TestLongestPalindrome2(t *testing.T) {
	s1 := "ab1238f321add"// 9
	s2 := "abc7cbfa"// 7

	l1 := algo.LongestPalindrome2([]byte(s1))
	if l1 != 9 {
		t.Errorf("s1:%s test wrong,real:%d,get:%d\n", s1, 9, l1)
	}

	l2 := algo.LongestPalindrome2([]byte(s2))
	if l2 != 7 {
		t.Errorf("s2:%s test wrong,real:%d,get:%d\n", s2, 7, l2)
	}
}

func TestLongestPalindrome3(t *testing.T) {
	s1 := "ab1238f321add"// 9
	s2 := "abc7cbfa"// 7

	l1 := algo.LongestPalindrome3([]byte(s1))
	if l1 != 9 {
		t.Errorf("s1:%s test wrong,real:%d,get:%d\n", s1, 9, l1)
	}

	l2 := algo.LongestPalindrome3([]byte(s2))
	if l2 != 7 {
		t.Errorf("s2:%s test wrong,real:%d,get:%d\n", s2, 7, l2)
	}
}

func TestLongestPalindromeDP1(t *testing.T) {
	s1 := "ab1238f321add"// 9

	l1 := algo.LongestPalindromeDP1([]byte(s1))
	if l1 != 9 {
		t.Errorf("s1:%s test wrong,real:%d,get:%d\n", s1, 9, l1)
	}
}

func TestLongestPalindromeDP3(t *testing.T) {
	s1 := "ab1238f321add"// 9

	l1 := algo.LongestPalindromeDP3([]byte(s1))
	if l1 != 9 {
		t.Errorf("s1:%s test wrong,real:%d,get:%d\n", s1, 9, l1)
	}
}

func BenchmarkLongestPalindrome1(b *testing.B) {
	s1 := []byte("ab1238f321add")
	for i := 0; i < b.N; i++ {
		algo.LongestPalindrome1(s1)
	}
}

func BenchmarkLongestPalindrome3(b *testing.B) {
	s1 := []byte("ab1238f321add")
	for i := 0; i < b.N; i++ {
		algo.LongestPalindrome3(s1)
	}
}

func BenchmarkLongestPalindromeDP1(b *testing.B) {
	s1 := []byte("ab1238f321add")
	for i := 0; i < b.N; i++ {
		algo.LongestPalindromeDP1(s1)
	}
}

func BenchmarkLongestPalindromeDP3(b *testing.B) {
	s1 := []byte("ab1238f321add")
	for i := 0; i < b.N; i++ {
		algo.LongestPalindromeDP3(s1)
	}
}