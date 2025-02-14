package reverse

import "github.com/fatih/color"

func ReverseArray[T int | string | float64](arr *[]T) *[]T {
	reverse := make([]T, len(*arr))
	for i := 0; i < len(*arr); i++ {
		reverse[i] = (*arr)[len(*arr)-1-i]
	}
	return &reverse
}
func OutputReverse[T int | string | float64](arr *[]T) {
	switch f := any(*arr).(type) {
	case []int:
		color.Red("%v", f)
	case []string:
		color.Green("%v", f)
	case []float64:
		color.Yellow("%v", f)
	default:
		color.Red("unknown type")
	}
}
