package handlers

import (
	tb "gopkg.in/telebot.v3"
)

func (h *HandlerV1) StartBot(c tb.Context) error {
	text := `Привет 🖐 
Этот бот для изучения японских
иероглифов для старта нажмите /challenge
Здесь иероглифы  JLPT 5 и 4 уровня `
	return c.Send(text)
}
