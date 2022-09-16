package logic

import (
	"context"
	"errors"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"go-travel/service/travel/model"
	"time"

	"go-travel/service/travel/cmd/rpc/internal/svc"
	"go-travel/service/travel/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateHomestayBussinessLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateHomestayBussinessLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateHomestayBussinessLogic {
	return &UpdateHomestayBussinessLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateHomestayBussinessLogic) UpdateHomestayBussiness(in *pb.UpdateHomestayBusinessReq) (*pb.UpdateHomestayBusinessResp, error) {
	// todo: add your logic here and delete this line
	_, err := l.svcCtx.HomestayBusinessModel.FindOne(l.ctx, in.Id)
	switch err {
	case nil:
	case sqlc.ErrNotFound:
		logx.WithContext(l.ctx).Errorf("homestayBusiness不存在,参数:%d,异常:%s", in.Id, err.Error())
		return nil, errors.New("homestayBusiness not exist")
	default:
		logx.WithContext(l.ctx).Errorf("homestayBusiness search操作未成功,参数:%d,异常:%s", in.Id, err.Error())
		return nil, err
	}

	//string->time.Time
	createtime, _ := time.Parse("2006-01-02 15:04:05", in.CreateTime)
	updatetime, _ := time.Parse("2006-01-02 15:04:05", in.UpdateTime)
	deletetime, _ := time.Parse("2006-01-02 15:04:05", in.DeleteTime)
	data := &model.HomestayBusiness{
		Id:          in.Id,
		CreateTime:  createtime,
		UpdateTime:  updatetime,
		DeleteTime:  deletetime,
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
	}
	_, err = l.svcCtx.HomestayBusinessModel.UpdateNoSess(l.ctx, data)
	if err != nil {
		logx.WithContext(l.ctx).Errorf("update homestayBusiness error(rpc)")
		return nil, errors.New("update homestayBusiness error(rpc)")
	}
	return &pb.UpdateHomestayBusinessResp{
		Pong: 1,
	}, nil
}
