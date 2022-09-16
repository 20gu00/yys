package logic

import (
	"context"
	"errors"

	"go-travel/service/travel/cmd/rpc/internal/svc"
	"go-travel/service/travel/cmd/rpc/pb"
	"go-travel/service/travel/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddHomestayBusinessLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddHomestayBusinessLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddHomestayBusinessLogic {
	return &AddHomestayBusinessLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// homestayBusiness
func (l *AddHomestayBusinessLogic) AddHomestayBusiness(in *pb.AddHomestayBusinessReq) (*pb.AddHomestayActivityResp, error) {
	// todo: add your logic here and delete this line

	_, err := l.svcCtx.HomestayBusinessModel.FindOne(l.ctx, in.Id)
	if err == nil {
		return nil, errors.New("homestaybusiness is exist")
	}
	//string->time.Time
	//createtime, _ := time.Parse("2006-01-02 15:04:05", in.CreateTime)
	//updatetime, _ := time.Parse("2006-01-02 15:04:05", in.UpdateTime)
	//deletetime, _ := time.Parse("2006-01-02 15:04:05", in.DeleteTime)
	_, err = l.svcCtx.HomestayBusinessModel.InsertById(l.ctx, &model.HomestayBusiness{
		Id: in.Id,
		//CreateTime:  createtime,
		//UpdateTime:  updatetime,
		//DeleteTime:  deletetime,
		DelState:    in.DelState,
		Title:       in.Title,
		UserId:      in.UserId,
		Info:        in.Info,
		BossInfo:    in.BossInfo,
		LicenseFron: in.LicenseFron,
		LicenseBack: in.LicenseBack,
		RowState:    in.RowState,
		Star:        float64(in.Star), //star
		Tags:        in.Tag,
		Cover:       in.Cover,
		HeaderImg:   in.HeaderImg,
		Version:     in.Version,
	})
	if err != nil {
		return nil, errors.New("add homestay bussiness error")
	}
	return &pb.AddHomestayActivityResp{
		Pong: 1,
	}, nil
}
