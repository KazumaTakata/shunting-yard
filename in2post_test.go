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
	operators = append(operators, Operator{Value: '|', Precedence: 0, IsLeftAssociative: true})
	operators = append(operators, Operator{Value: '.', Precedence: 1, IsLeftAssociative: true})
	operators = append(operators, Operator{Value: '+', Precedence: 2, IsLeftAssociative: true})
	operators = append(operators, Operator{Value: '-', Precedence: 2, IsLeftAssociative: true})
	operators = append(operators, Operator{Value: '*', Precedence: 3, IsLeftAssociative: true})
	operators = append(operators, Operator{Value: '/', Precedence: 3, IsLeftAssociative: true})
	operators = append(operators, Operator{Value: '^', Precedence: 4, IsLeftAssociative: false})

	i2p := NewIn2Post(operators, false)
	output := i2p.Parse("(1|3)")

	expect := []byte("13|")
	if !byte_array_equal(output, expect) {
		//t.Errorf("In2Post expected:%s, got:%s", expect, output)
	}

}
