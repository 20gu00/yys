package model

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlc"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ AdminModel = (*customAdminModel)(nil)

type (
	// AdminModel is an interface to be customized, add more methods here,
	// and implement the added methods in customAdminModel.
	AdminModel interface {
		adminModel
		FindOneByName(ctx context.Context, username string) (*Admin, error)
		//FindOneByUserId(ctx context.Context, userid int64) (*Admin, error)
		FindOneById(ctx context.Context, id int64) (*Admin, error)
	}

	customAdminModel struct {
		*defaultAdminModel
	}
)

// NewAdminModel returns a model for the database table.
func NewAdminModel(conn sqlx.SqlConn) AdminModel {
	return &customAdminModel{
		defaultAdminModel: newAdminModel(conn),
	}
}
func (m *defaultAdminModel) FindOneById(ctx context.Context, id int64) (*Admin, error) {
	var resp Admin
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", adminRows, m.table)
	err := m.conn.QueryRowCtx(ctx, &resp, query, id)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}
func (m *defaultAdminModel) FindOneByName(ctx context.Context, username string) (*Admin, error) {

	query1 := fmt.Sprintf("select %s from %s ", adminRows, m.table)
	query2 := fmt.Sprintf("where `user_name` = '%s' limit 1", username)
	query := query1 + query2
	//query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", adminRows, m.table)
	var resp Admin
	err := m.conn.QueryRowCtx(ctx, &resp, query)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}
