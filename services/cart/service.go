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

func (h *Handler) CreateOrderFromCart(products []types.Product, items []types.CartItem, userId int) (int, float64, error) {
	productMap := make(map[int]types.Product)

	for _, product := range products {
		productMap[product.ID] = product
	}

	// check if all product in stock
	if err := checkCartIsInStock(items, productMap); err != nil {
		return 0, 0, err
	}

	// calculate total price
	totalPrice := calculateTotalPrice(items, productMap)

	// reduce product stock in DB
	for _, item := range items {
		product := productMap[item.ProductID]
		product.Quantity -= item.Quantity

		h.productStore.UpdateProduct(product)
	}

	// create order
	orderId, err := h.store.CreateOrder(types.Order{
		UserID:  userId,
		Total:   totalPrice,
		Status:  "pending",
		Address: "Address here",
	})
	if err != nil {
		return 0, 0, err
	}
	// create order items
	for _, item := range items {
		h.store.CreateOrderItem(types.OrderItem{
			OrderID:   orderId,
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
			Price:     productMap[item.ProductID].Price,
		})
	}

	return orderId, totalPrice, nil
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
