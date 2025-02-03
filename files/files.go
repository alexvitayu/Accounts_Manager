package files

import (
	"demo/password-1/output"
	"fmt"
	"os"
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
		output.OutputErrors("не удалось прочитать файл")
		return nil, err
	}
	return data, nil
}

func (db *JsonDb) Write(content []byte) {
	file, err := os.Create(db.filename)
	if err != nil {
		output.OutputErrors("не удалось создать файл")
	}
	defer file.Close()
	_, err = file.Write(content)
	if err != nil {
		output.OutputErrors("не удалось записать файл")
		return
	}
	fmt.Println("Запись успешна")
}
