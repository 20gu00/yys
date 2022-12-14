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
	homestayOrderFieldNames          = builder.RawFieldNames(&HomestayOrder{})
	homestayOrderRows                = strings.Join(homestayOrderFieldNames, ",")
	homestayOrderRowsExpectAutoSet   = strings.Join(stringx.Remove(homestayOrderFieldNames, "`id`", "`create_time`", "`update_time`", "`create_at`", "`update_at`"), ",")
	homestayOrderRowsWithPlaceHolder = strings.Join(stringx.Remove(homestayOrderFieldNames, "`id`", "`create_time`", "`update_time`", "`create_at`", "`update_at`"), "=?,") + "=?"

	cacheLooklookOrderHomestayOrderIdPrefix    = "cache:looklookOrder:homestayOrder:id:"
	cacheLooklookOrderHomestayOrderSnPrefix    = "cache:looklookOrder:homestayOrder:sn:"
	cacheLooklookOrderHomestayOrderInfoPrefix  = "cache:looklookOrder:homestayOrder:info:"
	cacheLooklookOrderHomestayOrderCountPrefix = "cache:looklookOrder:homestayOrder:count:"
)

type (
	homestayOrderModel interface {
		Insert(ctx context.Context, session sqlx.Session, data *HomestayOrder) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*HomestayOrder, error)
		FindOneBySn(ctx context.Context, sn string) (*HomestayOrder, error)
		Update(ctx context.Context, session sqlx.Session, data *HomestayOrder) (sql.Result, error)
		UpdateWithVersion(ctx context.Context, session sqlx.Session, data *HomestayOrder) error
		Delete(ctx context.Context, session sqlx.Session, id int64) error
	}

	defaultHomestayOrderModel struct {
		sqlc.CachedConn
		table string
	}

	HomestayOrder struct {
		Id                  int64     `db:"id"`
		CreateTime          time.Time `db:"create_time"`
		UpdateTime          time.Time `db:"update_time"`
		DeleteTime          time.Time `db:"delete_time"`
		DelState            int64     `db:"del_state"`
		Version             int64     `db:"version"`               // ?????????
		Sn                  string    `db:"sn"`                    // ?????????
		UserId              int64     `db:"user_id"`               // ????????????id
		HomestayId          int64     `db:"homestay_id"`           // ??????id
		Title               string    `db:"title"`                 // ??????
		SubTitle            string    `db:"sub_title"`             // ?????????
		Cover               string    `db:"cover"`                 // ??????
		Info                string    `db:"info"`                  // ??????
		PeopleNum           int64     `db:"people_num"`            // ??????????????????
		RowType             int64     `db:"row_type"`              // ????????????0?????????????????? 1:???????????????
		NeedFood            int64     `db:"need_food"`             // 0:??????????????? 1:????????????
		FoodInfo            string    `db:"food_info"`             // ????????????
		FoodPrice           int64     `db:"food_price"`            // ????????????(???)
		HomestayPrice       int64     `db:"homestay_price"`        // ????????????(???)
		MarketHomestayPrice int64     `db:"market_homestay_price"` // ??????????????????(???)
		HomestayBusinessId  int64     `db:"homestay_business_id"`  // ??????id
		HomestayUserId      int64     `db:"homestay_user_id"`      // ????????????id
		LiveStartDate       time.Time `db:"live_start_date"`       // ??????????????????
		LiveEndDate         time.Time `db:"live_end_date"`         // ??????????????????
		LivePeopleNum       int64     `db:"live_people_num"`       // ??????????????????
		TradeState          int64     `db:"trade_state"`           // -1: ????????? 0:????????? 1:????????? 2:?????????  3:????????? 4:?????????
		TradeCode           string    `db:"trade_code"`            // ?????????
		Remark              string    `db:"remark"`                // ??????????????????
		OrderTotalPrice     int64     `db:"order_total_price"`     // ?????????????????????????????????+??????????????????(???)
		FoodTotalPrice      int64     `db:"food_total_price"`      // ???????????????(???)
		HomestayTotalPrice  int64     `db:"homestay_total_price"`  // ???????????????(???)
	}
)

func newHomestayOrderModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultHomestayOrderModel {
	return &defaultHomestayOrderModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`homestay_order`",
	}
}

