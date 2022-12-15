// Code generated by goctl. DO NOT EDIT!

package model

import (
	"context"
	"fmt"

	"database/sql"

	"github.com/SpectatorNan/gorm-zero/gormc"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"gorm.io/gorm"
)

var (
	cacheGormzeroUsersIdPrefix = "cache:gormzero:users:id:"
)

type (
	usersModel interface {
		Insert(ctx context.Context, tx *gorm.DB, data *Users) error

		FindOne(ctx context.Context, id int64) (*Users, error)
		Update(ctx context.Context, tx *gorm.DB, data *Users) error

		Delete(ctx context.Context, tx *gorm.DB, id int64) error
		Transaction(ctx context.Context, fn func(db *gorm.DB) error) error
	}

	defaultUsersModel struct {
		gormc.CachedConn
		table string
	}

	Users struct {
		Id       int64          `gorm:"column:id"`
		Account  sql.NullString `gorm:"column:account"`
		NickName sql.NullString `gorm:"column:nick_name"`
		Password sql.NullString `gorm:"column:password"`
		CreateAt sql.NullTime   `gorm:"column:create_at"`
		UpdateAt sql.NullTime   `gorm:"column:update_at"`
		DeleteAt sql.NullTime   `gorm:"column:delete_at"`
	}
)

func (Users) TableName() string {
	return "`users`"
}

func newUsersModel(conn *gorm.DB, c cache.CacheConf) *defaultUsersModel {
	return &defaultUsersModel{
		CachedConn: gormc.NewConn(conn, c),
		table:      "`users`",
	}
}

func (m *defaultUsersModel) Insert(ctx context.Context, tx *gorm.DB, data *Users) error {

	err := m.ExecCtx(ctx, func(conn *gorm.DB) error {
		db := conn
		if tx != nil {
			db = tx
		}
		return db.Save(&data).Error
	}, m.getCacheKeys(data)...)
	return err
}

func (m *defaultUsersModel) FindOne(ctx context.Context, id int64) (*Users, error) {
	gormzeroUsersIdKey := fmt.Sprintf("%s%v", cacheGormzeroUsersIdPrefix, id)
	var resp Users
	err := m.QueryCtx(ctx, &resp, gormzeroUsersIdKey, func(conn *gorm.DB, v interface{}) error {
		return conn.Model(&Users{}).Where("`id` = ?", id).First(&resp).Error
	})
	switch err {
	case nil:
		return &resp, nil
	case gormc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultUsersModel) Update(ctx context.Context, tx *gorm.DB, data *Users) error {
	old, err := m.FindOne(ctx, data.Id)
	if err != nil && err != ErrNotFound {
		return err
	}
	err = m.ExecCtx(ctx, func(conn *gorm.DB) error {
		db := conn
		if tx != nil {
			db = tx
		}
		return db.Save(data).Error
	}, m.getCacheKeys(old)...)
	return err
}

func (m *defaultUsersModel) getCacheKeys(data *Users) []string {
	if data == nil {
		return []string{}
	}
	gormzeroUsersIdKey := fmt.Sprintf("%s%v", cacheGormzeroUsersIdPrefix, data.Id)
	cacheKeys := []string{
		gormzeroUsersIdKey,
	}
	cacheKeys = append(cacheKeys, m.customCacheKeys(data)...)
	return cacheKeys
}

func (m *defaultUsersModel) Delete(ctx context.Context, tx *gorm.DB, id int64) error {
	data, err := m.FindOne(ctx, id)
	if err != nil {
		if err == ErrNotFound {
			return nil
		}
		return err
	}
	err = m.ExecCtx(ctx, func(conn *gorm.DB) error {
		db := conn
		if tx != nil {
			db = tx
		}
		return db.Delete(&Users{}, id).Error
	}, m.getCacheKeys(data)...)
	return err
}

func (m *defaultUsersModel) Transaction(ctx context.Context, fn func(db *gorm.DB) error) error {
	return m.TransactCtx(ctx, fn)
}

func (m *defaultUsersModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheGormzeroUsersIdPrefix, primary)
}

func (m *defaultUsersModel) queryPrimary(conn *gorm.DB, v, primary interface{}) error {
	return conn.Model(&Users{}).Where("`id` = ?", primary).Take(v).Error
}

func (m *defaultUsersModel) tableName() string {
	return m.table
}