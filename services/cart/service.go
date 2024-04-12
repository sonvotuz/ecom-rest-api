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

func checkCartIsInStock(items []types.CartItem, products map[int]types.Product) error {
	if len(items) == 0 {
		return fmt.Errorf("Cart is empty")
	}

	for _, item := range items {
		product, ok := products[item.ProductID]
		if !ok {
			return fmt.Errorf("Product %d is not available in the store, please check your cart", item.ProductID)
		}

		if product.Quantity < item.Quantity {
			return fmt.Errorf("Product %s is not available with the quantity requested", product.Name)
		}
	}

	return nil
}

func calculateTotalPrice(items []types.CartItem, productMap map[int]types.Product) float64 {
	var total float64

	for _, item := range items {
		product := productMap[item.ProductID]
		total += product.Price * float64(item.Quantity)
	}
	return total
}
