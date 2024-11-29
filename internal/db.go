package internal

import (
	"fmt"
	"ordering-microservice/protogen/golang/orders"
)

type DB struct {
	collections []*orders.Order
}

func NewDB() *DB {
	return &DB{collections: make([]*orders.Order, 0)}
}

func (db *DB) AddOrder(o *orders.Order) error {
	for _, v := range db.collections {
		if v.Id == o.Id {
			return fmt.Errorf("Duplicate order id %v", o.Id)
		}
	}
	db.collections = append(db.collections, o)
	return nil
}

func (db *DB) GetOrder(id uint64) (*orders.Order, error) {
	for _, v := range db.collections {
		if v.Id == id {
			return v, nil
		}
	}
	return nil, fmt.Errorf("Order not found")
}

func (db *DB) UpdateOrder(o *orders.Order) error {
	for i, v := range db.collections {
		if v.Id == o.Id {
			db.collections[i] = o
			return nil
		}
	}
	return fmt.Errorf("Order not found")
}

func (db *DB) DeleteOrder(id uint64) {
	filtered := make([]*orders.Order, len(db.collections)-1)
	for i, v := range db.collections {
		if v.Id != id {
			filtered[i] = db.collections[i]
		}
	}
	db.collections = filtered
}
