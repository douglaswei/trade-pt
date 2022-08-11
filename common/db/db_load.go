package db

import (
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
    "gorm.io/gorm/schema"
)

func LoadByDSN(dsn string) *gorm.DB {
    // 启动Gorm支持
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
        NamingStrategy: schema.NamingStrategy{
            SingularTable: true, // 使用单数表名
        },
    })
    // 如果出错就GameOver了
    if err != nil {
        panic(err)
    }
    return db
}
