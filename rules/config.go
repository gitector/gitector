package rules

import (
	"fmt"
	"github.com/bclicn/color"
	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/json"
	"github.com/knadh/koanf/providers/file"
	"gitlab.com/tbacompany/gitector/utils"
)

type ProjectConfig struct {
	emailDomains  []string
	maxCharacters int
	ticketRegexp  string
}

var k = koanf.New(".")

func defaultConfig() ProjectConfig {
	return ProjectConfig{
		emailDomains:  []string{},
		maxCharacters: 72,
		ticketRegexp:  "",
	}
}

func ReadConfig(directory string) ProjectConfig {
	configLocation := directory + "/.gitector.json"
	f := file.Provider(configLocation)
	if err := k.Load(f, json.Parser()); err != nil {
		fmt.Println(color.Yellow("Couldn't read config. Generating default one"))
		utils.InitConfig(configLocation)
		k.Load(f, json.Parser())
	}

	config := ProjectConfig{
		emailDomains:  k.Strings("allowed_domains"),
		maxCharacters: k.Int("title_max_characters"),
		ticketRegexp:  k.String("ticket_regexp"),
	}

	return config
}
