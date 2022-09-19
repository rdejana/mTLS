package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.TLS.VerifiedChains[0][0].Subject)
	// Write "Hello, world!" to the response body
	io.WriteString(w, "Hello, world!\n")
}

func main() {

	caCert, err := os.ReadFile("cert.pem")
	if err != nil {
		log.Fatal(err)
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	tlsConfig := &tls.Config{
		ClientCAs:  caCertPool,
		ClientAuth: tls.RequireAndVerifyClientCert,
	}
	//this is deprecated
	//tlsConfig.BuildNameToCertificate()
	server := &http.Server{
		Addr:      ":443",
		TLSConfig: tlsConfig,
	}
	// Set up a /hello resource handler
	http.HandleFunc("/hello", helloHandler)

	// Listen to port 80 and wait
	//log.Fatal(http.ListenAndServe(":80", nil))
	//TLS time
	//log.Fatal(server.ListenAndServeTLS("cert.pem", "key.pem"))
	//log.Fatal(http.ListenAndServeTLS(":443", "cert.pem", "key.pem", nil))
	log.Fatal(server.ListenAndServeTLS("cert.pem", "key.pem"))
}
