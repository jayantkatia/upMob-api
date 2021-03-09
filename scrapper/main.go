package main

// import (
// 	"fmt"
// 	"log"
// 	"net/http"
// )

// func repoHandler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
// }
// func indexHandler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
// }
// func main() {
// 	http.HandleFunc("/api/index", indexHandler)
// 	http.HandleFunc("/api/repo", repoHandler)

// 	log.Fatal(http.ListenAndServe("localhost:8003", nil))
// }
