package main

import "net/http"

type myHander struct {
}

func (m *myHander) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello go~"))
}

func main() {
	//http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
	//	writer.Write([]byte("hello go~"))
	//})

	mh := myHander{}
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: &mh,
	}
	server.ListenAndServe()
	//http.ListenAndServe("localhost:8080", nil) //路由
}
