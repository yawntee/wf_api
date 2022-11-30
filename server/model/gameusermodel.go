package model

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ GameUserModel = (*customGameUserModel)(nil)

type (
	// GameUserModel is an interface to be customized, add more methods here,
	// and implement the added methods in customGameUserModel.
	GameUserModel interface {
		gameUserModel
		UpdateOrInsert(ctx context.Context, user *GameUser) error
		GetInfoByUserId(ctx context.Context, userId int64) ([]GameUser, error)
		GetData(ctx context.Context, userId int64) ([]byte, error)
	}

	customGameUserModel struct {
		*defaultGameUserModel
	}
)

// NewGameUserModel returns a model for the database table.
func NewGameUserModel(conn sqlx.SqlConn) GameUserModel {
	return &customGameUserModel{
		defaultGameUserModel: newGameUserModel(conn),
	}
}

func (m *customGameUserModel) UpdateOrInsert(ctx context.Context, data *GameUser) error {
	query := fmt.Sprintf("update %s set %s where `id` = ? limit 1", m.table, gameUserRowsWithPlaceHolder)
	rs, err := m.conn.ExecCtx(ctx, query, data.User, data.Name, data.Channel, data.Data, data.Id)
	if err != nil {
		return errors.Wrapf(err, "%+v", data)
	}
	affected, err := rs.RowsAffected()
	if err != nil {
		return errors.WithStack(err)
	}
	if affected > 0 {
		return nil
	}
	query = fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?)", m.table, gameUserRowsExpectAutoSet)
	rs, err = m.conn.ExecCtx(ctx, query, data.Id, data.User, data.Name, data.Channel, data.Data)
	if err != nil {
		return err
	}
	return nil
}

func (m customGameUserModel) GetInfoByUserId(ctx context.Context, userId int64) ([]GameUser, error) {
	var gameUsers []GameUser
	query := fmt.Sprintf("select id, name, channel from %s where user = ?", m.table)
	err := m.conn.QueryRowsPartialCtx(ctx, &gameUsers, query, userId)
	if err != nil {
		return nil, err
	}
	return gameUsers, nil
}

func (m *customGameUserModel) GetData(ctx context.Context, gameUserId int64) ([]byte, error) {
	var gameUsers GameUser
	query := fmt.Sprintf("select data from %s where id = ? limit 1", m.table)
	err := m.conn.QueryRowPartialCtx(ctx, &gameUsers, query, gameUserId)
	if err != nil {
		return nil, err
	}
	return gameUsers.Data, nil
}
