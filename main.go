package main

import (
	"Bank/services"
	"fmt"
)

func main() {

	var username string
	var password string

	fmt.Print("Nhập username: ")
	fmt.Scanln(&username)

	fmt.Print("Nhập password: ")
	fmt.Scanln(&password)

	user := services.Login(username, password)

	if user == nil {
		fmt.Println("Đăng nhập thất bại")
		return
	}

	fmt.Println("Đăng nhập thành công:", user.Username)

	var choice int

	for {
		fmt.Println("\n===== BANK MENU =====")
		fmt.Println("1. Deposit")
		fmt.Println("2. Withdraw")
		fmt.Println("3. Transfer")
		fmt.Println("4. Update Account")
		fmt.Println("5. Exit")
		fmt.Print("Chọn chức năng: ")

		fmt.Scanln(&choice)

		switch choice {

		case 1:
			var amount float64
			fmt.Print("Nhập số tiền muốn nạp: ")
			fmt.Scanln(&amount)

			services.Deposit(user.Username, amount)
			fmt.Println("Nạp tiền thành công")

		case 2:
			var amount float64
			fmt.Print("Nhập số tiền muốn rút: ")
			fmt.Scanln(&amount)

			services.Withdraw(user.Username, amount)

		case 3:
			var toUser string
			var amount float64

			fmt.Print("Nhập username người nhận: ")
			fmt.Scanln(&toUser)

			fmt.Print("Nhập số tiền chuyển: ")
			fmt.Scanln(&amount)

			services.Transfer(user.Username, toUser, amount)

		case 4:
			var newPassword string

			fmt.Print("Nhập password mới: ")
			fmt.Scanln(&newPassword)

			services.UpdateAccount(user.Username, newPassword)

			fmt.Println("Cập nhật tài khoản thành công")

		case 5:
			fmt.Println("Thoát chương trình")
			return

		default:
			fmt.Println("Lựa chọn không hợp lệ")
		}
	}
}