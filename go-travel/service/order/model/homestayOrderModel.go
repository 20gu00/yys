package model

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"go-travel/common/globalkey"
	"go-travel/common/xerr"
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ HomestayOrderModel = (*customHomestayOrderModel)(nil)

type (
	// HomestayOrderModel is an interface to be customized, add more methods here,
	// and implement the added methods in customHomestayOrderModel.
	HomestayOrderModel interface {
		homestayOrderModel
		Trans(ctx context.Context, fn func(context context.Context, session sqlx.Session) error) error
		RowBuilder() squirrel.SelectBuilder
		CountBuilder(field string) squirrel.SelectBuilder
		SumBuilder(field string) squirrel.SelectBuilder
		DeleteSoft(ctx context.Context, session sqlx.Session, data *HomestayOrder) error
		FindOneByQuery(ctx context.Context, rowBuilder squirrel.SelectBuilder) (*HomestayOrder, error)
		FindSum(ctx context.Context, sumBuilder squirrel.SelectBuilder) (float64, error)
		FindCount(ctx context.Context, countBuilder squirrel.SelectBuilder) (int64, error)
		FindAll(ctx context.Context, rowBuilder squirrel.SelectBuilder, orderBy string) ([]*HomestayOrder, error)
		FindPageListByPage(ctx context.Context, rowBuilder squirrel.SelectBuilder, page, pageSize int64, orderBy string) ([]*HomestayOrder, error)
		FindPageListByIdDESC(ctx context.Context, rowBuilder squirrel.SelectBuilder, preMinId, pageSize int64) ([]*HomestayOrder, error)
		FindPageListByIdASC(ctx context.Context, rowBuilder squirrel.SelectBuilder, preMaxId, pageSize int64) ([]*HomestayOrder, error)
		//testrdo
		InsertById(ctx context.Context, data *HomestayOrder) (sql.Result, error)
		DeleteById(ctx context.Context, id int64) error
		FindAllForPage(ctx context.Context, Info string, Current int64, PageSize int64) ([]*HomestayOrder, error)
		Count(ctx context.Context) (int64, error)
		UpdateNoTrans(ctx context.Context, data *HomestayOrder) (sql.Result, error)
	}

	customHomestayOrderModel struct {
		*defaultHomestayOrderModel
	}
)

// NewHomestayOrderModel returns a model for the database table.
func NewHomestayOrderModel(conn sqlx.SqlConn, c cache.CacheConf) HomestayOrderModel {
	return &customHomestayOrderModel{
		defaultHomestayOrderModel: newHomestayOrderModel(conn, c),
	}
}

func (m *defaultHomestayOrderModel) DeleteSoft(ctx context.Context, session sqlx.Session, data *HomestayOrder) error {
	data.DelState = globalkey.DelStateYes
	data.DeleteTime = time.Now()
	if err := m.UpdateWithVersion(ctx, session, data); err != nil {
		return errors.Wrapf(xerr.NewErrMsg("删除数据失败"), "HomestayOrderModel delete err : %+v", err)
	}
	return nil
}

func (m *defaultHomestayOrderModel) FindOneByQuery(ctx context.Context, rowBuilder squirrel.SelectBuilder) (*HomestayOrder, error) {

	query, values, err := rowBuilder.Where("del_state = ?", globalkey.DelStateNo).ToSql()
	if err != nil {
		return nil, err
	}

	var resp HomestayOrder
	err = m.QueryRowNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return &resp, nil
	default:
		return nil, err
	}
}

func (m *defaultHomestayOrderModel) FindSum(ctx context.Context, sumBuilder squirrel.SelectBuilder) (float64, error) {

	query, values, err := sumBuilder.Where("del_state = ?", globalkey.DelStateNo).ToSql()
	if err != nil {
		return 0, err
	}

	var resp float64
	err = m.QueryRowNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return 0, err
	}
}

