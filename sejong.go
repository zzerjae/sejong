// Copyright © 2020 Jaeho Cho <jaeho8032@gmail.com>.
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

// sejong is an simple translate library.
// it was inspired by ruby's i18n module.

package sejong

import (
	"errors"
	"fmt"
	"regexp"
	"strings"

	"github.com/spf13/viper"
)

type Sejong struct {
	Locale string
	V *viper.Viper
	configured []string
}

var sj *Sejong
var Locale string

func init() {
	sj, _ = New("")
}

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

	var configured []string
	if locale != "" {
		v.SetConfigName(locale)
		if err := v.ReadInConfig(); err != nil {
			return nil, errors.New("sejong.New: locale file is not exist")
		}
		configured = append(configured, locale)

	}

	return &Sejong{
		V:          v,
		configured: configured,
		Locale:     locale,
	}, nil
}

func T(key string, words ...string) (string, error) {
	sj.Locale = Locale
	return sj.T(key, words...)
}

func (s *Sejong) T(key string, words ...string) (string, error) {
	if s.Locale == "" {
		return "", errors.New("sejong.T: locale is not set")
	}

	var ok bool
	for _, l := range s.configured {
		if s.Locale == l {
			ok = true
			break
		}
	}
	if !ok {
		s.V.SetConfigName(s.Locale)
		if err := s.V.ReadInConfig(); err != nil {
			panic("sejong.T: locale file is not exist")
		}
		s.configured = append(s.configured, s.Locale)
	}
	str := s.V.GetString(s.Locale + "." + key)

	if len(words)%2 == 1 {
		return str, errors.New("sejong.T: odd word count")
	}
	r := newReplacer(words)

	result := r.Replace(str)
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
