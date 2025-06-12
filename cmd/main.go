package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/Chayakorn2002/thb-amount-to-text-go/utils"
	"github.com/shopspring/decimal"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var amount string
	fmt.Println("Enter amount in THB (e.g., 1234.56):")
	if scanner.Scan() {
		amount = scanner.Text()
	}

	decimalAmount, err := decimal.NewFromString(amount)
	if err != nil {
		os.Stderr.WriteString("Invalid amount: " + err.Error() + "\n")
		return
	}

	amountText, err := utils.DecimalToBahtText(decimalAmount)
	if err != nil {
		os.Stderr.WriteString("Error converting amount: " + err.Error() + "\n")
		return
	}

	os.Stdout.WriteString("Amount in text: " + amountText + "\n")
}
