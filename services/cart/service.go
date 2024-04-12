package cart

import (
	"fmt"

	"github.com/vnsonvo/ecom-rest-api/types"
)

func getCartItemIds(items []types.CartItem) ([]int, error) {
	productIds := make([]int, len(items))
	for i, item := range items {
		if item.Quantity <= 0 {
			return nil, fmt.Errorf("Invalid quantity for product %d", item.ProductID)
		}
		productIds[i] = item.ProductID
	}

	return productIds, nil
}
