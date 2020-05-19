package example

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func BenchmarkRouter(b *testing.B) {
	r := http.NewServeMux()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
	ts := httptest.NewServer(r)
	defer ts.Close()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		http.Get(ts.URL)
	}
}
