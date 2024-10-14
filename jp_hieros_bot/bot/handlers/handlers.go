package handlers

import "gitlab.com/jp_hieros_bot/config"

type HandlerV1 struct {
	cfg config.Config
}

func NewHandlerV1(cfg config.Config) *HandlerV1 {
	return &HandlerV1{
		cfg: cfg,
	}
}

