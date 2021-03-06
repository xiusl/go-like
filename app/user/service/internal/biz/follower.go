package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/metadata"
	"strconv"
	"time"
)

// UserFollower is .
type UserFollower struct {
	Id         int64
	UserId     int64
	FollowerId int64
	CreatedAt  time.Time
}

// FollowerRepo is .
type FollowerRepo interface {
	Insert(ctx context.Context, user, follower int64) error
	PageFollower(ctx context.Context, uid int64, page, count int64) ([]*UserFollower, error)
	PageFollowing(ctx context.Context, uid int64, page, count int64) ([]*UserFollower, error)
	Delete(ctx context.Context, user, follower int64) error
	GetFollowing(ctx context.Context, uid int64) ([]int64, error)

	Tx(ctx context.Context, userRepo UserRepo, f func(ctx context.Context, repo FollowerRepo, userRepo UserRepo) error) error
}

// FollowUser is .
func (uc *UserUseCase) FollowUser(ctx context.Context, user, follower int64) error {
	u, err := uc.repo.Select(ctx, user)
	if err != nil {
		return err
	}

	err = uc.fRepo.Tx(ctx, uc.repo, func(ctx context.Context, repo FollowerRepo, userRepo UserRepo) error {
		err := repo.Insert(ctx, user, follower)
		if err != nil {
			return err
		}
		if err := userRepo.UpdateFollowerCount(ctx, u.Id, u.FollowerCount+1); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

// UnFollowUser is .
func (uc *UserUseCase) UnFollowUser(ctx context.Context, user, follower int64) error {
	u, err := uc.repo.Select(ctx, user)
	if err != nil {
		return err
	}

	err = uc.fRepo.Tx(ctx, uc.repo, func(ctx context.Context, repo FollowerRepo, userRepo UserRepo) error {
		err := repo.Delete(ctx, user, follower)
		if err != nil {
			return err
		}
		if err := userRepo.UpdateFollowerCount(ctx, u.Id, u.FollowerCount-1); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

// GetFollower is .
func (uc *UserUseCase) GetFollower(ctx context.Context, uid int64, page, count int64) ([]*User, error) {
	followers, err := uc.fRepo.PageFollower(ctx, uid, page, count)
	if err != nil {
		return nil, err
	}
	ids := getFollowerIds(followers)
	userMap, err := uc.repo.MapByIds(ctx, ids)
	if err != nil {
		return nil, err
	}

	currentId := getCurrentUser(ctx)
	followingMap := uc.getFollowingMap(ctx, currentId)

	users := make([]*User, 0)
	for _, id := range ids {
		if user, ex := userMap[id]; ex {
			if _, ok := followingMap[user.Id]; ok {
				user.IsFollowing = true
			}

			followingMap := uc.getFollowingMap(ctx, user.Id)
			if _, ok := followingMap[currentId]; ok {
				user.IsFollowed = true
			}
			users = append(users, user)
		}
	}
	return users, nil
}

// GetFollowing is .
func (uc *UserUseCase) GetFollowing(ctx context.Context, uid int64, page, count int64) ([]*User, error) {
	followers, err := uc.fRepo.PageFollowing(ctx, uid, page, count)
	if err != nil {
		return nil, err
	}
	ids := getFollowingIds(followers)
	userMap, err := uc.repo.MapByIds(ctx, ids)
	if err != nil {
		return nil, err
	}

	currentId := getCurrentUser(ctx)
	followingMap := uc.getFollowingMap(ctx, currentId)

	users := make([]*User, 0)
	for _, id := range ids {
		if user, ex := userMap[id]; ex {

			if _, ok := followingMap[user.Id]; ok {
				user.IsFollowing = true
			}

			followingMap := uc.getFollowingMap(ctx, user.Id)
			if _, ok := followingMap[currentId]; ok {
				user.IsFollowed = true
			}

			users = append(users, user)
		}
	}
	return users, nil
}

func (uc *UserUseCase) getFollowingMap(ctx context.Context, uid int64) map[int64]bool {
	followingMap := make(map[int64]bool)
	followings, err := uc.fRepo.GetFollowing(ctx, uid)
	if err != nil {
		return followingMap
	}

	for _, id := range followings {
		followingMap[id] = true
	}
	return followingMap
}

func getCurrentUser(ctx context.Context) int64 {
	if md, ok := metadata.FromServerContext(ctx); ok {
		idStr := md.Get("user_id")
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			return 0
		}
		return id
	}
	return 0
}

func sortMapUsers(ids []int64, userMap map[int64]*User) []*User {
	users := make([]*User, 0)
	for _, id := range ids {
		if user, ex := userMap[id]; ex {
			users = append(users, user)
		}
	}
	return users
}
func getFollowerIds(followers []*UserFollower) []int64 {
	ids := make([]int64, len(followers))
	for i, f := range followers {
		ids[i] = f.FollowerId
	}
	return ids
}
func getFollowingIds(followers []*UserFollower) []int64 {
	ids := make([]int64, len(followers))
	for i, f := range followers {
		ids[i] = f.UserId
	}
	return ids
}
