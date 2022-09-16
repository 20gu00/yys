// Code generated by goctl. DO NOT EDIT!

package model

import (
	"context"
	"database/sql"
	"fmt"
	"go-travel/common/globalkey"
	"go-travel/common/xerr"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	thirdPaymentFieldNames          = builder.RawFieldNames(&ThirdPayment{})
	thirdPaymentRows                = strings.Join(thirdPaymentFieldNames, ",")
	thirdPaymentRowsExpectAutoSet   = strings.Join(stringx.Remove(thirdPaymentFieldNames, "`id`", "`create_time`", "`update_time`", "`create_at`", "`update_at`"), ",")
	thirdPaymentRowsWithPlaceHolder = strings.Join(stringx.Remove(thirdPaymentFieldNames, "`id`", "`create_time`", "`update_time`", "`create_at`", "`update_at`"), "=?,") + "=?"

	cacheLooklookPaymentThirdPaymentIdPrefix = "cache:looklookPayment:thirdPayment:id:"
	cacheLooklookPaymentThirdPaymentSnPrefix = "cache:looklookPayment:thirdPayment:sn:"
	cacheLooklookUsercenterInfoPrefix        = "cache:looklookPayment:thirdPayment:info:"
	cacheLooklookUsercenterCountPrefix       = "cache:looklookPayment:thirdPayment:count:"
)

type (
	thirdPaymentModel interface {
		Insert(ctx context.Context, session sqlx.Session, data *ThirdPayment) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*ThirdPayment, error)
		FindOneBySn(ctx context.Context, sn string) (*ThirdPayment, error)
		Update(ctx context.Context, session sqlx.Session, data *ThirdPayment) (sql.Result, error)
		UpdateWithVersion(ctx context.Context, session sqlx.Session, data *ThirdPayment) error
		Delete(ctx context.Context, session sqlx.Session, id int64) error
	}

	defaultThirdPaymentModel struct {
		sqlc.CachedConn
		table string
	}

	ThirdPayment struct {
		Id             int64     `db:"id"`
		Sn             string    `db:"sn"` // 流水单号
		CreateTime     time.Time `db:"create_time"`
		UpdateTime     time.Time `db:"update_time"`
		DeleteTime     time.Time `db:"delete_time"`
		DelState       int64     `db:"del_state"`
		Version        int64     `db:"version"`          // 乐观锁版本号
		UserId         int64     `db:"user_id"`          // 用户id
		PayMode        string    `db:"pay_mode"`         // 支付方式 1:微信支付
		TradeType      string    `db:"trade_type"`       // 第三方支付类型
		TradeState     string    `db:"trade_state"`      // 第三方交易状态
		PayTotal       int64     `db:"pay_total"`        // 支付总金额(分)
		TransactionId  string    `db:"transaction_id"`   // 第三方支付单号
		TradeStateDesc string    `db:"trade_state_desc"` // 支付状态描述
		OrderSn        string    `db:"order_sn"`         // 业务单号
		ServiceType    string    `db:"service_type"`     // 业务类型
		PayStatus      int64     `db:"pay_status"`       // 平台内交易状态   -1:支付失败 0:未支付 1:支付成功 2:已退款
		PayTime        time.Time `db:"pay_time"`         // 支付成功时间
	}
)

func newThirdPaymentModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultThirdPaymentModel {
	return &defaultThirdPaymentModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`third_payment`",
	}
}

func (m *defaultThirdPaymentModel) Insert(ctx context.Context, session sqlx.Session, data *ThirdPayment) (sql.Result, error) {
	data.DeleteTime = time.Unix(0, 0)
	looklookPaymentThirdPaymentIdKey := fmt.Sprintf("%s%v", cacheLooklookPaymentThirdPaymentIdPrefix, data.Id)
	looklookPaymentThirdPaymentSnKey := fmt.Sprintf("%s%v", cacheLooklookPaymentThirdPaymentSnPrefix, data.Sn)
	return m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, thirdPaymentRowsExpectAutoSet)
		if session != nil {
			return session.ExecCtx(ctx, query, data.Sn, data.DeleteTime, data.DelState, data.Version, data.UserId, data.PayMode, data.TradeType, data.TradeState, data.PayTotal, data.TransactionId, data.TradeStateDesc, data.OrderSn, data.ServiceType, data.PayStatus, data.PayTime)
		}
		return conn.ExecCtx(ctx, query, data.Sn, data.DeleteTime, data.DelState, data.Version, data.UserId, data.PayMode, data.TradeType, data.TradeState, data.PayTotal, data.TransactionId, data.TradeStateDesc, data.OrderSn, data.ServiceType, data.PayStatus, data.PayTime)
	}, looklookPaymentThirdPaymentIdKey, looklookPaymentThirdPaymentSnKey)
}

