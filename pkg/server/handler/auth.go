package handler

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/google/uuid"

	"Golang-API-Game/pkg/repository/user"
	"Golang-API-Game/pkg/server/response"
)

//HandleAuthCreate 認証情報作成処理
func HandleAuthCreate() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		// RequestBodyのパース
		body, err := ioutil.ReadAll(request.Body)
		if err != nil {
			log.Println(err)
			response.BadRequest(writer, "Invalid Request Body")
			return
		}
		var requestBody authCreateRequest
		json.Unmarshal(body, &requestBody)
		//UUIDでユーザID作成
		userID, err := uuid.NewRandom()
		if err != nil {
			log.Println(err)
			response.InternalServerError(writer, "Internal Server Error")
			return
		}

		//UUIDで認証トークンを生成
		authToken, err := uuid.NewRandom()
		if err != nil {
			log.Println(err)
			response.InternalServerError(writer, "Internal Server Error")
			return
		}

		//user database string
		err := user.Insert(userID, String(), authToken.String(), requestBody.Name)
		if err != nil {
			log.Println(err)
			response.InternalServerError(writer, "Internal server error")
			return
		}
		response.Success(writer, authCreateResponse{Tokens: authToken.String()})
	}
}

type authCreateRequest struct {
	Name string `json:"name"`
}

type authCreateResponse struct {
	Token string `json:"token"`
}
