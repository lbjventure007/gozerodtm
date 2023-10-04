package logic

import (
	"awesomeProject7/orders/orderapi/internal/svc"
	"awesomeProject7/orders/orderapi/internal/types"
	"awesomeProject7/orders/orderrpc/orderrpc"
	"awesomeProject7/users/userrpc/userrpc"
	"context"
	"fmt"
	"github.com/dtm-labs/client/dtmgrpc"
	"github.com/dtm-labs/dtmdriver"
	"google.golang.org/protobuf/types/known/emptypb"

	_ "github.com/dtm-labs/dtmdriver-gozero"
	"github.com/zeromicro/go-zero/core/logx"
)

type CreateOrderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateOrderLogic {
	return &CreateOrderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateOrderLogic) CreateOrder(req *types.CreateOrderRequest) (resp *types.CreateOrderResponse, err error) {
	// todo: add your logic here and delete this line

	fmt.Println("000", err)
	orderRequest := orderrpc.CreateOrderRequest{}
	orderRequest.Id = req.Id
	orderRequest.Userid = req.Userid
	orderRequest.Payment = float32(req.Payment)
	orderRequest.Status = int32(req.Status)
	orderRequest.Postage = int32(req.Postage)
	orderRequest.Shoppingid = int32(req.ShoppingId)
	orderRequest.Paymenttype = int32(req.PaymentType)
	orderTarget, err := l.svcCtx.Config.OrderRpc.BuildTarget()
	if err != nil {
		return nil, err
	}
	userTarget, err := l.svcCtx.Config.UserRpc.BuildTarget()
	if err != nil {
		return nil, err
	}
	userRequest := userrpc.UpdateBalanceRequest{}
	userRequest.Id = req.Userid
	userRequest.Balance = float32(req.Balance)

	var dtmServer = "consul://localhost:8500/dtmservice"

	gid := dtmgrpc.MustGenGid(dtmServer)
	err = dtmdriver.Use("dtm-driver-gozero")
	fmt.Println("=++++", err)
	err1 := dtmgrpc.XaGlobalTransaction(dtmServer, gid, func(xa *dtmgrpc.XaGrpc) error {
		//
		r := &emptypb.Empty{}

		err3 := xa.CallBranch(&orderRequest, orderTarget+"/orderrpc.Orderrpc/CreateOrder", r)
		//err3 := xa.CallBranch(&orderRequest, orderTarget+"/CreateOrder", res)
		fmt.Println("---", r, "---", err3)
		if err3 != nil {
			return err3
		}
		r1 := &emptypb.Empty{}
		err2 := xa.CallBranch(&userRequest, userTarget+"/userrpc.Userrpc/UpdateBalance", r1)
		fmt.Println("---", r1, "-3--", err2)
		//if err2 != nil {
		//	return dtmcli.ErrFailure
		//}
		return err2
	})
	//order, err := l.svcCtx.OrderRpc.CreateOrder(l.ctx, &orderRequest)
	//if err != nil {
	//	fmt.Println("11111")
	//	return nil, err
	//}
	//if order.Affected != 1 {
	//	return nil, errors.New("插入订单失败")
	//}
	//
	//balance, err := l.svcCtx.UserRpc.UpdateBalance(l.ctx, &userRequest)
	//if err != nil {
	//	fmt.Println("22222")
	//	return nil, err
	//}
	//if balance.Message != "ok" {
	//	return nil, errors.New("更新用户余额失败")
	//}
	resp = &types.CreateOrderResponse{}
	if err1 != nil {
		resp.Message = err1.Error()
	} else {
		resp.Message = "创建订单成功"
	}
	return
}
