package main

import "net/http"

type helloHander struct {
}
type aboutHander struct {
}

func (m *helloHander) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello go~"))
}
func (m *aboutHander) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("about go~"))
}
func main() {
	//http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
	//	writer.Write([]byte("hello go~"))
	//})

	a := helloHander{}
	b := aboutHander{}
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: nil,
	}
	http.Handle("/hello", &a)
	http.Handle("/about", &b)
	server.ListenAndServe()
	//http.ListenAndServe("localhost:8080", nil) //路由
}
