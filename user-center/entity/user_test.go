package entity

import (
    c_config "github.com/douglaswei/go-common/consul-config"
    c_db "github.com/douglaswei/go-common/db"
    "testing"
    l_config "user-center/config"
)

func TestInit(t *testing.T) {
    cfg := c_config.GetConfig(l_config.ConsulPath)
    if cfg == nil {
        return
    }
    db, err := c_db.GetGormFromkey(cfg, l_config.DbKey)
    if err != nil {
        return
    }
    err = db.AutoMigrate(&User{})
    if err != nil {
        print(err)
    }
    err = db.AutoMigrate(&UserCategory{})
    if err != nil {
        print(err)
    }
    err = db.AutoMigrate(&UserRelation{})
    if err != nil {
        print(err)
    }
    err = db.AutoMigrate(&UserTag{})
    if err != nil {
        print(err)
    }
}
