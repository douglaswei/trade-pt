package entity

import "gorm.io/gorm"

type Sku struct {
    *gorm.Model
    SkuID   uint64 `gorm:" uniqueIndex;bigint(64);not null;comment:'全局自增唯一ID'"`
    SkuName uint64 `gorm:" varchar(128);default 'sku-dft';comment:'商品名称'"`
    Img string `gorm:" varchar(256);not null;comment:'商品主图'"`
}
