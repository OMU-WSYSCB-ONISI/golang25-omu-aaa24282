package main

import (
	"fmt"
	"html"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"time"
)

const saveFile = "public/memo.txt"
const historyDir = "public/history/"

func main() {
	fmt.Printf("Go version: %s\n", runtime.Version())

	http.HandleFunc("/hello", hellohandler)
	http.HandleFunc("/memo", memo)
	http.HandleFunc("/mwrite", mwrite)
	http.HandleFunc("/history", historyList)
	http.HandleFunc("/history/view", historyView)

	fmt.Println("Launch server...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Failed to launch server: %v", err)
	}
}

func hellohandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "こんにちは from Codespace !")
}

func memo(w http.ResponseWriter, r *http.Request) {
	text, err := os.ReadFile(saveFile)
	if err != nil {
		text = []byte("ここにメモを記入してください。")
	}
	htmlText := html.EscapeString(string(text))

	s := "<html>" +
		"<style>textarea { width:99%; height:200px; }</style>" +
		"<a href='/history'>履歴を見る</a><br><br>" +
		"<form method='post' action='/mwrite'>" +
		"<textarea name='text'>" + htmlText + "</textarea>" +
		"<input type='submit' value='保存' /></form></html>"

	w.Write([]byte(s))
}

func mwrite(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "POSTメソッドを使ってください", http.StatusMethodNotAllowed)
		return
	}

	r.ParseForm()
	text := r.FormValue("text")

	timestamp := time.Now().Format("20060102_150405")
	historyFile := filepath.Join(historyDir, "memo_"+timestamp+".txt")

	os.WriteFile(historyFile, []byte(text), 0644)

	os.WriteFile(saveFile, []byte(text), 0644)
	fmt.Println("save:", text)

	http.Redirect(w, r, "/memo", http.StatusSeeOther)
}

func historyList(w http.ResponseWriter, r *http.Request) {
	files, err := os.ReadDir(historyDir)
	if err != nil {
		http.Error(w, "履歴フォルダが読み込めません", 500)
		return
	}

	s := "<html><h2>履歴一覧</h2><ul>"

	for _, f := range files {
		name := f.Name()
		s += "<li><a href='/history/view?file=" + name + "'>" + name + "</a></li>"
	}

	s += "</ul><a href='/memo'>メモに戻る</a></html>"

	w.Write([]byte(s))
}

func historyView(w http.ResponseWriter, r *http.Request) {
	file := r.URL.Query().Get("file")
	if file == "" {
		http.Error(w, "ファイルが指定されていません", 400)
		return
	}

	content, err := os.ReadFile(filepath.Join(historyDir, file))
	if err != nil {
		http.Error(w, "ファイルが読み込めません", 500)
		return
	}

	htmlText := html.EscapeString(string(content))

	s := "<html><h2>" + file + "</h2><pre>" + htmlText +
		"</pre><br><a href='/history'>履歴一覧に戻る</a></html>"

	w.Write([]byte(s))
}
