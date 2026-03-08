package services

import "fmt"

func UpdateAccount(username string, newPassword string) {
	accounts := LoadAccounts()

	for i := range accounts {
		if accounts[i].Username == username {
			accounts[i].Password = newPassword
		}
	}

	SaveAccounts(accounts)
}

func Deposit(username string, amount float64) {
	accounts := LoadAccounts()

	for i := range accounts {
		if accounts[i].Username == username {
			accounts[i].Balance += amount
		}
	}

	SaveAccounts(accounts)
}

func Withdraw(username string, amount float64) {
	accounts := LoadAccounts()

	for i := range accounts {
		if accounts[i].Username == username {
			if accounts[i].Balance < amount {
				fmt.Println("Số dư không đủ")
				return
			}

			accounts[i].Balance -= amount
			fmt.Println("Rút tiền thành công")
		}
	}

	SaveAccounts(accounts)
}

func Transfer(fromUser string, toUser string, amount float64) {

	accounts := LoadAccounts()

	fromIndex := -1
	toIndex := -1

	for i := range accounts {

		if accounts[i].Username == fromUser {
			fromIndex = i
		}

		if accounts[i].Username == toUser {
			toIndex = i
		}
	}

	// check tài khoản tồn tại
	if fromIndex == -1 {
		fmt.Println("Tài khoản gửi không tồn tại")
		return
	}

	if toIndex == -1 {
		fmt.Println("Tài khoản nhận không tồn tại")
		return
	}

	// check số dư
	if accounts[fromIndex].Balance < amount {
		fmt.Println("Số dư không đủ")
		return
	}

	accounts[fromIndex].Balance -= amount
	accounts[toIndex].Balance += amount

	SaveAccounts(accounts)

	fmt.Println("Chuyển tiền thành công")
}