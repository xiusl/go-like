package data

import (
	"context"
	"database/sql"
	"fmt"
	"go-like/app/user/service/internal/biz"
	"go-like/pkg/xsql"
)

type userRepo struct {
	data *Data
}

func NewUserRepo(data *Data) biz.UserRepo {
	return &userRepo{
		data: data,
	}
}

var userTableSql = "-- 用户表\n" +
	"CREATE TABLE `user`(\n" +
	"    `id`         bigint(20)   NOT NULL COMMENT '唯一标识',\n" +
	"    `mobile`     varchar(20)  NOT NULL DEFAULT '' COMMENT '手机号',\n" +
	"    `name`       varchar(48)  NOT NULL DEFAULT '' COMMENT '昵称',\n" +
	"    `avatar`     varchar(255) NOT NULL DEFAULT '' COMMENT '头像，只存path',\n" +
	"    `follower_count` int      NOT NULL DEFAULT '0' COMMENT '粉丝数',\n" +
	"    `created_at` timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',\n" +
	"    `updated_at` timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',\n" +
	"    `del`        int(1)       NOT NULL DEFAULT '0' COMMENT '逻辑删除',\n" +
	"    `version`    int(1)       NOT NULL DEFAULT '0' COMMENT '乐观锁',\n" +
	"    PRIMARY KEY (`id`),\n" +
	"    KEY `mobile` (`mobile`) USING BTREE\n" +
	") ENGINE = InnoDB\n" +
	"  DEFAULT CHARSET = utf8;"

var (
	selectUserSql              = "select `id`, `mobile`, `name`, `follower_count` from user where id=? and del = 0"
	selectUserByMobileSql      = "select `id`, `mobile`, `name`, `follower_count` from user where mobile=? and del = 0"
	insertUserSql              = "insert into user (id, mobile, name) values (?,?,?)"
	updateUserFollowerCountSql = "update user set follower_count=? where id=?"
	listUserByIdsSql           = "select `id`, `mobile`, `name`, `follower_count` from user where id in (%s)"
)

func scanUser(row *sql.Row) (*biz.User, error) {
	var u biz.User
	err := row.Scan(
		&u.Id,
		&u.Mobile,
		&u.Name,
		&u.FollowerCount,
	)
	if err != nil {
		return nil, err
	}
	return &u, nil
}
func scanUsers(rows *sql.Rows) ([]*biz.User, error) {
	us := make([]*biz.User, 0)
	for rows.Next() {
		var u biz.User
		err := rows.Scan(
			&u.Id,
			&u.Mobile,
			&u.Name,
			&u.FollowerCount,
		)
		if err != nil {
			return nil, err
		}
		us = append(us, &u)
	}
	return us, nil
}

func (r *userRepo) Select(ctx context.Context, id int64) (*biz.User, error) {
	row := r.data.db.QueryRowContext(ctx, selectUserSql, id)
	return scanUser(row)
}

func (r *userRepo) SelectByMobile(ctx context.Context, mobile string) (*biz.User, error) {
	row := r.data.db.QueryRowContext(ctx, selectUserByMobileSql, mobile)
	return scanUser(row)
}

func (r *userRepo) Insert(ctx context.Context, u *biz.User) (int64, error) {
	u.Id = r.data.GenerateID()
	res, err := r.data.db.ExecContext(ctx, insertUserSql, u.Id, u.Mobile, u.Name)
	if err != nil {
		return 0, err
	}
	return u.Id, scanInsert(res)
}

func (r *userRepo) UpdateFollowerCount(ctx context.Context, uid int64, count int64) error {
	res, err := r.data.db.ExecContext(ctx, updateUserFollowerCountSql, count, uid)
	if err != nil {
		return err
	}
	return scanInsert(res)
}

func (r *userRepo) ListByIds(ctx context.Context, ids []int64) ([]*biz.User, error) {
	sql := fmt.Sprintf(listUserByIdsSql, xsql.IdsToStr(ids))
	rows, err := r.data.db.QueryContext(ctx, sql)
	if err != nil {
		return nil, err
	}
	return scanUsers(rows)
}

func (r *userRepo) MapByIds(ctx context.Context, ids []int64) (map[int64]*biz.User, error) {
	users, err := r.ListByIds(ctx, ids)
	if err != nil {
		return nil, err
	}
	m := make(map[int64]*biz.User)
	for _, user := range users {
		m[user.Id] = user
	}
	return m, nil
}
