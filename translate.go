package translate

import (
	"github.com/pkg/errors"
)

const (
	TextTokenType = iota
	VarTokenType
)

var EOFError = errors.New("EOF")

type Token interface {
	toString() string
}

type JinjaTranslation struct {
	original  string
	runes     []rune
	hasVars   bool
	tokens    []Token
	i         int
	currState uint
	buff      string
}

func NewJinjaTranslation(s string) *JinjaTranslation {
	t := JinjaTranslation{}
	t.tokenize(s)
	return &t
}

func (tr *JinjaTranslation) toString() string {
	s := ""
	for _, t := range tr.tokens {
		s += t.toString()
	}
	return s
}

func (tr *JinjaTranslation) peak() (rune, error) {
	return tr.getAt(tr.i)
}

func (tr *JinjaTranslation) get() (rune, error) {
	r, err := tr.getAt(tr.i)
	tr.i += 1
	return r, err
}

func (tr *JinjaTranslation) getAt(i int) (rune, error) {
	if i >= len(tr.runes) {
		return ' ', EOFError
	}
	return tr.runes[i], nil
}

func (tr *JinjaTranslation) thisAndNextAre(curr rune, expected rune) bool {
	if curr != expected {
		return false
	}
	n, _ := tr.peak()
	if n != expected {
		return false
	}
	return true
}

func (tr *JinjaTranslation) initTokenizer(s string) {
	tr.i = 0
	tr.original = s
	tr.runes = []rune(s)
	tr.currState = TextTokenType
	tr.buff = ""
}

func (tr *JinjaTranslation) flush() string {
	bufferContent := tr.buff
	tr.buff = ""
	return bufferContent
}

func (tr *JinjaTranslation) tokenize(s string) error {
	tr.initTokenizer(s)
	var err error
	var r rune
	for ; err == nil; r, err = tr.get() {
		if string(r) == "\x00" {
			continue
		}
		if err == EOFError {
			break
		}
		if tr.currState == TextTokenType {
			if tr.thisAndNextAre(r, '{') {
				tr.tokens = append(tr.tokens, StringToken{s: tr.flush()})
				tr.get()
				tr.currState = VarTokenType
				continue
			}
		}

		if tr.currState == VarTokenType {
			if tr.thisAndNextAre(r, '}') {
				tr.tokens = append(tr.tokens, VariableToken{s: tr.flush()})
				tr.get()
				tr.hasVars = true
				tr.currState = TextTokenType
				continue
			}
		}
		tr.buff += string(r)
	}
	if tr.currState != TextTokenType {
		return errors.New("unclosed variable")
	}
	tr.tokens = append(tr.tokens, StringToken{s: tr.flush()})
	return nil
}
