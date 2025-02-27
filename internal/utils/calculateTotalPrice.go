package utils

import (
	"fmt"
	"strconv"
)

// ticket_price has to be float-string format(example: "0.00")
func CalculateTotalPrice(ticket_price string, quantity int32, fees string) (string, error) {
	// convert string "0.00" to float 0.00
	conv_ticket_price, err := strconv.ParseFloat(ticket_price, 32)
	if err != nil {
		return "", err
	}
	// convert fees string to float
	conv_fees, err := strconv.ParseFloat(fees, 32)
	fmt.Println("ERROR AT THIS")
	if err != nil {
		return "", err
	}
	total := conv_ticket_price * float64(quantity) // example: 75*2
	total = total + total*conv_fees/100

	// finally convert total float-format into total string format
	conv_total := strconv.FormatFloat(total, 'f', 2, 64)
	return conv_total, nil

}
