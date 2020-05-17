package blog

import (
	"html/template"
	"net/http"
)

func (s *Service) UserInterface() http.HandlerFunc {
	t := template.Must(template.New("").Parse(htmlTemplates))
	return func(w http.ResponseWriter, r *http.Request) {
		result := make([]*Article, 10)
		n := s.blog.LoadArticles(result)
		t.ExecuteTemplate(w, "index", struct {
			BlogName string
			Articles []*Article
		}{
			"Spidermans blog",
			result[:n],
		})
	}
}

const htmlTemplates = `
{{define "index"}}<!DOCTYPE html>

<html>
<head>
<script>
function postArticle() {
 var xhr = new XMLHttpRequest()
  xhr.open("POST", "/articles")
  xhr.setRequestHeader('Content-Type', 'application/json; charset=UTF-8')

  // send the collected data as JSON
  var data = {};
  data["title"] = document.getElementById('title').value
  data["content"] = document.getElementById('content').value
  xhr.send(JSON.stringify(data))

  xhr.onloadend = function () {
    location.reload()
  }
}
</script>
<style>
label { display: inline-block; width: 100px; float: left }
input { width: 100% }
textarea { width: 100%; height: 100px }
</style>
</head>
<body>
<h1>{{.BlogName}}</h1>

<div>
<label>Title:</label><input id="title"><br>
<label>Content:</label><textarea id="content"></textarea><br>
<button onclick="postArticle()">Post</button>
</div>

<hr>

{{range .Articles}}{{template "article" . }}{{end}}
</body></html>{{end}}


{{define "article"}}<article>
<h2>{{.Title}}</h2>
<div>{{.Content}}</div>
</article>
{{end}}

`
