package io_utils

import (
	"strings"

	"github.com/MakeNowJust/heredoc"
)

// Indentation is the indentation applied to examples
const Indentation = `  `

type normalizer struct {
	string
}

// Examples normalizes a command's examples to follow the conventions.
func Examples(s string) string {
	if len(s) == 0 {
		return s
	}
	return normalizer{s}.trim().indent().string
}

func (s normalizer) heredoc() normalizer {
	s.string = heredoc.Doc(s.string)
	return s
}

// LongDesc normalizes a command's long description to follow the conventions.
func LongDesc(s string) string {
	if len(s) == 0 {
		return s
	}
	return normalizer{s}.heredoc().trim().string
}

// func (s normalizer) markdown() normalizer {
// 	bytes := []byte(s.string)
// 	formatted := blackfriday.Markdown(bytes, &ASCIIRenderer{Indentation: Indentation}, blackfriday.EXTENSION_NO_INTRA_EMPHASIS)
// 	s.string = string(formatted)
// 	return s
// }

func (s normalizer) trim() normalizer {
	s.string = strings.TrimSpace(s.string)
	return s
}

func (s normalizer) indent() normalizer {
	indentedLines := []string{}
	for _, line := range strings.Split(s.string, "\n") {
		trimmed := strings.TrimSpace(line)
		indented := Indentation + trimmed
		indentedLines = append(indentedLines, indented)
	}
	s.string = strings.Join(indentedLines, "\n")
	return s
}
