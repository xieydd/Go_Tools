package tls

import (
	"os/exec"
	"testing"
)

//curl --cacert server.crt https://localhost:70
func TestHttpsServer(t *testing.T) {
	cmd1 := exec.Command("openssl", "genrsa", "-out", "server.key", "2048")
	cmd1.Run()
	cmd2 := exec.Command("openssl", "req", "-nodes", "-new", "-key", "server.key", "-subj", "/CN=localhost", "-out", "server.csr")
	cmd2.Run()
	cmd3 := exec.Command("openssl", "x509", "-req", "-sha256", "-days", "365", "-in", "server.csr", "-signkey", "server.key", "-out", "server.crt")
	cmd3.Run()
	TlsServer()
}
