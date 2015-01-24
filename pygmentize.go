package pygmentize

import (
	"bufio"
	"bytes"
	"fmt"
	"html"
	"io"
	"os/exec"
	"strings"
)

type Token struct {
	Type    string
	Subtype string
	Detail  string
}

func (t Token) String() string {
	result := t.Type
	if t.Subtype != "" {
		result += "." + t.Subtype
	}
	if t.Detail != "" {
		result += "." + t.Detail
	}
	return result
}

type Formatter interface {
	Format(token Token, input string) (string, error)
}

func Highlight(code string, formatter Formatter) (string, error) {
	cmd := exec.Command("pygmentize", "-f", "raw")
	cmd.Stdin = strings.NewReader(code)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return "", err
	}

	err = cmd.Start()
	if err != nil {
		return "", err
	}

	out, err := parse(stdout, formatter)
	if err != nil {
		return "", err
	}

	err = cmd.Wait()
	if err != nil {
		return "", err
	}
	return out, nil
}

func parse(reader io.Reader, formatter Formatter) (string, error) {
	var out bytes.Buffer
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		parts := strings.SplitN(scanner.Text(), "\t", 2)

		// Decode token
		tokenParts := strings.Split(parts[0], ".")
		token := Token{Type: tokenParts[1]}
		if len(tokenParts) > 2 {
			token.Subtype = tokenParts[2]
		}
		if len(tokenParts) > 3 {
			token.Detail = tokenParts[3]
		}

		// Decode value
		if parts[1] == "" {
			continue
		}

		valueIn := parts[1][2 : len(parts[1])-1]
		var str bytes.Buffer
		n := len(valueIn)
		for i := 0; i < n; i++ {
			char := valueIn[i]
			if char == '\\' {
				next := valueIn[i+1]
				if next == 'n' {
					str.WriteString("\n")
					i += 1
				} else if next == 'x' {
					c := html.UnescapeString("&#" + valueIn[i+1:i+4] + ";")
					str.WriteString(c)
					i += 3
				} else if next == 'u' {
					c := html.UnescapeString("&#x" + valueIn[i+2:i+6] + ";")
					str.WriteString(c)
					i += 5
				} else {
					return "", fmt.Errorf("Unknown escape sequence: %s (%s)", string(next), valueIn)
				}
			} else {
				str.WriteString(string(char))
			}
		}

		formatted, err := formatter.Format(token, str.String())
		if err != nil {
			return "", err
		}
		out.WriteString(formatted)
	}

	err := scanner.Err()
	if err != nil {
		return "", err
	}

	return out.String(), nil
}

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

type HtmlFormatter struct {
	Classes map[string]string
	Prefix  string
	Strict  bool
}

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

var DebugFormatter = &debugFormatter{
	HtmlFormatter: NewHtmlFormatter(),
}

type debugFormatter struct {
	*HtmlFormatter
}

func (f *debugFormatter) Format(token Token, input string) (string, error) {
	return f.formatSpan(token.String(), input), nil
}
