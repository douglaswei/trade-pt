package entity

import "gorm.io/gorm"

type Order struct {
    *gorm.Model
    OrderId uint64 `gorm:" uniqueIndex;bigint(64);not null;comment:'全局自增唯一ID'"`
    UserId  uint64 `gorm:" uniqueIndex;bigint(64);not null;comment:'全局自增唯一ID'"`
    Status  uint8  `gorm:" bigint(8);not null;default 0;comment：'订单状态，0 创建，1 支付完成，2 取消， 3 作废'"`
    Fee     uint64 `gorm:" bigint(64);not null; default 0;comment:'订单金额'"`
}
