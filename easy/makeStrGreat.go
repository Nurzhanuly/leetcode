package easy

func MakeGood(s string) string {

	if len(s) < 0 {
		return s
	}
	var st stack
	st.push(s[0])
	for i := 1; i < len(s); i++ {
		if st.top() == s[i]-32 || s[i] == st.top()-32 {
			st.pop()
		} else {
			st.push(s[i])
		}
	}
	return string(st)
}

type stack []byte

func (s *stack) pop() {
	*s = (*s)[:len(*s)-1]
}

func (s *stack) top() byte {
	if len(*s) == 0 {
		return 0
	}
	return (*s)[len(*s)-1]
}

func (s *stack) push(x byte) {
	*s = append(*s, x)
	// if len(s.stack) ==0 || x <= stack
}
