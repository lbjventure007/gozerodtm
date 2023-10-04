package logic

import (
	"awesomeProject7/users/userrpc/internal/svc"
	"awesomeProject7/users/userrpc/userrpc"
	"context"
	"database/sql"
	"errors"
	"github.com/dtm-labs/client/dtmcli"
	"github.com/dtm-labs/client/dtmgrpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateBalanceLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateBalanceLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateBalanceLogic {
	return &UpdateBalanceLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateBalanceLogic) UpdateBalance(in *userrpc.UpdateBalanceRequest) (*userrpc.UpdateBalanceResponse, error) {
	// todo: add your logic here and delete this line

	id := in.Id

	dbconf := dtmcli.DBConf{Driver: "mysql", Host: "localhost", Port: 3306, User: "root", Password: "1234qwer", Db: "order"}
	err2 := dtmgrpc.XaLocalTransaction(l.ctx, dbconf, func(db *sql.DB, xa *dtmgrpc.XaGrpc) error {
		//这里最好也换成db来处理
		one, err := l.svcCtx.UserModel.FindOne(l.ctx, id)
		if err != nil {
			return err
		}
		userBalance := one.Balance.Float64
		if one.Id > 0 && userBalance >= 0 && userBalance >= float64(in.Balance) {
			one.Balance = sql.NullFloat64{one.Balance.Float64 - float64(in.Balance), true}
			err1 := l.svcCtx.UserModel.Update(l.ctx, one)
			if err1 != nil {
				return err1
			}
			return nil

		}

		return errors.New("条件不满足")
	})
	if err2 != nil {
		return nil, err2
	}
	return &userrpc.UpdateBalanceResponse{Message: "ok"}, nil

}
