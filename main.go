package main

import "fmt"

func main()  {

	revenue := getInputUser("Pemasukan Rp. ")
	expense := getInputUser("Modal Rp. ")
	taxRate := getInputUser("Pajak 11% ")

	ebt, profit, rasio := calculateFinancials(revenue,expense,taxRate)

	fmt.Println("EBT : Rp.", ebt)
	fmt.Println("Profit : Rp.", profit)
	fmt.Println("Rasio Profit : ", rasio)		
}

func calculateFinancials(revenue, expense, taxRate float64) (float64, float64, float64) {
	ebt := revenue - expense
	profit := ebt * (1 - taxRate / 100)
	rasio := ebt / profit
	return ebt, profit, rasio
}

func getInputUser(infotext string) float64 {
	var userInput float64
	fmt.Print(infotext)
	fmt.Scan(&userInput)
	return userInput
}