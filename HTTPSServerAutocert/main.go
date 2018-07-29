package main

import (
	"crypto/tls"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"golang.org/x/crypto/acme/autocert"
)

func createHTTPServerFromMux(mux *http.ServeMux) *http.Server {
	httpServ := &http.Server{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
		IdleTimeout:  120 * time.Second,
		Handler:      mux,
	}
	return httpServ
}

func createHTTPServer() *http.Server {
	mux := &http.ServeMux{}
	mux.HandleFunc("/", handleRequest)
	return createHTTPServerFromMux(mux)
}

func handleRequest(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	io.WriteString(w, "Hello from TLS Server")
}

func httpTOHttpsRedirectServer() *http.Server {
	mux := &http.ServeMux{}
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Println("Redirected")
		newURI := "https://" + req.Host + req.URL.String()
		http.Redirect(w, req, newURI, http.StatusFound)
	})

	return createHTTPServerFromMux(mux)
}
func main() {
	var m *autocert.Manager
	//var httpsSrv *http.Server
	http.HandleFunc("/", handleRequest)
	m = &autocert.Manager{
		Prompt: autocert.AcceptTOS,
		Cache:  autocert.DirCache("."),
	}

	srv := &http.Server{
		Addr: ":10443",
		TLSConfig: &tls.Config{
			GetCertificate: m.GetCertificate,
		},
	}

	//httpsSrv = createHTTPServer()
	//httpsSrv.Addr = ":10443"
	//httpsSrv.TLSConfig = &tls.Config{GetCertificate: m.GetCertificate}
	//httpsSrv.Handler = m.HTTPHandler(httpsSrv.Handler)
	//go log.Fatalln(httpsSrv.ListenAndServe())
	//go func() {
	fmt.Printf("Starting HTTPS server on %s\n", srv.Addr)
	err := srv.ListenAndServeTLS("", "")
	fmt.Println(err)
	if err != nil {
		log.Fatalf("httpsSrv.ListendAndServeTLS() failed with %s", err)
	}
	//}()
}
