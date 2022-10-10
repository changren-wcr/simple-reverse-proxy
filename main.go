package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
)

func main() {
	reverseProxyForSingleFile := http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		// extract host info
		r := regexp.MustCompile("/(.+?)/")
		host := strings.Trim(r.FindString(req.URL.Path), "/")
		if host == "" {
			rw.WriteHeader(http.StatusBadRequest)
			rw.Write([]byte("please append a web url"))
			return
		}

		// set request path and host
		req.Host = host
		req.URL.Host = host
		req.URL.Scheme = "http"
		req.URL.Path = strings.TrimPrefix(req.URL.Path, "/"+host)
		req.RequestURI = ""

		// do request
		log.Printf("sent url %s", req.URL.String())
		originServerResponse, err := http.DefaultClient.Do(req)
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(rw, err)
			return
		}

		// copy headers
		for k, v := range originServerResponse.Header {
			rw.Header().Set(k, v[0])
			for i := 1; i < len(v); i++ {
				rw.Header().Add(k, v[i])
			}
		}
		rw.WriteHeader(http.StatusOK)

		// copy body
		io.Copy(rw, originServerResponse.Body)
	})
	port := "8080"
	// The web process must listen for HTTP traffic on $PORT, which is set by Heroku
	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	}
	log.Fatal(http.ListenAndServe(":"+port, reverseProxyForSingleFile))
}
