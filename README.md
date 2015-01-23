# pygmentize

[![Build Status](https://travis-ci.org/rubenv/pygmentize.svg?branch=master)](https://travis-ci.org/rubenv/pygmentize) [![GoDoc](https://godoc.org/github.com/rubenv/pygmentize?status.png)](https://godoc.org/github.com/rubenv/pygmentize)

Work in progress module that wraps pygments for use with Go.

http://pygments.org/

## Installation
```
go get github.com/rubenv/pygmentize
```

Import into your application with:

```go
import "github.com/rubenv/pygmentize"
```

## Usage

#### func  Highlight

```go
func Highlight(code string, formatter Formatter) (string, error)
```

#### type Formatter

```go
type Formatter interface {
	Format(token Token, input string) string
}
```


#### type HtmlFormatter

```go
type HtmlFormatter struct {
}
```


#### func (*HtmlFormatter) Format

```go
func (f *HtmlFormatter) Format(token Token, input string) string
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
