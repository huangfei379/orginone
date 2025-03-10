// Code generated by goctl. DO NOT EDIT!

package entitys

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	flwRuBatchPartFieldNames          = builder.RawFieldNames(&FlwRuBatchPart{})
	flwRuBatchPartRows                = strings.Join(flwRuBatchPartFieldNames, ",")
	flwRuBatchPartRowsExpectAutoSet   = strings.Join(stringx.Remove(flwRuBatchPartFieldNames, "`create_time`", "`update_time`"), ",")
	flwRuBatchPartRowsWithPlaceHolder = strings.Join(stringx.Remove(flwRuBatchPartFieldNames, "`ID_`", "`create_time`", "`update_time`"), "=?,") + "=?"

	cacheAssetFlwRuBatchPartIDPrefix = "cache:asset:flwRuBatchPart:iD:"
)

type (
	flwRuBatchPartModel interface {
		Insert(ctx context.Context, data *FlwRuBatchPart) (sql.Result, error)
		FindOne(ctx context.Context, iD string) (*FlwRuBatchPart, error)
		Update(ctx context.Context, data *FlwRuBatchPart) error
		Delete(ctx context.Context, iD string) error
	}

	defaultFlwRuBatchPartModel struct {
		sqlc.CachedConn
		table string
	}

	FlwRuBatchPart struct {
		ID           string         `db:"ID_"`
		REV          sql.NullInt64  `db:"REV_"`
		BATCHID      sql.NullString `db:"BATCH_ID_"`
		TYPE         string         `db:"TYPE_"`
		SCOPEID      sql.NullString `db:"SCOPE_ID_"`
		SUBSCOPEID   sql.NullString `db:"SUB_SCOPE_ID_"`
		SCOPETYPE    sql.NullString `db:"SCOPE_TYPE_"`
		SEARCHKEY    sql.NullString `db:"SEARCH_KEY_"`
		SEARCHKEY2   sql.NullString `db:"SEARCH_KEY2_"`
		CREATETIME   time.Time      `db:"CREATE_TIME_"`
		COMPLETETIME sql.NullTime   `db:"COMPLETE_TIME_"`
		STATUS       sql.NullString `db:"STATUS_"`
		RESULTDOCID  sql.NullString `db:"RESULT_DOC_ID_"`
		TENANTID     string         `db:"TENANT_ID_"`
	}
)

func newFlwRuBatchPartModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultFlwRuBatchPartModel {
	return &defaultFlwRuBatchPartModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`flw_ru_batch_part`",
	}
}

func (m *defaultFlwRuBatchPartModel) Insert(ctx context.Context, data *FlwRuBatchPart) (sql.Result, error) {
	assetFlwRuBatchPartIDKey := fmt.Sprintf("%s%v", cacheAssetFlwRuBatchPartIDPrefix, data.ID)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, flwRuBatchPartRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.ID, data.REV, data.BATCHID, data.TYPE, data.SCOPEID, data.SUBSCOPEID, data.SCOPETYPE, data.SEARCHKEY, data.SEARCHKEY2, data.CREATETIME, data.COMPLETETIME, data.STATUS, data.RESULTDOCID, data.TENANTID)
	}, assetFlwRuBatchPartIDKey)
	return ret, err
}

func (m *defaultFlwRuBatchPartModel) FindOne(ctx context.Context, iD string) (*FlwRuBatchPart, error) {
	assetFlwRuBatchPartIDKey := fmt.Sprintf("%s%v", cacheAssetFlwRuBatchPartIDPrefix, iD)
	var resp FlwRuBatchPart
	err := m.QueryRowCtx(ctx, &resp, assetFlwRuBatchPartIDKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `ID_` = ? limit 1", flwRuBatchPartRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, iD)
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

func (m *defaultFlwRuBatchPartModel) Update(ctx context.Context, data *FlwRuBatchPart) error {
	assetFlwRuBatchPartIDKey := fmt.Sprintf("%s%v", cacheAssetFlwRuBatchPartIDPrefix, data.ID)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `ID_` = ?", m.table, flwRuBatchPartRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.REV, data.BATCHID, data.TYPE, data.SCOPEID, data.SUBSCOPEID, data.SCOPETYPE, data.SEARCHKEY, data.SEARCHKEY2, data.CREATETIME, data.COMPLETETIME, data.STATUS, data.RESULTDOCID, data.TENANTID, data.ID)
	}, assetFlwRuBatchPartIDKey)
	return err
}

func (m *defaultFlwRuBatchPartModel) Delete(ctx context.Context, iD string) error {
	assetFlwRuBatchPartIDKey := fmt.Sprintf("%s%v", cacheAssetFlwRuBatchPartIDPrefix, iD)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `ID_` = ?", m.table)
		return conn.ExecCtx(ctx, query, iD)
	}, assetFlwRuBatchPartIDKey)
	return err
}

func (m *defaultFlwRuBatchPartModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheAssetFlwRuBatchPartIDPrefix, primary)
}

func (m *defaultFlwRuBatchPartModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `ID_` = ? limit 1", flwRuBatchPartRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultFlwRuBatchPartModel) tableName() string {
	return m.table
}
