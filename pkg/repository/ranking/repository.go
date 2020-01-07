package ranking

import (
	"Golang-API-Game/pkg/repository"
	"database/sql"
	"log"
)

type Rank struct {
	Score  int
	UserID string
	Rank   int
}

// UpsertByPower if there were no data in ranking table insert userIDscore
// If there were already stored data update score,character_id to new character
func UpsertByPower(userID string, score int) error {
	stmt, err := repository.DB.Prepare("INSERT INTO ranking (user_id,score) VALUES (?,?) ON DUPLICATE KEY UPDATE score = (CASE WHEN score<values(score) THEN values(score) ELSE score END)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(userID, score)
	return err
}

// OrderByRank Make Ranking by userID in descending order
func OrderByRank() ([]Rank, error) {
	rows, err := repository.DB.Query("SELECT (SELECT count(b.score) FROM ranking b WHERE a.score < b.score) + 1 AS 'rank',a.user_id,a.score FROM ranking a ORDER BY rank")
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return convertToPower(rows)
}

// convert row data to powerdata
func convertToPower(rows *sql.Rows) ([]Rank, error) {
	var ranklist []Rank
	for rows.Next() {
		rank := Rank{}
		err := rows.Scan(&rank.Rank, &rank.UserID, &rank.Score) //userテーブルの三つのカラムを構造体に入れる
		if err != nil {
			log.Println("failed convert rows to ranklist ", err)
		}
		ranklist = append(ranklist, rank)
	}
	return ranklist, nil
}
