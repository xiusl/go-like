package data

import (
	"database/sql"
	"go-like/app/article/service/internal/conf"

	"github.com/bwmarrin/snowflake"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(NewData, NewArticleRepo)

// Data is .
type Data struct {
	db   *sql.DB
	node *snowflake.Node
	log  *log.Helper
}

// NewData is .
func NewData(conf *conf.Data, logger log.Logger) (*Data, func(), error) {
	node, err := snowflake.NewNode(1)
	if err != nil {
		panic(err)
	}

	db, err := sql.Open(conf.Database.Driver, conf.Database.Source)
	logHelper := log.NewHelper(log.With(logger, "module", "user/data"))

	d := &Data{
		db:   db,
		node: node,
		log:  logHelper,
	}

	return d, func() {
		if err := d.db.Close(); err != nil {
			logHelper.Error(err)
		}
	}, err
}

// GenerateID is .
func (data *Data) GenerateID() int64 {
	return data.node.Generate().Int64()
}
