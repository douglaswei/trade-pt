package entity

import "gorm.io/gorm"

type Fund struct {
    *gorm.Model
    FundId       uint64 `gorm:" uniqueIndex;bigint(64);not null;comment:'全局自增唯一ID'"`
    UserId       uint64 `gorm:" uniqueIndex;bigint(64);not null;comment:'全局自增唯一ID'"`
    Cash         uint64 `gorm:" bigint(64);not null;comment:'账户资金（非现金，可作为代币）'"`
}