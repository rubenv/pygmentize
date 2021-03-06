# pygmentize

[![Build Status](https://travis-ci.org/rubenv/pygmentize.svg?branch=master)](https://travis-ci.org/rubenv/pygmentize) [![GoDoc](https://godoc.org/github.com/rubenv/pygmentize?status.png)](https://godoc.org/github.com/rubenv/pygmentize)

Pygments wrapper for Go.

http://pygments.org/


### Quick Example

To highlight a piece of code, use:

    code := `fmt.Println("hello world")`
    out, err := pygmentize.HighlightLanguage(code, "go", NewHtmlFormatter())
    if err != nil {
    	return err
    }
    fmt.Println(out)

This outputs:

    <span class="n no">fmt</span><span class="p">.</span><span class="n no">Println</span><span class="p">(</span><span class="l ls">&#34;hello world&#34;</span><span class="p">)</span><span class="t">
    </span>

Apply CSS to your liking to get the desired visual effect.

## Installation
```
go get github.com/rubenv/pygmentize
```

Import into your application with:

```go
import "github.com/rubenv/pygmentize"
```

## Usage

```go
var DebugFormatter = &debugFormatter{
	HtmlFormatter: NewHtmlFormatter(),
}
```
Formatter that outputs HTML with token types as the class name.

```go
var DefaultClasses = map[string]string{
	"Comment":                "c",
	"Comment.Preproc":        "cp",
	"Comment.Single":         "cs",
	"Keyword":                "k",
	"Keyword.Constant":       "kc",
	"Keyword.Type":           "kt",
	"Literal":                "l",
	"Literal.Number":         "ln",
	"Literal.Number.Integer": "lni",
	"Literal.String":         "ls",
	"Literal.String.Double":  "lsd",
	"Literal.String.Single":  "lss",
	"Name":                   "n",
	"Name.Class":             "nc",
	"Name.Entity":            "ne",
	"Name.Function":          "nf",
	"Name.Other":             "no",
	"Name.Namespace":         "nn",
	"Name.Variable":          "nv",
	"Operator":               "o",
	"Punctuation":            "p",
	"Text":                   "t",
}
```
Default token -> class mapping.

#### func  Highlight

```go
func Highlight(code string, formatter Formatter) (string, error)
```
Highlight a piece of code.

#### func  HighlightLanguage

```go
func HighlightLanguage(code, language string, formatter Formatter) (string, error)
```
Highlight a piece of code, with a given language.

See http://pygments.org/docs/lexers/ for a list of languages (look under "Short
names").

#### type Formatter

```go
type Formatter interface {
	Format(token Token, input string) (string, error)
}
```


#### type HtmlFormatter

```go
type HtmlFormatter struct {
	// Maps of token types to class names
	Classes map[string]string

	// Prefix added to each class
	Prefix string

	// Fail on unmapped token types?
	Strict bool
}
```

Highlights by adding <span> tags.

#### func  NewHtmlFormatter

```go
func NewHtmlFormatter() *HtmlFormatter
```
Create a new HtmlFormatter with default class mapping.

#### func (*HtmlFormatter) Format

```go
func (f *HtmlFormatter) Format(token Token, input string) (string, error)
```

#### type Token

```go
type Token struct {
	Type    string
	Subtype string
	Detail  string
}
```


#### func (Token) String

```go
func (t Token) String() string
```

## License

    (The MIT License)

    Copyright (C) 2015 by Ruben Vermeersch <ruben@rocketeer.be>

    Permission is hereby granted, free of charge, to any person obtaining a copy
    of this software and associated documentation files (the "Software"), to deal
    in the Software without restriction, including without limitation the rights
    to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
    copies of the Software, and to permit persons to whom the Software is
    furnished to do so, subject to the following conditions:

    The above copyright notice and this permission notice shall be included in
    all copies or substantial portions of the Software.

    THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
    IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
    FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
    AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
    LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
    OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
    THE SOFTWARE.
