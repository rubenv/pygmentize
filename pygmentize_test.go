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
	_, err := Highlight(phpSample, &HtmlFormatter{})
	if err != nil {
		t.Error(err)
	}

	//log.Println(out)
}

func BenchmarkParse(b *testing.B) {
	cmd := exec.Command("pygmentize", "-f", "raw")
	cmd.Stdin = strings.NewReader(phpSample)
	out, err := cmd.Output()
	if err != nil {
		b.Error(err)
	}

	for n := 0; n < b.N; n++ {
		parse(bytes.NewReader(out), &HtmlFormatter{})
	}
}
