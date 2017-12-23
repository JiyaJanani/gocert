package main

import (
	"fmt"
	"log"
	"net/http"
)

func HelloServer(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	// w.Write([]byte("This is an example server.\n"))
	fmt.Fprintf(w, "This is an example server.\n")
	// io.WriteString(w, "This is an example server.\n")
}

func main() {
	http.HandleFunc("/hello", HelloServer)
	err := http.ListenAndServeTLS(":3001", "E:/cert/ssl.crt/server.crt", "E:/cert/ssl.key/server.key", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
