package files

import (
	"os"

	"github.com/fatih/color"
)

type JsonDb struct {
	filename string
}

func NewJsonDb(name string) *JsonDb {
	return &JsonDb{
		filename: name,
	}
}

func (db *JsonDb) Read() ([]byte, error) {
	data, err := os.ReadFile(db.filename)
	if err != nil {
		color.Red("не удалось прочитать файл")
		return nil, err
	}
	return data, nil
}
func (db *JsonDb) Write(content []byte) {
	file, err := os.Create(db.filename)
	if err != nil {
		color.Red("не удалось создать файл")
		return
	}
	_, err = file.Write(content)
	defer file.Close()
	if err != nil {
		color.Red("не удалось записать файл")
		return
	}
	color.Green("Запись успешна")
}
