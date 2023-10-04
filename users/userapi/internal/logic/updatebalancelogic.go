package logic

import (
	"context"

	"awesomeProject7/users/userapi/internal/svc"
	"awesomeProject7/users/userapi/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateBalanceLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateBalanceLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateBalanceLogic {
	return &UpdateBalanceLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateBalanceLogic) UpdateBalance(req *types.UpdateBalanceRequest) (resp *types.UpdateBalanceResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
