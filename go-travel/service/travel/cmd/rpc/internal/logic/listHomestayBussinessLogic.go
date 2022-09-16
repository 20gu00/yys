package logic

import (
	"context"
	"encoding/json"
	"go-travel/service/travel/cmd/rpc/travel"

	"go-travel/service/travel/cmd/rpc/internal/svc"
	"go-travel/service/travel/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListHomestayBussinessLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListHomestayBussinessLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListHomestayBussinessLogic {
	return &ListHomestayBussinessLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ListHomestayBussinessLogic) ListHomestayBussiness(in *pb.ListHomestayBusinessReq) (*pb.ListHomestayBusinessResp, error) {
	// todo: add your logic here and delete this line
	countbuilder := l.svcCtx.HomestayBusinessModel.CountBuilder("id")
	count, err := l.svcCtx.HomestayBusinessModel.FindCount(l.ctx, countbuilder)
	if err != nil {
		logx.WithContext(l.ctx).Errorf("查询user信息count失败,参数,异常:%s", err.Error())
		return nil, err
	}
	rowBuilder := l.svcCtx.HomestayBusinessModel.RowBuilder()
	all, err := l.svcCtx.HomestayBusinessModel.FindPageListByPage(l.ctx, rowBuilder, in.Page, in.PageSize, "id") //Pageno Current

	if err != nil {
		reqStr, _ := json.Marshal(in)
		logx.WithContext(l.ctx).Errorf("分页查询所有信息失败,参数:%s,异常:%s", reqStr, err.Error())
		return nil, err
	}

	var outList []*travel.ListHomestayBusinessItemResp
	for _, homestayBusinesslist := range all {
		createtime := homestayBusinesslist.CreateTime.Format("2006-01-02 15:04:05")
		updatetime := homestayBusinesslist.UpdateTime.Format("2006-01-02 15:04:05")
		deletetime := homestayBusinesslist.DeleteTime.Format("2006-01-02 15:04:05")

		outItem := &travel.ListHomestayBusinessItemResp{
			Id:         homestayBusinesslist.Id,
			CreateTime: createtime,
			UpdateTime: updatetime,
			DeleteTime: deletetime,
			DelState:   homestayBusinesslist.DelState,
			RowState:   homestayBusinesslist.RowState,
			Version:    homestayBusinesslist.Version,

			Title:       homestayBusinesslist.Title,
			UserId:      homestayBusinesslist.UserId,
			Info:        homestayBusinesslist.Info,
			BossInfo:    homestayBusinesslist.BossInfo,
			LicenseFron: homestayBusinesslist.LicenseFron,
			LicenseBack: homestayBusinesslist.LicenseBack,
			Star:        float32(homestayBusinesslist.Star), //star
			Tag:         homestayBusinesslist.Tags,
			Cover:       homestayBusinesslist.Cover,
			HeaderImg:   homestayBusinesslist.HeaderImg,
		}
		outList = append(outList, outItem)
	}
	reqStr, _ := json.Marshal(in)
	listStr, _ := json.Marshal(outList)
	logx.WithContext(l.ctx).Infof("分页查询所有homestayBusiness信息,参数：%s,响应：%s", reqStr, listStr)

	return &pb.ListHomestayBusinessResp{
		Total: count,
		List:  outList,
	}, nil
}
