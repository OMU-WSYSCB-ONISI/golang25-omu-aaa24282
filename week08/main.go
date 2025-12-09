package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("public/")))
	http.HandleFunc("/avg", averageandgraphhandler)

	fmt.Println("Server running on http://localhost:8080/")
	http.ListenAndServe(":8080", nil)
}

func averageandgraphhandler(w http.ResponseWriter, r *http.Request) {
	var sum int

	if err := r.ParseForm(); err != nil {
		fmt.Fprintln(w, "error: フォームを読み込めませんでした")
		return
	}

	scores := strings.Split(r.FormValue("dd"), ",")
	intScores := make([]int, 0)

	for _, s := range scores {
		n, err := strconv.Atoi(strings.TrimSpace(s))
		if err == nil {
			intScores = append(intScores, n)
			sum += n
		}
	}

	avg := float64(sum) / float64(len(intScores))

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	fmt.Fprintf(w, "平均 %0.2f\n\n", avg)

	dis := make([]int, 10)
	for i := 0; i < 10; i++ {
		border := 90 - 10*i
		dis[i] = returnDV(intScores, border)
	}

	for i := 0; i < 10; i++ {
		border := 90 - 10*i
		fmt.Fprintf(w, "%d以上: ", border)
		printAsterisk(w, dis[i])
	}
}

func returnDV(scores []int, border int) int {
	count := 0
	for _, s := range scores {
		if s >= border {
			count++
		}
	}
	return count
}

func printAsterisk(w http.ResponseWriter, n int) {
	for i := 0; i < n; i++ {
		fmt.Fprint(w, "*")
	}
	fmt.Fprintln(w)
}
