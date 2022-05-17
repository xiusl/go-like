package data

import (
	"context"
	"database/sql"
	"fmt"
	"go-like/app/user/service/internal/biz"
	"strconv"
)

var (
	insertFollowerSql = "insert into user_follower (`user`, `follower`) value(?, ?)"
	pageFollowerSql   = "select `user`, `follower`, `created_at` from user_follower " +
		"where user=? order by created_at desc limit ?, ?"
	pageFollowingSql = "select `id`, `user`, `follower`, `created_at` from user_follower " +
		"where follower=? order by created_at desc limit ?, ?"
	listFollowingSql = "select `id`, `user`, `follower`, `created_at` from user_follower " +
		"where follower=? order by created_at desc"
	deleteFollowingSql = "delete from user_follower where user=? and follower=?"
)

type followerRepo struct {
	data *Data
}

// NewFollowerRepo is .
func NewFollowerRepo(data *Data) biz.FollowerRepo {
	return &followerRepo{
		data: data,
	}
}

// Insert is .
func (r *followerRepo) Insert(ctx context.Context, user, follower int64) error {
	res, err := r.data.db.ExecContext(ctx, insertFollowerSql, user, follower)
	if err != nil {
		return err
	}
	return scanInsert(res)
}

// PageFollower is .
func (r *followerRepo) PageFollower(ctx context.Context, uid int64, page, count int64) ([]*biz.UserFollower, error) {
	start := (page - 1) * count
	rows, err := r.data.db.QueryContext(ctx, pageFollowerSql, uid, start, count)
	if err != nil {
		return nil, err
	}
	return scanUserFollowers(rows)
}

// PageFollowing is .
func (r *followerRepo) PageFollowing(ctx context.Context, uid int64, page, count int64) ([]*biz.UserFollower, error) {
	start := (page - 1) * count
	rows, err := r.data.db.QueryContext(ctx, pageFollowingSql, uid, start, count)
	if err != nil {
		return nil, err
	}
	return scanUserFollowers(rows)
}

// Delete is .
func (r *followerRepo) Delete(ctx context.Context, user, follower int64) error {
	res, err := r.data.db.ExecContext(ctx, deleteFollowingSql, user, follower)
	if err != nil {
		return err
	}
	return scanInsert(res)
}

// Tx is .
func (r *followerRepo) Tx(ctx context.Context, userRepo biz.UserRepo, f func(ctx context.Context, repo biz.FollowerRepo, userRepo biz.UserRepo) error) error {
	tx, err := r.data.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	err = f(ctx, r, userRepo)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx error: %v, rb error: %v", err, rbErr)
		}
		return err
	}
	return nil
}

func (r *followerRepo) GetFollowing(ctx context.Context, uid int64) ([]int64, error) {
	key := fmt.Sprintf("user:following:%d", uid)
	if r.data.rdb.Exists(ctx, key).Val() == 1 {
		cmd := r.data.rdb.SMembers(ctx, key)
		res, err := cmd.Result()
		if err != nil {
			return nil, err
		}
		ids := make([]int64, len(res))
		for i, str := range res {
			id, err := strconv.ParseInt(str, 10, 64)
			if err != nil {
				return nil, err
			}
			ids[i] = id
		}
		return ids, nil
	}
	rows, err := r.data.db.QueryContext(ctx, listFollowingSql, uid)
	if err != nil {
		return nil, err
	}
	res, err := scanUserFollowers(rows)
	if err != nil {
		return nil, err
	}
	ids := make([]int64, len(res))
	for i, uf := range res {
		ids[i] = uf.UserId
	}
	r.data.rdb.SAdd(ctx, key, ids)
	return ids, nil
}

func scanUserFollowers(rows *sql.Rows) ([]*biz.UserFollower, error) {
	res := make([]*biz.UserFollower, 0)
	for rows.Next() {
		var uf biz.UserFollower
		err := rows.Scan(
			&uf.Id,
			&uf.UserId,
			&uf.FollowerId,
			&uf.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		res = append(res, &uf)
	}
	return res, nil
}

var followerTableSql = "-- 用户表\n" +
	"CREATE TABLE `user_follower`(\n" +
	"    `id`         bigint      NOT NULL AUTO_INCREMENT COMMENT '唯一标识',\n" +
	"    `user`       bigint      NOT NULL COMMENT '唯一标识',\n" +
	"    `follower`   bigint      NOT NULL COMMENT '唯一标识',\n" +
	"    `created_at` timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',\n" +
	"    `updated_at` timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',\n" +
	"    `del`        int(1)       NOT NULL DEFAULT '0' COMMENT '逻辑删除',\n" +
	"    `version`    int(1)       NOT NULL DEFAULT '0' COMMENT '乐观锁',\n" +
	"    PRIMARY KEY (`id`),\n" +
	"    KEY `user_follower` (`user`, `follower`) USING BTREE\n" +
	") ENGINE = InnoDB\n" +
	"  DEFAULT CHARSET = utf8;"
