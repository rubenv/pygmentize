package pygmentize

import (
	"bytes"
	"os/exec"
	"strings"
	"testing"
)

var phpSample = `<?php
// Keys can be obtained in the Mollom site manager.
$public = "your-public-key";
$private = 'yoür-private-key';
$mollom = new Zend_Service_Mollom($public, $private);

// Mandarin: 官話
`

func TestPhp(t *testing.T) {
	f := NewHtmlFormatter()
	f.Strict = true
	out, err := Highlight(phpSample, f)
	if err != nil {
		t.Error(err)
	}

	expected := `<span class="c cp"><?php</span><span class="t">
</span><span class="c cs">// Keys can be obtained in the Mollom site manager.
</span><span class="n nv">$public</span><span class="t"> </span><span class="o">=</span><span class="t"> </span><span class="l ls lsd">"</span><span class="l ls lsd">your-public-key</span><span class="l ls lsd">"</span><span class="p">;</span><span class="t">
</span><span class="n nv">$private</span><span class="t"> </span><span class="o">=</span><span class="t"> </span><span class="l ls lss">'yoür-private-key'</span><span class="p">;</span><span class="t">
</span><span class="n nv">$mollom</span><span class="t"> </span><span class="o">=</span><span class="t"> </span><span class="k">new</span><span class="t"> </span><span class="n no">Zend_Service_Mollom</span><span class="p">(</span><span class="n nv">$public</span><span class="p">,</span><span class="t"> </span><span class="n nv">$private</span><span class="p">);</span><span class="t">

</span><span class="c cs">// Mandarin: 官話
</span>`

	if out != expected {
		t.Errorf("Bad formatting, expected:\n%s\n\nGot:\n%s", expected, out)
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
