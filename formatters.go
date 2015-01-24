package pygmentize

import (
	"bytes"
	"fmt"
	"strings"
)

// Default token -> class mapping.
var DefaultClasses = map[string]string{
	"Comment":               "c",
	"Comment.Preproc":       "cp",
	"Comment.Single":        "cs",
	"Keyword":               "k",
	"Literal":               "l",
	"Literal.String":        "ls",
	"Literal.String.Double": "lsd",
	"Literal.String.Single": "lss",
	"Name":                  "n",
	"Name.Other":            "no",
	"Name.Variable":         "nv",
	"Operator":              "o",
	"Punctuation":           "p",
	"Text":                  "t",
}

// Highlights by adding <span> tags.
type HtmlFormatter struct {
	// Maps of token types to class names
	Classes map[string]string

	// Prefix added to each class
	Prefix string

	// Fail on unmapped token types?
	Strict bool
}

// Create a new HtmlFormatter with default class mapping.
func NewHtmlFormatter() *HtmlFormatter {
	return &HtmlFormatter{
		Classes: DefaultClasses,
	}
}

func (f *HtmlFormatter) Format(token Token, input string) (string, error) {
	var key bytes.Buffer
	var err error
	c := ""

	key.WriteString(token.Type)
	c, err = f.tryTokenClass(key.String(), c)
	if err != nil {
		return "", err
	}

	key.WriteString(".")
	key.WriteString(token.Subtype)
	c, err = f.tryTokenClass(key.String(), c)
	if err != nil {
		return "", err
	}

	key.WriteString(".")
	key.WriteString(token.Detail)
	c, err = f.tryTokenClass(key.String(), c)
	if err != nil {
		return "", err
	}

	return f.formatSpan(c, input), nil
}

func (f *HtmlFormatter) formatSpan(c, input string) string {
	if c == "" {
		return input
	} else {
		return fmt.Sprintf(`<span class="%s%s">%s</span>`, f.Prefix, c, input)
	}
}

func (f *HtmlFormatter) tryTokenClass(tokenType, prev string) (string, error) {
	c, exists := f.Classes[tokenType]
	if f.Strict && !exists && !strings.HasSuffix(tokenType, ".") {
		return "", fmt.Errorf("Unknown token type: %s", tokenType)
	}

	if exists && prev != "" {
		return prev + " " + c, nil
	} else if exists {
		return c, nil
	} else {
		return prev, nil
	}
}

// Formatter that outputs HTML with token types as the class name.
var DebugFormatter = &debugFormatter{
	HtmlFormatter: NewHtmlFormatter(),
}

type debugFormatter struct {
	*HtmlFormatter
}

func (f *debugFormatter) Format(token Token, input string) (string, error) {
	return f.formatSpan(token.String(), input), nil
}
