package services

import "Bank/models"

func Login(username, password string) *models.Account {
	accounts := LoadAccounts()

	for i := range accounts {
		if accounts[i].Username == username && accounts[i].Password == password {
			return &accounts[i]
		}
	}
	return nil
}
