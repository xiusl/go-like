package data

import (
	"context"
	nacosAPI "github.com/go-kratos/kratos/contrib/registry/nacos/v2"
	"github.com/go-kratos/kratos/v2/log"
	mmd "github.com/go-kratos/kratos/v2/middleware/metadata"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/google/wire"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	artv1 "go-like/api/article/service/v1"
	userv1 "go-like/api/user/service/v1"
	"go-like/app/interface/service/internal/conf"
)

var ProviderSet = wire.NewSet(
	NewData,
	NewDiscovery,
	NewUserClient,
	NewUserRepo,
	NewArticleRepo,
	NewArticleClient,
)

type Data struct {
	log *log.Helper

	uc userv1.UserClient
	ac artv1.ArticleClient
}

func NewData(conf *conf.Data, uc userv1.UserClient, ac artv1.ArticleClient, logger log.Logger) *Data {
	data := &Data{
		uc:  uc,
		ac:  ac,
		log: log.NewHelper(log.With(logger, "module", "interface/data")),
	}
	return data
}

func NewDiscovery(conf *conf.Registry) registry.Discovery {
	sc := []constant.ServerConfig{
		*constant.NewServerConfig(conf.GetNacos().Address, conf.GetNacos().Port),
	}
	cc := &constant.ClientConfig{
		NamespaceId:         "public",
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogLevel:            "debug",
	}
	client, err := clients.NewNamingClient(
		vo.NacosClientParam{
			ServerConfigs: sc,
			ClientConfig:  cc,
		},
	)
	if err != nil {
		panic(err)
	}
	return nacosAPI.New(client)
}

func NewUserClient(r registry.Discovery) userv1.UserClient {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("discovery:///go-like.user.service.grpc"),
		grpc.WithDiscovery(r),
		grpc.WithMiddleware(
			recovery.Recovery(),
			mmd.Client(),
		),
	)
	if err != nil {
		panic(err)
	}
	c := userv1.NewUserClient(conn)
	return c
}

func NewArticleClient(r registry.Discovery) artv1.ArticleClient {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("discovery:///go-like.article.service.grpc"),
		grpc.WithDiscovery(r),
		grpc.WithMiddleware(
			recovery.Recovery(),
			mmd.Client(),
		),
	)
	if err != nil {
		panic(err)
	}
	c := artv1.NewArticleClient(conn)
	return c
}
