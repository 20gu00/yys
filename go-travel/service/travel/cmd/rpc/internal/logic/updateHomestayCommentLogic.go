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

type UpdateHomestayCommentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateHomestayCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateHomestayCommentLogic {
	return &UpdateHomestayCommentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateHomestayCommentLogic) UpdateHomestayComment(in *pb.UpdateHomestaycommentReq) (*pb.UpdateHomestaycommentResp, error) {
	// todo: add your logic here and delete this line
	_, err := l.svcCtx.HomestayCommentModel.FindOne(l.ctx, in.Id)
	switch err {
	case nil:
	case sqlc.ErrNotFound:
		logx.WithContext(l.ctx).Errorf("homestaycomment不存在,参数:%d,异常:%s", in.Id, err.Error())
		return nil, errors.New("homestaycomment not exist")
	default:
		logx.WithContext(l.ctx).Errorf("homestaycomment search操作未成功,参数:%d,异常:%s", in.Id, err.Error())
		return nil, err
	}

	//string->time.Time
	createtime, _ := time.Parse("2006-01-02 15:04:05", in.CreateTime)
	updatetime, _ := time.Parse("2006-01-02 15:04:05", in.UpdateTime)
	deletetime, _ := time.Parse("2006-01-02 15:04:05", in.DeleteTime)
	data := &model.HomestayComment{
		Id:         in.Id,
		CreateTime: createtime,
		UpdateTime: updatetime,
		DeleteTime: deletetime,
		Version:    in.Version,
		DelState:   in.DelState,
		HomestayId: in.HomestayId,
		UserId:     in.UserId,
		Star:       float64(in.Star),
		Content:    in.Content,
	}
	err = l.svcCtx.HomestayCommentModel.Update(l.ctx, data)
	if err != nil {
		logx.WithContext(l.ctx).Errorf("update homestayComment error(rpc)")
		return nil, errors.New("update homestaycOMMENT error(rpc)")
	}
	return &pb.UpdateHomestaycommentResp{
		Pong: 1,
	}, nil
}
