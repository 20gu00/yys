package logic

import (
	"context"
	"errors"
	"go-travel/service/travel/model"

	"go-travel/service/travel/cmd/rpc/internal/svc"
	"go-travel/service/travel/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddHomestayCommentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddHomestayCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddHomestayCommentLogic {
	return &AddHomestayCommentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddHomestayCommentLogic) AddHomestayComment(in *pb.AddHomestaycommentReq) (*pb.AddHomestaycommentResp, error) {
	// todo: add your logic here and delete this line
	_, err := l.svcCtx.HomestayCommentModel.FindOne(l.ctx, in.Id)
	if err == nil {
		return nil, errors.New("homestaycomment is exist")
	}

	//string->time.Time
	//createtime, _ := time.Parse("2006-01-02 15:04:05", in.CreateTime)
	//updatetime, _ := time.Parse("2006-01-02 15:04:05", in.UpdateTime)
	//deletetime, _ := time.Parse("2006-01-02 15:04:05", in.DeleteTime)

	_, err = l.svcCtx.HomestayCommentModel.Insert(l.ctx, &model.HomestayComment{
		Id: in.Id,
		//CreateTime :
		//UpdateTime time.Time `db:"update_time"`
		//DeleteTime time.Time `db:"delete_time"`
		DelState:   in.DelState,
		HomestayId: in.HomestayId,
		UserId:     in.UserId,
		Content:    in.Content,
		Star:       float64(in.Star),
		Version:    in.Version,
	})

	if err != nil {
		return nil, errors.New("add homestay comment error")
	}
	return &pb.AddHomestaycommentResp{
		Pong: 1,
	}, nil

}
