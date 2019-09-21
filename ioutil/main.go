package main

import (
	"fmt"
	"io/ioutil"
)

type Page struct {
	Title string
	Body  []byte
}

// 上記のPage構造体をもつ関数
func (p *Page) save() error {
	filename := p.Title + ".txt"
	// filenameで定義したファイル名で0600(所有者のみ読み書き可能)にしてファイル作成
	return ioutil.WriteFile(filename, p.Body, 0600)
}

// Pageのポインタ(?)とエラーを返す
func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	// Pageのアドレスを＆で指定し、変数で中身を特定する？
	return &Page{Title: title, Body: body}, nil
}

func main() {
	// Page構造体に文字列を入れてp1に格納
	p1 := &Page{Title: "test", Body: []byte("This is a sample test..")}
	p1.save()

	// エラーも返すよう定義されているので、_でomitしている
	p2, _ := loadPage(p1.Title)
	fmt.Println(string(p2.Body))

}
