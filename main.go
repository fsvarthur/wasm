package main

import (
    "flag"
    "fmt"
    "log"
    "net/http"
)

var (
    listen = flag.String("listen", ":8080", "listen address")
    dir    = flag.String("dir", ".", "directory to serve")
)

func main() {
    flag.Parse()
    log.Printf("listening on %q...", *listen)
    err := http.ListenAndServe(*listen, http.FileServer(http.Dir(*dir)))
    log.Fatalln(err)
    fmt.Println("Hello, WebAssembly!")
}