func (m *defaultThirdPaymentModel) FindOne(ctx context.Context, id int64) (*ThirdPayment, error) {
	looklookPaymentThirdPaymentIdKey := fmt.Sprintf("%s%v", cacheLooklookPaymentThirdPaymentIdPrefix, id)
	var resp ThirdPayment
	err := m.QueryRowCtx(ctx, &resp, looklookPaymentThirdPaymentIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? and del_state = ? limit 1", thirdPaymentRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, id, globalkey.DelStateNo)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultThirdPaymentModel) FindOneBySn(ctx context.Context, sn string) (*ThirdPayment, error) {
	looklookPaymentThirdPaymentSnKey := fmt.Sprintf("%s%v", cacheLooklookPaymentThirdPaymentSnPrefix, sn)
	var resp ThirdPayment
	err := m.QueryRowIndexCtx(ctx, &resp, looklookPaymentThirdPaymentSnKey, m.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) (i interface{}, e error) {
		query := fmt.Sprintf("select %s from %s where `sn` = ? and del_state = ? limit 1", thirdPaymentRows, m.table)
		if err := conn.QueryRowCtx(ctx, &resp, query, sn, globalkey.DelStateNo); err != nil {
			return nil, err
		}
		return resp.Id, nil
	}, m.queryPrimary)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultThirdPaymentModel) Update(ctx context.Context, session sqlx.Session, data *ThirdPayment) (sql.Result, error) {
	looklookPaymentThirdPaymentIdKey := fmt.Sprintf("%s%v", cacheLooklookPaymentThirdPaymentIdPrefix, data.Id)
	looklookPaymentThirdPaymentSnKey := fmt.Sprintf("%s%v", cacheLooklookPaymentThirdPaymentSnPrefix, data.Sn)
	return m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, thirdPaymentRowsWithPlaceHolder)
		if session != nil {
			return session.ExecCtx(ctx, query, data.Sn, data.DeleteTime, data.DelState, data.Version, data.UserId, data.PayMode, data.TradeType, data.TradeState, data.PayTotal, data.TransactionId, data.TradeStateDesc, data.OrderSn, data.ServiceType, data.PayStatus, data.PayTime, data.Id)
		}
		return conn.ExecCtx(ctx, query, data.Sn, data.DeleteTime, data.DelState, data.Version, data.UserId, data.PayMode, data.TradeType, data.TradeState, data.PayTotal, data.TransactionId, data.TradeStateDesc, data.OrderSn, data.ServiceType, data.PayStatus, data.PayTime, data.Id)
	}, looklookPaymentThirdPaymentSnKey, looklookPaymentThirdPaymentIdKey)
}

func (m *defaultThirdPaymentModel) UpdateWithVersion(ctx context.Context, session sqlx.Session, data *ThirdPayment) error {

	oldVersion := data.Version
	data.Version += 1

	var sqlResult sql.Result
	var err error

	looklookPaymentThirdPaymentIdKey := fmt.Sprintf("%s%v", cacheLooklookPaymentThirdPaymentIdPrefix, data.Id)
	looklookPaymentThirdPaymentSnKey := fmt.Sprintf("%s%v", cacheLooklookPaymentThirdPaymentSnPrefix, data.Sn)
	sqlResult, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ? and version = ? ", m.table, thirdPaymentRowsWithPlaceHolder)
		if session != nil {
			return session.ExecCtx(ctx, query, data.Sn, data.DeleteTime, data.DelState, data.Version, data.UserId, data.PayMode, data.TradeType, data.TradeState, data.PayTotal, data.TransactionId, data.TradeStateDesc, data.OrderSn, data.ServiceType, data.PayStatus, data.PayTime, data.Id, oldVersion)
		}
		return conn.ExecCtx(ctx, query, data.Sn, data.DeleteTime, data.DelState, data.Version, data.UserId, data.PayMode, data.TradeType, data.TradeState, data.PayTotal, data.TransactionId, data.TradeStateDesc, data.OrderSn, data.ServiceType, data.PayStatus, data.PayTime, data.Id, oldVersion)
	}, looklookPaymentThirdPaymentSnKey, looklookPaymentThirdPaymentIdKey)
	if err != nil {
		return err
	}
	updateCount, err := sqlResult.RowsAffected()
	if err != nil {
		return err
	}
	if updateCount == 0 {
		return xerr.NewErrCode(xerr.DB_UPDATE_AFFECTED_ZERO_ERROR)
	}

	return nil
}

func (m *defaultThirdPaymentModel) Delete(ctx context.Context, session sqlx.Session, id int64) error {
	data, err := m.FindOne(ctx, id)
	if err != nil {
		return err
	}

	looklookPaymentThirdPaymentIdKey := fmt.Sprintf("%s%v", cacheLooklookPaymentThirdPaymentIdPrefix, id)
	looklookPaymentThirdPaymentSnKey := fmt.Sprintf("%s%v", cacheLooklookPaymentThirdPaymentSnPrefix, data.Sn)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		if session != nil {
			return session.ExecCtx(ctx, query, id)
		}
		return conn.ExecCtx(ctx, query, id)
	}, looklookPaymentThirdPaymentIdKey, looklookPaymentThirdPaymentSnKey)
	return err
}

func (m *defaultThirdPaymentModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheLooklookPaymentThirdPaymentIdPrefix, primary)
}
func (m *defaultThirdPaymentModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? and del_state = ? limit 1", thirdPaymentRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary, globalkey.DelStateNo)
}

func (m *defaultThirdPaymentModel) tableName() string {
	return m.table
}
