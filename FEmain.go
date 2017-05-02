package main
import "fmt"
import "net/http"

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, Web")
}

func main() {
	http.HandleFunc("/",handler)
	http.ListenAndServe(":8090",nil)
}