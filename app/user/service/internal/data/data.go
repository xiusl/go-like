package data

import (
	"database/sql"
	"github.com/bwmarrin/snowflake"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/v8"
	"github.com/google/wire"
	"go-like/app/user/service/internal/conf"

	_ "github.com/go-sql-driver/mysql"
)

var ProviderSet = wire.NewSet(
	NewData,
	NewUserRepo,
	NewVerifyCodeRepo)

type Data struct {
	db   *sql.DB
	rdb  *redis.Client
	node *snowflake.Node
	log  *log.Helper
}

func NewData(conf *conf.Data, logger log.Logger) (*Data, func(), error) {
	node, err := snowflake.NewNode(1)
	if err != nil {
		panic(err)
	}

	db, err := sql.Open(conf.Database.Driver, conf.Database.Source)
	logHelper := log.NewHelper(log.With(logger, "module", "user/data"))

	rdb := redis.NewClient(&redis.Options{
		Addr:     conf.Redis.Addr,
		Password: conf.Redis.Password,
		DB:       int(conf.Redis.Db),
	})

	d := &Data{
		db:   db,
		rdb:  rdb,
		node: node,
		log:  logHelper,
	}

	return d, func() {
		if err := d.db.Close(); err != nil {
			logHelper.Error(err)
		}
	}, err
}

func (data *Data) GenerateID() int64 {
	return data.node.Generate().Int64()
}
