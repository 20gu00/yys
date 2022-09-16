package homestay

import (
	"context"

	"go-travel/common/tool"
	"go-travel/common/xerr"
	"go-travel/service/travel/cmd/api/internal/svc"
	"go-travel/service/travel/cmd/api/internal/types"

	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type GuessListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGuessListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GuessListLogic {
	return &GuessListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GuessListLogic) GuessList(req *types.GuessListReq) (resp *types.GuessListResp, err error) {
	// todo: add your logic here and delete this line
	var res []types.Homestay

	list, err := l.svcCtx.HomestayModel.FindPageListByIdDESC(l.ctx, l.svcCtx.HomestayModel.RowBuilder(), 0, 5)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "GuessList db err req : %+v , err : %v", req, err)
	}

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

	return &types.GuessListResp{
		List: res,
	}, nil
}
