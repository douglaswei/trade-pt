package entity

import "gorm.io/gorm"

type Spu struct {
    *gorm.Model
    SpuID uint64 `gorm:" uniqueIndex;bigint(64);not null;comment:'全局自增唯一ID'"`
    SkuID uint64 `gorm:" bigint(64);not null;comment:'skuID'"`
    Price uint64 `gorm:" bigint(32);not null;comment:'价格，单位分'"`
    Fund  uint64 `gorm:" bigint(32);not null;comment:'虚拟货币数'"`
}
