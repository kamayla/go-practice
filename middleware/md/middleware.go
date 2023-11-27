package md

import (
	"fmt"
	"net/http"
)

func MiddlewareOne(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("一つ目のMiddlewareだよ")
		next.ServeHTTP(writer, request)
	})
}

func MiddlewareTwo(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("2つ目のMiddlewareだよ")
		next.ServeHTTP(writer, request)
	})
}
