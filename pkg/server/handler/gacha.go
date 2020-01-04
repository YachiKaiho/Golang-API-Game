package handler

import (
	"Golang-API-Game/pkg/dcontext"
	"Golang-API-Game/pkg/repository/characters"
	"Golang-API-Game/pkg/repository/user_characters"
	"Golang-API-Game/pkg/server/response"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/google/uuid"
)

type gachaDrawRequest struct {
	Times int `json:"times"`
}

type gachaDrawResult struct {
	CharacterID string `json:"characterID"`
	Name        string `json:"name"`
}

type userGetResponses struct {
	Result []gachaDrawResult `json:"result"` //何個も入れられる
}

//HandleGachaUpdate User gachainfo update
func HandleGachaUpdate() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		body, err := ioutil.ReadAll(request.Body)
		log.Println(string(body))
		if err != nil {
			log.Println(err)
			response.BadRequest(writer, "Invalid Request Body")
			return
		}
		var requestBody gachaDrawRequest
		json.Unmarshal(body, &requestBody)
		log.Println(requestBody)
		// Get userID from Context
		ctx := request.Context()
		userID := dcontext.GetUserIDFromContext(ctx)
		if len(userID) == 0 {
			log.Println(errors.New("userID is empty"))
			response.InternalServerError(writer, "Internal Server Error")
			return
		}
		var sum, randoms int
		//Odds sum
		sum, err = gacha_gacha_odds.OddsSum()

		var gacha_odds *dojo_gacha_odds.GachaOdds
		//Make is used for only data is known,prevent from append repeated
		characterget := make([]gachaDrawResult, 0, requestBody.Times)
		rand.Seed(time.Now().UnixNano())
		for i := 0; i < requestBody.Times; i++ {
			randoms = rand.Intn(sum + 1)
			gacha_odds, err = dojo_gacha_odds.SelectByRandomNumber(randoms)
			// Make usercharacterID by uuid
			userCharacterID, err := uuid.NewRandom()
			character, err := characters.SelectByCharacterID(gacha_odds.CharacterID)
			if err != nil {
				log.Println("append(characterID,name)failed")
				return
			}
			var list gachaDrawResult
			list.CharacterID = gacha_odds.CharacterID
			list.Name = character.Name
			characterget = append(characterget, list)
			user_characters.Insert(userID, userCharacterID.String(), gacha_odds.CharacterID)
			characterinfo, err := dojo_character.SelectByCharacterID(gacha_odds.CharacterID)
			if err != nil {
				log.Println(err)
				return
			}
			score := characterinfo.Power
			err = dojo_ranking.UpsertByPower(userID, score)
			if err != nil {
				log.Println(err)
				return
			}
		}
		response.Success(writer, userGetResponses{Result: characterget})
	}
}
