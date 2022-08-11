package service

import (
    "user-center/entity"
)

type User entity.User

type UserService interface {
    // Register 用户注册
    Register(user User) error
    // GetByUserID 用户自我查询
    GetByUserID(userID int64) User
    // GetOtherByUserID 获取其他用户信息
    GetOtherByUserID(curUserID int64, getUserID int64) User
    // UpdateByUserID 用户信息更新
    UpdateByUserID(userID int64, user User) error
    // UpdatePasswd 密码重置
    UpdatePasswd(user User) error
    // Follow 关注取关
    Follow(uerId uint64, followUserID uint64, isFollow bool) error
}