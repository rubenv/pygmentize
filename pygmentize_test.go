package pygmentize

import (
	"bytes"
	"log"
	"os/exec"
	"strings"
	"testing"
)

func newStrictHtmlFormatter() *HtmlFormatter {
	f := NewHtmlFormatter()
	f.Strict = true
	return f
}

var phpSample = `<?php
// Keys can be obtained in the Mollom site manager.
$public = "your-public-key";
$private = 'yoür-private-key';
$mollom = new Zend_Service_Mollom($public, $private);

// Mandarin: 官話
`

func TestGo(t *testing.T) {
	out, err := HighlightLanguage(`fmt.Println("hello world")`, "go", newStrictHtmlFormatter())
	if err != nil {
		t.Error(err)
	}
	log.Println(out)
}

func TestPhp(t *testing.T) {
	out, err := Highlight(phpSample, newStrictHtmlFormatter())
	if err != nil {
		t.Error(err)
	}

	expected := `<span class="c cp">&lt;?php</span>
<span class="c cs">// Keys can be obtained in the Mollom site manager.</span>
<span class="n nv">$public</span><span class="t"> </span><span class="o">=</span><span class="t"> </span><span class="l ls lsd">&#34;</span><span class="l ls lsd">your-public-key</span><span class="l ls lsd">&#34;</span><span class="p">;</span>
<span class="n nv">$private</span><span class="t"> </span><span class="o">=</span><span class="t"> </span><span class="l ls lss">&#39;yoür-private-key&#39;</span><span class="p">;</span>
<span class="n nv">$mollom</span><span class="t"> </span><span class="o">=</span><span class="t"> </span><span class="k">new</span><span class="t"> </span><span class="n no">Zend_Service_Mollom</span><span class="p">(</span><span class="n nv">$public</span><span class="p">,</span><span class="t"> </span><span class="n nv">$private</span><span class="p">);</span><span class="t">
</span>
<span class="c cs">// Mandarin: 官話</span>
`

	if out != expected {
		t.Errorf("Bad formatting, expected:\n%s\n\nGot:\n%s", expected, out)
	}
}

func TestLanguage(t *testing.T) {
	out, err := HighlightLanguage(`console.log("Hello");`, "js", newStrictHtmlFormatter())
	if err != nil {
		t.Error(err)
	}

	expected := `<span class="n no">console</span><span class="p">.</span><span class="n no">log</span><span class="p">(</span><span class="l ls lsd">&#34;Hello&#34;</span><span class="p">)</span><span class="p">;</span>
`

	if out != expected {
		t.Errorf("Bad formatting, expected:\n%s\n\nGot:\n%s", expected, out)
	}
}

func TestErlang(t *testing.T) {
	_, err := HighlightLanguage(`-module(factorial).
-export([fact/1]).
 
fact(0) -> 1;
fact(N) -> N * fact(N-1).`, "erlang", newStrictHtmlFormatter())
	if err != nil {
		t.Error(err)
	}
}

func TestScala(t *testing.T) {
	_, err := HighlightLanguage(`import scala.actors.Actor
import scala.actors.Actor._
 
case class Inc(amount: Int)
case class Value
 
class Counter extends Actor {
    var counter: Int = 0;
 
    def act() = {
        while (true) {
            receive {
                case Inc(amount) =>
                    counter += amount
                case Value =>
                    println("Value is "+counter)
                    exit()
            }
        }
    }
}
 
object ActorTest extends Application {
    val counter = new Counter
    counter.start()
 
    for (i <- 0 until 100000) {
        counter ! Inc(1)
    }
    counter ! Value
    // Output: Value is 100000
}`, "scala", newStrictHtmlFormatter())
	if err != nil {
		t.Error(err)
	}
}

func TestDebug(t *testing.T) {
	out, err := Highlight(phpSample, DebugFormatter)
	if err != nil {
		t.Error(err)
	}

	t.Log(out)
}

func BenchmarkParse(b *testing.B) {
	cmd := exec.Command("pygmentize", "-f", "raw")
	cmd.Stdin = strings.NewReader(phpSample)
	out, err := cmd.Output()
	if err != nil {
		b.Error(err)
	}

	for n := 0; n < b.N; n++ {
		parse(bytes.NewReader(out), NewHtmlFormatter())
	}
}