func (m *defaultHomestayOrderModel) FindCount(ctx context.Context, countBuilder squirrel.SelectBuilder) (int64, error) {

	query, values, err := countBuilder.Where("del_state = ?", globalkey.DelStateNo).ToSql()
	if err != nil {
		return 0, err
	}

	var resp int64
	err = m.QueryRowNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return 0, err
	}
}

func (m *defaultHomestayOrderModel) FindAll(ctx context.Context, rowBuilder squirrel.SelectBuilder, orderBy string) ([]*HomestayOrder, error) {

	if orderBy == "" {
		rowBuilder = rowBuilder.OrderBy("id DESC")
	} else {
		rowBuilder = rowBuilder.OrderBy(orderBy)
	}

	query, values, err := rowBuilder.Where("del_state = ?", globalkey.DelStateNo).ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*HomestayOrder
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *defaultHomestayOrderModel) FindPageListByPage(ctx context.Context, rowBuilder squirrel.SelectBuilder, page, pageSize int64, orderBy string) ([]*HomestayOrder, error) {

	if orderBy == "" {
		rowBuilder = rowBuilder.OrderBy("id DESC")
	} else {
		rowBuilder = rowBuilder.OrderBy(orderBy)
	}

	if page < 1 {
		page = 1
	}
	offset := (page - 1) * pageSize

	query, values, err := rowBuilder.Where("del_state = ?", globalkey.DelStateNo).Offset(uint64(offset)).Limit(uint64(pageSize)).ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*HomestayOrder
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *defaultHomestayOrderModel) FindPageListByIdDESC(ctx context.Context, rowBuilder squirrel.SelectBuilder, preMinId, pageSize int64) ([]*HomestayOrder, error) {

	if preMinId > 0 {
		rowBuilder = rowBuilder.Where(" id < ? ", preMinId)
	}

	query, values, err := rowBuilder.Where("del_state = ?", globalkey.DelStateNo).OrderBy("id DESC").Limit(uint64(pageSize)).ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*HomestayOrder
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

//按照id升序分页查询数据，不支持排序
func (m *defaultHomestayOrderModel) FindPageListByIdASC(ctx context.Context, rowBuilder squirrel.SelectBuilder, preMaxId, pageSize int64) ([]*HomestayOrder, error) {

	if preMaxId > 0 {
		rowBuilder = rowBuilder.Where(" id > ? ", preMaxId)
	}

	query, values, err := rowBuilder.Where("del_state = ?", globalkey.DelStateNo).OrderBy("id ASC").Limit(uint64(pageSize)).ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*HomestayOrder
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

// export logic
func (m *defaultHomestayOrderModel) Trans(ctx context.Context, fn func(ctx context.Context, session sqlx.Session) error) error {

	return m.TransactCtx(ctx, func(ctx context.Context, session sqlx.Session) error {
		return fn(ctx, session)
	})

}

// export logic
func (m *defaultHomestayOrderModel) RowBuilder() squirrel.SelectBuilder {
	return squirrel.Select(homestayOrderRows).From(m.table)
}

// export logic
func (m *defaultHomestayOrderModel) CountBuilder(field string) squirrel.SelectBuilder {
	return squirrel.Select("COUNT(" + field + ")").From(m.table)
}

// export logic
func (m *defaultHomestayOrderModel) SumBuilder(field string) squirrel.SelectBuilder {
	return squirrel.Select("IFNULL(SUM(" + field + "),0)").From(m.table)
}

//testdo
func (m *defaultHomestayOrderModel) InsertById(ctx context.Context, data *HomestayOrder) (sql.Result, error) {
	data.DeleteTime = time.Unix(0, 0)
	looklookOrderHomestayOrderIdKey := fmt.Sprintf("%s%v", cacheLooklookOrderHomestayOrderIdPrefix, data.Id)
	//looklookOrderHomestayOrderSnKey := fmt.Sprintf("%s%v", cacheLooklookOrderHomestayOrderSnPrefix, data.Sn)
	return m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, homestayOrderRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.DeleteTime, data.DelState, data.Version, data.Sn, data.UserId, data.HomestayId, data.Title, data.SubTitle, data.Cover, data.Info, data.PeopleNum, data.RowType, data.NeedFood, data.FoodInfo, data.FoodPrice, data.HomestayPrice, data.MarketHomestayPrice, data.HomestayBusinessId, data.HomestayUserId, data.LiveStartDate, data.LiveEndDate, data.LivePeopleNum, data.TradeState, data.TradeCode, data.Remark, data.OrderTotalPrice, data.FoodTotalPrice, data.HomestayTotalPrice)
	}, looklookOrderHomestayOrderIdKey)
}

