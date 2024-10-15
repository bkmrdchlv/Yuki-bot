package storage

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"time"

	"gitlab.com/jp_hieros_bot/models"
)

func ReadKanjiFromFile(filename string) (map[int]models.Kanji, error) {
	// Open the JSON file
	file, err := os.Open("src/kanji_char/" + filename)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %w", err)
	}
	defer file.Close()

	// Decode the JSON data
	var kanjis []models.Kanji
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&kanjis); err != nil {
		return nil, fmt.Errorf("error decoding JSON: %w", err)
	}

	// Create a map to hold ID as key and Kanji as value
	kanjiMap := make(map[int]models.Kanji)
	for _, kanji := range kanjis {
		kanjiMap[kanji.ID] = kanji
	}

	return kanjiMap, nil
}

func GenRandom() int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	randomNum := r.Intn(247) + 1

	return randomNum
}
