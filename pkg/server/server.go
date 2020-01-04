package server

import (
	"log"

	"GOLANG-API-GAME/pkg/handler"
	"GOLANG-API-GAME/pkg/middleware"
	"net/http"
)

//Serve HTTP server run
func Serve(addr string) {

	//URLマッピングを行う
	http.HandleFunc("/auth/create", post(handler.HandleAuthCreate()))
	//middlewares
	http.Handle("/user/get",
		get(middleware.Authenticate(handler.HandleUserGet())))

	http.HandleFunc("/user/update",
		post(middleware.Authenticate(handler.HandleUserUpdate())))

	http.HandleFunc("/gacha/draw",
		post(middleware.Authenticate(handler.HandleGachaUpdate())))

	http.HandleFunc("/character/list",
		get(middleware.Authenticate(handler.HandleCharacterList())))

	http.HandleFunc("/ranking/get",
		get(middleware.Authenticate(handler.HandleRankingGet())))

	http.HandleFunc("/ranking/list",
		get(middleware.Authenticate(handler.HandleRankingListGet())))

	http.HandleFunc("/power/get",
		get(middleware.Authenticate(handler.HandlePowerGet())))

	/* ===== server run ===== */
	log.Println("Server running...")
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatalf("Listen and serve failed. %+v", err)
	}
}

// get GETリクエストを処理する
func get(apiFunc http.HandlerFunc) http.HandlerFunc {
	return httpMethod(apiFunc, http.MethodGet)
}

// post POSTリクエストを処理する
func post(apiFunc http.HandlerFunc) http.HandlerFunc {
	return httpMethod(apiFunc, http.MethodPost)
}

// httpMethod 指定したHTTPメソッドでAPIの処理を実行する
func httpMethod(apiFunc http.HandlerFunc, method string) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {

		// CORS対応
		writer.Header().Add("Access-Control-Allow-Origin", "*")
		writer.Header().Add("Access-Control-Allow-Headers", "Content-Type,Accept,Origin,x-token")

		// プリフライトリクエストは処理を通さない(optionメソッドに対応する)
		if request.Method == http.MethodOptions {
			return
		}
		// 指定のHTTPメソッドでない場合はエラー
		if request.Method != method {
			writer.WriteHeader(http.StatusMethodNotAllowed)
			writer.Write([]byte("Method Not Allowed"))
			return
		}

		// 共通のレスポンスヘッダを設定
		writer.Header().Add("Content-Type", "application/json")
		apiFunc(writer, request)
	}
}
