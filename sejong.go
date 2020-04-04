// Copyright © 2020 Jaeho Cho <jaeho8032@gmail.com>.
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

// sejong is an simple translate library.
// it was inspired by ruby's i18n module.
package sejong

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/pkg/errors"

	"github.com/spf13/viper"
)

// Sejong handles locale configurations for translation.
type Sejong struct {
	// Locale contains the language to be translate.
	Locale string
	// v gets the location of the locale file from the environment variable.
	// v is also used to get the sentence using the key in each locale file.
	v *viper.Viper
	// locales is used when using translations in multiple languages
	// ​​at the same time to determine whether each locale file is loaded.
	locales []string
}

// Locale contains the language to be translate. It is used only if you don't create your Sejong instance.
var Locale string

var sj *Sejong

func init() {
	sj, _ = New("")
}

// New returns an initialized sejong instance.
func New(locale string) (*Sejong, error) {
	v := viper.New()
	v.SetEnvPrefix("SEJONG")
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	v.AutomaticEnv()

	path := v.GetString("locale.directory")
	if path == "" {
		path = "."
	}
	v.AddConfigPath(path)
	v.SetConfigType("yml")

	var locales []string
	if locale != "" {
		v.SetConfigName(locale)
		if err := v.ReadInConfig(); err != nil {
			return nil, errors.New("sejong.New: locale file is not exist")
		}
		locales = append(locales, locale)

	}

	return &Sejong{
		v:       v,
		locales: locales,
		Locale:  locale,
	}, nil
}

// T translate a sentence found by key.
// In words, Each word in turn becomes key, value, key, value, ...
// In the sentence, each %{key} is replaced with a value.
func T(key string, words ...string) (string, error) {
	sj.Locale = Locale
	return sj.T(key, words...)
}

// T translate a sentence found by key.
// Each word in turn becomes key, value, key, value, ...
// In the sentence, each %{key} is replaced with a value.
func (s *Sejong) T(key string, words ...string) (string, error) {
	if s.Locale == "" {
		return "", errors.New("sejong.T: locale is not set")
	}

	var ok bool
	for _, l := range s.locales {
		if s.Locale == l {
			ok = true
			break
		}
	}
	if !ok {
		s.v.SetConfigName(s.Locale)
		if err := s.v.ReadInConfig(); err != nil {
			panic("sejong.T: locale file is not exist")
		}
		s.locales = append(s.locales, s.Locale)
	}

	dict, err := getDict(words)
	if err != nil {
		return "", errors.Wrap(err, "sejong.T")
	}

	key = s.Locale + "." + key
	entry := s.v.Get(key)
	if entry == nil {
		return "", errors.New("sejong.T: not exist key")
	}

	sentence, err := getSentence(entry, dict)
	if err != nil {
		return "", errors.Wrap(err, "sejong.T")
	}

	r := newReplacer(words)

	result := r.Replace(sentence)
	if matched, _ := regexp.MatchString("%{\\w+}", result); matched {
		return result, errors.New("sejong.T: translate is not completed")
	}

	return result, nil
}

func newReplacer(words []string) *strings.Replacer {
	var keywords []string
	for i, word := range words {
		if i%2 == 0 {
			keywords = append(keywords, fmt.Sprintf("%%{%s}", word))
		} else {
			keywords = append(keywords, word)
		}
	}
	return strings.NewReplacer(keywords...)
}

func getDict(words []string) (map[string]string, error) {
	if len(words) == 0 || len(words)%2 == 1 {
		return nil, errors.New("sejong.getDict: odd word count")
	}
	dict := make(map[string]string, len(words)/2)
	for i := 0; i < len(words)-1; i += 2 {
		key := words[i]
		value := words[i+1]
		dict[key] = value
	}

	return dict, nil
}

func getSentence(entry interface{}, dict map[string]string) (string, error) {
	var sentence string
	switch entry := entry.(type) {
	case string:
		sentence = entry
	case map[string]interface{}:
		if dict != nil && dict["count"] != "" {
			count, err := strconv.ParseInt(dict["count"], 10, 64)
			if err != nil {
				return "", errors.Wrap(err, "sejong.T")
			}
			sentence, err = pluralize(entry, count)
			if err != nil {
				return "", errors.Wrap(err, "sejong.T")
			}
		} else {
			return "", errors.New("sejong.T: count should be provided")
		}
	}

	return sentence, nil
}

func pluralize(entryMap map[string]interface{}, count int64) (string, error) {
	strMap := make(map[string]string, 4)
	for k, v := range entryMap {
		strMap[k] = v.(string)
	}
	if len(strMap) == 0 {
		return "", errors.New("failed to pluralize")
	}
	key := pluralizationKey(strMap, count)

	if strMap[key] != "" {
		return strMap[key], nil
	} else {
		return "", errors.New("failed to pluralize")
	}
}

func pluralizationKey(strMap map[string]string, count int64) string {
	if count == 0 && strMap["zero"] != "" {
		return "zero"
	} else if count == 1 && strMap["one"] != "" {
		return "one"
	} else {
		return "other"
	}
}
