package logic

import (
	"fmt"
	"strings"
)

const (
	// Типы ответных сообщений
	TypeStr   = "str"
	TypeImg   = "img"
	TypeImgs  = "imgs"
	TypeAudio = "audio"

	// Файлы в зависимости от этапа
	ImgFileKrDom  = "files/krdom.jpg"
	ImgFileOblast = "files/oblast.png"
	ImgNews1      = "files/news1.jpg"
	ImgNews2      = "files/news2.jpg"
	ImgNews3      = "files/news3.jpg"
	ImgNews4      = "files/news4.png"
	ImgNews5      = "files/news5.png"
	ImgNews6      = "files/news6.png"
	ImgNews7      = "files/news7.png"

	AudioFileDedMoroz     = "files/ded.mp3"
	AudioFileDedMorozName = "Это еще не конец"
	AudioFileShturval     = "files/shturval.mp3"
	AudioFileShturvalName = "Очень горячо"

	// Ответные сообщения

	// Текст при повторном старте бота
	RespDubleOpening = "Ты уже начал игру"

	RespStage1  = "Не доброго вечера вам\\! Поздравляю, вы даже не заметили как пропал ваш друг, а с ним и все подарки\\.\n\nЧтобы спасти друга и подарки, вы должны использовать свои мозги и ловкость, если сможете, конечно☠️☠️☠️\\. И помните, чем дольше вы действуете, тем сильнее замерзает ваш друг🥶🥶🥶\\. Удачи\\!\n\nПервым делом вы должны отправится по следующим координатам\\:"
	RespStage11 = " \n_Координаты копируются кликом\\. \nСоберите из всех точек один код_ 💀 🧱 ⚡️\\."
	ReqStage1   = "473"

	RespStage2 = "У вас получилось, что же, это было просто\\! Теперь поиграем в горячо\\-холодно\\: вы должны отправлять мне свое местоположение, а я буду говорить близко вы или далеко, в целевой локации найдите один код\\. \n\n_Местоположение нужно отправлять через вложения \\-\\> геопозиция, для удобства можно выбрать слой спутник_\\."
	ReqStage2  = "6"

	RespStage3 = "Хм, у вас получилось\\.\nТеперь найдите еще 3 кода в одной локации и соберите из них фразу\\.\n\n_Коды красным маркером с префиксом HNY\\.\nКоординаты\\:_" // Начало второго этапа
	ReqStage3  = "ваш друг рядом"

	RespStage4 = "А теперь, если хотите спасти своего друга, вы должны прочувствовать всю глубину страха и отчаяния\\."
	ReqStage4  = "моя прелесть"

	ReqStage5 = "подарки мои"

	RespStage6   = "Понравились напитки\\? Да или нет\\?"
	ReqStage6Net = "нет"
	ReqStage6Da  = "да"

	RespStage7Pidor   = "Пидора ответ\\! \nПонравились напитки\\? Да или нет\\?"
	RespStage7NePidor = "Замечательно\\! А теперь проверим, насколько хорошо вы помните этот год:"
	ReqStage7         = "8086930"

	RespStage8 = "Говорят, что у истории о трех свиньях счастливый конец, но на самом деле волк не оставляет попыток найти их\\. До него дошли слухи, что каждый свин спрятался в разных местах\\: \n\n" +
		"\\- Первый уехал в жаркие страны\n\n" +
		"\\- Второй устроился работать на очистные сооружения\n\n" +
		"\\- Третий стал гримёром"
	ReqStage8  = "366"
	RespStage9 = "Вы показали себя стойкими, упорными и находчивыми\\. Теперь у вас есть все, чтобы отыскать подарки\\.\nПоздравляю\\!\n\n" +
		"Подарки ищите в темном, холодном месте, за стальной стеной\\."

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
	FileName     string
	Stage        int
	ReferenceMsg string
	Images       []string
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
				FileName: AudioFileDedMorozName,
				Stage:    5,
			}
		case 6:
			return RespMsg{
				Message: RespStage6,
				Type:    TypeStr,
				Stage:   6,
			}
		case 7:
			return RespMsg{
				Message: RespStage7NePidor,
				Type:    TypeImgs,
				Stage:   7,
				Images:  []string{ImgNews4, ImgNews1, ImgNews7, ImgNews5, ImgNews2, ImgNews6, ImgNews3},
			}
		case 8:
			return RespMsg{
				Message: RespStage8,
				Type:    TypeStr,
				Stage:   8,
			}
		case 9:
			return RespMsg{
				Message: RespStage9,
				Type:    TypeStr,
				Stage:   9,
			}
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
				FileName: AudioFileDedMorozName,
				Stage:    5,
			}
		}

	case ReqStage5:
		if pullStage == 5 {
			return RespMsg{
				Message: RespStage6,
				Type:    TypeStr,
				Stage:   6,
			}
		}

	case ReqStage6Net:
		if pullStage == 6 {
			return RespMsg{
				Message: RespStage7Pidor,
				Type:    TypeStr,
				Stage:   6,
			}
		}

	case ReqStage6Da:
		if pullStage == 6 {
			return RespMsg{
				Message: RespStage7NePidor,
				Type:    TypeImgs,
				Stage:   7,
				Images:  []string{ImgNews4, ImgNews1, ImgNews7, ImgNews5, ImgNews2, ImgNews6, ImgNews3},
			}
		}

	case ReqStage7:
		if pullStage == 7 {
			return RespMsg{
				Message: RespStage8,
				Type:    TypeStr,
				Stage:   8,
			}
		}

	case ReqStage8:
		if pullStage == 8 {
			return RespMsg{
				Message: RespStage9,
				Type:    TypeStr,
				Stage:   9,
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
