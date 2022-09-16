package payment

import (
	"context"
	"go-travel/service/payment/cmd/rpc/payment"
	"go-travel/service/payment/cmd/rpc/pb"

	"go-travel/service/admin/cmd/api/internal/svc"
	"go-travel/service/admin/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PaymentListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPaymentListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PaymentListLogic {
	return &PaymentListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

type PaymentListApiResp struct {
	Total int64
	List  []*pb.ListPaymentItemResp
}

func (l *PaymentListLogic) PaymentList(req *types.ListPaymentReq) (resp *PaymentListApiResp, err error) {
	// todo: add your logic here and delete this line
	res, err := l.svcCtx.PaymentRpc.ListPayment(l.ctx, &payment.ListPaymentReq{
		//Info:     req.Info,
		Page:     req.Page,
		PageSize: req.PageSize,
	})

	if err != nil {
		//logx.WithContext(l.ctx).Errorf("page: %s and pagesize:%s 查询用户异常:%s", req.Page, req.PageSize, err.Error())
		return nil, err //errors.New("payment list error")
	}
	return &PaymentListApiResp{
		Total: res.Total,
		List:  res.List,
	}, nil
}
