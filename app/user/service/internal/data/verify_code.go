package data

import (
	"context"
	"database/sql"
	"errors"
	"go-like/app/user/service/internal/biz"
)

type verifyCodeRepo struct {
	data *Data
}

func NewVerifyCodeRepo(data *Data) biz.VerifyCodeRepo {
	return &verifyCodeRepo{
		data: data,
	}
}

var verifyCodeTableSql = "-- 用户短信表\n" +
	"CREATE TABLE `verify_code` (\n" +
	"    `id`         bigint      NOT NULL AUTO_INCREMENT COMMENT '唯一标识',\n" +
	"    `key`        varchar(64) NOT NULL DEFAULT '' COMMENT '手机号',\n" +
	"    `code`       varchar(10) NOT NULL DEFAULT '' COMMENT '验证码',\n" +
	"    `biz_type`   int(1)      NOT NULL DEFAULT '0' COMMENT '业务类型，0: 登录/注册;',\n" +
	"    `expired_at` timestamp   NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '过期时间',\n" +
	"    `created_at` timestamp   NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',\n" +
	"    `updated_at` timestamp   NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',\n" +
	"    `del`        int(1)      NOT NULL DEFAULT '0' COMMENT '逻辑删除',\n" +
	"    `version`    int(1)      NOT NULL DEFAULT '0' COMMENT '乐观锁',\n" +
	"    PRIMARY KEY (`id`)\n" +
	") ENGINE = InnoDB\n" +
	"  DEFAULT CHARSET = utf8;"
var (
	selectVerifyCodeSql      = "select `id`, `key`, `code`, `biz_type`, `expired_at` from verify_code where id=?"
	selectVerifyCodeByKeySql = "select `id`, `key`, `code`, `biz_type`, `expired_at` from verify_code " +
		"where `key`=? and `biz_type`=? and del=0 order by expired_at desc limit 1"
	insertVerifyCodeSql = "insert into verify_code (`key`, `code`, biz_type, expired_at) values (?,?,?,?)"
)

func scanVerifyCode(row *sql.Row) (*biz.VerifyCode, error) {
	var c biz.VerifyCode
	err := row.Scan(
		&c.Id,
		&c.Key,
		&c.Code,
		&c.BizType,
		&c.ExpiredAt,
	)
	if err != nil {
		return nil, err
	}
	return &c, nil
}

func scanInsert(res sql.Result) error {
	row, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if row == 0 {
		return errors.New("insert none")
	}
	return nil
}

func (r *verifyCodeRepo) Select(ctx context.Context, id int64) (*biz.VerifyCode, error) {
	row := r.data.db.QueryRowContext(ctx, selectVerifyCodeSql, id)
	return scanVerifyCode(row)
}

func (r *verifyCodeRepo) SelectByKey(ctx context.Context, key string, bizType int64) (*biz.VerifyCode, error) {
	row := r.data.db.QueryRowContext(ctx, selectVerifyCodeByKeySql, key, bizType)
	return scanVerifyCode(row)
}

func (r *verifyCodeRepo) Insert(ctx context.Context, c *biz.VerifyCode) error {
	res, err := r.data.db.ExecContext(ctx, insertVerifyCodeSql, c.Key, c.Code, c.BizType, c.ExpiredAt)
	if err != nil {
		return err
	}
	return scanInsert(res)
}
