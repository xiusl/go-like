package server

import (
    "github.com/go-kratos/kratos/v2/log"
    "github.com/go-kratos/kratos/v2/middleware/ratelimit"
    "github.com/go-kratos/kratos/v2/middleware/recovery"
    "github.com/go-kratos/kratos/v2/transport/http"
    "github.com/gorilla/handlers"
    "go-like/app/interface/service/internal/conf"
    "go-like/app/interface/service/internal/service"
)

// NewHTTPServer new HTTP server.
func NewHTTPServer(c *conf.Server, s *service.InterfaceService, logger log.Logger) *http.Server {
    var opts = []http.ServerOption{
        http.Middleware(
            recovery.Recovery(),
            ratelimit.Server(),
        ),
        http.Filter(handlers.CORS(
            handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
            handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "DELETE", "PATCH", "OPTIONS"}),
            handlers.AllowedOrigins([]string{"*"}),
            handlers.MaxAge(600),
        )),
        http.Logger(logger),
    }
    if c.Http.Network != "" {
        opts = append(opts, http.Network(c.Http.Network))
    }
    if c.Http.Addr != "" {
        opts = append(opts, http.Address(c.Http.Addr))
    }
    if c.Http.Timeout != nil {
        opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
    }
    srv := http.NewServer(opts...)
    srv.HandlePrefix("/", s)
    return srv
}
