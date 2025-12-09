package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("public/")))

	http.HandleFunc("/calpm", calpmhandler)

	fmt.Println("Server running on http://localhost:8080/")
	http.ListenAndServe(":8080", nil)
}

func calpmhandler(w http.ResponseWriter, r *http.Request) {
	// フォームの読み込み
	if err := r.ParseForm(); err != nil {
		fmt.Fprintln(w, "error: formを読み込めません")
		return
	}

	// 入力値取得
	x, _ := strconv.Atoi(r.FormValue("x"))
	y, _ := strconv.Atoi(r.FormValue("y"))
	op := r.FormValue("cal0")

	// 計算
	switch op {
	case "+":
		fmt.Fprintln(w, x+y)
	case "-":
		fmt.Fprintln(w, x-y)
	case "*":
		fmt.Fprintln(w, x*y)
	case "/":
		if y == 0 {
			fmt.Fprintln(w, "エラー：0で割れません")
			return
		}
		fmt.Fprintln(w, float64(x)/float64(y))
	default:
		fmt.Fprintln(w, "演算子を選んでください")
	}
}
