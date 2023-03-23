package main

func IsPalindrome(s string, l, r int) bool {
	if len(s) == 0 || l < 0 || r < 0 {
		return false
	}
	if len(s) == 1 || l >= r {
		return true
	}
	rn := []rune(s)
	if rn[l] == rn[r] {
		return IsPalindrome(s, l+1, r-1)
	}
	return false
}

func IsPalindromeLoop(s string) bool {
	l := 0
	r := len(s) - 1
	if len(s) == 1 || l >= r {
		return true
	}
	rn := []rune(s)
	for ; l <= r; l, r = l+1, r-1 {
		if rn[l] != rn[r] {
			return false
		}
	}
	return true
}

func swap(r []rune, i, j int) {
	tmp := r[i]
	r[i] = r[j]
	r[j] = tmp
}

func Reverse(s []byte) {
	/*if len(s) == 0 {
		return
	}
	Reverse(s[1:])
	c := s[0]
	for i := 0; i < len(s)-1; i++ {
		s[i] = s[i+1]
	}
	s[len(s)-1] = c
	*/
	if len(s) == 0 {
		return
	}
	for i, j := 0, len(s)-1; j > i; i, j = i+1, j-1 {
		tmp := s[i]
		s[i] = s[j]
		s[j] = tmp
	}

}
