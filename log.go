package log

import (
	"io"
	"net/http"
	"os"

	"gopkg.in/h2non/apachelog.v0"
	"gopkg.in/vinxi/layer.v0"
)

// Default provides a default logger middleware who writes into stdout.
var Default = New(os.Stdout)

// Logger represents the log layer who writes to the given io.Writer.
type Logger struct {
	w io.Writer
}

// New creates a new log middleware.
func New(w io.Writer) *Logger {
	return &Logger{w: w}
}

// Register registers the log middleware handler.
func (l *Logger) Register(mw layer.Middleware) {
	mw.UsePriority("request", layer.TopHead, l.LogHTTP)
}

// LogHTTP instruments and logs an incoming HTTP request and response.
func (l *Logger) LogHTTP(h http.Handler) func(w http.ResponseWriter, r *http.Request) {
	log := apachelog.New(h, l.w)
	return func(w http.ResponseWriter, r *http.Request) {
		log.ServeHTTP(w, r)
	}
}
