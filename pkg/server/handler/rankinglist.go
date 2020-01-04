package handler

import (
	"Golang-API-Game/pkg/dcontext"
	gacha_ranking "Golang-API-Game/pkg/repository/ranking"
	gacha_user "Golang-API-Game/pkg/repository/user"
	"Golang-API-Game/pkg/server/response"
	"errors"
	"log"
	"net/http"
	"strconv"
)

type rankingListGetResponse struct {
	Name  string `json:"name`
	Rank  int    `json:"rank"`
	Score int    `json:"score"`
}

type Result struct {
	Result []rankingListGetResponse `json:"result"`
}

func HandleRankingListGet() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		// Get userID from Context
		ctx := request.Context()
		userID := dcontext.GetUserIDFromContext(ctx)
		if len(userID) == 0 {
			log.Println(errors.New("userID is empty"))
			response.InternalServerError(writer, "Internal Server Error")
			return
		}
		//Get Ranking from ranking table repository
		rankinglist, err := gacha_ranking.OrderByRank()
		//rankinglist = rankinglist[0:9]→これだと数字が変わった時に対応できない
		log.Println(rankinglist)
		if err != nil {
			log.Println("error receiving rankinglist")
			response.InternalServerError(writer, "Internal Server Error")
			return
		}
		//request.URL.Queryでurl全体を取れるので、その中の必要なwordを受け取る
		var startStr, endStr string
		startStr = request.URL.Query().Get("start")
		endStr = request.URL.Query().Get("end")
		//ifunc Atoi(s string) (i int, err error)関数を使って文字列を数値に変換する
		start, err := strconv.Atoi(startStr)
		if err != nil {
			log.Println(err)
			response.InternalServerError(writer, "Internal Server Error")
			return
		}
		end, err := strconv.Atoi(endStr)
		if err != nil {
			log.Println(err)
			response.InternalServerError(writer, "Internal Server Error")
			return
		}
		//endまで取得可能
		start--
		if start >= len(rankinglist) {
			response.BadRequest(writer, "start is out of range")
			return
		}
		if end > len(rankinglist) {
			end = len(rankinglist)
		}
		//Set ranking get range of start to end in slice
		ranklist := rankinglist[start:end]
		log.Println(ranklist)
		var rankResponse []rankingListGetResponse
		for _, rankingget := range ranklist {
			log.Println(rankingget)
			// Names := *gacha_user.User
			Names, err := gacha_user.SelectByPrimaryKey(rankingget.UserID)
			if err != nil {
				response.InternalServerError(writer, "Internal Server Error") //クライアント側
				log.Println(Names)
				return
			}
			log.Println(Names)
			rankingresponse := rankingListGetResponse{Name: Names.Name, Rank: rankingget.Rank, Score: rankingget.Score}
			rankResponse = append(rankResponse, rankingresponse)
		}
		log.Println(rankResponse)
		response.Success(writer, Result{Result: rankResponse})
	}
}
