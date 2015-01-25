/*
Work in progress module that wraps pygments for use with Go.

http://pygments.org/

Quick Example

To highlight a piece of code, use:

	code := `fmt.Println("hello world")`
	out, err := pygmentize.HighlightLanguage(code, "go", NewHtmlFormatter())
	if err != nil {
		return err
	}
	fmt.Println(out)

This outputs:

	<span class="n no">fmt</span><span class="p">.</span><span class="n no">Println</span><span class="p">(</span><span class="l ls">"hello world"</span><span class="p">)</span><span class="t">
	</span>

Apply CSS to your liking to get the desired visual effect.
*/
package pygmentize

//go:generate godocdown -output README.md