func (m *defaultHomestayOrderModel) Insert(ctx context.Context, session sqlx.Session, data *HomestayOrder) (sql.Result, error) {
	data.DeleteTime = time.Unix(0, 0)
	looklookOrderHomestayOrderIdKey := fmt.Sprintf("%s%v", cacheLooklookOrderHomestayOrderIdPrefix, data.Id)
	looklookOrderHomestayOrderSnKey := fmt.Sprintf("%s%v", cacheLooklookOrderHomestayOrderSnPrefix, data.Sn)
	return m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, homestayOrderRowsExpectAutoSet)
		if session != nil {
			return session.ExecCtx(ctx, query, data.DeleteTime, data.DelState, data.Version, data.Sn, data.UserId, data.HomestayId, data.Title, data.SubTitle, data.Cover, data.Info, data.PeopleNum, data.RowType, data.NeedFood, data.FoodInfo, data.FoodPrice, data.HomestayPrice, data.MarketHomestayPrice, data.HomestayBusinessId, data.HomestayUserId, data.LiveStartDate, data.LiveEndDate, data.LivePeopleNum, data.TradeState, data.TradeCode, data.Remark, data.OrderTotalPrice, data.FoodTotalPrice, data.HomestayTotalPrice)
		}
		return conn.ExecCtx(ctx, query, data.DeleteTime, data.DelState, data.Version, data.Sn, data.UserId, data.HomestayId, data.Title, data.SubTitle, data.Cover, data.Info, data.PeopleNum, data.RowType, data.NeedFood, data.FoodInfo, data.FoodPrice, data.HomestayPrice, data.MarketHomestayPrice, data.HomestayBusinessId, data.HomestayUserId, data.LiveStartDate, data.LiveEndDate, data.LivePeopleNum, data.TradeState, data.TradeCode, data.Remark, data.OrderTotalPrice, data.FoodTotalPrice, data.HomestayTotalPrice)
	}, looklookOrderHomestayOrderIdKey, looklookOrderHomestayOrderSnKey)
}

