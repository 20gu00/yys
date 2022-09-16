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

var _ HomestayBusinessModel = (*customHomestayBusinessModel)(nil)

type (
	// HomestayBusinessModel is an interface to be customized, add more methods here,
	// and implement the added methods in customHomestayBusinessModel.
	HomestayBusinessModel interface {
		homestayBusinessModel
		Trans(ctx context.Context, fn func(context context.Context, session sqlx.Session) error) error
		RowBuilder() squirrel.SelectBuilder
		CountBuilder(field string) squirrel.SelectBuilder
		SumBuilder(field string) squirrel.SelectBuilder
		DeleteSoft(ctx context.Context, session sqlx.Session, data *HomestayBusiness) error
		FindOneByQuery(ctx context.Context, rowBuilder squirrel.SelectBuilder) (*HomestayBusiness, error)
		FindSum(ctx context.Context, sumBuilder squirrel.SelectBuilder) (float64, error)
		FindCount(ctx context.Context, countBuilder squirrel.SelectBuilder) (int64, error)
		FindAll(ctx context.Context, rowBuilder squirrel.SelectBuilder, orderBy string) ([]*HomestayBusiness, error)
		FindPageListByPage(ctx context.Context, rowBuilder squirrel.SelectBuilder, page, pageSize int64, orderBy string) ([]*HomestayBusiness, error)
		FindPageListByIdDESC(ctx context.Context, rowBuilder squirrel.SelectBuilder, preMinId, pageSize int64) ([]*HomestayBusiness, error)
		FindPageListByIdASC(ctx context.Context, rowBuilder squirrel.SelectBuilder, preMaxId, pageSize int64) ([]*HomestayBusiness, error)

		//provide to admin use
		InsertById(ctx context.Context, data *HomestayBusiness) (sql.Result, error)
		FindAllForPage(ctx context.Context, Info string, Current int64, PageSize int64) ([]*HomestayBusiness, error)
		Count(ctx context.Context) (int64, error)
		UpdateNoSess(ctx context.Context, data *HomestayBusiness) (sql.Result, error)
	}

	customHomestayBusinessModel struct {
		*defaultHomestayBusinessModel
	}
)

// NewHomestayBusinessModel returns a model for the database table.
func NewHomestayBusinessModel(conn sqlx.SqlConn, c cache.CacheConf) HomestayBusinessModel {
	return &customHomestayBusinessModel{
		defaultHomestayBusinessModel: newHomestayBusinessModel(conn, c),
	}
}

func (m *defaultHomestayBusinessModel) DeleteSoft(ctx context.Context, session sqlx.Session, data *HomestayBusiness) error {
	data.DelState = globalkey.DelStateYes
	data.DeleteTime = time.Now()
	if err := m.UpdateWithVersion(ctx, session, data); err != nil {
		return errors.Wrapf(xerr.NewErrMsg("删除数据失败"), "HomestayBusinessModel delete err : %+v", err)
	}
	return nil
}

func (m *defaultHomestayBusinessModel) FindOneByQuery(ctx context.Context, rowBuilder squirrel.SelectBuilder) (*HomestayBusiness, error) {

	query, values, err := rowBuilder.Where("del_state = ?", globalkey.DelStateNo).ToSql()
	if err != nil {
		return nil, err
	}

	var resp HomestayBusiness
	err = m.QueryRowNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return &resp, nil
	default:
		return nil, err
	}
}

