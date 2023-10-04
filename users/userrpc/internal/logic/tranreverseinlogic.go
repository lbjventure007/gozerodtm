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

type TranReverseInLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewTranReverseInLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TranReverseInLogic {
	return &TranReverseInLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *TranReverseInLogic) TranReverseIn(in *userrpc.TranInRequest) (*userrpc.TranInReverseResponse, error) {
	// todo: add your logic here and delete this line
	fmt.Println("-----inreservse----")
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
		//if one.Balance.Float64 < float64(in.Balance) {
		//	return nil, errors.New("用户余额不足")
		//}

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
		fmt.Println("---inreverse--success")
		return nil
	})
	if err != nil {
		return nil, err
	}
	//fmt.Println("======", "tranresverin")
	//if in.Id <= 0 {
	//	return nil, errors.New("请求参数错误")
	//}
	//one, err := l.svcCtx.UserModel.FindOne(l.ctx, in.Id)
	//if err != nil {
	//	return nil, err
	//}
	//if one.Id <= 0 {
	//	return nil, errors.New("参数异常")
	//}
	////if one.Balance.Float64 < float64(in.Balance) {
	////	return nil, errors.New("用户余额不足")
	////}
	//
	//one.Balance = sql.NullFloat64{one.Balance.Float64 - float64(in.Balance), true}
	//err = l.svcCtx.UserModel.Update(l.ctx, one)
	//if err != nil {
	//	return nil, err
	//}
	return &userrpc.TranInReverseResponse{Message: "ok"}, nil
}
