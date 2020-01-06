package user

import (
	"Golang-API-Game/pkg/repository"
	"database/sql"
	"log"
)

// User usertable data
type User struct {
	UserID    string
	AuthToken string
	Name      string
}

// Insert database
func Insert(userID, authToken, name string) error {
	stmt, err := repository.DB.Prepare("INSERT INTO user(user_id, auth_token, name) VALUES(?, ?, ?)") //prepare 少し早くなる
	if err != nil {
		return err
	}
	_, err = stmt.Exec(userID, authToken, name)
	return err
}

// SelectByAuthToken conditon:auth_token
func SelectByAuthToken(authToken string) (*User, error) {
	row := repository.DB.QueryRow("SELECT * FROM user WHERE auth_token=?", authToken)
	return convertToUser(row)
}

// UpdateByPrimaryKey
func UpdateByPrimaryKey(userID string, name string) error {
	stmt, err := repository.DB.Prepare("UPDATE user SET name=? WHERE user_id=?")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(name, userID)
	return err
}

// SelectByPrimaryKey get auserIDauth_token,name from user table
func SelectByPrimaryKey(userID string) (*User, error) {
	row := repository.DB.QueryRow("SELECT * FROM user WHERE user_id=?", userID)
	return convertToUser(row)
}

// convertToUser convert rowdata to User data
func convertToUser(row *sql.Row) (*User, error) {
	user := User{}
	err := row.Scan(&user.UserID, &user.AuthToken, &user.Name)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		log.Println(err)
		return nil, err
	}
	return &user, nil
}
