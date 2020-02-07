package shunting

func isNumber(ch byte) bool {
	if ch >= '0' && ch <= '9' {
		return true
	}

	return false
}

func isAlphabet(ch byte) bool {
	if (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') {
		return true
	}

	return false

}

func isWord(ch byte) bool {
	if isNumber(ch) || isAlphabet(ch) || ch == '_' {
		return true
	}

	return false
}

type stack struct {
	stack []byte
}

func (s *stack) pop() byte {
	last := s.stack[len(s.stack)-1]
	s.stack = s.stack[:len(s.stack)-1]
	return last
}

func (s *stack) push(ch byte) {
	s.stack = append(s.stack, ch)
}

func (s *stack) top() byte {
	return s.stack[len(s.stack)-1]
}

func (s *stack) empty() bool {
	if len(s.stack) == 0 {
		return true
	}

	return false
}

type In2Post struct {
	Precedence       map[byte]int
	rightAssociative []byte
	add_paren        bool
}

func NewIn2Post(operators []Operator, add_paren bool) In2Post {
	precedence := map[byte]int{}
	rightAssociative := []byte{}
	for _, operator := range operators {
		precedence[operator.Value] = operator.Precedence
		if !operator.IsLeftAssociative {
			rightAssociative = append(rightAssociative, operator.Value)
		}
	}

	in2post := In2Post{Precedence: precedence, rightAssociative: rightAssociative, add_paren: add_paren}
	return in2post

}

type Operator struct {
	Value             byte
	Precedence        int
	IsLeftAssociative bool
}

func (in2post *In2Post) isLeftAssociative(ch byte) bool {
	for _, right_op := range in2post.rightAssociative {
		if right_op == ch {
			return false
		}
	}
	return true
}

func (in2post *In2Post) isOperator(ch byte) bool {
	for op, _ := range in2post.Precedence {
		if op == ch {
			return true
		}

	}
	return false
}
func (in2post *In2Post) Parse(input string) []byte {

	stack := stack{}
	output := []byte{}

	for len(input) > 0 {
		if isWord(input[0]) {
			output = append(output, input[0])
			input = input[1:]
		} else if in2post.isOperator(input[0]) {
			for !stack.empty() && (in2post.Precedence[input[0]] < in2post.Precedence[stack.top()] || (in2post.Precedence[input[0]] == in2post.Precedence[stack.top()] && in2post.isLeftAssociative(input[0]))) {
				if stack.top() == '(' {
					break
				}
				output = append(output, stack.pop())
			}

			stack.push(input[0])
			input = input[1:]
		} else if input[0] == '(' {
			stack.push(input[0])
			input = input[1:]
			if in2post.add_paren {
				output = append(output, '(')
			}
		} else if input[0] == ')' {
			token := stack.pop()
			for token != '(' {
				output = append(output, token)
				token = stack.pop()
			}
			input = input[1:]
			if in2post.add_paren {
				output = append(output, ')')
			}
		} else if input[0] == '\\' {
			output = append(output, input[:2]...)
			input = input[2:]
		}

	}

	for !stack.empty() {
		output = append(output, stack.pop())
	}

	return output
}
