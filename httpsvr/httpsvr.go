package main

import "net/http"

//golang启动一个http服务
func main() {

	mux := http.NewServeMux()
	mux.Handle("/hello", http.HandlerFunc(handler))
	http.ListenAndServe(":5688", mux)
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world"))
}
