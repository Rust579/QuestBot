package logic

const (
	// Messages
	MsgStageOne = "1234"

	//Command
	CommandGetCode   = "Получить код"
	CommandStart     = "start"
	CommandReference = "Справка"
)

func ProcessMessagesText(txt string) string {

	switch txt {
	case CommandGetCode:
		return "Вот тебе код"
	case CommandReference:
		return "Сам думай"
	case MsgStageOne:
		return "Молодец, вот тебе следующие координаты: 666"
	}
	return ""
}

func ProcessMessagesCommand(com string) string {

	switch com {
	case CommandStart:
		return "Здарова, ты попал в квест чувааааааак"
	}
	return ""

}
