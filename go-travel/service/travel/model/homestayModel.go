package model

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"go-travel/common/globalkey"
	"go-travel/common/xerr"
	"time"
)

var _ HomestayModel = (*customHomestayModel)(nil)

type (
	// HomestayModel is an interface to be customized, add more methods here,
	// and implement the added methods in customHomestayModel.
	HomestayModel interface {
		homestayModel
		Trans(ctx context.Context, fn func(context context.Context, session sqlx.Session) error) error
		RowBuilder() squirrel.SelectBuilder
		CountBuilder(field string) squirrel.SelectBuilder
		SumBuilder(field string) squirrel.SelectBuilder
		DeleteSoft(ctx context.Context, session sqlx.Session, data *Homestay) error
		FindOneByQuery(ctx context.Context, rowBuilder squirrel.SelectBuilder) (*Homestay, error)
		FindSum(ctx context.Context, sumBuilder squirrel.SelectBuilder) (float64, error)
		FindCount(ctx context.Context, countBuilder squirrel.SelectBuilder) (int64, error)
		FindAll(ctx context.Context, rowBuilder squirrel.SelectBuilder, orderBy string) ([]*Homestay, error)
		FindPageListByPage(ctx context.Context, rowBuilder squirrel.SelectBuilder, page, pageSize int64, orderBy string) ([]*Homestay, error)
		FindPageListByIdDESC(ctx context.Context, rowBuilder squirrel.SelectBuilder, preMinId, pageSize int64) ([]*Homestay, error)
		FindPageListByIdASC(ctx context.Context, rowBuilder squirrel.SelectBuilder, preMaxId, pageSize int64) ([]*Homestay, error)

		//provide to admin use
		InsertById(ctx context.Context, data *Homestay) (sql.Result, error)
		FindAllForPage(ctx context.Context, Info string, Current int64, PageSize int64) ([]*Homestay, error)
		Count(ctx context.Context) (int64, error)
		UpdateNoSess(ctx context.Context, data *Homestay) (sql.Result, error)
		DeleteById(ctx context.Context, id int64) error
	}

	customHomestayModel struct {
		*defaultHomestayModel
	}
)

// NewHomestayModel returns a model for the database table.
func NewHomestayModel(conn sqlx.SqlConn, c cache.CacheConf) HomestayModel {
	return &customHomestayModel{
		defaultHomestayModel: newHomestayModel(conn, c),
	}
}

func (m *defaultHomestayModel) DeleteSoft(ctx context.Context, session sqlx.Session, data *Homestay) error {
	data.DelState = globalkey.DelStateYes
	data.DeleteTime = time.Now()
	if err := m.UpdateWithVersion(ctx, session, data); err != nil {
		return errors.Wrapf(xerr.NewErrMsg("删除数据失败"), "HomestayModel delete err : %+v", err)
	}
	return nil
}

func (m *defaultHomestayModel) FindOneByQuery(ctx context.Context, rowBuilder squirrel.SelectBuilder) (*Homestay, error) {

	query, values, err := rowBuilder.Where("del_state = ?", globalkey.DelStateNo).ToSql()
	if err != nil {
		return nil, err
	}

	var resp Homestay
	err = m.QueryRowNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return &resp, nil
	default:
		return nil, err
	}
}

