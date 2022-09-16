package logic

import (
	"context"
	"encoding/json"

	"go-travel/service/usercenter/cmd/rpc/internal/svc"
	"go-travel/service/usercenter/cmd/rpc/pb"
	"go-travel/service/usercenter/cmd/rpc/usercenter"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserListLogic {
	return &UserListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserListLogic) UserList(in *pb.UserListReq) (*pb.UserListResp, error) {
	// todo: add your logic here and delete this line

	//[]* *[]
	countbuilder := l.svcCtx.UserModel.CountBuilder("id")
	count, err := l.svcCtx.UserModel.FindCount(l.ctx, countbuilder)
	if err != nil {
		logx.WithContext(l.ctx).Errorf("查询user信息count失败,参数,异常:%s", err.Error())
		return nil, err
	}

	rowBuilder := l.svcCtx.UserModel.RowBuilder()

	//all, err := l.svcCtx.UserRpcModel.FindAll(l.ctx, in.PageNo, in.PageSize) //Pageno Current
	all, err := l.svcCtx.UserModel.FindPageListByPage(l.ctx, rowBuilder, in.PageNo, in.PageSize, "id")
	if err != nil {
		//reqStr, _ := json.Marshal(in.Info)
		// rpc terminal can use
		logx.WithContext(l.ctx).Errorf("分页查询user信息失败,参数,异常:%s", err.Error())
		return nil, err
	}

	var outList []*usercenter.UserListItemResp
	for _, userlist := range all { //*all

		//string->time
		//createtime, _ := time.Parse("2006-01-02 15:04:05", userlist.CreateTime)
		//updatetime, _ := time.Parse("2006-01-02 15:04:05", userlist.UpdateTime)
		//deletetime, _ := time.Parse("2006-01-02 15:04:05", userlist.DeleteTime)

		//time->string
		createtime := userlist.CreateTime.Format("2006-01-02 15:04:05")
		updatetime := userlist.UpdateTime.Format("2006-01-02 15:04:05")
		deletetime := userlist.DeleteTime.Format("2006-01-02 15:04:05")

		outItem := &usercenter.UserListItemResp{ //& not&
			ID: userlist.Id,

			//string
			CreateTime: createtime, //time.Unix(userlist.CreateTime, 0),  //time.Unix()函数用于产生本地时间  UTC,+in.LiveEndTime second
			UpdateTime: updatetime,
			DeleteTime: deletetime,

			DelState: userlist.DelState,
			Version:  userlist.Version,
			Mobile:   userlist.Mobile,
			Password: userlist.Password,
			NickName: userlist.Nickname,
			Sex:      userlist.Sex,
			Avatar:   userlist.Avatar,
			Info:     userlist.Info,
		}
		outList = append(outList, outItem)
	}

	reqStr, _ := json.Marshal(in)
	listStr, _ := json.Marshal(outList)
	logx.WithContext(l.ctx).Infof("分页查询所有user信息,参数：%s,响应：%s", reqStr, listStr)

	return &pb.UserListResp{
		Total: count,
		List:  outList,
	}, nil
}
