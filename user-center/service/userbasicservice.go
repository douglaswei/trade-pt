package service

import (
    "errors"
    "gorm.io/gorm"
)

type UserBasicService struct {
    db *gorm.DB
}

// Register 用户注册
func (service UserBasicService) Register(user User) error {
    if user.UserName == "" {
        return errors.New("empty userName")
    }

    // 未找到
    var eUser User
    err := service.db.Where("userName = >", user.UserName).First(&eUser).Error
    if errors.Is(err, gorm.ErrRecordNotFound) {
    }

    return nil
}

// GetByUserID 用户自我查询
func (service UserBasicService) GetByUserID(userID int64) User {
    return User{}
}

// GetOtherByUserID 获取其他用户信息
func (service UserBasicService) GetOtherByUserID(curUserID int64, getUserID int64) User {
    return User{}
}

// UpdateByUserID 用户信息更新
func (service UserBasicService) UpdateByUserID(userID int64, user User) error {
    return nil
}

// UpdatePasswd 密码重置
func (service UserBasicService) UpdatePasswd(user User) error {
    return nil
}

// Follow 关注取关
func (service UserBasicService) Follow(uerId uint64, followUserID uint64, isFollow bool) error {
    return nil
}
