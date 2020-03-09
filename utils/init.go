package utils

import (
	"io/ioutil"
	"strings"
)

var defaultConfig = []byte(`
{
  "allowed_domains": [],
  "title_max_characters": 72,
  "ticket_regexp": "",
   "no_trailing_punctuation_in_title": true
}
`)

func InitConfig(configLocation string) {
	ioutil.WriteFile(configLocation, defaultConfig, 0644)
}

var MaxOutputWidth = 72

func Dash() string {
	return strings.Repeat("-", MaxOutputWidth)
}
