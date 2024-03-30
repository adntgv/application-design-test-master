package records

import "applicationDesignTest/internal/business/domains"

func (o *Order) ToDomain() *domains.OrderDomain {
	return &domains.OrderDomain{}
}

func FromOrderDomain(o *domains.OrderDomain) *Order {
	return &Order{}
}
