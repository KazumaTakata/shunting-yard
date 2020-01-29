package shunting

import (
	"testing"
)

func byte_array_equal(a, b []byte) bool {

	if len(a) != len(b) {
		return false
	}

	for i, a_e := range a {
		if b[i] != a_e {
			return false
		}
	}
	return true
}

func TestInfix2Post(t *testing.T) {

	operators := []Operator{}
	operators = append(operators, Operator{value: '.', precedence: 1, isLeftAssociative: true})
	operators = append(operators, Operator{value: '+', precedence: 2, isLeftAssociative: true})
	operators = append(operators, Operator{value: '-', precedence: 2, isLeftAssociative: true})
	operators = append(operators, Operator{value: '*', precedence: 3, isLeftAssociative: true})
	operators = append(operators, Operator{value: '/', precedence: 3, isLeftAssociative: true})
	operators = append(operators, Operator{value: '^', precedence: 4, isLeftAssociative: false})

	i2p := NewIn2Post(operators)
	output := i2p.parse("s+.a+")

	expect := []byte("s+a+.")
	if !byte_array_equal(output, expect) {
		//t.Errorf("In2Post expected:%s, got:%s", expect, output)
	}

}
