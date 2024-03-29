// Code generated by goctl. DO NOT EDIT!
// Source: user.proto

package server

import (
    "context"

    "github.com/douglaswei/trade-pt/user/internal/logic"
    "github.com/douglaswei/trade-pt/user/internal/svc"
    "github.com/douglaswei/trade-pt/user/user"
)

type UserServer struct {
    svcCtx *svc.ServiceContext
    user.UnimplementedUserServer
}

func NewUserServer(svcCtx *svc.ServiceContext) *UserServer {
    return &UserServer{
        svcCtx: svcCtx,
    }
}

func (s *UserServer) Ping(ctx context.Context, in *user.Request) (*user.Response, error) {
    l := logic.NewPingLogic(ctx, s.svcCtx)
    return l.Ping(in)
}
