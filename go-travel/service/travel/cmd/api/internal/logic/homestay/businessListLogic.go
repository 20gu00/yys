package homestay

import (
	"context"

	"go-travel/common/tool"
	"go-travel/common/xerr"
	"go-travel/service/travel/cmd/api/internal/svc"
	"go-travel/service/travel/cmd/api/internal/types"

	"github.com/Masterminds/squirrel"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type BusinessListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewBusinessListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BusinessListLogic {
	return &BusinessListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *BusinessListLogic) BusinessList(req *types.BusinessListReq) (resp *types.BusinessListResp, err error) {
	// todo: add your logic here and delete this line
	whereBuilder := l.svcCtx.HomestayModel.RowBuilder().Where(squirrel.Eq{"homestay_business_id": req.HomestayBusinessId})
	list, err := l.svcCtx.HomestayModel.FindPageListByIdDESC(l.ctx, whereBuilder, req.LastId, req.PageSize)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "HomestayBusinessId: %d ,err : %v", req.HomestayBusinessId, err)
	}

	var res []types.Homestay
	if len(list) > 0 {
		for _, homestay := range list {

			var typeHomestay types.Homestay
			_ = copier.Copy(&typeHomestay, homestay)

			typeHomestay.FoodPrice = tool.Fen2Yuan(homestay.FoodPrice)
			typeHomestay.HomestayPrice = tool.Fen2Yuan(homestay.HomestayPrice)
			typeHomestay.MarketHomestayPrice = tool.Fen2Yuan(homestay.MarketHomestayPrice)

			res = append(res, typeHomestay)
		}
	}

	return &types.BusinessListResp{
		List: res,
	}, nil
}
