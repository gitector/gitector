package printer

import (
	"fmt"
	"strings"

	"gitlab.com/tbacompany/gitector/rules"
)

// GenerateHTML generates list of errors in HTML format
func GenerateHTML(errors []rules.GitError) string {
	var builder strings.Builder

	if len(errors) > 0 {
		fmt.Fprintf(&builder, "    <h1>Errors found: %d</h1>", len(errors))
		for _, element := range errors {
			fmt.Fprintf(&builder, "\n    <p>%s</p>", element.Title)
		}
	} else {
		fmt.Fprint(&builder, "No errors")
	}
	return fmt.Sprintf(template, builder.String())
}

const template = `<!doctype html>

<html lang="en">
<head>
  <meta charset="utf-8">
  <title>Error report</title>
</head>
<body>
%s
</body>
</html>`
