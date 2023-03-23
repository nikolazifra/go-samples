package main

import (
	"reflect"
	"testing"
)

func Test_IsPalindromeEmpty(t *testing.T) {
	want := false
	var s string
	if got := IsPalindrome(s, 0, 1); got != want {
		t.Errorf("IsPalindrome want %v but got %v", want, got)
	}
}

func Test_IsPalindromeSingle(t *testing.T) {
	want := true
	s := "a"
	if got := IsPalindrome(s, 1, 1); got != want {
		t.Errorf("IsPalindrome want %v but got %v", want, got)
	}
}

func Test_IsPalindromeFirstAndLast(t *testing.T) {
	want := true
	s := "aa"
	if got := IsPalindrome(s, 0, 1); got != want {
		t.Errorf("IsPalindrome want %v but got %v", want, got)
	}
}

func Test_IsPalindromeLoopFirstAndLast(t *testing.T) {
	want := true
	s := "aa"
	if got := IsPalindromeLoop(s); got != want {
		t.Errorf("IsPalindrome want %v but got %v", want, got)
	}
}

func Test_IsPalindromeThreeSuccess(t *testing.T) {
	want := true
	s := "aba"
	if got := IsPalindrome(s, 0, len(s)-1); got != want {
		t.Errorf("IsPalindrome want %v but got %v", want, got)
	}
}

func Test_IsPalindromeLoopThreeSuccess(t *testing.T) {
	want := true
	s := "aba"
	if got := IsPalindromeLoop(s); got != want {
		t.Errorf("IsPalindromeLoop want %v but got %v", want, got)
	}
}

func Test_IsPalindromeThreeFail(t *testing.T) {
	want := false
	s := "abc"
	if got := IsPalindrome(s, 0, len(s)-1); got != want {
		t.Errorf("IsPalindrome want %v but got %v", want, got)
	}
}

func Test_IsPalindromeLoopThreeFail(t *testing.T) {
	want := false
	s := "abc"
	if got := IsPalindromeLoop(s); got != want {
		t.Errorf("IsPalindromeLoop want %v but got %v", want, got)
	}

}

func TestSwap(t *testing.T) {
	want := []rune("ba")
	in := []rune("ab")
	swap(in, 0, 1)
	if !reflect.DeepEqual(in, want) {
		t.Errorf("Reverse want %v but got %v", want, in)
	}
}

func TestReverseSingle(t *testing.T) {
	in := []byte("a")
	want := []byte("a")
	Reverse(in)
	if !reflect.DeepEqual(in, want) {
		t.Errorf("Reverse in %s but want %s", in, want)
	}
}

func TestReverseTwo(t *testing.T) {
	in := []byte("ab")
	want := []byte("ba")
	Reverse(in)
	if !reflect.DeepEqual(in, want) {
		t.Errorf("Reverse in %s but want %s", in, want)
	}
}

func TestReverseThree(t *testing.T) {
	in := []byte("abc")
	want := []byte("cba")
	Reverse(in)
	if !reflect.DeepEqual(in, want) {
		t.Errorf("Reverse in %s but want %s", in, want)
	}
}
