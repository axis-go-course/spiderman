package blog

import (
	"html/template"
	"net/http"
)

const indexHtml = `<!DOCTYPE html>

<html><body>
<h1>{{.BlogName}}</h1>
<form method="POST">

</form>
</body></html>
`

func (s *Service) UserInterface() http.HandlerFunc {
	t, err := template.New("foo").Parse(indexHtml)
	if err != nil {
		panic(err)
	}
	return func(w http.ResponseWriter, r *http.Request) {
		t.Execute(w, struct {
			BlogName string
		}{
			"<script>alert('you have been pwned')</script>",
		})
	}
}
