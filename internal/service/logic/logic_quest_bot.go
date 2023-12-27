package logic

import (
	"fmt"
	"strings"
)

const (
	// Типы ответных сообщений
	TypeStr   = "str"
	TypeImg   = "img"
	TypeAudio = "audio"

	// Файлы в зависимости от этапа
	ImgFileKrDom      = "files/krdom.png"
	AudioFileDedMoroz = "files/ded.mp3"
	ImgFileOblast     = "files/oblast.png"

	// Ответные сообщения

	// Текст при повторном старте бота
	RespDubleOpening = "Ты уже начал игру"

	RespStage1  = "Приветствую, дорогие друзья\\! Подготовка к праздникам испорчена: кое\\-кто похитил подарки, а заодно и вашего друга\\! Теперь вы должны используя свои мозги и ловкость найти эти подарки и спасти своего друга, если сможете, конечно☠️☠️☠️\\. И помните, чем дольше вы действуете, тем сильнее замерзает ваш друг🥶🥶🥶\\. Удачи\\! \nПервым делом вы должны отправится по следующим координатам\\:"
	RespStage11 = " \nКоординаты копируются кликом\\. \nСоберите из всех точек один код 💀 🧱 ⚡️\\."
	ReqStage1   = "473"

	RespStage2 = "У вас получилось, поздравляю\\! Теперь поиграем в горячо\\-холодно\\: вы должны отправлять мне свое местоположение, а я буду говорить близко вы или далеко, в целевой локации найдите один код\\. \nМестоположение нужно отправлять через вложения \\-\\> геопозиция, для удобства можно выбрать слой спутник\\."
	ReqStage2  = "6"

	RespStage3 = "Молодцы, теперь найдите еще 3 кода в одной локации и соберите из них фразу\\. Коды черным маркером с префиксом HNY\\. Координаты\\:" // Начало второго этапа
	ReqStage3  = "ваш друг рядом"

	RespStage4 = "А теперь, если хотите спасти своего друга, вы должны прочувствовать всю глубину страха и отчаяния\\."
	ReqStage4  = "моя прелесть"

	RespStage5 = "stage 4"
	ReqStage5  = "5"

	CommandStart = "start"
	ReqReference = "текущее задание"

	// Координаты первой точки для юзера
	gameLocTochka1 = "54.596840, 55.801314" // столб кирпич
	gameLocTochka2 = "54.595245, 55.802220" // столб молния
	gameLocTochka3 = "54.594835, 55.801694" // столб череп
	gameLocMagaz   = "54.594884, 55.800410" // недостроенный магазин
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

	txt = strings.ToLower(txt)

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
				FilePath: ImgFileOblast,
				Stage:    2,
			}
		case 3:
			return RespMsg{
				Message: RespStage3 + fmt.Sprintf("\n\n`%v`\n", gameLocMagaz),
				Type:    TypeStr,
				Stage:   3,
			}
		case 4:
			return RespMsg{
				Message:  RespStage4,
				Type:     TypeImg,
				FilePath: ImgFileKrDom,
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
				FilePath: ImgFileOblast,
				Stage:    2,
			}
		}

	case ReqStage2:
		if pullStage == 2 {
			return RespMsg{
				Message: RespStage3 + fmt.Sprintf("\n\n`%v`\n", gameLocMagaz),
				Type:    TypeStr,
				Stage:   3,
			}
		}

	case ReqStage3:
		if pullStage == 3 {
			return RespMsg{
				Message:  RespStage4,
				Type:     TypeImg,
				FilePath: ImgFileKrDom,
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
