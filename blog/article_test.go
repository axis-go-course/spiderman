package blog

import (
	"bytes"
	"io"
	"strings"
	"testing"
)

func Test_article(t *testing.T) {
	A := &Article{Title: "first", Content: "first"}
	var buf bytes.Buffer
	io.Copy(&buf, A.Reader())
	got := buf.String()
	if !strings.Contains(got, "title\":") {
		t.Fatal("missing title", got)
	}
}
