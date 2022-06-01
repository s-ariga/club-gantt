/**
 * Seiichi Ariga <seiichi.ariga@gmail.com>
 */

package main

import (
	"os"

	"github.com/s-ariga/club-gantt/csv"
	"github.com/s-ariga/club-gantt/draw"
)

func checkFileExists(filePath string) bool {
	_, err := os.Stat(filePath)
	return err == nil
}

func main() {
	if len(os.Args) < 2 {
		panic("Usage: go run main.go <file>")
	}
	filePath := os.Args[1]
	if !checkFileExists(filePath) {
		panic("File not found: " + filePath)
	}

	events, err := csv.ReadCsv(filePath)
	if err != nil {
		panic("Csv読み込みエラー: " + err.Error())
	}

	if draw.DrawEventsGantt(events, "gantt.svg") != nil {
		panic("Drawing error: " + err.Error())
	}
}
