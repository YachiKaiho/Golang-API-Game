package power

import (
	"Golang-API-Game/pkg/repository"
	"database/sql"
	"log"
)

// powerテーブルデータ
type Power struct {
	UserCharacterID string
	Power           int
}

type Totalpower struct {
	sum int
}

//powerを合計したrowデータをとる
func OddsSum() (int, error) {
	row := repository.DB.QueryRow("SELECT SUM(odds) FROM power")
	return convertToInt(row)
}

// convertToInt rowデータを整数値に変換する
func convertToInt(row *sql.Row) (int, error) { //row型からint型へ変換する必要がある。
	total := Total{} //varのように構造体の宣言方法(ゼロ値)
	err := row.Scan(&total.sum)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, nil //int型だからゼロ値
		}
		log.Println(err)
		return 0, err
	}
	return total.sum, nil //int型だから,intのポイント型を返す場合
}
