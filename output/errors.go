package output

import "github.com/fatih/color"

func OutputErrors(value any) {
	switch n := value.(type) {
	case string:
		color.Red(n)
	case int:
		color.Red("Код ошибки: %d", n)
	case error:
		color.Red(n.Error())
	default:
		color.Red("неизвестный тип ошибки")
	}
}

func OutputErrorsNew(value any) {
	intValue, ok := value.(int)
	if ok {
		color.Red("Код ошибки: %d", intValue)
		return
	}
	strValue, ok := value.(string)
	if ok {
		color.Red(strValue)
		return
	}
	errValue, ok := value.(error)
	if ok {
		color.Red(errValue.Error())
		return
	}
	color.Red("неизвестный тип ошибки")
}

func OutputErrorsHack[T any](value T) {
	intValue, ok := any(value).(int)
	if ok {
		color.Red("Код ошибки: %d", intValue)
		return
	}
	strValue, ok := any(value).(string)
	if ok {
		color.Red(strValue)
		return
	}
	errValue, ok := any(value).(error)
	if ok {
		color.Red(errValue.Error())
		return
	}
	color.Red("неизвестный тип ошибки")

}
