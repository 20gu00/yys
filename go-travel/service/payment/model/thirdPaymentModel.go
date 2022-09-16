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

var _ ThirdPaymentModel = (*customThirdPaymentModel)(nil)

type (
	// ThirdPaymentModel is an interface to be customized, add more methods here,
	// and implement the added methods in customThirdPaymentModel.
	ThirdPaymentModel interface {
		thirdPaymentModel
		Trans(ctx context.Context, fn func(context context.Context, session sqlx.Session) error) error
		RowBuilder() squirrel.SelectBuilder
		CountBuilder(field string) squirrel.SelectBuilder
		SumBuilder(field string) squirrel.SelectBuilder
		DeleteSoft(ctx context.Context, session sqlx.Session, data *ThirdPayment) error
		FindOneByQuery(ctx context.Context, rowBuilder squirrel.SelectBuilder) (*ThirdPayment, error)
		FindSum(ctx context.Context, sumBuilder squirrel.SelectBuilder) (float64, error)
		FindCount(ctx context.Context, countBuilder squirrel.SelectBuilder) (int64, error)
		FindAll(ctx context.Context, rowBuilder squirrel.SelectBuilder, orderBy string) ([]*ThirdPayment, error)
		FindPageListByPage(ctx context.Context, rowBuilder squirrel.SelectBuilder, page, pageSize int64, orderBy string) ([]*ThirdPayment, error)
		FindPageListByIdDESC(ctx context.Context, rowBuilder squirrel.SelectBuilder, preMinId, pageSize int64) ([]*ThirdPayment, error)
		FindPageListByIdASC(ctx context.Context, rowBuilder squirrel.SelectBuilder, preMaxId, pageSize int64) ([]*ThirdPayment, error)
		//test
		FindAllForPage(ctx context.Context, Info string, Current int64, PageSize int64) ([]*ThirdPayment, error)
		Count(ctx context.Context) (int64, error)
	}

	customThirdPaymentModel struct {
		*defaultThirdPaymentModel
	}
)

// NewThirdPaymentModel returns a model for the database table.
func NewThirdPaymentModel(conn sqlx.SqlConn, c cache.CacheConf) ThirdPaymentModel {
	return &customThirdPaymentModel{
		defaultThirdPaymentModel: newThirdPaymentModel(conn, c),
	}
}

func (m *defaultThirdPaymentModel) DeleteSoft(ctx context.Context, session sqlx.Session, data *ThirdPayment) error {
	data.DelState = globalkey.DelStateYes
	data.DeleteTime = time.Now()
	if err := m.UpdateWithVersion(ctx, session, data); err != nil {
		return errors.Wrapf(xerr.NewErrMsg("删除数据失败"), "ThirdPaymentModel delete err : %+v", err)
	}
	return nil
}

func (m *defaultThirdPaymentModel) FindOneByQuery(ctx context.Context, rowBuilder squirrel.SelectBuilder) (*ThirdPayment, error) {

	query, values, err := rowBuilder.Where("del_state = ?", globalkey.DelStateNo).ToSql()
	if err != nil {
		return nil, err
	}

	var resp ThirdPayment
	err = m.QueryRowNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return &resp, nil
	default:
		return nil, err
	}
}

func (m *defaultThirdPaymentModel) FindSum(ctx context.Context, sumBuilder squirrel.SelectBuilder) (float64, error) {

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

func (m *defaultThirdPaymentModel) FindCount(ctx context.Context, countBuilder squirrel.SelectBuilder) (int64, error) {

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

func (m *defaultThirdPaymentModel) FindAll(ctx context.Context, rowBuilder squirrel.SelectBuilder, orderBy string) ([]*ThirdPayment, error) {

	if orderBy == "" {
		rowBuilder = rowBuilder.OrderBy("id DESC")
	} else {
		rowBuilder = rowBuilder.OrderBy(orderBy)
	}

	query, values, err := rowBuilder.Where("del_state = ?", globalkey.DelStateNo).ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*ThirdPayment
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *defaultThirdPaymentModel) FindPageListByPage(ctx context.Context, rowBuilder squirrel.SelectBuilder, page, pageSize int64, orderBy string) ([]*ThirdPayment, error) {

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

	var resp []*ThirdPayment
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *defaultThirdPaymentModel) FindPageListByIdDESC(ctx context.Context, rowBuilder squirrel.SelectBuilder, preMinId, pageSize int64) ([]*ThirdPayment, error) {

	if preMinId > 0 {
		rowBuilder = rowBuilder.Where(" id < ? ", preMinId)
	}

	query, values, err := rowBuilder.Where("del_state = ?", globalkey.DelStateNo).OrderBy("id DESC").Limit(uint64(pageSize)).ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*ThirdPayment
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

//按照id升序分页查询数据，不支持排序
func (m *defaultThirdPaymentModel) FindPageListByIdASC(ctx context.Context, rowBuilder squirrel.SelectBuilder, preMaxId, pageSize int64) ([]*ThirdPayment, error) {

	if preMaxId > 0 {
		rowBuilder = rowBuilder.Where(" id > ? ", preMaxId)
	}

	query, values, err := rowBuilder.Where("del_state = ?", globalkey.DelStateNo).OrderBy("id ASC").Limit(uint64(pageSize)).ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*ThirdPayment
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

// export logic
func (m *defaultThirdPaymentModel) Trans(ctx context.Context, fn func(ctx context.Context, session sqlx.Session) error) error {

	return m.TransactCtx(ctx, func(ctx context.Context, session sqlx.Session) error {
		return fn(ctx, session)
	})

}

// export logic
func (m *defaultThirdPaymentModel) RowBuilder() squirrel.SelectBuilder {
	return squirrel.Select(thirdPaymentRows).From(m.table)
}

// export logic
func (m *defaultThirdPaymentModel) CountBuilder(field string) squirrel.SelectBuilder {
	return squirrel.Select("COUNT(" + field + ")").From(m.table)
}

// export logic
func (m *defaultThirdPaymentModel) SumBuilder(field string) squirrel.SelectBuilder {
	return squirrel.Select("IFNULL(SUM(" + field + "),0)").From(m.table)
}

//test
func (m *defaultThirdPaymentModel) FindAllForPage(ctx context.Context, Info string, Current int64, PageSize int64) ([]*ThirdPayment, error) { //[]*
	looklookUsercenterInfoKey := fmt.Sprintf("%s%v:%v:%v", cacheLooklookUsercenterInfoPrefix, Info, Current, PageSize)

	var resp []*ThirdPayment
	query1 := fmt.Sprintf("select %s from %s ", thirdPaymentRows, m.table)
	query2 := fmt.Sprintf("where id like '%s' ", Info) //like need ''
	query3 := fmt.Sprintf("order by `id` ASC limit ?,?")
	//fmt.Sprintf

	err := m.QueryRowCtx(ctx, &resp, looklookUsercenterInfoKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
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

func (m *defaultThirdPaymentModel) Count(ctx context.Context) (int64, error) {
	looklookUsercenterCountKey := fmt.Sprintf("%s%v", cacheLooklookUsercenterCountPrefix, "no")
	var count int64

	//err := m.QueryRowCtx(ctx, &resp, goGatewayGatewayServiceInfoCountKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
	err := m.QueryRowNoCacheCtx(ctx, &count, looklookUsercenterCountKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
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
