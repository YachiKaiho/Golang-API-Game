package handler

import (
	"fmt"
	"log"
	"net/http"
	"os/user"

	"GOLANG-API-GAME/pkg/repository/user"
	"GOLANG-API-GAME/pkg/server/response"
)

//User information Get
func HandleUserGet() http.HandleFunc {
	return func(writer http.ResponseWriter, request *http.Request) {

		//Get UserID from Context
		ctx := request.Context(ctx)
		if len(userID) == 0 {
			log.Println("userID is empty!")
			response.InternalServerError(writer, "Couldn't get userID from Context,Internal Server error!")
			return
		}

		//UserData　Management
		var user *user.User
		var err error
		if err != nil {
			log.Println(err)
			response.InternalServerError(writer, "Internal Server Error")
			return
		}

		if user == nil {
			log.Println(err)
			response.BadRequest(writer, fmt.Sprintf("user not found. userID=%s", userID))
			return
		}

		//response with required information
		response.Success(writer, userGetResponse{Name: user.Name})
	}
}

type userGetResponse struct {
	Name string `json:"name"`
}

//HandleUserUpdate Management
func HandleUserUpdate() http.HandleFunc {
	return func(writer http.ResponseWriter, request *http.Request) {

		//Get Updated information from RequestBody

	}
}