package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type Page struct {
	Title string
	Body  []byte
}

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"

	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	return &Page{Title: title, Body: body}, nil
}

// 引数にresponseWriterとhttp.requestを持つviewHandler
func viewHandler(w http.ResponseWriter, r *http.Request) {
	// responseWriterのURL.Pathでviewを除いた値をtitleに格納, [len("/view/"):]という記述になる
	title := r.URL.Path[len("/view/"):]
	// titleを引数にloadPage関数をpに格納
	p, _ := loadPage(title)

	// responsdeWritermでhtmlを記述し、pのtitleとbodyを指定
	// Fprintfじゃないとブラウザに出力できない
	fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
}

func main() {
	// HandleFunc
	http.HandleFunc("/view/", viewHandler)
	// ListenAndServe
	http.ListenAndServe(":8000", nil)
}
