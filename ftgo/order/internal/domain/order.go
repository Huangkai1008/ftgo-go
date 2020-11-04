package domain

import (
	"time"

	metav1 "github.com/Huangkai1008/micro-kit/pkg/meta/v1"
)

type OrderState int

const (
	ApprovalPending OrderState = iota + 1
	Approved
	Rejected
	CancelPending
	Cancelled
	RevisionPending
)

// Order represents an order resource.
type Order struct {
	metav1.ObjectMeta

	// Required: true
	ConsumerID uint64 `json:"consumer_id"`

	// Required: true
	RestaurantID uint64 `json:"restaurant_id"`

	// Required: true
	State OrderState `json:"state"`

	// Required: true
	LineItems []OrderLineItem

	// Required: true
	DeliveryAddress Address

	// Required: true
	DeliveryTime time.Time `json:"delivery_time"`
}
