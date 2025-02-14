package output

import "github.com/fatih/color"

func OutputErrorsBySwitchCase(value any) {
	switch f := value.(type) {
	case int:
		color.Red("Код ошибки:", f)
	case string:
		color.Red(f)
	case error:
		color.Red(f.Error())
	default:
		color.Red("unknown error type")
	}
}

func OutputErrorsByTypes(value any) {
	intValue, ok := value.(int)
	if ok {
		color.Red("Код ошибки:", intValue)
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
	color.Red("unknown error type")

}

func OutputErrorHack[T any](value T) {
	intValue, ok := any(value).(int)
	if ok {
		color.Red("Код ошибки:", intValue)
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
	color.Red("unknown error type")

}
