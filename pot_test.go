package translate

import (
	"io/ioutil"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPOT_ToString(t *testing.T) {
	for input, expected := range getTestGoldenContent(t) {
		translations := strings.Split(input, "=== translation ===\n")
		pot := POT{}
		for _, translation := range translations {
			translation = strings.Replace(translation, "ORIGINAL:", "", 1)
			translation = strings.TrimRight(translation, "\n")
			parts := strings.Split(translation, "\nTRANS:")
			assert.Equal(t, len(parts), 2)
			original := parts[0]
			translated := parts[1]
			pot.AddTrans(original, translated)
		}

		actual, err := pot.ToString()
		assert.NoError(t, err)
		assert.Equal(t, expected, actual)
	}
}

func getTestGoldenContent(t *testing.T) map[string]string {
	tests := map[string]string{}

	files, err := ioutil.ReadDir("./golden/input")
	if err != nil {
		t.Fatalf("unable to find golden file inputs")
	}
	for _, fileInfo := range files {
		inputPath := "./golden/input/" + fileInfo.Name()
		expectedPath := "./golden/expected/" + fileInfo.Name()
		input, err := ioutil.ReadFile(inputPath)
		if err != nil {
			t.Fatalf("unable to read input file: %s", inputPath)
		}
		expected, err := ioutil.ReadFile(expectedPath)
		if err != nil {
			t.Fatalf("unable to read output file: %s", inputPath)
		}
		tests[string(input)] = string(expected)
	}

	return tests
}
