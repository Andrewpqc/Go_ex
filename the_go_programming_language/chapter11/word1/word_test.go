package word

import(
	"testing"
)


func TestIsPalindrome(t *testing.T){
	if !IsPalindrome("abccba"){
		t.Error(`IsPalindrome("abccba") = false`)
	}
	if !IsPalindrome("ccc"){
		t.Error(`IsPalindrome("ccc") = false`)
	}
}


func TestNonPalindrome(t *testing.T){
	if IsPalindrome("adv"){
		t.Error(`IsPalindrome("adv")=true`)
	}
}