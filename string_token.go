package translate

type StringToken struct {
	s string
}

func (st StringToken) toString() string {
	return st.s
}
