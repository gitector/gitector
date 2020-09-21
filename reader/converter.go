package reader

import (
	"strings"
)

func StringToModel(message string, signature Signature, filesCount int, isMerge bool) GitCommit {
	var res = strings.SplitN(message, "\n", 2)
	if len(res) == 1 {
		return GitCommit{Title: message, Description: "", Signature: signature, RawMessage: message, IsMerge: isMerge}
	} else {
		return GitCommit{Title: res[0], Description: res[1], Signature: signature, RawMessage: message, FilesCount: filesCount, IsMerge: isMerge}
	}
}

func GitVarToSignature(message string) Signature {
	nameAndEmail := strings.Split(message, "<")
	name := strings.TrimSuffix(nameAndEmail[0], " ")
	email := strings.Split(nameAndEmail[1], ">")[0]
	return Signature{
		Name:  name,
		Email: email,
	}
}
