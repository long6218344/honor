package http

import "net/http"

const (
	noWritten     = -1
	defaultStatus = http.StatusOK
)

// ResponseWriter .
type ResponseWriter interface {
	http.ResponseWriter
	http.Hijacker
	http.Flusher
	http.CloseNotifier

	// Status the HTTP response status code of the current request.
	Status() int

	// Size returns the number of bytes already written into the response http body.
	Size() int

	// Writes the string into the response body.
	WriteString(string) (int, error)

	// WriteHeaderNow force to wirte the http header (status code + headers).
	WriteHeaderNow()

	// get the http.Pusher for server push
	Pusher() http.Pusher
}

type responseWriter struct {
	http.ResponseWriter
	size   int
	status int
}

func (w *responseWriter) reset(writer http.ResponseWriter) {
	w.ResponseWriter = writer
	w.size = noWritten
	w.status = defaultStatus
}
