
package main

import (
    "log"
    "net/http"
    "fmt"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "")
}

func ServeFile(w http.ResponseWriter, r *http.Request, name string){
    http.ServeFile(w, r, "") // r.URL.Path[1:])
}

func main() {
    //http.HandleFunc("/", homeHandler)
    
    fs := http.FileServer(http.Dir("./web"))
    http.Handle("/", http.StripPrefix("/", fs))

    log.Fatal(http.ListenAndServe(":8080", nil))
}
