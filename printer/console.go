package printer

import (
	"fmt"
	"strings"

	"github.com/bclicn/color"
	"gitlab.com/tbacompany/gitector/rules"
)

// GenerateConsoleErrors generates list of errors in ASCII format
func GenerateConsoleErrors(errors []rules.GitError) string {
	var builder strings.Builder

	if len(errors) > 0 {
		fmt.Fprintln(&builder, "==== Found errors: ====")
	}

	for _, element := range errors {
		fmt.Fprint(&builder, color.Red(element.Title), ": ")
		fmt.Fprintln(&builder, element.Commit.Title)
	}
	fmt.Fprintln(&builder, "==== Summary: ====")
	if len(errors) > 0 {
		fmt.Fprintf(&builder, color.Red("%#v errors found \n"), len(errors))
	} else {
		fmt.Fprintf(&builder, color.Green("%#v errors found \n"), len(errors))
	}
	return builder.String()
}
