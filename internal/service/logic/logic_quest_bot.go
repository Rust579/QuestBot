package logic

import "fmt"

const (
	// Типы ответных сообщений
	TypeStr   = "str"
	TypeImg   = "img"
	TypeAudio = "audio"

	// Файлы в зависимости от этапа
	ImgFileMagaz      = "files/Magaz.png"
	AudioFileDedMoroz = "files/ded.mp3"

	// Сообщения Справки в зависимости от этапа
	ReferStage1 = "Тебе нужно 111"
	ReferStage2 = "Тебе нужно 222"
	ReferStage3 = "Тебе нужно 333"
	ReferStage4 = "Тебе нужно 444"
	ReferStage5 = "Тебе нужно 555"
	ReferStage6 = "Тебе нужно 666"
	ReferStage7 = "Тебе нужно 777"
	ReferStage8 = "Тебе нужно 888"

	// Ответные сообщения

	// Текст при повторном старте бота
	RespDubleOpening = "Ты уже начал игру"

	RespStage1  = "Приветствую, дорогие друзья\\! Подготовка к праздникам испорчена: кое\\-кто похитил подарки, а заодно и вашего друга\\! Теперь вы должны используя свои мозги и ловкость найти эти подарки и спасти своего друга, если сможете, конечно☠️☠️☠️\\. И помните, чем дольше вы действуете, тем сильнее замерзает ваш друг🥶🥶🥶\\. Удачи\\! \nПервым делом вы должны отправится по следующим координатам\\:"
	RespStage11 = " \nКоординаты копируются кликом\\. \nВ каждой точке вам необходимо найти по 3 кода, а дальше вы должны разобраться сами"
	ReqStage1   = "любой код"

	RespStage2 = "Молодцы, теперь найди 3 qr кода в одной локации и собери из них фразу\\. Координаты\\:" // Начало второго этапа
	ReqStage2  = "любая фраза"

	RespStage3 = "А теперь, если хотите спасти своего друга, вы должны прочувствовать всю глубину страха и отчаяния\\. Координаты\\:"
	ReqStage3  = "Ебать холодно"

	RespStage4 = "stage 4"
	ReqStage4  = "5"

	RespStage5 = "Пришли координаты"

	CommandStart = "start"
	ReqReference = "Справка"

	// Координаты первой точки для юзера
	gameLocation1 = "54.596341, 55.800177"
	gameLocation2 = "54.000000, 55.000000"
	gameLocKrDom  = "54.595539, 55.800322"
	gameLocMagaz  = "54.594908, 55.800422"

	// Эталонные координаты первой точки для сравнения
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

// Обработка сообщений от пользователя
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

	case ReqStage1:
		if pullStage == 1 {
			return RespMsg{
				Message: RespStage2 + fmt.Sprintf("\n\n`%v`\n", gameLocKrDom),
				Type:    TypeStr,
				Stage:   2,
			}
		}

	case ReqStage2:
		if pullStage == 2 {
			return RespMsg{
				Message:  RespStage3,
				Type:     TypeImg,
				FilePath: ImgFileMagaz,
				Stage:    3,
			}
		}

	case ReqStage3:
		if pullStage == 3 {
			return RespMsg{
				//Message:  RespStage4,
				Type:     TypeAudio,
				FilePath: AudioFileDedMoroz,
				Stage:    4,
			}
		}
	case ReqStage4:
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

// Обработка команд от пользователя (только латиница)
func ProcessMessagesCommand(com string, pullStage int) RespMsg {

	switch com {
	case CommandStart:
		if pullStage == 0 {
			return RespMsg{
				Message: RespStage1 + fmt.Sprintf("\n\n`%v`\n", gameLocation1) + fmt.Sprintf("\n`%v`\n", gameLocation2) + RespStage11,
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

// Обработка местоположения
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
