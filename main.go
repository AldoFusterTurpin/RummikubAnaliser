package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

type DayInfo struct {
	Date    string  `json:"date"`
	Matches []Match `json:"matches"`
}

type Match struct {
	MatchScores []PlayerScore `json:"match_scores"`
}

type PlayerScore struct {
	PlayerName string `json:"player_name"`
	Score      int    `json:"score"`
}

func main() {
	// filePath := "rummikub_data.json"
	filePath := "simple_data.json"

	daysInfo := readData(filePath)

	result := analyseData(daysInfo)
	result.Print()
}

func readData(filePath string) []DayInfo {
	jsonFile, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer jsonFile.Close()

	bytes, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		log.Fatal(err)
	}

	var daysInfo []DayInfo
	json.Unmarshal(bytes, &daysInfo)

	return daysInfo
}

func analyseData(daysInfo []DayInfo) *Result {
	r := &Result{
		TotalScore: 0,
		Scores:     make(map[string]int),
		MatchesWon: make(map[string]int),
	}

	for _, dayInfo := range daysInfo {
		for _, match := range dayInfo.Matches {
			analyseMatch(match, r)
		}
	}
	return r
}

func analyseMatch(match Match, r *Result) {
	totalScoreOfThatMatch := 0
	playerThatWonTheMatch := ""

	for _, playerScore := range match.MatchScores {
		score := playerScore.Score
		name := playerScore.PlayerName

		// Needed if the input score comes in negative form.
		if score <= 0 {
			score *= -1
		}

		totalScoreOfThatMatch += score

		r.TotalScore -= score

		currentPlayerWonTheMatch := score == 0
		if currentPlayerWonTheMatch {
			playerThatWonTheMatch = name
			r.MatchesWon[name]++
		} else {
			// This does not need to be in "else" as x -= 0 has no effect,
			// but IMHO it is more clear this way.
			r.Scores[name] -= score
		}
	}
	r.Scores[playerThatWonTheMatch] += totalScoreOfThatMatch
}
