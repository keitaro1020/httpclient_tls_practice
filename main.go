package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"os"
)

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		// https://dev.classmethod.jp/articles/check-your-client-side-ssl-tls-support-via-api/
		req, _ := http.NewRequest(http.MethodGet, "https://www.howsmyssl.com/a/check", nil)
		client := http.Client{}
		res, _ := client.Do(req)
		defer res.Body.Close()

		dumpResp, _ := httputil.DumpResponse(res, true)
		fmt.Printf("dumpResp: %v", dumpResp)

		resBody, _ := ioutil.ReadAll(res.Body)
		writer.Write(resBody)
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	log.Printf("Listening on port %s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}