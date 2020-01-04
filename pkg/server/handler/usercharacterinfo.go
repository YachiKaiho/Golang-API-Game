package handler

import (
	"Golang-API-Game/pkg/dcontext"
	"Golang-API-Game/pkg/repository/characters"
	"Golang-API-Game/pkg/server/response"
	"errors"
	"log"
	"net/http"
)

type UserCharacters struct {
	UserID          string
	UserCharacterID string
	CharacterID     string
}

type CharactersResponse struct {
	UserCharacterID string `json:"userCharacterID"`
	CharacterID     string `json:"characterID"`
	Name            string `json:"name"`
}

type userCharacterListResponses struct {
	Characters []CharactersResponse `json:"characters"` //何個も入れられる
}

//HandleUserCharacterInfoUpdate ユーザ情報更新処理
func HandleCharacterList() http.HandlerFunc { //すでにuser.goで使われているのでUserUpdateではなく違うものにした方がわかりやすい
	return func(writer http.ResponseWriter, request *http.Request) { //ヘッダーとボディをrequestは持っている()
		// Contextから認証済みのユーザIDを取得
		ctx := request.Context()
		userID := dcontext.GetUserIDFromContext(ctx)
		if userID == "" {
			log.Println(errors.New("userID is empty"))
			response.InternalServerError(writer, "Internal Server Error")
			return
		}
		//userIDによってuserCharacter型のスライスを返す
		userCharacterList, err := gachaUserCharacter.SelectByUserID(userID)
		if err != nil {
			log.Println(err)
			response.InternalServerError(writer, "Internal Server Error")
			return
		}
		var charaResponse []CharactersResponse
		for _, userCharacter := range userCharacterList {
			log.Println("failed to get")
			var character *characters.Character
			character, err = characters.SelectByCharacterName(userCharacter.CharacterID)
			if err != nil {
				response.InternalServerError(writer, "Internal Server Error") //クライアント側
				log.Println(character)
				return
			}
			charaRes := CharactersResponse{userCharacter.UserCharacterID, userCharacter.CharacterID, character.Name}
			charaResponse = append(charaResponse, charaRes)
		}
		response.Success(writer, userCharacterListResponses{Characters: charaResponse})
	}
}
