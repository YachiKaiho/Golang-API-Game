package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os/user"

	"Golang-API-Game/pkg/dcontext"
	"Golang-API-Game/pkg/repository/user"
	"Golang-API-Game/pkg/server/response"
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
		body, err := ioutil.ReadAll(request.Body)
		if err != nil {
			log.Println(err)
			response.BadRequest(writer, "Invalid Request Body")
			return
		}
		//json conversion including array
		var requestBody userUpdateRequest
		json.Unmarshal(body, &requestBody)

		////Get UserID from Context
		ctx := request.Context()
		userID := dcontext.GetUserIDFromContext(ctx)
		if len(userID) == 0 {
			log.Println(errors.New("userID is empty"))
			response.InternalServerError(writer, "Internal server error")
			return
		}

		//userTableupdate
		if error != nil {
			log.Println(err)
			response.InternalServerError(writer, "Internal Server Error")
			return
		}

		response.Success(writer, "")
	}
}

type userUpdateRequest struct {
	Name string `json:"name"`
}
