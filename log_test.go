package apachelog

import (
	"net/http"
	"net/url"
	"strings"
	"testing"

	"github.com/nbio/st"
	"gopkg.in/vinxi/utils.v0"
)

func TestClient(t *testing.T) {
	w := &writerStub{}
	log := New(w)

	var called bool
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		called = true
	})

	rw := utils.NewWriterStub()
	url := &url.URL{Host: "foo.com"}
	req := &http.Request{
		RemoteAddr: "127.0.0.1",
		Method:     "GET",
		RequestURI: "/bar",
		Proto:      "HTTP/1.1",
		Header:     make(http.Header),
		URL:        url,
	}

	log.LogHTTP(handler)(rw, req)
	rw.WriteHeader(200)
	rw.Write([]byte("foo"))

	st.Expect(t, called, true)
	st.Expect(t, rw.Code, 200)
	st.Expect(t, string(rw.Body), "foo")
	st.Expect(t, strings.Contains(w.data, "127.0.0.1 -"), true)
	st.Expect(t, strings.Contains(w.data, "GET /bar HTTP/1.1"), true)
	st.Expect(t, strings.Contains(w.data, " 200 "), true)
}

type writerStub struct {
	code int
	data string
}

func (w *writerStub) WriteHeader(code int) {
	w.code = code
}

func (w *writerStub) Write(data []byte) (int, error) {
	w.data = string(data)
	return 0, nil
}
