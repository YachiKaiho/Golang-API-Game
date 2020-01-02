package server

import (
	"log"

	"net/http"
	"GOLANG-API-GAME/pkg/handler"
）

//Serve HTTPサーバ起動
func Serve(addr string) {

	//URLマッピングを行う
	http.HandleFunc("/auth/create", post(handler.HandleAuthCreate()))

	http.Handle("/user/get",
		get(handler.HandleUserGet()))
	http.Handle()
)
}