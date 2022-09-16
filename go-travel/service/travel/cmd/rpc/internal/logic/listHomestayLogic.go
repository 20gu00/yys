package logic

import (
	"context"
	"encoding/json"
	"go-travel/service/travel/cmd/rpc/travel"

	"go-travel/service/travel/cmd/rpc/internal/svc"
	"go-travel/service/travel/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListHomestayLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListHomestayLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListHomestayLogic {
	return &ListHomestayLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ListHomestayLogic) ListHomestay(in *pb.ListHomestayReq) (*pb.ListHomestayResp, error) {
	// todo: add your logic here and delete this line
	countbuilder := l.svcCtx.HomestayModel.CountBuilder("id")
	count, err := l.svcCtx.HomestayModel.FindCount(l.ctx, countbuilder)
	if err != nil {
		logx.WithContext(l.ctx).Errorf("查询user信息count失败,参数,异常:%s", err.Error())
		return nil, err
	}
	rowBuilder := l.svcCtx.HomestayModel.RowBuilder()
	all, err := l.svcCtx.HomestayModel.FindPageListByPage(l.ctx, rowBuilder, in.Page, in.PageSize, "id") //Pageno Current

	if err != nil {
		reqStr, _ := json.Marshal(in)
		logx.WithContext(l.ctx).Errorf("分页查询所有信息失败,参数:%s,异常:%s", reqStr, err.Error())
		return nil, err
	}

	var outList []*travel.ListHomestayItemResp
	for _, homestaylist := range all {
		createtime := homestaylist.CreateTime.Format("2006-01-02 15:04:05")
		updatetime := homestaylist.UpdateTime.Format("2006-01-02 15:04:05")
		deletetime := homestaylist.DeleteTime.Format("2006-01-02 15:04:05")

		outItem := &travel.ListHomestayItemResp{
			Id:                  homestaylist.Id,
			CreateTime:          createtime,
			UpdateTime:          updatetime,
			DelTime:             deletetime,
			Version:             homestaylist.Version,
			Title:               homestaylist.Title,
			SubTitle:            homestaylist.SubTitle,
			Banner:              homestaylist.Banner,
			Info:                homestaylist.Info,
			PeopleNum:           homestaylist.PeopleNum,
			HomestayBusinessId:  homestaylist.HomestayBusinessId,
			UserId:              homestaylist.UserId,
			RowState:            homestaylist.RowState,
			FoodInfo:            homestaylist.FoodInfo,
			FoodPrice:           homestaylist.FoodPrice,
			HomestayPrice:       homestaylist.HomestayPrice,
			MarketHomestayPrice: homestaylist.MarketHomestayPrice,
		}
		outList = append(outList, outItem)
	}
	reqStr, _ := json.Marshal(in)
	listStr, _ := json.Marshal(outList)
	logx.WithContext(l.ctx).Infof("分页查询所有homestay信息,参数：%s,响应：%s", reqStr, listStr)

	return &pb.ListHomestayResp{
		Total: count,
		List:  outList,
	}, nil
}
