package domain

// OrderLineItem represents an order line item resource.
type OrderLineItem struct {
	// Required: true
	MenuItemId string `json:"menu_item_id"`
}
