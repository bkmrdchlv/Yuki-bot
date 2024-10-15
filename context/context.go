package context

type Answer struct {
	ChatId int64
	On     []string
	Kun    []string
}

var AnswerByUserContext = make(map[int64]Answer)
