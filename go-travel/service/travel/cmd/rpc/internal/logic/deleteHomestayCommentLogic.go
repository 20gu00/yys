package logic

import (
	"context"
	"errors"

	"go-travel/service/travel/cmd/rpc/internal/svc"
	"go-travel/service/travel/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteHomestayCommentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteHomestayCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteHomestayCommentLogic {
	return &DeleteHomestayCommentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteHomestayCommentLogic) DeleteHomestayComment(in *pb.DeleteHomestaycommentReq) (*pb.DeleteHomestaycommentResp, error) {
	// todo: add your logic here and delete this line
	_, err := l.svcCtx.HomestayCommentModel.FindOne(l.ctx, in.Id)
	if err != nil {
		return nil, errors.New("homestay comment is not exist")
	}
	err = l.svcCtx.HomestayCommentModel.Delete(l.ctx, in.Id)
	if err != nil {
		return nil, errors.New("delete homestay comment error")
	}
	return &pb.DeleteHomestaycommentResp{
		Pong: 1,
	}, nil
}
