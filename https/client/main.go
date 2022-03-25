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
	clientKeyPath = "../cert/server.key"
	clientCrtPath = "../cert/server.crt"
)

func main() {
	pool := x509.NewCertPool()

	caCrt, err := ioutil.ReadFile(caCrtPath)
	if err != nil {
		fmt.Println("ReadFile err:", err)
		return
	}
	pool.AppendCertsFromPEM(caCrt)
	cliCrt, err := tls.LoadX509KeyPair(clientCrtPath, clientKeyPath)
	if err != nil {
		fmt.Println("Load X509 keyPair err:", err)
		return
	}

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			RootCAs:      pool,
			Certificates: []tls.Certificate{cliCrt},
		},
	}
	client := &http.Client{Transport: tr}
	//resp, err := k8s.Get("https://115.238.145.60:6443/apis/apiregistration.k8s.io/v1/apiservices/v1alpha1.custom-nginx.nginx.k8s.io")
	resp, err := client.Get("https://115.238.145.60:6443/apis/apiregistration.k8s.io/v1/apiservices/v1.custom-gateway")
	if err != nil {
		fmt.Println("Get error:", err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}
