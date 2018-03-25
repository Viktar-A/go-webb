package main

import (
	"html/template"
	"image"
	"image/jpeg"
	"log"
	"net/http"
)

type page struct {
	Title string
	Msg   string
}

func index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html")

	title := r.URL.Path[len("/"):]

	if title != "exec/" {
		t, _ := template.ParseFiles("index.html")
		t.Execute(w, &page{Title: "Convert Image"})
	} else {
		imgfile, fhead, _ := r.FormFile("imgfile")

		img, ext, _ := image.Decode(imgfile)

		w.Header().Set("Content-type", "image/jpeg")
		w.Header().Set("Content-Disposition", "filename=\""+fhead.Filename+"."+ext+"\"")
		//w.Header().Set("Content-Disposition", "filename=\""+fhead.Filename+"\"")
		im := 50 // Quality ranges from 1 to 100 inclusive, higher is better.
		jpeg.Encode(w, img, &jpeg.Options{Quality: im})
	}
}

func main() {

	http.HandleFunc("/", index)
	http.Handle("/img/", http.StripPrefix("/img/", http.FileServer(http.Dir("./img/"))))
	//log.Println("ListenAndServe....0.0.0.0:3000")
	if err := http.ListenAndServe(":3000", nil); err != nil {
		log.Fatal("failed to start server", err)
	}
}
