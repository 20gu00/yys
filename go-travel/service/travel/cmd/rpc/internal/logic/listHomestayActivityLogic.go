package logic

import (
	"context"
	"encoding/json"
	"go-travel/service/travel/cmd/rpc/travel"

	"go-travel/service/travel/cmd/rpc/internal/svc"
	"go-travel/service/travel/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListHomestayActivityLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListHomestayActivityLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListHomestayActivityLogic {
	return &ListHomestayActivityLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ListHomestayActivityLogic) ListHomestayActivity(in *pb.ListHomestayActivityReq) (*pb.ListHomestayActivityResp, error) {
	// todo: add your logic here and delete this line
	countbuilder := l.svcCtx.HomestayActivityModel.CountBuilder("id")
	count, err := l.svcCtx.HomestayActivityModel.FindCount(l.ctx, countbuilder)
	if err != nil {
		restr, _ := json.Marshal(in)
		logx.WithContext(l.ctx).Errorf("查询homestay_order信息count失败,参数%s,异常:%s", restr, err.Error())
		return nil, err
	}

	rowBuilder := l.svcCtx.HomestayActivityModel.RowBuilder()

	all, err := l.svcCtx.HomestayActivityModel.FindPageListByPage(l.ctx, rowBuilder, in.Page, in.PageSize, "id") //Pageno Current
	if err != nil {
		reqStr, _ := json.Marshal(in)
		logx.WithContext(l.ctx).Errorf("分页查询所有信息失败,参数:%s,异常:%s", reqStr, err.Error())
		return nil, err
	}

	var outList []*travel.ListHomestayAcitivityItemResp
	for _, homestayActivitylist := range all {
		//string->time
		createtime := homestayActivitylist.CreateTime.Format("2006-01-02 15:04:05")
		updatetime := homestayActivitylist.UpdateTime.Format("2006-01-02 15:04:05")
		deletetime := homestayActivitylist.DeleteTime.Format("2006-01-02 15:04:05")

		//can not all
		outItem := &travel.ListHomestayAcitivityItemResp{
			Id:         homestayActivitylist.Id,
			CreateTime: createtime,
			UpdateTime: updatetime,
			DeleteTime: deletetime,
			DelState:   homestayActivitylist.DelState,
			RowType:    homestayActivitylist.RowType,
			DataId:     homestayActivitylist.DataId,
			RowStatus:  homestayActivitylist.RowStatus,
			Version:    homestayActivitylist.Version,
		}
		outList = append(outList, outItem)
	}
	reqStr, _ := json.Marshal(in)
	listStr, _ := json.Marshal(outList)
	logx.WithContext(l.ctx).Infof("分页查询所有homestayActivity信息,参数：%s,响应：%s", reqStr, listStr)

	return &pb.ListHomestayActivityResp{
		Total: count,
		List:  outList,
	}, nil
}