func (m *defaultHomestayOrderModel) FindOne(ctx context.Context, id int64) (*HomestayOrder, error) {
	looklookOrderHomestayOrderIdKey := fmt.Sprintf("%s%v", cacheLooklookOrderHomestayOrderIdPrefix, id)
	var resp HomestayOrder
	err := m.QueryRowCtx(ctx, &resp, looklookOrderHomestayOrderIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? and del_state = ? limit 1", homestayOrderRows, m.table)
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

func (m *defaultHomestayOrderModel) FindOneBySn(ctx context.Context, sn string) (*HomestayOrder, error) {
	looklookOrderHomestayOrderSnKey := fmt.Sprintf("%s%v", cacheLooklookOrderHomestayOrderSnPrefix, sn)
	var resp HomestayOrder
	err := m.QueryRowIndexCtx(ctx, &resp, looklookOrderHomestayOrderSnKey, m.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) (i interface{}, e error) {
		query := fmt.Sprintf("select %s from %s where `sn` = ? and del_state = ? limit 1", homestayOrderRows, m.table)
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

func (m *defaultHomestayOrderModel) Update(ctx context.Context, session sqlx.Session, data *HomestayOrder) (sql.Result, error) {
	looklookOrderHomestayOrderIdKey := fmt.Sprintf("%s%v", cacheLooklookOrderHomestayOrderIdPrefix, data.Id)
	looklookOrderHomestayOrderSnKey := fmt.Sprintf("%s%v", cacheLooklookOrderHomestayOrderSnPrefix, data.Sn)
	return m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, homestayOrderRowsWithPlaceHolder)
		if session != nil {
			return session.ExecCtx(ctx, query, data.DeleteTime, data.DelState, data.Version, data.Sn, data.UserId, data.HomestayId, data.Title, data.SubTitle, data.Cover, data.Info, data.PeopleNum, data.RowType, data.NeedFood, data.FoodInfo, data.FoodPrice, data.HomestayPrice, data.MarketHomestayPrice, data.HomestayBusinessId, data.HomestayUserId, data.LiveStartDate, data.LiveEndDate, data.LivePeopleNum, data.TradeState, data.TradeCode, data.Remark, data.OrderTotalPrice, data.FoodTotalPrice, data.HomestayTotalPrice, data.Id)
		}
		return conn.ExecCtx(ctx, query, data.DeleteTime, data.DelState, data.Version, data.Sn, data.UserId, data.HomestayId, data.Title, data.SubTitle, data.Cover, data.Info, data.PeopleNum, data.RowType, data.NeedFood, data.FoodInfo, data.FoodPrice, data.HomestayPrice, data.MarketHomestayPrice, data.HomestayBusinessId, data.HomestayUserId, data.LiveStartDate, data.LiveEndDate, data.LivePeopleNum, data.TradeState, data.TradeCode, data.Remark, data.OrderTotalPrice, data.FoodTotalPrice, data.HomestayTotalPrice, data.Id)
	}, looklookOrderHomestayOrderIdKey, looklookOrderHomestayOrderSnKey)
}

func (m *defaultHomestayOrderModel) UpdateWithVersion(ctx context.Context, session sqlx.Session, data *HomestayOrder) error {

	oldVersion := data.Version
	data.Version += 1

	var sqlResult sql.Result
	var err error

	looklookOrderHomestayOrderIdKey := fmt.Sprintf("%s%v", cacheLooklookOrderHomestayOrderIdPrefix, data.Id)
	looklookOrderHomestayOrderSnKey := fmt.Sprintf("%s%v", cacheLooklookOrderHomestayOrderSnPrefix, data.Sn)
	sqlResult, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ? and version = ? ", m.table, homestayOrderRowsWithPlaceHolder)
		if session != nil {
			return session.ExecCtx(ctx, query, data.DeleteTime, data.DelState, data.Version, data.Sn, data.UserId, data.HomestayId, data.Title, data.SubTitle, data.Cover, data.Info, data.PeopleNum, data.RowType, data.NeedFood, data.FoodInfo, data.FoodPrice, data.HomestayPrice, data.MarketHomestayPrice, data.HomestayBusinessId, data.HomestayUserId, data.LiveStartDate, data.LiveEndDate, data.LivePeopleNum, data.TradeState, data.TradeCode, data.Remark, data.OrderTotalPrice, data.FoodTotalPrice, data.HomestayTotalPrice, data.Id, oldVersion)
		}
		return conn.ExecCtx(ctx, query, data.DeleteTime, data.DelState, data.Version, data.Sn, data.UserId, data.HomestayId, data.Title, data.SubTitle, data.Cover, data.Info, data.PeopleNum, data.RowType, data.NeedFood, data.FoodInfo, data.FoodPrice, data.HomestayPrice, data.MarketHomestayPrice, data.HomestayBusinessId, data.HomestayUserId, data.LiveStartDate, data.LiveEndDate, data.LivePeopleNum, data.TradeState, data.TradeCode, data.Remark, data.OrderTotalPrice, data.FoodTotalPrice, data.HomestayTotalPrice, data.Id, oldVersion)
	}, looklookOrderHomestayOrderIdKey, looklookOrderHomestayOrderSnKey)
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

func (m *defaultHomestayOrderModel) Delete(ctx context.Context, session sqlx.Session, id int64) error {
	data, err := m.FindOne(ctx, id)
	if err != nil {
		return err
	}

	looklookOrderHomestayOrderIdKey := fmt.Sprintf("%s%v", cacheLooklookOrderHomestayOrderIdPrefix, id)
	looklookOrderHomestayOrderSnKey := fmt.Sprintf("%s%v", cacheLooklookOrderHomestayOrderSnPrefix, data.Sn)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		if session != nil {
			return session.ExecCtx(ctx, query, id)
		}
		return conn.ExecCtx(ctx, query, id)
	}, looklookOrderHomestayOrderIdKey, looklookOrderHomestayOrderSnKey)
	return err
}

func (m *defaultHomestayOrderModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheLooklookOrderHomestayOrderIdPrefix, primary)
}
func (m *defaultHomestayOrderModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? and del_state = ? limit 1", homestayOrderRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary, globalkey.DelStateNo)
}

func (m *defaultHomestayOrderModel) tableName() string {
	return m.table
}
