package handler

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"github.com/google/uuid"

	"GOLANG-API-GAME/repository/user"
	"GOLANG-API-GAME/server/response"
)

func HandleAuthCreate() http.HandleFunc{
	  return func(writer http.ResponseWriter, request *http.Request){
	       //RequestsBody Parse
	       body,err := ioutil.ReadAll(request.Body)
	       if err != nillog.Println(err)
	       response.BadRequest(writer, "Invalid RequestBody")
		   return
		   
	  var requestBody authCreateRequest
	  json.Unmarshal(body, &requestBody)

	  //UUIDでユーザID作成
	  userID,err := uuid.NewRandom()
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
	  err := user.Insert(userID,String(), authToken.String(), requestBody.Name)
	  if err != nil{
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