package main

import (
	"fmt"
	"io/ioutil"
)

// Page構造体。文字列のTitleとByteのBodyをもつ
type Page struct {
	Title string
	Body  []byte
}

// 上記のPage構造体をもつsave関数。エラーを返す
func (p *Page) save() error {

	// Page構造体のTitleに.txtを付け加えてfilenameに格納
	filename := p.Title + ".txt"
	// filenameで定義したファイル名でp.Bodyを指定し0600(所有者のみ読み書き可能)にしてioutil.WhiteFileでファイル作成
	return ioutil.WriteFile(filename, p.Body, 0600)
}

// titleを引数にPageのポインタ(?)とエラーを返すloadPage関数
func loadPage(title string) (*Page, error) {
	// 引数のtitleに".txt"を加えて filenameに格納
	filename := title + ".txt"
	// ioutil.ReadFileで上記のfilenameを引数にし、body, errに格納
	body, err := ioutil.ReadFile(filename)
	// エラー処理
	if err != nil {
		return nil, err
	}
	// Pageのアドレスを＆で指定し、変数で中身を特定する？
	//関数定義時に*Pageとerrを返すようにしているので、&でPageを指定し、{}で囲んでTitleとBodyを設定。エラーはnilを返す。もしnilじゃないなら上のエラー処理に引っかかっている
	return &Page{Title: title, Body: body}, nil
}
func main() {
	// Page構造体に文字列を入れてp1に格納、Bodyは[]byteを指定
	p1 := &Page{Title: "test", Body: []byte("this is test.")}
	// p1をsave()
	p1.save()
	// p2変数に上のp1をloadPageして格納
	// エラーも返すよう定義されているので、_でomitするか、エラー処理も書く
	p2, _ := loadPage(p1.Title)
	// p2のbodyをprintlnで出力、string()で包む必要がある
	fmt.Println(string(p2.Body))
}
