package entity

import "gorm.io/gorm"

type UserCategory struct {
    *gorm.Model
    UserID uint64 `gorm:" type:bigint(64);comment:'UserID'"`
    Category string `gorm:" type:bigint(10);default:0;comment:'用户分类；0，普通用户'"`
}

type UserTag struct {
    *gorm.Model
    UserID uint64 `gorm:" type:bigint(64);comment:'UserID'"`
    Tag string `gorm:" type:varchar(20);not null;comment:'用户标签'"`
}

type UserRelation struct {
    *gorm.Model
    UserID uint64 `gorm:" type:bigint(64);comment:'UserID'"`
    ForkUserID uint64 `gorm:" type:bigint(64);comment:'关注的UserID'"`
}