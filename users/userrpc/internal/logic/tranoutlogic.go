package logic

import (
	"awesomeProject7/users/userrpc/internal/svc"
	"awesomeProject7/users/userrpc/userrpc"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/dtm-labs/client/dtmcli"
	"github.com/dtm-labs/client/dtmgrpc"
	"github.com/zeromicro/go-zero/core/logx"
)

type TranOutLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewTranOutLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TranOutLogic {
	return &TranOutLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *TranOutLogic) TranOut(in *userrpc.TranOutRequest) (*userrpc.TranOutResponse, error) {
	// todo: add your logic here and delete this line
	fmt.Println("-----out----")
	barrier, err := dtmgrpc.BarrierFromGrpc(l.ctx)
	db, err := l.svcCtx.Conn2.RawDB()
	if err != nil {
		return nil, err
	}
	err = barrier.CallWithDB(db, func(tx *sql.Tx) error {
		if in.Id <= 0 {
			return errors.New("请求参数错误")
		}
		one, err := l.svcCtx.UserModel2.FindOne(l.ctx, in.Id)
		if err != nil {
			return err
		}
		if one.Id <= 0 {
			return errors.New("参数异常")
		}
		if one.Balance.Float64 < float64(in.Balance) {
			fmt.Println("---out--fail")
			return dtmcli.ErrFailure
			return errors.New("用户余额不足")
		}

		one.Balance = sql.NullFloat64{one.Balance.Float64 - float64(in.Balance), true}

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
		return nil
	})
	if err != nil {
		return nil, err
	}
	//if in.Id <= 0 {
	//	return nil, errors.New("请求参数错误")
	//}
	//one, err := l.svcCtx.UserModel.FindOne(l.ctx, in.Id)
	//fmt.Println("----11--", in.Id, err)
	//if err != nil {
	//	return nil, err
	//}
	//if one.Id <= 0 {
	//	return nil, errors.New("参数异常")
	//}
	//if one.Balance.Float64 < float64(in.Balance) {
	//
	//	fmt.Println("----222---")
	//	return nil, errors.New("FAILURE")
	//}
	//fmt.Println("----333---")
	//one.Balance = sql.NullFloat64{one.Balance.Float64 - float64(in.Balance), true}
	//err = l.svcCtx.UserModel.Update(l.ctx, one)
	//if err != nil {
	//	fmt.Println("----4444---")
	//	return nil, err
	//}
	//fmt.Println("----5555---")
	return &userrpc.TranOutResponse{Message: "ok"}, nil
}
