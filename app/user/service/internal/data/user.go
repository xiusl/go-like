package data

import (
	"context"
	"database/sql"
	"go-like/app/user/service/internal/biz"
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
	"    `created_at` timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',\n" +
	"    `updated_at` timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',\n" +
	"    `del`        int(1)       NOT NULL DEFAULT '0' COMMENT '逻辑删除',\n" +
	"    `version`    int(1)       NOT NULL DEFAULT '0' COMMENT '乐观锁',\n" +
	"    PRIMARY KEY (`id`),\n" +
	"    KEY `mobile` (`mobile`) USING BTREE\n" +
	") ENGINE = InnoDB\n" +
	"  DEFAULT CHARSET = utf8;"

var (
	selectUserSql         = "select `id`, `mobile`, `name` from user where id=? and del = 0"
	selectUserByMobileSql = "select `id`, `mobile`, `name` from user where mobile=? and del = 0"
	insertUserSql         = "insert into user (id, mobile, name) values (?,?,?)"
)

func scanUser(row *sql.Row) (*biz.User, error) {
	var u biz.User
	err := row.Scan(
		&u.Id,
		&u.Mobile,
		&u.Name,
	)
	if err != nil {
		return nil, err
	}
	return &u, nil
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
