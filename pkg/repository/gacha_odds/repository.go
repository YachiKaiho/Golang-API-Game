package gacha_odds

import (
	"Golang-API-Game/pkg/repository"
	"database/sql"
	"log"
)

// gacha_oddsテーブルデータ
type GachaOdds struct {
	CharacterID string
	Odds        int
}

type Total struct {
	sum int
}

//gacha_OddsTableのOddsを合計したrowデータをとる
func OddsSum() (int, error) { //selectにするとoddsのsumを取ってくるという意味になってしまう
	row := repository.DB.QueryRow("SELECT SUM(odds) FROM gacha_odds") //*だと全カラム,?は引数を受け取るときに必要
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

// convertToGachaOdds rowデータをGachaOddsデータへ変換する
func convertToGachaOdds(row *sql.Row) (*GachaOdds, error) {
	gacha_odds := GachaOdds{}
	err := row.Scan(&gacha_odds.CharacterID, &gacha_odds.Odds)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		log.Println(err)
		return nil, err
	}
	return &gacha_odds, nil //*GachaOdds,ポイント型で返す必要があるから
}

func SelectByRandomNumber(random int) (*GachaOdds, error) {
	rows, err := repository.DB.Query("SELECT * FROM  gacha_odds") //ROWをつけると単一の行のみになってしまう
	if err != nil {
		log.Println("failed to gacha", err)
	}
	gacha_odds := GachaOdds{} //メモリ消費を抑える
	for rows.Next() {
		err = rows.Scan(&gacha_odds.CharacterID, &gacha_odds.Odds) //scanによってgacha_oddsに格納しているからエラーを返すだけでいい
		if err != nil {
			log.Println(gacha_odds)
		}
		random -= gacha_odds.Odds
		if random <= 0 {
			break //
		}
	}
	return &gacha_odds, nil
}
