package logic

import (
	"awesomeProject7/orders/model"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/dtm-labs/client/dtmcli"
	"github.com/dtm-labs/client/dtmgrpc"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	//"github.com/go-sql-driver/mysql"
	"time"

	"awesomeProject7/orders/orderrpc/internal/svc"
	"awesomeProject7/orders/orderrpc/orderrpc"

	//_ "github.com/dtm-labs/driver-gozero"
	"github.com/zeromicro/go-zero/core/logx"
)

type CreateOrderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateOrderLogic {
	return &CreateOrderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateOrderLogic) CreateOrder(in *orderrpc.CreateOrderRequest) (*orderrpc.CreateOrderResponse, error) {

	// todo: add your logic here and delete this line
	//l.svcCtx.Db.TransactCtx()
	dbconf := dtmcli.DBConf{Driver: "mysql", Host: "localhost", Port: 3306, User: "root", Password: "1234qwer", Db: "order"}

	err2 := dtmgrpc.XaLocalTransaction(l.ctx, dbconf, func(db *sql.DB, xa *dtmgrpc.XaGrpc) error {
		orders := model.Orders{}
		orders.Id = in.Id
		orders.Userid = 1
		orders.Status = int64(in.Status)
		orders.Postage = 1
		orders.Shoppingid = 1
		orders.CreateTime = time.Now()
		orders.Payment = 1
		var dia gorm.Dialector
		dia = mysql.New(mysql.Config{Conn: db})
		gdb, err2 := gorm.Open(dia, &gorm.Config{})
		if err2 != nil {
			return err2
		}
		m := model.Orders{}

		tx1 := gdb.Table("orders").Where("id", orders.Id).First(&m)
		if tx1.Error != nil {
			if tx1.Error.Error() != "record not found" {
				return tx1.Error
			}
		}
		if m.Id != "" {
			return errors.New("订单已经存在")
		}

		tx := gdb.Exec("insert into `orders`(`id`,`userid`,`status`,`postage`,`shoppingid`,`payment`) values (?,?,?,?,?,?)",
			orders.Id, orders.Userid, orders.Status, orders.Postage, orders.Shoppingid, orders.Payment)
		//--------
		//tx := gdb.Exec("update `orders` set status = ? where id =?", orders.Status, orders.Id)
		//exec, err := db.Exec("update `orders` set status = ? where id =?", 2, "1")
		//------- 这里两个都可以
		//fmt.Println("---", exec, err, "aaa")
		fmt.Println("----", tx.Error, tx.RowsAffected)
		if tx.Error != nil {
			return tx.Error
		}
		if tx.RowsAffected == 0 {
			return errors.New("受影响的为0,修改失败")
		}

		//if err != nil {
		//	return dtmcli.ErrFailure
		//}
		//affected, err1 := exec.RowsAffected()
		//if err1 != nil {
		//	return dtmcli.ErrFailure
		//}
		//if affected == 0 {
		//	return dtmcli.ErrFailure
		//}

		return nil
	})

	if err2 != nil {
		return nil, err2
	}
	return &orderrpc.CreateOrderResponse{
		Affected: 1,
	}, nil
}
