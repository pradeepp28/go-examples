package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/vehicle/images", func(rw http.ResponseWriter, r *http.Request) {
		if err := r.ParseMultipartForm(1 << 10); err != nil {
			http.Error(rw, err.Error(), http.StatusBadRequest)
			return
		}
		files := r.MultipartForm.File["images"]
		for _, file := range files {
			rf, err := file.Open()
			if err != nil {
				http.Error(rw, err.Error(), http.StatusBadRequest)
				return
			}
			wf, err := os.Create(file.Filename)
			if err != nil {
				http.Error(rw, err.Error(), http.StatusBadRequest)
				return
			}
			io.Copy(wf, rf)
			wf.Close()
			rf.Close()
		}
	})
	srv := &http.Server{
		Addr:    ":9110",
		Handler: mux,
	}
	if err := srv.ListenAndServe(); err != nil {
		log.Fatal("Failed to start service", err)
	}
}
