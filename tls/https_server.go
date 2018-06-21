package tls

import (
	"io"
	"log"
	"net/http"
	"os"
)

func TlsServer() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hello World!\n")
	})
	if e := http.ListenAndServeTLS(":70", "server.crt", "server.key", nil); e != nil {
		log.Fatal("ListenAndServer: ", e)
		os.Exit(-1)
	}
}
