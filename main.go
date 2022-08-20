package main

import (
	"fmt"
	"log"
	"net/http"
)

const (
	videosDir = "videos"
	port      = 8080
)

func main() {
	http.HandleFunc("/", serveIndex())
	http.HandleFunc("/videos/", addHeaders(serveVideos()))

	fmt.Printf("Starting server on %v\n", port)
	log.Printf("Serving %s on HTTP port: %v\n", videosDir, port)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}

func addHeaders(h http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		h.ServeHTTP(w, r)
	}
}

func serveIndex() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("serving index")
		http.ServeFile(w, r, "public/index.html")
	}
}

func serveVideos() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/x-mpegURL")
		w.Header().Set("Content-Type", "video/MP2T")
		vid := r.URL.Path[len("/videos/"):]
		log.Println("serving", vid)
		vidPath := fmt.Sprintf("%s/%s", videosDir, vid)
		http.ServeFile(w, r, vidPath)
	}
}
