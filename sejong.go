// Copyright © 2020 Jaeho Cho <jaeho8032@gmail.com>.
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

// sejong is an simple translate library.
// it was inspired by ruby's i18n module.

package sejong

import (
	"fmt"
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
	sj = New("")
}

func New(locale string) *Sejong {
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
			panic("sejong.T: locale file is not exist")
		}
		configured = append(configured, locale)

	}

	return &Sejong{
		V: v,
		configured: configured,
		Locale: locale,
	}
}

func T(key string, words ...string) string {
	sj.Locale = Locale
	return sj.T(key, words...)
}

func (s *Sejong) T(key string, words ...string) string {
	if s.Locale == "" {
		panic("sejong.T: locale is not set")
	}
	if len(words)%2 == 1 {
		panic("sejong.T: odd word count")
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
	r := newReplacer(words)

	return r.Replace(str)
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
