package logic

import (
	"awesomeProject7/users/userrpc/userrpc"
	"context"
	"fmt"
	"github.com/dtm-labs/client/dtmcli"
	"github.com/dtm-labs/client/dtmgrpc"
	"github.com/dtm-labs/dtmdriver"

	"awesomeProject7/users/userapi/internal/svc"
	"awesomeProject7/users/userapi/internal/types"

	_ "github.com/dtm-labs/dtmdriver-gozero"
	"github.com/zeromicro/go-zero/core/logx"
)

type TranInOutBalanceLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTranInOutBalanceLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TranInOutBalanceLogic {
	return &TranInOutBalanceLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TranInOutBalanceLogic) TranInOutBalance(req *types.TranInOutRequest) (resp *types.TranInOutResponse, err error) {
	// todo: add your logic here and delete this line
	var dtmServer = "consul://localhost:8500/dtmservice"
	//var dtmServer = "localhost:36790"
	gid := dtmgrpc.MustGenGid(dtmServer)
	err = dtmdriver.Use("dtm-driver-gozero")
	fmt.Println("=++++", err)
	userTarget, err := l.svcCtx.Config.UserRpc.BuildTarget()
	if err != nil {
		return nil, err
	}
	inReq := userrpc.TranInRequest{Id: req.InId, Balance: float32(req.Balance)}
	outReq := userrpc.TranInRequest{Id: req.OutId, Balance: float32(req.Balance)}

	saga := dtmgrpc.NewSagaGrpc(dtmServer, gid).
		Add(userTarget+"/userrpc.Userrpc/TranIn", userTarget+"/userrpc.Userrpc/TranReverseIn", &inReq).
		Add(userTarget+"/userrpc.Userrpc/TranOut", userTarget+"/userrpc.Userrpc/TranReverseOut", &outReq)
	saga.WaitResult = true
	saga.WithRetryLimit(1)
	//这里如果有某个分子 异常 则会重试1一次  也就是异常分子 最多会执行两次
	//如果这里不限制 dtm 会一直重试 ，即便我返回了FAILURE(dtmcli.ErrFailure)

	err = saga.Submit()
	fmt.Println(err, "saga----")
	resp = &types.TranInOutResponse{}
	if err != nil {
		resp.Message = err.Error()
		return resp, dtmcli.ErrFailure
		//
		//return nil, errors.New("fail-" + gid)
	}

	resp.Message = "ok" + gid
	return
}
