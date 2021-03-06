package handlers

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

type LoggerHandler struct {
	DebugMode bool
}

func (l *LoggerHandler) Logging(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("\n%s %s %s %s", r.RemoteAddr, r.Method, r.URL, r.Proto)
		if l.DebugMode {
			body, err := ioutil.ReadAll(r.Body)
			if err != nil {
				panic(err)
			}
			reqStr := ioutil.NopCloser(bytes.NewBuffer(body))
			fmt.Printf("%v\n", reqStr)
			r.Body = reqStr
		}
		h.ServeHTTP(w, r)
	})
}
