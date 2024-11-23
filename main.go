package main

import "fmt"

func main()  {
	var revenue float64
	var expense float64
	const taxRate float64 = 11

	fmt.Print("Pemasukan Rp. ")
	fmt.Scan(&revenue)
	fmt.Print("Modal Rp. ")
	fmt.Scan(&expense)
	fmt.Print(`Pajak 11% = `, taxRate)	

	earningBeforeTax := revenue - expense
	earningAfterTax := earningBeforeTax - (earningBeforeTax * taxRate / 100)
	ratio := (earningBeforeTax / revenue) * 100
	ppn := earningBeforeTax - earningAfterTax

	ebt := revenue - expense
	profit := ebt * (1 - taxRate / 100)
	rasio := ebt / profit

	fmt.Println("\nLaba Sebelum Pajak : Rp.", earningBeforeTax)
	fmt.Println("Laba Setelah Pajak : Rp.", earningAfterTax)
	fmt.Println("PPn : Rp.", ppn)
	fmt.Println("Rasio Laba : ", ratio, "%")
	fmt.Println("EBT : Rp.", ebt)
	fmt.Println("Profit : Rp.", profit)
	fmt.Println("Rasio Profit : ", rasio)	
	
}