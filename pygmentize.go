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

// Highlight a piece of code.
func Highlight(code string, formatter Formatter) (string, error) {
	return HighlightLanguage(code, "", formatter)
}

// Highlight a piece of code, with a given language.
//
// See http://pygments.org/docs/lexers/ for a list of languages (look under "Short names").
func HighlightLanguage(code, language string, formatter Formatter) (string, error) {
	args := []string{
		"-f", "raw",
	}
	if language != "" {
		args = append(args, "-l", language)
	}

	cmd := exec.Command("pygmentize", args...)
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
