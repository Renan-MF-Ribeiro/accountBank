package main

import (
	"accountBank/accounts"
	"accountBank/client"
	"fmt"
	"math"
	"math/rand"
)

func main() {
	client1 := client.User{Name: "Renan", Idetifer: "123.456.789.11", Occupation: "dev"}
	client2 := client.NewUser("Samuel", "321.654.987.11", "Student")

	currentAccount := accounts.Current{User: client1}
	currentAccount2 := accounts.Current{User: client2}

	fmt.Println(currentAccount)
	fmt.Println(currentAccount2)
	fmt.Println()

	TestActionsAccount(&currentAccount, &currentAccount2)
	fmt.Println()
	TestActionsAccount(&currentAccount2, &currentAccount)

	fmt.Println()
	fmt.Println(currentAccount.Balance())
	fmt.Println(currentAccount2.Balance())

}

// Function to generate a random number between min and max values
func valueRandom(min float64, max float64) float64 {
	random := rand.Float64()*max + min
	return math.Round(random*100) / 100

}

// function to test deposit
func TestActionsAccount(ac *accounts.Current, otherAc *accounts.Current) {
	fmt.Println("Iniciando testes ....")
	fmt.Println()
	fmt.Println("Saldo Atual", ac.Balance())

	fmt.Println("Teste de Deposito")
	deposit := valueRandom(1.0, 1000.0)
	fmt.Println("Deposit", deposit)
	fmt.Println(ac.Deposit(deposit))
	fmt.Println("Saldo Atual", ac.Balance())
	fmt.Println()
	fmt.Println()

	fmt.Println("Teste de Saque")
	withdrawal := valueRandom(1.0, 100.0)
	fmt.Println("withdrawal", withdrawal)
	fmt.Println(ac.Withdrawal(withdrawal))
	fmt.Println("Saldo Atual", ac.Balance())
	fmt.Println()
	fmt.Println()

	fmt.Println("Teste de Transferência")
	fmt.Println(ac.User.Name, " -> ", otherAc.User.Name)
	transfer := valueRandom(1.0, 100.0)
	fmt.Println("transfer", transfer)
	fmt.Println("Saldo Atual", ac.Balance())
	fmt.Println(ac.Transfer(*otherAc, transfer))
	fmt.Println("Saldo Atual", ac.Balance())
}
