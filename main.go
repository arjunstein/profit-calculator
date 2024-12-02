package main

import (
	"errors"
	"fmt"
	"os"
)

func main()  {

	revenue, err := getInputUser("Pemasukan Rp. ")
	if err != nil {
		fmt.Println(err)
		return
		// panic(err)
	}

	expense, err := getInputUser("Modal Rp. ")
	if err != nil {
		fmt.Println(err)
		return
	}
	taxRate, err := getInputUser("Pajak 11% ")
	if err != nil {
		fmt.Println(err)
		return
	}

	ebt, profit, rasio := calculateFinancials(revenue,expense,taxRate)

	fmt.Println("EBT : Rp.", ebt)
	fmt.Println("Profit : Rp.", profit)
	fmt.Println("Rasio Profit : ", rasio)		

	storeResults(ebt, profit, rasio)
}

func storeResults(ebt, profit, rasio float64)  {
	results := fmt.Sprintf("EBT: %.1f\nProfit: %.1f\nRasio: %.3f\n", ebt, profit, rasio)
	os.WriteFile("results.txt", []byte(results), 0644)
}

func calculateFinancials(revenue, expense, taxRate float64) (float64, float64, float64) {
	ebt := revenue - expense
	profit := ebt * (1 - taxRate / 100)
	rasio := ebt / profit
	return ebt, profit, rasio
}

func getInputUser(infotext string) (float64, error) {
	var userInput float64
	fmt.Print(infotext)
	fmt.Scan(&userInput)

	if userInput <= 0 {
		return 0, errors.New("Input tidak boleh kurang dari 0")
	}

	return userInput, nil
}