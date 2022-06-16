package main

import (
	"flag"
	"fmt"
	"net/http"
	"reflect"

	"cyntra/common/errorx"
	"cyntra/gateway/api/internal/config"
	"cyntra/gateway/api/internal/handler"
	"cyntra/gateway/api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/rest/httpx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var configFile = flag.String("f", "etc/public-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c, conf.UseEnv())

	ctx := svc.NewServiceContext(c)
	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	handler.RegisterHandlers(server, ctx)

	httpx.SetErrorHandler(func(err error) (int, interface{}) {

		fmt.Println("Error type", reflect.TypeOf(err))
		st, ok := status.FromError(err)

		if ok {
			codeResponse := errorx.CodeResponse{Code: st.Code(), Msg: st.Message(), Data: nil}
			return http.StatusOK, codeResponse.Response()
		} else {
			switch v := err.(type) {
			case *errorx.CodeResponse:
				return http.StatusUnprocessableEntity, v.Response()
			default:
				codeResponse := errorx.CodeResponse{Code: codes.Internal, Msg: codes.Internal.String(), Data: "Internal error occurred"}
				return http.StatusInternalServerError, codeResponse.Response()
			}
			// return http.StatusBadRequest, err
		}
	})

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
