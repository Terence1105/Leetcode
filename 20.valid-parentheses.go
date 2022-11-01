/*
* @lc app=leetcode id=20 lang=golang
*
  - [20] Valid Parentheses
  - Input: string s containing just the characters '(', ')', '{', '}', '[' and ']'
  - Output: Is valid(true or false)
  - 解法:
    step1: 透過Slice實作Stack
    step2: 將字元從頭push到slice裡面
    step3: 如果與前一個字元有對應到一組就pop出來({}/()/[]這樣是一組)
    時間複雜度: O(n)
    空間複雜度: O(n)
*/

// @lc code=start
func isValid(s string) bool {
	stack := &Stack{}

	for i := 0; i < len(s); i++ {
		if len(stack.str) == 0 {
			stack.Push(s[i])
		} else if stack.Peek() == Parentheses[s[i]] {
			stack.Pop()
		} else {
			stack.Push(s[i])
		}
	}

	if len(stack.str) == 0 {
		return true
	} else {
		return false
	}
}

var Parentheses = map[byte]byte{
	')': '(',
	'}': '{',
	']': '[',
}

type Stack struct {
	str []byte
}

func (s *Stack) Push(char byte) {
	s.str = append(s.str, char)
}

func (s *Stack) Pop() byte {
	lastChar := s.str[len(s.str)-1]
	s.str = s.str[:len(s.str)-1]
	return lastChar
}

func (s *Stack) Peek() byte {
	return s.str[len(s.str)-1]
}

// @lc code=end
