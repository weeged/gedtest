package main

import (
	"fmt"
	"net/http"
)

func defaultHandler(w http.ResponseWriter, r *http.Request) {
<<<<<<< HEAD
    fmt.Fprintf(w, "<h1>Hello from Cisco Shipped!  Shmersh!!!</h1>\n")
=======
    fmt.Fprintf(w, "<h1>Hello from Cisco Shipped!  Shmersh!!</h1>\n")
>>>>>>> 16b5679ae7334e9c4dde4acc2301180331f8f5fe
}

func main() {
    http.HandleFunc("/", defaultHandler)
    fmt.Println("Example app listening at http://localhost:8888")
    http.ListenAndServe(":8888", nil)
}
