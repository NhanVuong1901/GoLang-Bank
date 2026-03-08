package services

import (
	"Bank/models"
	"encoding/json"
	"os"
)

const dataFile = "data/users.json"

func LoadAccounts() []models.Account {
	if _, err := os.Stat(dataFile); os.IsNotExist(err) {
		return []models.Account{}
	}

	bytes, _ := os.ReadFile(dataFile)

	var accounts []models.Account

	_ = json.Unmarshal(bytes, &accounts)
	return accounts
}

func SaveAccounts(accounts []models.Account) error {
	bytes, _ := json.Marshal(accounts)
	return os.WriteFile(dataFile, bytes, 0777)
}
