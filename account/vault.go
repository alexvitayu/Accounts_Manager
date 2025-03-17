package account

import (
	"encoding/json"
	"revision/part-1/encryptor"
	"strings"
	"time"

	"github.com/fatih/color"
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
	Accounts  []AccountWithTimeStamps
	UpdatedAt time.Time
}

type VaultWithDb struct {
	Vault
	db  Db
	enc encryptor.Encryptor
}

func NewMyVault(db Db, enc *encryptor.Encryptor) *VaultWithDb {
	data, err := db.Read()
	if err != nil {
		return &VaultWithDb{
			Vault: Vault{
				Accounts:  []AccountWithTimeStamps{},
				UpdatedAt: time.Now(),
			},
			db:  db,
			enc: *enc,
		}
	}
	decryptedData := enc.Decrypt(data)
	var myVault Vault
	err = json.Unmarshal(decryptedData, &myVault)
	if err != nil {
		return &VaultWithDb{
			Vault: Vault{
				Accounts:  []AccountWithTimeStamps{},
				UpdatedAt: time.Now(),
			},
			db:  db,
			enc: *enc,
		}
	}
	return &VaultWithDb{
		Vault: myVault,
		db:    db,
		enc:   *enc,
	}
}

func (vault *VaultWithDb) AddAccount(acc *AccountWithTimeStamps) {
	vault.Accounts = append(vault.Accounts, *acc)
	vault.save()
}

func (vault *VaultWithDb) save() {
	vault.UpdatedAt = time.Now()
	data, err := json.MarshalIndent(vault, "", "")
	if err != nil {
		color.Red("не удалось преобразовать в json")
	}
	encData := vault.enc.Encrypt(data)
	vault.db.Write(encData)
}

func (vault *VaultWithDb) FindAccount(str string, checker func(string, AccountWithTimeStamps) bool) *[]AccountWithTimeStamps {
	var accounts []AccountWithTimeStamps
	for _, account := range vault.Accounts {
		isMatched := checker(str, account)
		if isMatched {
			accounts = append(accounts, account)
		}
	}
	return &accounts
}

func (vault *VaultWithDb) DeleteAccountByUrl(url string) bool {
	isCompared := false
	for index, account := range vault.Accounts {
		isMatched := strings.Contains(account.Url, url)
		if isMatched {
			vault.Accounts = append(vault.Accounts[:index], vault.Accounts[index+1:]...)
			isCompared = true
		}
	}
	vault.save()
	return isCompared
}