func (m *defaultHomestayModel) FindSum(ctx context.Context, sumBuilder squirrel.SelectBuilder) (float64, error) {

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

func (m *defaultHomestayModel) FindCount(ctx context.Context, countBuilder squirrel.SelectBuilder) (int64, error) {

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

func (m *defaultHomestayModel) FindAll(ctx context.Context, rowBuilder squirrel.SelectBuilder, orderBy string) ([]*Homestay, error) {

	if orderBy == "" {
		rowBuilder = rowBuilder.OrderBy("id DESC")
	} else {
		rowBuilder = rowBuilder.OrderBy(orderBy)
	}

	query, values, err := rowBuilder.Where("del_state = ?", globalkey.DelStateNo).ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*Homestay
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *defaultHomestayModel) FindPageListByPage(ctx context.Context, rowBuilder squirrel.SelectBuilder, page, pageSize int64, orderBy string) ([]*Homestay, error) {

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

	var resp []*Homestay
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *defaultHomestayModel) FindPageListByIdDESC(ctx context.Context, rowBuilder squirrel.SelectBuilder, preMinId, pageSize int64) ([]*Homestay, error) {

	if preMinId > 0 {
		rowBuilder = rowBuilder.Where(" id < ? ", preMinId)
	}

	query, values, err := rowBuilder.Where("del_state = ?", globalkey.DelStateNo).OrderBy("id DESC").Limit(uint64(pageSize)).ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*Homestay
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

//按照id升序分页查询数据，不支持排序
func (m *defaultHomestayModel) FindPageListByIdASC(ctx context.Context, rowBuilder squirrel.SelectBuilder, preMaxId, pageSize int64) ([]*Homestay, error) {

	if preMaxId > 0 {
		rowBuilder = rowBuilder.Where(" id > ? ", preMaxId)
	}

	query, values, err := rowBuilder.Where("del_state = ?", globalkey.DelStateNo).OrderBy("id ASC").Limit(uint64(pageSize)).ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*Homestay
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

// export logic
func (m *defaultHomestayModel) Trans(ctx context.Context, fn func(ctx context.Context, session sqlx.Session) error) error {

	return m.TransactCtx(ctx, func(ctx context.Context, session sqlx.Session) error {
		return fn(ctx, session)
	})

}

// export logic
func (m *defaultHomestayModel) RowBuilder() squirrel.SelectBuilder {
	return squirrel.Select(homestayRows).From(m.table)
}

// export logic
func (m *defaultHomestayModel) CountBuilder(field string) squirrel.SelectBuilder {
	return squirrel.Select("COUNT(" + field + ")").From(m.table)
}

// export logic
func (m *defaultHomestayModel) SumBuilder(field string) squirrel.SelectBuilder {
	return squirrel.Select("IFNULL(SUM(" + field + "),0)").From(m.table)
}

func (m *defaultHomestayModel) InsertById(ctx context.Context, data *Homestay) (sql.Result, error) {
	data.DeleteTime = time.Unix(0, 0)
	looklookTravelHomestayIdKey := fmt.Sprintf("%s%v", cacheLooklookTravelHomestayIdPrefix, data.Id)
	return m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, homestayRowsExpectAutoSet)

		return conn.ExecCtx(ctx, query, data.DeleteTime, data.DelState, data.Version, data.Title, data.SubTitle, data.Banner, data.Info, data.PeopleNum, data.HomestayBusinessId, data.UserId, data.RowState, data.RowType, data.FoodInfo, data.FoodPrice, data.HomestayPrice, data.MarketHomestayPrice)
	}, looklookTravelHomestayIdKey)
}

func (m *defaultHomestayModel) FindAllForPage(ctx context.Context, Info string, Current int64, PageSize int64) ([]*Homestay, error) {
	looklookTravelHomestayInfoKey := fmt.Sprintf("%s%v", cacheLooklookTravelHomestayInfoPrefix, Info)

	var resp []*Homestay
	query1 := fmt.Sprintf("select %s from %s ", homestayRows, m.table)
	query2 := fmt.Sprintf("where id like '%s' ", Info) //like need ''   or and
	query3 := fmt.Sprintf("order by `id` ASC limit ?,?")

	err := m.QueryRowCtx(ctx, &resp, looklookTravelHomestayInfoKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
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

func (m *defaultHomestayModel) Count(ctx context.Context) (int64, error) {
	looklookTravelHomestayCountKey := fmt.Sprintf("%s%v", cacheLooklookTravelHomestayCountPrefix, "no")
	var count int64

	//err := m.QueryRowCtx(ctx, &resp, goGatewayGatewayServiceInfoCountKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
	err := m.QueryRowNoCacheCtx(ctx, &count, looklookTravelHomestayCountKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
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

func (m *defaultHomestayModel) UpdateNoSess(ctx context.Context, data *Homestay) (sql.Result, error) {
	//time.Time
	data.DeleteTime = time.Unix(0, 0)
	looklookTravelHomestayIdKey := fmt.Sprintf("%s%v", cacheLooklookTravelHomestayIdPrefix, data.Id)
	return m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, homestayRowsWithPlaceHolder)

		return conn.ExecCtx(ctx, query, data.DeleteTime, data.DelState, data.Version, data.Title, data.SubTitle, data.Banner, data.Info, data.PeopleNum, data.HomestayBusinessId, data.UserId, data.RowState, data.RowType, data.FoodInfo, data.FoodPrice, data.HomestayPrice, data.MarketHomestayPrice, data.Id)
	}, looklookTravelHomestayIdKey)
}

func (m *defaultHomestayModel) DeleteById(ctx context.Context, id int64) error {
	looklookTravelHomestayIdKey := fmt.Sprintf("%s%v", cacheLooklookTravelHomestayIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, looklookTravelHomestayIdKey)
	return err
}
