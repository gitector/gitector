package rules

import (
	"github.com/gobuffalo/packr"
	"gitlab.com/tbacompany/gitector/utils"
	"strings"
)

var wordsBox = packr.NewBox("./words")

func StartsWithBaseVerb(title string) bool {
	firstWord := utils.ExtractFirstWord(title)

	if endsWithEdOrIng(firstWord) || oneOfIrregularVerbs(firstWord) {
		return false
	} else {
		return true
	}
}

func oneOfIrregularVerbs(word string) bool {
	irregularVerbs, _ := wordsBox.FindString("irregular_verbs.txt")
	irregularVerbsList := strings.Split(irregularVerbs, "\n")

	return utils.Contains(word, irregularVerbsList)
}

func endsWithEdOrIng(word string) bool {
	return strings.HasSuffix(word, "ed") || strings.HasSuffix(word, "ing")
}
