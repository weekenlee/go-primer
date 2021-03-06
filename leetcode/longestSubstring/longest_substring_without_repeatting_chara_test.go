package main
import (
	"testing"
	"github.com/stretchr/testify/assert"
)

type para struct {
	one string
}

type ans struct {
	one int
}

type question struct {
	p para
	a ans
}

func Test_OK(t *testing.T) {
	ast := assert.New(t)

	qs := []question {
		question {
			p : para {
				one : "abcabccb",	
			}, 
			a : ans {
				one : 3,
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.one, lengthOfLongestSubstring(p.one), "输入:%v", p)
	}
}
