package main

import "net/http"

func main() {

	http.ListenAndServe(":8080", nil)
}

func index(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("Hello from ${APP_NAME}"))
}
