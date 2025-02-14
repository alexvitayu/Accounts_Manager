package files

import (
	"os"
	"revision/part-1/output"

	"github.com/fatih/color"
)

type JsonDb struct {
	namefield string
}

func NewJsondb(name string) *JsonDb {
	return &JsonDb{
		namefield: name,
	}
}

func (db *JsonDb) Read() ([]byte, error) {
	data, err := os.ReadFile(db.namefield)
	if err != nil {
		output.OutputErrorHack("не удалось прочитать файл")
		return nil, err
	}
	return data, nil
}

func (db *JsonDb) Write(content []byte) {
	file, err := os.Create(db.namefield)
	if err != nil {
		output.OutputErrorsBySwitchCase("не удалось создать файл")
	}
	defer file.Close()
	_, err = file.Write(content)
	if err != nil {
		output.OutputErrorsByTypes("не удалось записать в файл")
		return
	}
	color.Green("Запись успешна!")
}
