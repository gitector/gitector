package rules

import (
	"fmt"
	"github.com/gobuffalo/packr"
	"strings"
)

var wordsBox = packr.NewBox("./words")

func StartsWithVerb(title string) bool {
	irregularVerbs, _ := wordsBox.FindString("irregular_verbs.txt")
	irregularVerbsList := strings.Split(irregularVerbs, "\n")
}
