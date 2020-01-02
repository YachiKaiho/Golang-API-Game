package middleware

import (
	"context"
	"log"
	"net/http"

	"Golang-API-Game/pkg/dcontext"
	"Golang-API-Game/pkg/repository/user"
	"Golang-API-Game/pkg/server/response"
)

// Authenticate userID and Store UserID info to Context
func Authenticate(nextFunc http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {

		ctx := request.Context()
		if ctx == nil {
			ctx = context.Background()
		}

		// get x-token from request header
		token := request.Header.Get("x-token")
		if len(token) == 0 {
			log.Println("x-token is empty")
			return
		}

		// get user information linked to auth token from database
		// input SELECT query
		user, err := user.SelectByAuthToken(token)
		if err != nil {
			log.Println(err)
			response.InternalServerError(writer, "Invalid token")
			return
		}
		if user == nil {
			log.Printf("user not found. token=%s", token)
			response.BadRequest(writer, "Invalid token")
			return
		}

		// Store userId to Context and use to next process
		ctx = dcontext.SetUserID(ctx, user.UserID)

		// the next process
		nextFunc(writer, request.WithContext(ctx))
	}
}
