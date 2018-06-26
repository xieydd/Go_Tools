package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	h := http.Server{
		Addr: ":70",
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "Hello World!")
		}),
		TLSConfig: &tls.Config{
			ClientCAs:  loadCA("ca.crt"),
			ClientAuth: tls.RequireAndVerifyClientCert,
		},
	}
	if e := h.ListenAndServeTLS("server.crt", "server.key"); e != nil {
		log.Fatal("ListenAndServer: ", e)
		os.Exit(-1)
	}
}

func loadCA(caFile string) *x509.CertPool {
	pool := x509.NewCertPool()

	if ca, e := ioutil.ReadFile(caFile); e != nil {
		log.Fatal("ReadFile: ", e)
	} else {
		pool.AppendCertsFromPEM(ca)
	}
	return pool

}
