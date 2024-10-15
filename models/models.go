package models

type Kanji struct {
	ID          int      `json:"id"`
	Kanji       string   `json:"kanji"`
	Strokes     int      `json:"strokes"`
	JLPTNew     int      `json:"jlpt_new"`
	Meanings    []string `json:"meanings"`
	ReadingsOn  []string `json:"readings_on"`
	ReadingsKun []string `json:"readings_kun"`
}
