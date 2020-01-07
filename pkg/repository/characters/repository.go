package characters

import (
	"Golang-API-Game/pkg/repository"
	"database/sql"
	"log"
)

type Character struct {
	CharacterID string
	Name        string
	Power       int
}

// CharacterID条件にレコードを取得する
func SelectByCharacterName(characterID string) (*Character, error) {
	row := repository.DB.QueryRow("SELECT * FROM characters WHERE character_id=?", characterID)
	return convertToUserCharacters(row)
}

// convertToUser rowデータをUsercharacterデータへ変換する
func convertToUserCharacters(row *sql.Row) (*Character, error) {
	character := Character{}
	err := row.Scan(&character.CharacterID, &character.Name, &character.Power) //userテーブルの三つのカラムを構造体に入れる
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		log.Println(err)
		return nil, err
	}
	return &character, nil
}

func SelectByCharacterID(characterID string) (*Character, error) {
	character := Character{}
	row := repository.DB.QueryRow("SELECT * FROM characters WHERE character_id=?", characterID)
	err := row.Scan(&character.CharacterID, &character.Name, &character.Power) //where は条件この場合gacha.goのCharacterIDを指す
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		log.Println("failed to pick Character")
		return nil, err
	}
	return &character, nil
}
