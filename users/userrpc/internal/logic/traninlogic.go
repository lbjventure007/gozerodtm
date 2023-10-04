package logic

import (
	"awesomeProject7/users/userrpc/internal/svc"
	"awesomeProject7/users/userrpc/userrpc"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/dtm-labs/client/dtmgrpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type TranInLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewTranInLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TranInLogic {
	return &TranInLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *TranInLogic) TranIn(in *userrpc.TranInRequest) (*userrpc.TranInResponse, error) {
	// todo: add your logic here and delete this line
	fmt.Println("-----in----")
	barrier, err := dtmgrpc.BarrierFromGrpc(l.ctx)
	db, err := l.svcCtx.Conn1.RawDB()
	if err != nil {
		return nil, err
	}
	err = barrier.CallWithDB(db, func(tx *sql.Tx) error {
		if in.Id <= 0 {
			return errors.New("请求参数错误")
		}
		one, err := l.svcCtx.UserModel.FindOne(l.ctx, in.Id)
		if err != nil {
			return err
		}
		if one.Id <= 0 {
			return errors.New("参数异常")
		}

		one.Balance = sql.NullFloat64{one.Balance.Float64 + float64(in.Balance), true}

		exec, err := tx.Exec("update user set balance =? where id =?", one.Balance, one.Id)
		if err != nil {
			return err
		}
		affected, err := exec.RowsAffected()
		if err != nil {
			return err
		}
		if affected == 0 {
			return errors.New("失败")
		}

		fmt.Println("---in--success")
		return nil
	})
	if err != nil {
		return nil, err
	}
	return &userrpc.TranInResponse{}, nil
}
