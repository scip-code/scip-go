package line_directives

import "strings"

func MidFile(value string) string {
	before := strings.TrimSpace(value)
//line target.go:3:1
	after := strings.TrimSpace(value)
	return before + after
}
