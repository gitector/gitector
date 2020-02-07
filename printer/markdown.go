package printer

import (
	"fmt"
	"strings"

	"gitlab.com/tbacompany/gitector/rules"
)

// GenerateMarkdown is generating a list of errors in the Markdown format
func GenerateMarkdown(errors []rules.GitError) string {
	if len(errors) == 0 {
		return ""
	}

	var builder strings.Builder
	fmt.Fprintf(&builder, "## Errors found: %d\n\n", len(errors))
	for _, element := range errors {
		fmt.Fprintf(&builder, "- %s\n", element.Title)
	}
	return builder.String()
}
