package main

import (
	"context"
	"cyntra/user/rpc/internal/config"
	"cyntra/user/rpc/internal/server"
	"cyntra/user/rpc/internal/svc"
	"cyntra/user/rpc/types/user"
	"flag"
	"fmt"
	"time"

	"github.com/bxcodec/faker/v3"
	"github.com/zeromicro/go-zero/core/conf"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type FakeUser struct {
	Username  string `faker:"email"`
	Password  string
	FirstName string `faker:"first_name"`
	LastName  string `faker:"last_name"`
	Gender    string `faker:"oneof: MALE, FEMALE"`
}

var configFile = flag.String("f", "etc/user.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c, conf.UseEnv())
	ctx := svc.NewServiceContext(c)
	svr := server.NewUserServer(ctx)

	fakeUser := FakeUser{}

	err := faker.FakeData(&fakeUser)

	req := &user.RegisterRequest{
		Username:  fakeUser.Username,
		Password:  "secret",
		FirstName: fakeUser.FirstName,
		LastName:  fakeUser.LastName,
		Gender:    fakeUser.Gender,
		Dob:       timestamppb.New(time.Now()),
	}

	resp, err := svr.RegisterUser(context.Background(), req)

	if err != nil {
		fmt.Println(fmt.Errorf("HelloTest(%v) got unexpected error", err))
	} else {
		fmt.Println(resp.Username, resp.FirstName, resp.LastName)
	}
}
