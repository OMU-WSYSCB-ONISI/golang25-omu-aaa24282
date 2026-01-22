package main
import (
	"fmt"
	"net/http"
	"time"
        "math/rand" //追加
)
func main() {
	http.HandleFunc("/hello", hellohandler)
	http.HandleFunc("/now", nowhandler)
	http.HandleFunc("/dice", dicehandler)

	if err := http.ListenAndServe(":8080", nil); err != nil {
    fmt.Println("server error:", err)
}
}
func hellohandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "こんにちは from Cocespace !")
}
func nowhandler(w http.ResponseWriter, r *http.Request) {
	jst, _ := time.LoadLocation("Asia/Tokyo")
	fmt.Fprintln(w, (time.Now().In(jst)).Format("2006年01月02日 15:04:05"))
}
/* 以下，関数を追加 */
func dicehandler(w http.ResponseWriter, r *http.Request) {
	seed := time.Now().UnixNano()
	rnd := rand.New(rand.NewSource(seed))

fortunes :=[]string{"大吉","中吉","吉","凶"}
result:=fortunes[rnd.Intn(len(fortunes))]

fmt.Fprintf(w, "今の運勢は %s です！", result)

}
