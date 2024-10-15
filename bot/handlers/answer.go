package handlers

import (
	"fmt"

	"gitlab.com/jp_hieros_bot/context"
	tb "gopkg.in/telebot.v3"
)

func (h *HandlerV1) Answer(c tb.Context) error {
	var checkAnswer bool
	fmt.Println(c.Message().Text)

	// get Answer from context data
	answer := context.AnswerByUserContext[c.Chat().ID]


	if answer.Kun == nil && answer.On == nil {
		return h.StartBot(c)
	}

	for _, val := range answer.On {
		if c.Message().Text == val {
			checkAnswer = true
			rightAnswer := "🌟 Великолепно , что это от \nОн:\t" + val

			c.Send(rightAnswer)
			return h.Challenge(c)
		}
		fmt.Println("ON -->" + val)
	}

	for _, val := range answer.Kun {
		if c.Message().Text == val {
			checkAnswer = true
			rightAnswer := "🌟 Великолепно , что это от \nКун:\t" + val

			c.Send(rightAnswer)
			return h.Challenge(c)
		}
		fmt.Println("KUN -->" + val)
	}

	if !checkAnswer {
		return c.Send("Wrong answer")

	}
	return nil
}

func (h *HandlerV1) ReturnAnswer(c tb.Context) error {
	// get Answer from context data
	answer := context.AnswerByUserContext[c.Chat().ID]

	if answer.Kun == nil && answer.On == nil {
		return h.StartBot(c)
	}

	onOn := "Он: \n"
	onKun := "\nКун: \n"

	for _, val := range answer.On {
		onOn += val + "\n"
	}
	for _, val := range answer.Kun {
		onKun += val + "\n"
	}

	textAns := onOn + onKun

	c.Send(textAns)

	return nil
}
