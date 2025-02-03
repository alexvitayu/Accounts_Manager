package account

import (
	"demo/password-1/output"
	"encoding/json"
	"strings"
	"time"
)

type Db interface {
	Read() ([]byte, error)
	Write([]byte)
}

type Vault struct {
	Accounts  []Account
	UpdatedAt time.Time
}

type VaultWithDb struct {
	Vault
	db Db
}

func NewVault(db Db) *VaultWithDb {
	data, err := db.Read()
	if err != nil {
		return &VaultWithDb{
			Vault: Vault{
				Accounts:  []Account{},
				UpdatedAt: time.Now(),
			},
			db: db,
		}
	}
	var myVault Vault
	err = json.Unmarshal(data, &myVault)
	if err != nil {
		return &VaultWithDb{
			Vault: Vault{
				Accounts:  []Account{},
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

func (vault *VaultWithDb) AddAccount(acc Account) {
	vault.Accounts = append(vault.Accounts, acc)
	vault.save()
}

func (vault *VaultWithDb) FindAccountByUrl(url string) []Account {
	var accounts []Account
	for _, account := range vault.Accounts {
		isMatched := strings.Contains(account.Url, url)
		if isMatched {
			accounts = append(accounts, account)
		}
	}
	return accounts
}

func (vault *VaultWithDb) DeleteAccountByUrl(url string) bool {
	isDeleted := false
	for index, account := range vault.Accounts {
		isMatched := strings.Contains(account.Url, url)
		if isMatched {
			vault.Accounts = append(vault.Accounts[:index], vault.Accounts[index+1:]...)
			isDeleted = true
		}
		vault.save()
	}
	return isDeleted
}

func (vault *Vault) ToBytes() ([]byte, error) {
	data, err := json.MarshalIndent(vault, "", "")
	if err != nil {
		output.OutputErrors("не удалось преобразовать в json")
		return nil, err
	}
	return data, nil
}

func (vault *VaultWithDb) save() {
	vault.UpdatedAt = time.Now()
	data, err := vault.Vault.ToBytes()
	if err != nil {
		output.OutputErrors("не удалось преобразовать в json")
	}
	vault.db.Write(data)
}
