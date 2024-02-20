package app

const (
	HelloMsg        = "Программа для добавления url в список"
	QuitHelpMsg     = "Для выхода и приложения нажмите 'q'"
	CommandsHelpMsg = `Введите 'a', 'l', 'r', 's' или 'q' для выбора команды
- a (add) - добавить url в список
- l (list) - вывести список url
- r (remove) - удалить url из списка
- s (search) - поиск url
- q (quit) - выйти из приложения
`
	BadCommandErrTemplate = "Команда '%s' не верная, повторите\n\n"
	ErrMsgTemplate        = "Ошибка: %s\n"
	AddCmdHelpMsq         = "Введите новую запись в формате <url описание теги>"
	RemoveCmdHelpMsq      = "Введите имя ссылки, которое нужно удалить"
	SearchCmdHelpMsq      = "Введите поисковый запрос: "
	ListCmdTemplate       = "Имя: %s\nURL: %s\nТеги: %v\nДата: %s\n"

	AppExitedMsg       = "Завершение по требованию пользователя"
	AppKilledMsg       = "Завершение по системному сигналу"
	DefaultErrTemplate = "Ошибка: %v\n"

	ByeMsg = "До новых встреч!"
)
