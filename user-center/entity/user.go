package entity

import "gorm.io/gorm"

type User struct {
    *gorm.Model
    UserID       uint64 `gorm:" uniqueIndex;bigint(64);not null;comment:'全局自增唯一ID'"`
    UserName     string `gorm:" type:varchar(128);index;not null;default:Douglas;comment:'用户名，登录使用'"`
    NikeName     string `gorm:" type:varchar(128);not null;default:Douglas;comment:'昵称'"`
    Avatar       string `gorm:" type:varchar(128);comment:'头像'"`
    Gender       uint8  `gorm:" type:bigint(4);default:1;comment:'性别，0，男；1，女；2，未知；'"`
    DistrictCode uint64 `gorm:" type:bigint(64);default:1;comment:'地区码，留作快速查找'"`
    Country      string `gorm:" type:varchar(128);default:China;comment:'国家'"`
    Province     string `gorm:" type:varchar(128);default:Shanghai;comment:'省份'"`
    City         string `gorm:" type:varchar(128);default:Shanghai;comment:'城市'"`
    Region       string `gorm:" type:varchar(128);default:Xuhui;comment:'地区'"`
    Email        string `gorm:" type:varchar(128);comment:'电子邮件'"`
    Mobile       string `gorm:" type:varchar(20);comment:'手机号'"`
    Age          int8   `gorm:" type:bigint(8);default:14;comment:'年龄'"`
    PassWd       string `gorm:" type:varchar(64);comment:'密码'"`
}
