package proxy

import (
	"bytes"
	"net/http"
)

type responseRecorder struct {
	http.ResponseWriter
	body    *bytes.Buffer
	headers map[string]string
}

func newResponseRecorder(w http.ResponseWriter) *responseRecorder {
	return &responseRecorder{
		ResponseWriter: w,
		body:           bytes.NewBuffer(nil),
		headers:        make(map[string]string),
	}
}

func (r *responseRecorder) Write(b []byte) (int, error) {
	r.body.Write(b)
	return r.ResponseWriter.Write(b)
}

func (r *responseRecorder) Header() http.Header {
	return r.ResponseWriter.Header()
}
