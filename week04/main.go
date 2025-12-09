package main

import (
    "fmt"
    "net/http"
    "time"
)

func main() {
    http.HandleFunc("/info", infohandler)
    http.ListenAndServe(":8080", nil)
}

/* 以下，関数を追加 */
func infohandler(w http.ResponseWriter, r *http.Request) {
    jst, _ := time.LoadLocation("Asia/Tokyo")
    now := time.Now().In(jst).Format("2006年01月02日 15:04:05")

    userAgent := r.Header.Get("User-Agent")

    fmt.Fprintf(w,"今の時刻は", now, "で、利用しているブラウザは", userAgent, "ですね。")

}
