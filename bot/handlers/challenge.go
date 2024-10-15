package handlers

import (
	"fmt"
	"strings"

	"gitlab.com/jp_hieros_bot/context"
	"gitlab.com/jp_hieros_bot/storage"
	tb "gopkg.in/telebot.v3"
)

func (h *HandlerV1) Challenge(c tb.Context) error {
	// read all characters
	kanji, err := storage.ReadKanjiFromFile("kanji_4_and_5.json")
	if err != nil {
		fmt.Printf("Failed to read kanji data: %v", err)
	}

	// get random character
	char := kanji[storage.GenRandom()]
	fmt.Println(char)

	// set answer
	context.AnswerByUserContext[c.Chat().ID] = context.Answer{
		ChatId: c.Chat().ID,
		On:     char.ReadingsOn,
		Kun:    char.ReadingsKun,
	}
	fmt.Println("ANSWER::", context.AnswerByUserContext)

	// get photo by random character
	var charName = fmt.Sprintf("src/kanji_media/static/150x150/%v.png", char.Kanji)
	fmt.Println(charName)

	// complete message
	captionText := fmt.Sprintf("%v \nУровень Jlpt: %v \nЗначения: %v\nПожалуйста, напишите  иероглифы на Он или Кун❗",
		char.Kanji,
		char.JLPTNew,
		joinStringsOptimized(char.Meanings))

	photo := &tb.Photo{File: tb.FromDisk(charName), Caption: captionText}

	// answer button
	textOnButton := "🔎 Ответ"

	// send
	return c.Send(photo,
		&tb.ReplyMarkup{InlineKeyboard: [][]tb.InlineButton{{{Text: textOnButton, Data: "#Answer"}}}})

}

func joinStringsOptimized(stringArray []string) string {
	// Preallocate memory for performance
	var builder strings.Builder
	builder.Grow(len(stringArray) * 2) // Rough estimate for ", \n"

	for i, str := range stringArray {
		if i > 0 {
			builder.WriteString(", \n") // Add your separator here
		}
		builder.WriteString(str)
	}

	return builder.String()
}
