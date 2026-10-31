package main

import (
    "embed"
    "log"
    "net/http"
)

//go:embed index.html AWS_logo.png
var assets embed.FS

func main() {
    http.Handle("/", http.FileServer(http.FS(assets)))
    log.Println("Retro Minicon app listening on :8081")
    log.Fatal(http.ListenAndServe(":8081", nil))
}
