package user_characters

import (
	"Golang-API-Game/pkg/repository"
	"database/sql"
	"log"
)

//UserCharacter table data
type UserCharacter struct {
	UserID          string
	UserCharacterID string
	CharacterID     string
}

type User struct {
	UserID string
	Result string
}

type Character struct {
	name string
}

//Insert table data
func Insert(userID, userCharacterID, characterID string) error {
	stmt, err := repository.DB.Prepare("INSERT INTO user_characters(user_id, user_character_id, character_id) VALUES(?, ?, ?)") //prepare 少し早くなる
	if err != nil {
		return err
	}
	_, err = stmt.Exec(userID, userCharacterID, characterID)
	return err
}

// SelectByPrimaryKeuser Get character_id,character_id conditon:userID
func SelectByPrimaryKey(userID string) (*UserCharacter, error) {
	row := repository.DB.QueryRow("SELECT * FROM user_characters WHERE user_id=?", userID)
	return convertToUserCharacter(row)
}

// convert row data to Usercharacter data
func convertToUserCharacter(row *sql.Row) (*UserCharacter, error) {
	usercharacter := UserCharacter{}
	err := row.Scan(&usercharacter.UserID, &usercharacter.UserCharacterID, &usercharacter.CharacterID) //userテーブルの三つのカラムを構造体に入れる
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		log.Println(err)
		return nil, err
	}
	return &usercharacter, nil
}

// UpdateByPrimaryKey
func UpdateByPrimaryKey(userCharacterID string, random string) error {
	stmt, err := repository.DB.Prepare("UPDATE user_characters SET user_character_id=? WHERE user_id=?")
	if err != nil {
		log.Println(err)

		return err
	}
	_, err = stmt.Exec(userCharacterID, random)
	return err
}

// SelectByUserID
func SelectByUserID(userID string) ([]UserCharacter, error) {
	rows, err := repository.DB.Query("SELECT * FROM user_characters WHERE user_id=?", userID)
	if err != nil {
		log.Println(err)

		return nil, err
	}
	return convertToUserCharacterID(rows)
}

// convert row data to Usercharacter data
func convertToUserCharacterID(rows *sql.Rows) ([]UserCharacter, error) {
	var userCharacterList []UserCharacter
	for rows.Next() {
		userCharacter := UserCharacter{}
		err := rows.Scan(&userCharacter.UserID, &userCharacter.UserCharacterID, &userCharacter.CharacterID) //userテーブルの三つのカラムを構造体に入れる
		if err != nil {
			log.Println("failed convert rows to CharacterID ", err)
		}
		userCharacterList = append(userCharacterList, userCharacter)
	}
	return userCharacterList, nil
}
