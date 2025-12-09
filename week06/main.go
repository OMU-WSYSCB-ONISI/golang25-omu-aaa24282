package main

import (
	"fmt"
	"net/http"
	"runtime"
  "strconv"
)


func main(){
        fmt.Printf("Go version: %s\n", runtime.Version())
		http.Handle("/", http.FileServer(http.Dir("public/")))
		http.HandleFunc("/calBMI",calBMIhandler)
		if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Failed to launch server: %v", err)
	}

}

func calBMIhandler(w http.ResponseWriter, r *http.Request){
    if err := r.ParseForm(); err != nil {
        fmt.Println("errorだよ")
    }

    weight, _ := strconv.ParseFloat(r.FormValue("weight"), 64)
    height, _ := strconv.ParseFloat(r.FormValue("height"), 64)

    h := height / 100.0

    bmi := weight / (h * h)


fmt.Fprintf(w, "<h1>BMIは %.2f です。</h1>", bmi)

}