func (m *defaultHomestayBusinessModel) FindSum(ctx context.Context, sumBuilder squirrel.SelectBuilder) (float64, error) {

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

func (m *defaultHomestayBusinessModel) FindCount(ctx context.Context, countBuilder squirrel.SelectBuilder) (int64, error) {

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

func (m *defaultHomestayBusinessModel) FindAll(ctx context.Context, rowBuilder squirrel.SelectBuilder, orderBy string) ([]*HomestayBusiness, error) {

	if orderBy == "" {
		rowBuilder = rowBuilder.OrderBy("id DESC")
	} else {
		rowBuilder = rowBuilder.OrderBy(orderBy)
	}

	query, values, err := rowBuilder.Where("del_state = ?", globalkey.DelStateNo).ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*HomestayBusiness
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *defaultHomestayBusinessModel) FindPageListByPage(ctx context.Context, rowBuilder squirrel.SelectBuilder, page, pageSize int64, orderBy string) ([]*HomestayBusiness, error) {

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

	var resp []*HomestayBusiness
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *defaultHomestayBusinessModel) FindPageListByIdDESC(ctx context.Context, rowBuilder squirrel.SelectBuilder, preMinId, pageSize int64) ([]*HomestayBusiness, error) {

	if preMinId > 0 {
		rowBuilder = rowBuilder.Where(" id < ? ", preMinId)
	}

	query, values, err := rowBuilder.Where("del_state = ?", globalkey.DelStateNo).OrderBy("id DESC").Limit(uint64(pageSize)).ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*HomestayBusiness
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

//按照id升序分页查询数据，不支持排序
func (m *defaultHomestayBusinessModel) FindPageListByIdASC(ctx context.Context, rowBuilder squirrel.SelectBuilder, preMaxId, pageSize int64) ([]*HomestayBusiness, error) {

	if preMaxId > 0 {
		rowBuilder = rowBuilder.Where(" id > ? ", preMaxId)
	}

	query, values, err := rowBuilder.Where("del_state = ?", globalkey.DelStateNo).OrderBy("id ASC").Limit(uint64(pageSize)).ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*HomestayBusiness
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

// export logic
func (m *defaultHomestayBusinessModel) Trans(ctx context.Context, fn func(ctx context.Context, session sqlx.Session) error) error {

	return m.TransactCtx(ctx, func(ctx context.Context, session sqlx.Session) error {
		return fn(ctx, session)
	})

}

// export logic
func (m *defaultHomestayBusinessModel) RowBuilder() squirrel.SelectBuilder {
	return squirrel.Select(homestayBusinessRows).From(m.table)
}

// export logic
func (m *defaultHomestayBusinessModel) CountBuilder(field string) squirrel.SelectBuilder {
	return squirrel.Select("COUNT(" + field + ")").From(m.table)
}

// export logic
func (m *defaultHomestayBusinessModel) SumBuilder(field string) squirrel.SelectBuilder {
	return squirrel.Select("IFNULL(SUM(" + field + "),0)").From(m.table)
}

func (m *defaultHomestayBusinessModel) InsertById(ctx context.Context, data *HomestayBusiness) (sql.Result, error) {
	data.DeleteTime = time.Unix(0, 0)
	//many key
	looklookTravelHomestayBusinessIdKey := fmt.Sprintf("%s%v", cacheLooklookTravelHomestayBusinessIdPrefix, data.Id)
	return m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, homestayBusinessRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.DeleteTime, data.DelState, data.Title, data.UserId, data.Info, data.BossInfo, data.LicenseFron, data.LicenseBack, data.RowState, data.Star, data.Tags, data.Cover, data.HeaderImg, data.Version)
	}, looklookTravelHomestayBusinessIdKey)
}

func (m *defaultHomestayBusinessModel) FindAllForPage(ctx context.Context, Info string, Current int64, PageSize int64) ([]*HomestayBusiness, error) {
	looklookTravelHomestayBusinessInfoKey := fmt.Sprintf("%s%v", cacheLooklookTravelHomestayBusinessInfoPrefix, Info)

	var resp []*HomestayBusiness
	query1 := fmt.Sprintf("select %s from %s ", homestayBusinessRows, m.table)
	query2 := fmt.Sprintf("where id like '%s' ", Info) //like need ''   or and
	query3 := fmt.Sprintf("order by `id` ASC limit ?,?")

	err := m.QueryRowCtx(ctx, &resp, looklookTravelHomestayBusinessInfoKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
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

func (m *defaultHomestayBusinessModel) Count(ctx context.Context) (int64, error) {
	looklookTravelHomestayBusinessCountKey := fmt.Sprintf("%s%v", cacheLooklookTravelHomestayBusinessCountPrefix, "no")
	var count int64

	//err := m.QueryRowCtx(ctx, &resp, goGatewayGatewayServiceInfoCountKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
	err := m.QueryRowNoCacheCtx(ctx, &count, looklookTravelHomestayBusinessCountKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
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

func (m *defaultHomestayBusinessModel) UpdateNoSess(ctx context.Context, data *HomestayBusiness) (sql.Result, error) {
	data.DeleteTime = time.Unix(0, 0)
	looklookTravelHomestayBusinessIdKey := fmt.Sprintf("%s%v", cacheLooklookTravelHomestayBusinessIdPrefix, data.Id)
	return m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, homestayBusinessRowsWithPlaceHolder)

		return conn.ExecCtx(ctx, query, data.DeleteTime, data.DelState, data.Title, data.UserId, data.Info, data.BossInfo, data.LicenseFron, data.LicenseBack, data.RowState, data.Star, data.Tags, data.Cover, data.HeaderImg, data.Version, data.Id)
	}, looklookTravelHomestayBusinessIdKey)
}
