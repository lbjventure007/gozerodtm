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

type TranReverseOutLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewTranReverseOutLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TranReverseOutLogic {
	return &TranReverseOutLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *TranReverseOutLogic) TranReverseOut(in *userrpc.TranOutRequest) (*userrpc.TranOutReverseResponse, error) {
	// todo: add your logic here and delete this line
	fmt.Println("-----outreservse----")
	barrier, err := dtmgrpc.BarrierFromGrpc(l.ctx)
	db, err := l.svcCtx.Conn2.RawDB()
	if err != nil {
		return nil, err
	}
	fmt.Println("-----outreservse----1")
	err = barrier.CallWithDB(db, func(tx *sql.Tx) error {
		fmt.Println("-----outreservse----2")
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
		fmt.Println("-----outreservse----3")
		//if one.Balance.Float64 < float64(in.Balance) {
		//	return nil, errors.New("用户余额不足")
		//}

		one.Balance = sql.NullFloat64{one.Balance.Float64 + float64(in.Balance), true}

		exec, err := tx.Exec("update user set balance =? where id =?", one.Balance, one.Id)
		if err != nil {
			return err
		}
		fmt.Println("-----outreservse----4")
		affected, err := exec.RowsAffected()
		if err != nil {
			return err
		}
		if affected == 0 {
			return errors.New("失败")
		}
		fmt.Println("-----outreservse----5")
		return nil
	})
	if err != nil {
		return nil, err
	}

	return &userrpc.TranOutReverseResponse{Message: "ok"}, nil
}
