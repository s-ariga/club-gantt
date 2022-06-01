/**
 * Seiichi Ariga <seiichi.ariga@gmail.com>
 */

// 射座割りのCSVファイルを読み込む
package csv

import (
	"encoding/csv"
	"os"
	"strconv"
	"time"
)

// Event構造体がCSVファイルの並び順を決定する
type Event struct {
	Name        string
	StartTime   time.Time
	EndTime     time.Time
	TargetStart int
	TargetEnd   int
}

// CSVファイルの読み込み
func ReadCsv(filePath string) ([]Event, error) {

	records, err := readData(filePath)
	if err != nil {
		panic(err)
	}
	/** CSVファイルにはEvent構造体がそのまま入っているので、
	 * それをEventに変換
	 */
	events := make([]Event, len(records))
	for i, record := range records {
		events[i].Name = record[0]
		events[i].StartTime, err = time.Parse("2006-01-02 01:02", record[1])
		if err != nil {
			return []Event{}, err
		}
		events[i].EndTime, err = time.Parse("2006-01-02 01:02", record[2])
		if err != nil {
			return []Event{}, err
		}
		events[i].TargetStart, err = strconv.Atoi(record[3])
		if err != nil {
			return []Event{}, err
		}
		events[i].TargetEnd, err = strconv.Atoi(record[4])
		if err != nil {
			return []Event{}, err
		}
	}
	return events, nil
}

// Csvの実際の読み込み
func readData(filePath string) ([][]string, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return [][]string{}, err
	}
	defer f.Close()

	r := csv.NewReader(f)

	if _, err := r.Read(); err != nil {
		return [][]string{}, err
	}
	records, err := r.ReadAll()
	if err != nil {
		return [][]string{}, err
	}
	return records, nil
}