func (m *defaultHomestayOrderModel) DeleteById(ctx context.Context, id int64) error {

	looklookOrderHomestayOrderIdKey := fmt.Sprintf("%s%v", cacheLooklookOrderHomestayOrderIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, looklookOrderHomestayOrderIdKey)
	return err
}

func (m *defaultHomestayOrderModel) FindAllForPage(ctx context.Context, Info string, Current int64, PageSize int64) ([]*HomestayOrder, error) { //[]*
	looklookOrderHomestayOrderInfoKey := fmt.Sprintf("%s%v:%v:%v", cacheLooklookOrderHomestayOrderInfoPrefix, Info, Current, PageSize)

	var resp []*HomestayOrder
	query1 := fmt.Sprintf("select %s from %s ", homestayOrderRows, m.table)
	query2 := fmt.Sprintf("where id like '%s' ", Info) //like need ''   or and
	query3 := fmt.Sprintf("order by `id` ASC limit ?,?")

	err := m.QueryRowCtx(ctx, &resp, looklookOrderHomestayOrderInfoKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := query1 + query2 + query3
		return conn.QueryRowCtx(ctx, v, query, (Current-1)*PageSize, PageSize)
	})
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultHomestayOrderModel) Count(ctx context.Context) (int64, error) {
	looklookOrderHomestayOrderCountKey := fmt.Sprintf("%s%v", cacheLooklookOrderHomestayOrderCountPrefix, "no")
	var count int64

	//err := m.QueryRowCtx(ctx, &resp, goGatewayGatewayServiceInfoCountKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
	err := m.QueryRowNoCacheCtx(ctx, &count, looklookOrderHomestayOrderCountKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select count(*) as count from %s", m.table) // where `is_delete`=0
		return conn.QueryRow(&count, query)
	})
	switch err {
	case nil:
		return count, nil
	case sqlc.ErrNotFound:
		return 0, ErrNotFound
	default:
		return 0, err
	}
}

func (m *defaultHomestayOrderModel) UpdateNoTrans(ctx context.Context, data *HomestayOrder) (sql.Result, error) {
	data.DeleteTime = time.Unix(0, 0)
	//`delete_time`=1970-01-01 08:00:00 +0800 CST
	looklookOrderHomestayOrderIdKey := fmt.Sprintf("%s%v", cacheLooklookOrderHomestayOrderIdPrefix, data.Id)
	return m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, homestayOrderRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.DeleteTime, data.DelState, data.Version, data.Sn, data.UserId, data.HomestayId, data.Title, data.SubTitle, data.Cover, data.Info, data.PeopleNum, data.RowType, data.NeedFood, data.FoodInfo, data.FoodPrice, data.HomestayPrice, data.MarketHomestayPrice, data.HomestayBusinessId, data.HomestayUserId, data.LiveStartDate, data.LiveEndDate, data.LivePeopleNum, data.TradeState, data.TradeCode, data.Remark, data.OrderTotalPrice, data.FoodTotalPrice, data.HomestayTotalPrice, data.Id)
	}, looklookOrderHomestayOrderIdKey)
}
