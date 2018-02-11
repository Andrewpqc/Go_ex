package word2

import(
	"testing"
)


func TestIsPalindrome(t *testing.T){
	var tests =[]struct{
		input string
		want bool
	}{
		{"",true},
		{"a",true},
		{"aa",true},
		{"ab",false},
		{"kayak",true},
		{"detartrated",true},
		{"A man, a plan, a canal: Panama",true},
		{"A mansadfgsadfg, a plan, a canal: Panama",false},
	}

	for _,test:=range tests{
		if got:=IsPalindrome(test.input);got!=test.want{
			t.Errorf("IsPalindrome(%s)=%v,want %v",test.input,got,test.want)
		}
	}
}