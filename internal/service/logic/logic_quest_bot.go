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
	ImgFileOblast200  = "files/Location200.png"

	// Ответные сообщения

	// Текст при повторном старте бота
	RespDubleOpening = "Ты уже начал игру"

	RespStage1  = "Приветствую, дорогие друзья\\! Подготовка к праздникам испорчена: кое\\-кто похитил подарки, а заодно и вашего друга\\! Теперь вы должны используя свои мозги и ловкость найти эти подарки и спасти своего друга, если сможете, конечно☠️☠️☠️\\. И помните, чем дольше вы действуете, тем сильнее замерзает ваш друг🥶🥶🥶\\. Удачи\\! \nПервым делом вы должны отправится по следующим координатам\\:"
	RespStage11 = " \nКоординаты копируются кликом\\. \nВ каждой точке вам необходимо найти по 1 коду, а дальше вы должны разобраться сами"
	ReqStage1   = "любой код"

	RespStage2 = "У вас получилось, поздравляю\\! Теперь поиграем в горячо\\-холодно\\: вы должны отправлять мне свое местоположение, а я буду говорить близко вы или далеко, в целевой локации найдите один код\\. \nМестоположение нужно отправлять через вложения \\(значок скрепки\\)\\."
	ReqStage2  = "еще код"

	RespStage3 = "Молодцы, теперь найди 3 qr кода в одной локации и собери из них фразу\\. Координаты\\:" // Начало второго этапа
	ReqStage3  = "любая фраза"

	RespStage4 = "А теперь, если хотите спасти своего друга, вы должны прочувствовать всю глубину страха и отчаяния\\."
	ReqStage4  = "Ебать холодно"

	RespStage5 = "stage 4"
	ReqStage5  = "5"

	CommandStart = "start"
	ReqReference = "Справка"

	// Координаты первой точки для юзера
	gameLocTochka1 = "54.596341, 55.800177"
	gameLocTochka2 = "54.000000, 55.000000"
	gameLocTochka3 = "54.111111, 55.111111"
	gameLocKrDom   = "54.595539, 55.800322"
	gameLocMagaz   = "54.594908, 55.800422"

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

// Обработка сообщений от пользователя
func ProcessMessagesText(txt string, pullStage int) RespMsg {

	switch txt {
	case ReqReference:
		var message string
		switch pullStage {
		case 1:
			return RespMsg{
				Message: RespStage1 + fmt.Sprintf("\n\n`%v`\n", gameLocTochka1) + fmt.Sprintf("\n`%v`\n", gameLocTochka2) + fmt.Sprintf("\n`%v`\n", gameLocTochka3) + RespStage11,
				Type:    TypeStr,
				Stage:   1,
			}
		case 2:
			return RespMsg{
				Message:  RespStage2,
				Type:     TypeImg,
				FilePath: ImgFileOblast200,
				Stage:    2,
			}
		case 3:
			return RespMsg{
				Message: RespStage3 + fmt.Sprintf("\n\n`%v`\n", gameLocKrDom),
				Type:    TypeStr,
				Stage:   3,
			}
		case 4:
			return RespMsg{
				Message:  RespStage4,
				Type:     TypeImg,
				FilePath: ImgFileMagaz,
				Stage:    4,
			}
		case 5:
			return RespMsg{
				//Message:  RespStage4,
				Type:     TypeAudio,
				FilePath: AudioFileDedMoroz,
				Stage:    5,
			}
		case 6:
			message = "Э"
		case 7:
			message = "Э"
		case 8:
			message = "Э"
		}

		return RespMsg{
			Message: message,
			Type:    TypeStr,
			Stage:   0,
		}

	case ReqStage1:
		if pullStage == 1 {
			return RespMsg{
				Message:  RespStage2,
				Type:     TypeImg,
				FilePath: ImgFileOblast200,
				Stage:    2,
			}
		}

	case ReqStage2:
		if pullStage == 2 {
			return RespMsg{
				Message: RespStage3 + fmt.Sprintf("\n\n`%v`\n", gameLocKrDom),
				Type:    TypeStr,
				Stage:   3,
			}
		}

	case ReqStage3:
		if pullStage == 3 {
			return RespMsg{
				Message:  RespStage4,
				Type:     TypeImg,
				FilePath: ImgFileMagaz,
				Stage:    4,
			}
		}

	case ReqStage4:
		if pullStage == 4 {
			return RespMsg{
				//Message:  RespStage4,
				Type:     TypeAudio,
				FilePath: AudioFileDedMoroz,
				Stage:    5,
			}
		}
	case ReqStage5:
		if pullStage == 5 {
			return RespMsg{
				Message: "Э",
				Type:    TypeStr,
				Stage:   6,
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
				Message: RespStage1 + fmt.Sprintf("\n\n`%v`\n", gameLocTochka1) + fmt.Sprintf("\n`%v`\n", gameLocTochka2) + fmt.Sprintf("\n`%v`\n", gameLocTochka3) + RespStage11,
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
