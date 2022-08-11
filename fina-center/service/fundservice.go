package service

type FundService interface {
    // Depoist 充值
    Depoist(count int64)
    // PayOff 支付
    PayOff(count int64)
}