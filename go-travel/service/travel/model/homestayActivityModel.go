package model

import (
	"context"
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

var _ HomestayActivityModel = (*customHomestayActivityModel)(nil)

type (
	// HomestayActivityModel is an interface to be customized, add more methods here,
	// and implement the added methods in customHomestayActivityModel.
	HomestayActivityModel interface {
		homestayActivityModel
		Trans(ctx context.Context, fn func(context context.Context, session sqlx.Session) error) error
		RowBuilder() squirrel.SelectBuilder
		CountBuilder(field string) squirrel.SelectBuilder
		SumBuilder(field string) squirrel.SelectBuilder
		DeleteSoft(ctx context.Context, session sqlx.Session, data *HomestayActivity) error
		FindOneByQuery(ctx context.Context, rowBuilder squirrel.SelectBuilder) (*HomestayActivity, error)
		FindSum(ctx context.Context, sumBuilder squirrel.SelectBuilder) (float64, error)
		FindCount(ctx context.Context, countBuilder squirrel.SelectBuilder) (int64, error)
		FindAll(ctx context.Context, rowBuilder squirrel.SelectBuilder, orderBy string) ([]*HomestayActivity, error)
		FindPageListByPage(ctx context.Context, rowBuilder squirrel.SelectBuilder, page, pageSize int64, orderBy string) ([]*HomestayActivity, error)
		FindPageListByIdDESC(ctx context.Context, rowBuilder squirrel.SelectBuilder, preMinId, pageSize int64) ([]*HomestayActivity, error)
		FindPageListByIdASC(ctx context.Context, rowBuilder squirrel.SelectBuilder, preMaxId, pageSize int64) ([]*HomestayActivity, error)

		//provide to admin use
		FindAllForPage(ctx context.Context, Info string, Current int64, PageSize int64) ([]*HomestayActivity, error)
		Count(ctx context.Context) (int64, error)
	}

	customHomestayActivityModel struct {
		*defaultHomestayActivityModel
	}
)

// NewHomestayActivityModel returns a model for the database table.
func NewHomestayActivityModel(conn sqlx.SqlConn, c cache.CacheConf) HomestayActivityModel {
	return &customHomestayActivityModel{
		defaultHomestayActivityModel: newHomestayActivityModel(conn, c),
	}
}

func (m *defaultHomestayActivityModel) DeleteSoft(ctx context.Context, session sqlx.Session, data *HomestayActivity) error {
	data.DelState = globalkey.DelStateYes
	data.DeleteTime = time.Now()
	if err := m.UpdateWithVersion(ctx, session, data); err != nil {
		return errors.Wrapf(xerr.NewErrMsg("删除数据失败"), "HomestayActivityModel delete err : %+v", err)
	}
	return nil
}

func (m *defaultHomestayActivityModel) FindOneByQuery(ctx context.Context, rowBuilder squirrel.SelectBuilder) (*HomestayActivity, error) {

	query, values, err := rowBuilder.Where("del_state = ?", globalkey.DelStateNo).ToSql()
	if err != nil {
		return nil, err
	}

	var resp HomestayActivity
	err = m.QueryRowNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return &resp, nil
	default:
		return nil, err
	}
}

func (m *defaultHomestayActivityModel) FindSum(ctx context.Context, sumBuilder squirrel.SelectBuilder) (float64, error) {

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

func (m *defaultHomestayActivityModel) FindCount(ctx context.Context, countBuilder squirrel.SelectBuilder) (int64, error) {

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

func (m *defaultHomestayActivityModel) FindAll(ctx context.Context, rowBuilder squirrel.SelectBuilder, orderBy string) ([]*HomestayActivity, error) {

	if orderBy == "" {
		rowBuilder = rowBuilder.OrderBy("id DESC")
	} else {
		rowBuilder = rowBuilder.OrderBy(orderBy)
	}

	query, values, err := rowBuilder.Where("del_state = ?", globalkey.DelStateNo).ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*HomestayActivity
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *defaultHomestayActivityModel) FindPageListByPage(ctx context.Context, rowBuilder squirrel.SelectBuilder, page, pageSize int64, orderBy string) ([]*HomestayActivity, error) {

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

	var resp []*HomestayActivity
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *defaultHomestayActivityModel) FindPageListByIdDESC(ctx context.Context, rowBuilder squirrel.SelectBuilder, preMinId, pageSize int64) ([]*HomestayActivity, error) {

	if preMinId > 0 {
		rowBuilder = rowBuilder.Where(" id < ? ", preMinId)
	}

	query, values, err := rowBuilder.Where("del_state = ?", globalkey.DelStateNo).OrderBy("id DESC").Limit(uint64(pageSize)).ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*HomestayActivity
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

//按照id升序分页查询数据，不支持排序
func (m *defaultHomestayActivityModel) FindPageListByIdASC(ctx context.Context, rowBuilder squirrel.SelectBuilder, preMaxId, pageSize int64) ([]*HomestayActivity, error) {

	if preMaxId > 0 {
		rowBuilder = rowBuilder.Where(" id > ? ", preMaxId)
	}

	query, values, err := rowBuilder.Where("del_state = ?", globalkey.DelStateNo).OrderBy("id ASC").Limit(uint64(pageSize)).ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*HomestayActivity
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

// export logic
func (m *defaultHomestayActivityModel) Trans(ctx context.Context, fn func(ctx context.Context, session sqlx.Session) error) error {

	return m.TransactCtx(ctx, func(ctx context.Context, session sqlx.Session) error {
		return fn(ctx, session)
	})

}

// export logic
func (m *defaultHomestayActivityModel) RowBuilder() squirrel.SelectBuilder {
	return squirrel.Select(homestayActivityRows).From(m.table)
}

// export logic
func (m *defaultHomestayActivityModel) CountBuilder(field string) squirrel.SelectBuilder {
	return squirrel.Select("COUNT(" + field + ")").From(m.table)
}

// export logic
func (m *defaultHomestayActivityModel) SumBuilder(field string) squirrel.SelectBuilder {
	return squirrel.Select("IFNULL(SUM(" + field + "),0)").From(m.table)
}

func (m *defaultHomestayActivityModel) FindAllForPage(ctx context.Context, Info string, Current int64, PageSize int64) ([]*HomestayActivity, error) {
	looklookTravelHomestayActivityInfoKey := fmt.Sprintf("%s%v", cacheLooklookTravelHomestayActivityInfoPrefix, Info)

	var resp []*HomestayActivity
	query1 := fmt.Sprintf("select %s from %s ", homestayActivityRows, m.table)
	query2 := fmt.Sprintf("where id like '%s' ", Info) //like need ''   or and
	query3 := fmt.Sprintf("order by `id` ASC limit ?,?")

	err := m.QueryRowCtx(ctx, &resp, looklookTravelHomestayActivityInfoKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
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

func (m *defaultHomestayActivityModel) Count(ctx context.Context) (int64, error) {
	looklookTravelHomestayActivityCountKey := fmt.Sprintf("%s%v", cacheLooklookTravelHomestayActivityCountPrefix, "no")
	var count int64

	//err := m.QueryRowCtx(ctx, &resp, goGatewayGatewayServiceInfoCountKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
	err := m.QueryRowNoCacheCtx(ctx, &count, looklookTravelHomestayActivityCountKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
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
