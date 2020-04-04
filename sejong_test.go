package sejong

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_getDict(t *testing.T) {
	dict, err := getDict([]string{})
	assert.Nil(t, dict)
	assert.Error(t, err)

	dict, err = getDict([]string{"foo", "bar", "john", "doe"})
	assert.Equal(t, map[string]string{"foo": "bar", "john": "doe"}, dict)
	assert.Nil(t, err)

	dict, err = getDict([]string{"foo"})
	assert.Nil(t, dict)
	assert.Error(t, err)
}

func Test_getSentence(t *testing.T) {
	var entry interface{}
	var dict map[string]string

	// just string
	entry = "something to be translated"
	dict = nil
	sentence, err := getSentence(entry, dict)
	assert.Equal(t, "something to be translated", sentence)
	assert.Nil(t, err)

	entry = map[string]interface{}{
		"zero":  "there's nothing",
		"one":   "there's something",
		"other": "there're %{count} somethings",
	}

	// zero
	dict = map[string]string{"count": "0"}
	sentence, err = getSentence(entry, dict)
	assert.Equal(t, "there's nothing", sentence)
	assert.Nil(t, err)
	// one
	dict = map[string]string{"count": "1"}
	sentence, err = getSentence(entry, dict)
	assert.Equal(t, "there's something", sentence)
	assert.Nil(t, err)
	// more
	dict = map[string]string{"count": "2"}
	sentence, err = getSentence(entry, dict)
	assert.Equal(t, "there're %{count} somethings", sentence)
	assert.Nil(t, err)
	// nil dict
	dict = nil
	sentence, err = getSentence(entry, dict)
	assert.Empty(t, sentence)
	assert.Error(t, err)
	// no count
	dict = map[string]string{"foo": "bar"}
	sentence, err = getSentence(entry, dict)
	assert.Empty(t, sentence)
	assert.Error(t, err)
}
