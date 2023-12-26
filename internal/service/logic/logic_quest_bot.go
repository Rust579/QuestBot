package logic

const (
	// Massages types
	TypeStr   = "str"
	TypeImg   = "img"
	TypeAudio = "audio"

	// FilePaths
	ImgFile1   = "files/123.jpg"
	AudioFile1 = "files/alt.mp3"

	// Reference messages
	ReferStage1 = "Тебе нужно 111"
	ReferStage2 = "Тебе нужно 222"
	ReferStage3 = "Тебе нужно 333"
	ReferStage4 = "Тебе нужно 444"
	ReferStage5 = "Тебе нужно 555"
	ReferStage6 = "Тебе нужно 666"
	ReferStage7 = "Тебе нужно 777"
	ReferStage8 = "Тебе нужно 888"

	// Response messages
	RespOpening      = "Здарова, ты попал в квест"
	RespDubleOpening = "Ты уже начал игру"
	RespStage2       = "Молодец, вот тебе следующие координаты: 666"
	RespStage3       = "stage 3"
	RespStage4       = "stage 4"
	RespStage5       = "Пришли координаты"

	//Command from user
	CommandStart = "start"
	ReqReference = "Справка"
	ReqStage2    = "2"
	ReqStage3    = "3"
	ReqStage4    = "4"
	ReqStage5    = "5"

	refLat1 = 54.596341
	refLon1 = 55.800177
)

type RespMsg struct {
	Message      string
	Type         string
	FilePath     string
	Stage        int
	ReferenceMsg string
}

type RefLocation struct {
	Longitude    float64
	Latitude     float64
	CorrectMsg   string
	IncorrectMsg string
}

func ProcessMessagesText(txt string, pullStage int) RespMsg {

	switch txt {
	case ReqReference:
		var message string
		switch pullStage {
		case 1:
			message = ReferStage1
		case 2:
			message = ReferStage2
		case 3:
			message = ReferStage3
		case 4:
			message = ReferStage4
		case 5:
			message = ReferStage5
		case 6:
			message = ReferStage6
		case 7:
			message = ReferStage7
		case 8:
			message = ReferStage8
		}

		return RespMsg{
			Message: message,
			Type:    TypeStr,
			Stage:   0,
		}

	case ReqStage2:
		if pullStage == 1 {
			return RespMsg{
				Message: RespStage2,
				Type:    TypeStr,
				Stage:   2,
			}
		}

	case ReqStage3:
		if pullStage == 2 {
			return RespMsg{
				Message:  RespStage3,
				Type:     TypeImg,
				FilePath: ImgFile1,
				Stage:    3,
			}
		}

	case ReqStage4:
		if pullStage == 3 {
			return RespMsg{
				Message:  RespStage4,
				Type:     TypeAudio,
				FilePath: AudioFile1,
				Stage:    4,
			}
		}
	case ReqStage5:
		if pullStage == 4 {
			return RespMsg{
				Message: RespStage5,
				Type:    TypeStr,
				Stage:   5,
			}
		}
	}
	return RespMsg{}
}

func ProcessMessagesCommand(com string, pullStage int) RespMsg {

	switch com {
	case CommandStart:
		if pullStage == 0 {
			return RespMsg{
				Message: RespOpening,
				Type:    TypeStr,
				Stage:   1,
			}
		} else {
			return RespMsg{
				Message: RespDubleOpening,
				Type:    TypeStr,
				Stage:   1,
			}
		}
	}
	return RespMsg{}
}

func ProcessLocation(pullStage int) RefLocation {

	switch pullStage {
	case 5:
		return RefLocation{
			Latitude:     refLat1,
			Longitude:    refLon1,
			CorrectMsg:   "Молодец! Теперь ты должен 555",
			IncorrectMsg: "Координаты неверные",
		}
	}

	return RefLocation{}

}
