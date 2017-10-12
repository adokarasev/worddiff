package worddiff

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var strings = []struct {
	str1  string
	str2  string
	match string
}{
	{"baoocd", "gykaobocd", "aoocd"},
	{"dc", "bacd", "d"},
	{"abcd", "abcd", "abcd"},
	{"odkpqkdpqopjfwjioejfvsvspkopkc", "oekqowekcwoeicjewcwcoiejwoicjejwcojejeowjcwc", "okqkojwjioejoc"},
	{"qodkpqkdpqopjfwjioejfvsvspkopkc", "oekqowekcwoeicjewcwcoiejwoicjejwcojejeowjcwc", "okqkojwjioejoc"},
}

func TestWordDiff(t *testing.T) {
	for _, in := range strings {
		out := worddiff(in.str1, in.str2)
		assert.Equal(t, in.match, out, "they should be equal")
	}
}

func BenchmarkWordDiff(b *testing.B) {
	for i := 0; i < b.N; i++ {
		worddiff("odkpqkdpqopjfwjioejfvsvspkopkc", "oekqowekcwoeicjewcwcoiejwoicjejwcojejeowjcwc")
	}
}
