package internal

import (
	"context"
	"fmt"
	"log"
	"ordering-microservice/protogen/golang/orders"
)

type OrderService struct {
	db *DB
	orders.UnimplementedOrdersServer
}

func NewOrderService(db *DB) OrderService {
	return OrderService{db: db}
}

func (o *OrderService) AddOrder(_ context.Context, req *orders.PayloadWithSingleOrder) (*orders.Empty, error) {
	log.Println("Adding order!")
	err := o.db.AddOrder(req.GetOrder())
	return &orders.Empty{}, err
}

func (o *OrderService) GetOrder(_ context.Context, req *orders.PayloadWithOrderId) (*orders.PayloadWithSingleOrder, error) {
	log.Println("Get order received")

	order, err := o.db.GetOrder(req.GetId())
	if err != nil {
		return nil, fmt.Errorf("Error on ger order by id %v", err)
	}
	return &orders.PayloadWithSingleOrder{Order: order}, nil
}

func (o *OrderService) UpdateOrder(_ context.Context, req *orders.PayloadWithSingleOrder) (*orders.Order, error) {
	log.Println("Update order received")

	updateOrder := req.GetOrder()

	err := o.db.UpdateOrder(updateOrder)
	if err != nil {
		return nil, fmt.Errorf("Error on update order by id %v", err)
	}

	return updateOrder, nil
}

func (o *OrderService) DeleteOrder(_ context.Context, req *orders.PayloadWithOrderId) (*orders.Empty, error) {
	log.Println("Delete order received")

	o.db.DeleteOrder(req.GetId())

	return &orders.Empty{}, nil
}
