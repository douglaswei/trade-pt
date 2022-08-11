package service

type OrderService interface {
    // CreateOrder 创建订单
    CreateOrder(count int64)
    // PayOrder 支付订单
    PayOrder(orderId int64)
}
