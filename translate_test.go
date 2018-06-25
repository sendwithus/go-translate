package translate

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTranslations(t *testing.T) {
	checkTrans(t, "hello,", "hello,", false)
	checkTrans(t, "hello {{ name  }},", "hello %(name)s,", true)
	checkTrans(t, "hello {{ name  }} and {{ otherName }},", "hello %(name)s and %(otherName)s,", true)
}

func checkTrans(t *testing.T, input, expected string, containsVar bool) {
	trans := NewJinjaTranslation(input)
	assert.Equal(t, expected, trans.toString())
	assert.Equal(t, containsVar, trans.hasVars)
}
