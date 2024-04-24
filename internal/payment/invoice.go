package payment

import (
	"errors"
)

func CalculateAmount(itemsPurchased map[string]int, itemsPrices map[string]float64) (amount float64, err error) {
	for k, v := range itemsPurchased {
		price, isFound := itemsPrices[k]
		if !isFound {
			err = errors.New("item no longer in catalog")
			return
		}
		if price <= 0.0 {
			err = errors.New("price cannot be zero or less")
			return
		}
		rowsAmount := float64(v) * price
		amount += rowsAmount
	}
	return
}
