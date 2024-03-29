package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	caCrtPath     = "../cert/ca.crt"
	serverKeyPath = "../cert/server.key"
	serverCrtPath = "../cert/server.crt"
)

type myhandler struct {
}

func (h *myhandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi, This is an example of http service in golang!\n")
}

func main() {
	pool := x509.NewCertPool()
	caCrt, err := ioutil.ReadFile(caCrtPath)
	if err != nil {
		fmt.Println("ReadFile err:", err)
		return
	}
	pool.AppendCertsFromPEM(caCrt)

	s := &http.Server{
		Addr:    ":443",
		Handler: &myhandler{},
		TLSConfig: &tls.Config{
			ClientCAs:  pool,
			ClientAuth: tls.RequireAndVerifyClientCert,
		},
	}

	err = s.ListenAndServeTLS(serverCrtPath, serverKeyPath)
	if err != nil {
		fmt.Println("ListenAndServeTLS err:", err)
	}
}
