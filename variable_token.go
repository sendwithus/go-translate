package translate

import (
	"fmt"
	"strings"
)

type VariableToken struct {
	s string
}

func (st VariableToken) toString() string {
	s := strings.TrimLeft(st.s, " ")
	s = strings.TrimRight(s, " ")
	return fmt.Sprintf("%%(%s)s", s)
}
