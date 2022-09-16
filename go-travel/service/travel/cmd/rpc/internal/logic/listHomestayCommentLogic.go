package logic

import (
	"context"
	"encoding/json"
	"go-travel/service/travel/cmd/rpc/travel"

	"go-travel/service/travel/cmd/rpc/internal/svc"
	"go-travel/service/travel/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListHomestayCommentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListHomestayCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListHomestayCommentLogic {
	return &ListHomestayCommentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// homestayComment
func (l *ListHomestayCommentLogic) ListHomestayComment(in *pb.ListHomestaycommentReq) (*pb.ListHomestaycommentResp, error) {
	// todo: add your logic here and delete this line
	countbuilder := l.svcCtx.HomestayCommentModel.CountBuilder("id")
	count, err := l.svcCtx.HomestayCommentModel.FindCount(l.ctx, countbuilder)
	if err != nil {
		restr, _ := json.Marshal(in)
		logx.WithContext(l.ctx).Errorf("查询homestay_order信息count失败,参数%s,异常:%s", restr, err.Error())
		return nil, err
	}

	rowBuilder := l.svcCtx.HomestayCommentModel.RowBuilder()

	all, err := l.svcCtx.HomestayCommentModel.FindPageListByPage(l.ctx, rowBuilder, in.Page, in.PageSize, "id") //Pageno Current
	if err != nil {
		reqStr, _ := json.Marshal(in)
		logx.WithContext(l.ctx).Errorf("分页查询所有信息失败,参数:%s,异常:%s", reqStr, err.Error())
		return nil, err
	}

	var outList []*travel.ListHomestaycommentItemResp
	for _, homestaycommentlist := range all {
		//string->time
		createtime := homestaycommentlist.CreateTime.Format("2006-01-02 15:04:05")
		updatetime := homestaycommentlist.UpdateTime.Format("2006-01-02 15:04:05")
		deletetime := homestaycommentlist.DeleteTime.Format("2006-01-02 15:04:05")

		//can not all
		outItem := &travel.ListHomestaycommentItemResp{
			Id:         homestaycommentlist.Id,
			CreateTime: createtime,
			UpdateTime: updatetime,
			DeleteTime: deletetime,
			Version:    homestaycommentlist.Version,
			DelState:   homestaycommentlist.DelState,
			HomestayId: homestaycommentlist.HomestayId,
			UserId:     homestaycommentlist.UserId,
			Star:       float32(homestaycommentlist.Star),
			Content:    homestaycommentlist.Content,
		}
		outList = append(outList, outItem)
	}
	reqStr, _ := json.Marshal(in)
	listStr, _ := json.Marshal(outList)
	logx.WithContext(l.ctx).Infof("分页查询所有homestayActivity信息,参数：%s,响应：%s", reqStr, listStr)

	return &pb.ListHomestaycommentResp{
		Total: count,
		List:  outList,
	}, nil
}
