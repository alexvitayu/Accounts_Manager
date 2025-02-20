package account

import (
	"encoding/json"
	"errors"
	"revision/part-1/output"
	"strings"
	"time"
)

type ByteReader interface {
	Read() ([]byte, error)
}

type ByteWriter interface {
	Write([]byte)
}

type Db interface {
	ByteReader
	ByteWriter
}

type Vault struct {
	Accounts  []AccountWithTimeStamps `json:"MyAccounts:"`
	UpdatedAt time.Time               `json:"UpdatedAt"`
}

type VaultWithDb struct {
	Vault
	db Db
}

func NewMyVault(db Db) *VaultWithDb {
	file, err := db.Read()
	if err != nil {
		return &VaultWithDb{
			Vault: Vault{
				Accounts:  []AccountWithTimeStamps{},
				UpdatedAt: time.Now(),
			},
			db: db,
		}
	}
	var myVault Vault
	err = json.Unmarshal(file, &myVault)
	if err != nil {
		output.OutputErrorsByTypes("не удалось преобразовать из json")
		return &VaultWithDb{
			Vault: Vault{
				Accounts:  []AccountWithTimeStamps{},
				UpdatedAt: time.Now(),
			},
			db: db,
		}
	}
	return &VaultWithDb{
		Vault: myVault,
		db:    db,
	}
}

func (vault *VaultWithDb) AddAccount(acc AccountWithTimeStamps) error {
	for _, account := range vault.Accounts {
		isMatched := strings.Contains(account.Url, acc.Url)
		if isMatched {
			return errors.New("такой аккаунт уже существует")
		}
	}
	vault.Accounts = append(vault.Accounts, acc)
	vault.toBytesAndSave()
	return nil
}

func (vault *VaultWithDb) FindAccounts(strName string, searching func(str string, acc AccountWithTimeStamps) bool) *[]AccountWithTimeStamps {
	var accounts []AccountWithTimeStamps
	for _, account := range vault.Accounts {
		isMatched := searching(strName, account)
		if isMatched {
			accounts = append(accounts, account)
		}
	}
	return &accounts
}

func (vault *VaultWithDb) DeleteAccountByUrl(url string) bool {
	isDeleted := false
	for index, account := range vault.Accounts {
		isMatched := strings.Contains(account.Url, url)
		if isMatched {
			vault.Accounts = append(vault.Accounts[:index], vault.Accounts[index+1:]...)
			isDeleted = true
		}
	}
	vault.toBytesAndSave()
	return isDeleted
}

/*func (vault *Vault) ToBytes() ([]byte, error) {
	data, err := json.MarshalIndent(vault, "", "")
	if err != nil {
		return nil, errors.New("не удалось преобразовать в json")
	}
	return data, nil
}*/

/*func (vault *VaultWithDb) save() {
	vault.UpdatedAt = time.Now()
	data, err := vault.ToBytes()
	if err != nil {
		output.OutputErrorHack(err)
	}
	vault.db.Write(data)
}*/

func (vault *VaultWithDb) toBytesAndSave() {
	data, err := json.MarshalIndent(vault, "", "")
	if err != nil {
		output.OutputErrorHack("не удалось преобразовать в json")
		return
	}
	vault.db.Write(data)
}
