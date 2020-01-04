package handler

import (
	"Golang-API-Game/pkg/dcontext"
	gacha_ranking "Golang-API-Game/pkg/repository/ranking"
	"Golang-API-Game/pkg/server/response"
	"errors"
	"log"
	"net/http"
)

type rankingGetResponse struct {
	Rank  int `json:"rank"`
	Score int `json:"score"`
}

func HandleRankingGet() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		// Contextから認証済みのユーザIDを取得
		ctx := request.Context()
		userID := dcontext.GetUserIDFromContext(ctx)
		if len(userID) == 0 {
			log.Println(errors.New("userID is empty"))
			response.InternalServerError(writer, "Internal Server Error")
			return
		}
		//ranking table repositoryで作成したランキング受け取る
		rankinglist, err := gacha_ranking.OrderByRank()
		if err != nil {
			log.Println("error receiving rankinglist")
			response.InternalServerError(writer, "Internal Server Error")
			return
		}
		for _, rankingget := range rankinglist {
			if userID == rankingget.UserID {
				rankresponse := rankingGetResponse{Rank: rankingget.Rank, Score: rankingget.Score}
				response.Success(writer, rankresponse)
			}
		}
	}
}
