package main

// import (
// 	"fmt"
// 	"strconv"
// 	"strings"
// )

import (
	"fmt"
	"strconv"
	"strings"
)

func InputItemCashier() {
	// var restoName, dateReceipt, cashierName, item string
	// priceTag := 0
	cont := true
	billOut := false
	items := make(map[string]int)
	fmt.Println("Enter Restaurant Name :")
	var restoName string
	fmt.Scanln(&restoName)
	fmt.Println("Enter Cashier Name :")
	var cashierName string
	fmt.Scanln(&cashierName)
	fmt.Println("Enter Date of Receipt : ")
	var dateReceipt string
	fmt.Scanln(&dateReceipt)
	fmt.Println("Enter hours of Receipt : ")
	var hourReceipt string
	fmt.Scanln(&hourReceipt)
	for cont {
		fmt.Println("Enter Item : ")
		var item string
		fmt.Scanln(&item)
		fmt.Println("Enter Price : ")
		if item == "exit" {
			billOut = true
			cont = false
			break
		}
		var priceTag int
		fmt.Scanln(&priceTag)
		fmt.Println("\n")
		items[item] = priceTag
	}
	if billOut {
		total := 0
		dateState := "Tanggal : " + dateReceipt + " " + hourReceipt
		cashierState := "Nama Kasir : " + cashierName
		if len(restoName) >= 30 {
			fmt.Printf("%s\n%s\n", restoName[0:30], restoName[30:])
		} else {
			fmt.Printf("%s\n", restoName)
		}
		if len(dateState) >= 30 {
			fmt.Printf("%s\n%s\n", dateState[0:30], dateState[30:])
		} else {
			fmt.Printf("%s\n", dateState)
		}
		if len(cashierState) >= 30 {
			fmt.Printf("%s\n%s\n", cashierState[0:30], cashierState[30:])
		} else {
			fmt.Printf("%s\n", cashierState)
		}
		fmt.Println("==============================")
		for key, value := range items {
			total += value
			valueString := "Rp" + strconv.Itoa(value)
			lengthTotal := len(key + valueString)
			dotTotal := 30 - lengthTotal
			if dotTotal >= 0 {
				fmt.Printf("%s%s%s\n", key, strings.Repeat(".", dotTotal), valueString)
			} else {
				if len(key) == 30 {
					fmt.Printf("%s\n%s%s\n", key, strings.Repeat(".", int(30-len(valueString))), valueString)
				} else if len(key) > 30 {
					fmt.Printf("%s\n%s%s\n", key[:30], key[30:]+strings.Repeat(".", int(30-len(key[30:])-len(valueString))), valueString)
				} else if lengthTotal > 30 {
					fmt.Printf("%s%s\n%s%s\n", key[:30], strings.Repeat(".", int(30-len(key))), strings.Repeat(".", int(30-len(valueString))), valueString)
				}
			}
		}
		totalString := "Rp" + strconv.Itoa(total)
		dotTotalState := 30 - len("Total"+totalString)
		totalState := "Total" + strings.Repeat(".", dotTotalState) + totalString
		fmt.Printf("%s\n", totalState)
	}
}

func main() {
	fmt.Println("Start Cashier Machine\n")
	InputItemCashier()
	fmt.Println("\nShutdown Cashier Machine")
}
