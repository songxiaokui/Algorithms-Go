package homework

import "testing"

/*
@Time    : 2021/3/10 20:37
@Author  : austsxk
@Email   : austsxk@163.com
@File    : 8. valid-anagram_test.go.go
@Software: GoLand
*/

func TestIsAnagram(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"normal", struct {
			s string
			t string
		}{s: "abbabba", t: "bbabbaa"},
			true,
		},
		{"length", struct {
			s string
			t string
		}{s: "abba", t: "bbabbaa"},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsAnagram(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("IsAnagram() = %v, want %v", got, tt.want)
			}
		})
	}
}
